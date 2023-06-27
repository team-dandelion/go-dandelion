package swagger

import (
	routing "github.com/gly-hub/fasthttp-routing"
	"html/template"
	"net/http"
	"path/filepath"
	"regexp"
	"sync"

	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/swag"
)

// Config stores echoSwagger configuration variables.
type Config struct {
	// The url pointing to API definition (normally swagger.json or swagger.yaml). Default is `mockedSwag.json`.
	URL                  string
	InstanceName         string
	DocExpansion         string
	DomID                string
	DeepLinking          bool
	PersistAuthorization bool
}

// URL presents the url pointing to API definition (normally swagger.json or swagger.yaml).
func URL(url string) func(c *Config) {
	return func(c *Config) {
		c.URL = url
	}
}

// DeepLinking true, false.
func DeepLinking(deepLinking bool) func(c *Config) {
	return func(c *Config) {
		c.DeepLinking = deepLinking
	}
}

// DocExpansion list, full, none.
func DocExpansion(docExpansion string) func(c *Config) {
	return func(c *Config) {
		c.DocExpansion = docExpansion
	}
}

// DomID #swagger-ui.
func DomID(domID string) func(c *Config) {
	return func(c *Config) {
		c.DomID = domID
	}
}

// InstanceName set the instance name that was used to generate the swagger documents
// Defaults to swag.Name ("swagger").
func InstanceName(name string) func(*Config) {
	return func(c *Config) {
		c.InstanceName = name
	}
}

// PersistAuthorization Persist authorization information over browser close/refresh.
// Defaults to false.
func PersistAuthorization(persistAuthorization bool) func(c *Config) {
	return func(c *Config) {
		c.PersistAuthorization = persistAuthorization
	}
}

// WrapHandler wraps swaggerFiles.Handler and returns echo.HandlerFunc
var WrapHandler = RoutingWrapHandler()

// FiberWrapHandler wraps `http.Handler` into `fiber.Handler`.
func RoutingWrapHandler(configFns ...func(c *Config)) routing.Handler {
	var once sync.Once

	handler := swaggerFiles.Handler

	config := Config{
		URL:          "doc.json",
		DocExpansion: "list",
		DomID:        "swagger-ui",
		InstanceName: swag.Name,
		DeepLinking:  true,
	}

	for _, configFn := range configFns {
		configFn(&config)
	}

	// create a template with name
	index, _ := template.New("swagger_index.html").Parse(indexTemplate)

	var re = regexp.MustCompile(`^(.*/)([^?].*)?[?|.]*$`)

	return func(ctx *routing.Context) error {
		matches := re.FindStringSubmatch(string(ctx.Request.URI().Path()))
		path := matches[2]

		once.Do(func() {
			handler.Prefix = matches[1]
		})

		fileExt := filepath.Ext(path)
		switch path {
		case "":
			return ctx.Redirect(filepath.Join(handler.Prefix, "index.html"))

		case "index.html":
			ctx.Type(fileExt[0:], "utf-8")

			return index.Execute(ctx, config)
		case "doc.json":
			doc, err := swag.ReadDoc(config.InstanceName)
			if err != nil {
				_, err := ctx.Status(http.StatusInternalServerError).WriteString(http.StatusText(http.StatusInternalServerError))

				return err
			}

			ctx.Type(fileExt[0:], "utf-8")
			return ctx.SendString(doc)
		default:
			switch fileExt {
			case ".css":
				ctx.Type(fileExt[0:], "utf-8")
			case ".png", ".js":
				ctx.Type(fileExt[0:])
			}

			return nil
		}
	}
}

const indexTemplate = `<html>

<head>

<title>API Docs</title>

<!-- needed for mobile devices -->

<meta name="viewport" content="width=device-width, initial-scale=1">

</head>

<body>

<redoc spec-url="./doc.json"></redoc>

<script src="https://rebilly.github.io/ReDoc/releases/latest/redoc.min.js"></script>

</body>

</html>
`
