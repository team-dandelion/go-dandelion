package config

type Tracer struct {
	OpenTrace bool   `json:"open_trace" yaml:"openTrace"`
	TraceName string `json:"trace_name" yaml:"traceName"`
	Host      string `json:"host" yaml:"host"`
}
