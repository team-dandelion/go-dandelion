package analysis

import (
	"bytes"
	routing "github.com/gly-hub/fasthttp-routing"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

var defaultMetricPath = "/metrics"

var reqCnt = &Metric{
	ID:          "reqCnt",
	Name:        "requests_total",
	Description: "How many HTTP requests processed, partitioned by status code and HTTP method.",
	Type:        "counter_vec",
	Args:        []string{"code", "method", "handler", "host", "url"}}

var reqDur = &Metric{
	ID:          "reqDur",
	Name:        "request_duration_seconds",
	Description: "The HTTP request latencies in seconds.",
	Type:        "histogram_vec",
	Args:        []string{"code", "method", "url"},
}

var resSz = &Metric{
	ID:          "resSz",
	Name:        "response_size_bytes",
	Description: "The HTTP response sizes in bytes.",
	Type:        "summary"}

var reqSz = &Metric{
	ID:          "reqSz",
	Name:        "request_size_bytes",
	Description: "The HTTP request sizes in bytes.",
	Type:        "summary"}

var standardMetrics = []*Metric{
	reqCnt,
	reqDur,
	resSz,
	reqSz,
}

type RequestCounterURLLabelMappingFn func(c *routing.Context) string

type Metric struct {
	MetricCollector prometheus.Collector
	ID              string
	Name            string
	Description     string
	Type            string
	Args            []string
}

type Prometheus struct {
	reqCnt                  *prometheus.CounterVec
	reqDur                  *prometheus.HistogramVec
	reqSz, resSz            prometheus.Summary
	Ppg                     PrometheusPushGateway
	MetricsList             []*Metric
	MetricsPath             string
	ReqCntURLLabelMappingFn RequestCounterURLLabelMappingFn
	URLLabelFromContext     string
}

type PrometheusPushGateway struct {
	PushIntervalSeconds time.Duration
	PushGatewayURL      string
	MetricsURL          string
	Job                 string
}

func NewPrometheus(subsystem string, customMetricsList ...[]*Metric) *Prometheus {

	var metricsList []*Metric

	if len(customMetricsList) > 1 {
		panic("Too many args. NewPrometheus( string, <optional []*Metric> ).")
	} else if len(customMetricsList) == 1 {
		metricsList = customMetricsList[0]
	}

	for _, metric := range standardMetrics {
		metricsList = append(metricsList, metric)
	}

	p := &Prometheus{
		MetricsList: metricsList,
		MetricsPath: defaultMetricPath,
		ReqCntURLLabelMappingFn: func(c *routing.Context) string {
			return string(c.Path()) // i.e. by default do nothing, i.e. return URL as is
		},
	}

	p.registerMetrics(subsystem)

	return p
}

func (p *Prometheus) registerMetrics(subsystem string) {

	for _, metricDef := range p.MetricsList {
		metric := NewMetric(metricDef, subsystem)
		if err := prometheus.Register(metric); err != nil {
			log.WithError(err).Errorf("%s could not be registered in Prometheus", metricDef.Name)
		}
		switch metricDef {
		case reqCnt:
			p.reqCnt = metric.(*prometheus.CounterVec)
		case reqDur:
			p.reqDur = metric.(*prometheus.HistogramVec)
		case resSz:
			p.resSz = metric.(prometheus.Summary)
		case reqSz:
			p.reqSz = metric.(prometheus.Summary)
		}
		metricDef.MetricCollector = metric
	}
}

func NewMetric(m *Metric, subsystem string) prometheus.Collector {
	var metric prometheus.Collector
	switch m.Type {
	case "counter_vec":
		metric = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case "counter":
		metric = prometheus.NewCounter(
			prometheus.CounterOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	case "gauge_vec":
		metric = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case "gauge":
		metric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	case "histogram_vec":
		metric = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case "histogram":
		metric = prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	case "summary_vec":
		metric = prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case "summary":
		metric = prometheus.NewSummary(
			prometheus.SummaryOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	}
	return metric
}

func (p *Prometheus) Middleware() routing.Handler {
	return func(c *routing.Context) error {
		if string(c.Path()) == p.MetricsPath {
			return c.Next()
		}

		start := time.Now()
		reqSz := computeApproximateRequestSize(c.RequestCtx)

		_ = c.Next()

		status := strconv.Itoa(c.Response.StatusCode())
		elapsed := float64(time.Since(start)) / float64(time.Second)
		resSz := float64(len(c.Response.Body()))

		url := p.ReqCntURLLabelMappingFn(c)
		if len(p.URLLabelFromContext) > 0 {
			u := c.Get(p.URLLabelFromContext)
			if u == nil {
				u = "unknown"
			}
			url = u.(string)
		}
		p.reqDur.WithLabelValues(status, string(c.Request.Header.Method()), url).Observe(elapsed)
		p.reqCnt.WithLabelValues(status, string(c.Request.Header.Method()), c.Request.String(), string(c.Request.Header.Host()), url).Inc()
		p.reqSz.Observe(float64(reqSz))
		p.resSz.Observe(resSz)
		return nil
	}
}

func (p *Prometheus) HandlerFunc() http.Handler {
	return promhttp.Handler()
}

func (p *Prometheus) DefaultPath() string {
	return p.MetricsPath
}

func (p *Prometheus) SetPushGateway(pushGatewayURL, metricsURL string, pushIntervalSeconds time.Duration) {
	p.Ppg.PushGatewayURL = pushGatewayURL
	p.Ppg.MetricsURL = metricsURL
	p.Ppg.PushIntervalSeconds = pushIntervalSeconds
	p.startPushTicker()
}

func (p *Prometheus) SetPushGatewayJob(j string) {
	p.Ppg.Job = j
}

func (p *Prometheus) getPushGatewayURL() string {
	h, _ := os.Hostname()
	if p.Ppg.Job == "" {
		p.Ppg.Job = "http-server"
	}
	return p.Ppg.PushGatewayURL + "/metrics/job/" + p.Ppg.Job + "/instance/" + h
}

func (p *Prometheus) sendMetricsToPushGateway(metrics []byte) {
	req, err := http.NewRequest("POST", p.getPushGatewayURL(), bytes.NewBuffer(metrics))
	client := &http.Client{}
	if _, err = client.Do(req); err != nil {
		log.WithError(err).Errorln("Error sending to push gateway")
	}
}

func (p *Prometheus) startPushTicker() {
	ticker := time.NewTicker(time.Second * p.Ppg.PushIntervalSeconds)
	go func() {
		for range ticker.C {
			p.sendMetricsToPushGateway(p.getMetrics())
		}
	}()
}

func (p *Prometheus) getMetrics() []byte {
	response, _ := http.Get(p.Ppg.MetricsURL)

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	return body
}

func computeApproximateRequestSize(r *fasthttp.RequestCtx) int {
	s := 0
	if r.URI() != nil {
		s = len(string(r.URI().Path()))
	}

	s += len(string(r.Request.Header.Method()))
	s += len(string(r.Request.Header.Protocol()))

	headers := make(map[string]string)
	r.Request.Header.VisitAll(func(k, v []byte) {
		headers[string(k)] = string(v)
	})

	for name, values := range headers {
		s += len(name)
		s += len(values)
	}
	s += len(string(r.Host()))

	// N.B. r.Form and r.MultipartForm are assumed to be included in r.URL.

	if r.Request.Header.ContentLength() != -1 {
		s += r.Request.Header.ContentLength()
	}
	return s
}
