package swagger

import (
	"bufio"
	"bytes"
	"fmt"
	routing "github.com/gly-hub/fasthttp-routing"
	"github.com/stretchr/testify/assert"
	"github.com/swaggo/swag"
	"github.com/valyala/fasthttp"
	"net"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strconv"
	"testing"
	"time"
)

type mockedSwag struct{}

func (s *mockedSwag) ReadDoc() string {
	return `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample server Petstore server.",
        "title": "Swagger Example API",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "petstore.swagger.io",
    "basePath": "/v2",
    "paths": {
        "/file/upload": {
            "post": {
                "description": "Upload file",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Upload file",
                "operationId": "file.upload",
                "parameters": [
                    {
                        "type": "file",
                        "description": "this is a test file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "We need ID!!",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/web.APIError"
                        }
                    },
                    "404": {
                        "description": "Can not find ID",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/web.APIError"
                        }
                    }
                }
            }
        },
        "/testapi/get-string-by-int/{some_id}": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add a new pet to the store",
                "operationId": "get-string-by-int",
                "parameters": [
                    {
                        "type": "int",
                        "description": "Some ID",
                        "name": "some_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Some ID",
                        "name": "some_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/web.Pet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "We need ID!!",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/web.APIError"
                        }
                    },
                    "404": {
                        "description": "Can not find ID",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/web.APIError"
                        }
                    }
                }
            }
        },
        "/testapi/get-struct-array-by-string/{some_id}": {
            "get": {
                "description": "get struct array by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "operationId": "get-struct-array-by-string",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Some ID",
                        "name": "some_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "int",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "int",
                        "description": "Offset",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "We need ID!!",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/web.APIError"
                        }
                    },
                    "404": {
                        "description": "Can not find ID",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/web.APIError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "web.APIError": {
            "type": "object",
            "properties": {
                "CreatedAt": {
                    "type": "string",
                    "format": "date-time"
                },
                "ErrorCode": {
                    "type": "integer"
                },
                "ErrorMessage": {
                    "type": "string"
                }
            }
        },
        "web.Pet": {
            "type": "object",
            "properties": {
                "Category": {
                    "type": "object"
                },
                "ID": {
                    "type": "integer"
                },
                "Name": {
                    "type": "string"
                },
                "PhotoUrls": {
                    "type": "array"
                },
                "Status": {
                    "type": "string"
                },
                "Tags": {
                    "type": "array"
                }
            }
        }
    }
}`
}

func TestWrapHandler(t *testing.T) {
	router := routing.New()

	router.Group("").Get("/*", RoutingWrapHandler(DocExpansion("none"), DomID("swagger-ui")))

	w1 := performRequest(http.MethodGet, "/index.html", router)
	assert.Equal(t, http.StatusOK, w1.StatusCode)
	assert.Equal(t, w1.Header.Get("Content-Type"), "text/html; charset=utf-8")

	w2 := performRequest(http.MethodGet, "/doc.json", router)
	assert.Equal(t, http.StatusInternalServerError, w2.StatusCode)

	swag.Register(swag.Name, &mockedSwag{})
	w2 = performRequest(http.MethodGet, "/doc.json", router)
	assert.Equal(t, http.StatusOK, w2.StatusCode)
	assert.Equal(t, w2.Header.Get("Content-Type"), "application/json; charset=utf-8")

	w3 := performRequest(http.MethodGet, "/favicon-16x16.png", router)
	assert.Equal(t, http.StatusOK, w3.StatusCode)
	assert.Equal(t, w3.Header.Get("Content-Type"), "image/png")

	w4 := performRequest(http.MethodGet, "/swagger-ui.css", router)
	assert.Equal(t, http.StatusOK, w4.StatusCode)
	assert.Equal(t, w4.Header.Get("Content-Type"), "text/css; charset=utf-8")

	w5 := performRequest(http.MethodGet, "/swagger-ui-bundle.js", router)
	assert.Equal(t, http.StatusOK, w5.StatusCode)
	assert.Equal(t, w5.Header.Get("Content-Type"), "application/javascript")

	assert.Equal(t, 302, performRequest(http.MethodGet, "/swagger/", router).StatusCode)
}

func performRequest(method, target string, h *routing.Router) *http.Response {
	r := httptest.NewRequest(method, target, nil)
	w, _ := do(r, h)
	return w
}

func do(req *http.Request, h *routing.Router, msTimeout ...int) (resp *http.Response, err error) {
	timeout := 1000
	if len(msTimeout) > 0 {
		timeout = msTimeout[0]
	}

	// Add Content-Length if not provided with body
	if req.Body != http.NoBody && req.Header.Get("Content-Length") == "" {
		req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))
	}

	// Dump raw http request
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		return nil, err
	}

	// Create test connection
	conn := new(readWriter)

	// Write raw http request
	if _, err = conn.r.Write(dump); err != nil {
		return nil, err
	}

	// Serve conn to server
	channel := make(chan error)
	go func() {
		channel <- fasthttp.ServeConn(conn, h.HandleRequest)
	}()

	// Wait for callback
	if timeout >= 0 {
		// With timeout
		select {
		case err = <-channel:
		case <-time.After(time.Duration(timeout) * time.Millisecond):
			return nil, fmt.Errorf("test: timeout error %vms", timeout)
		}
	} else {
		// Without timeout
		err = <-channel
	}

	// Check for errors
	if err != nil && err != fasthttp.ErrGetOnly {
		return nil, err
	}

	// Read response
	buffer := bufio.NewReader(&conn.w)

	// Convert raw http response to *http.Response
	return http.ReadResponse(buffer, req)
}

func TestURL(t *testing.T) {
	expected := "https://github.com/swaggo/fasthttp-swagger"
	cfg := Config{}
	configFunc := URL(expected)
	configFunc(&cfg)
	assert.Equal(t, expected, cfg.URL)
}

func TestDeepLinking(t *testing.T) {
	expected := true
	cfg := Config{}
	configFunc := DeepLinking(expected)
	configFunc(&cfg)
	assert.Equal(t, expected, cfg.DeepLinking)
}

func TestDocExpansion(t *testing.T) {
	expected := "https://github.com/swaggo/docs"
	cfg := Config{}
	configFunc := DocExpansion(expected)
	configFunc(&cfg)
	assert.Equal(t, expected, cfg.DocExpansion)
}

func TestDomID(t *testing.T) {
	expected := "#swagger-ui"
	cfg := Config{}
	configFunc := DomID(expected)
	configFunc(&cfg)
	assert.Equal(t, expected, cfg.DomID)
}

func TestPersistAuthorization(t *testing.T) {
	expected := true
	cfg := Config{}
	configFunc := PersistAuthorization(expected)
	configFunc(&cfg)
	assert.Equal(t, expected, cfg.PersistAuthorization)
}

func TestInstanceName(t *testing.T) {
	var cfg Config

	assert.Equal(t, "", cfg.InstanceName)

	expected := swag.Name
	configFunc := InstanceName(expected)
	configFunc(&cfg)
	assert.Equal(t, expected, cfg.InstanceName)

	expected = "custom_name"
	configFunc = InstanceName(expected)
	configFunc(&cfg)
	assert.Equal(t, expected, cfg.InstanceName)
}

type readWriter struct {
	net.Conn
	r bytes.Buffer
	w bytes.Buffer
}

func (rw *readWriter) Close() error {
	return nil
}

func (rw *readWriter) Read(b []byte) (int, error) {
	return rw.r.Read(b)
}

func (rw *readWriter) Write(b []byte) (int, error) {
	return rw.w.Write(b)
}

func (rw *readWriter) SetDeadline(t time.Time) error {
	return nil
}

func (rw *readWriter) SetReadDeadline(t time.Time) error {
	return nil
}

func (rw *readWriter) SetWriteDeadline(t time.Time) error {
	return nil
}
