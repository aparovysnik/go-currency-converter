// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/currency/convert": {
            "get": {
                "description": "convert an amount of money from one currency to another.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Converts the given amount of money.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Amount to convert",
                        "name": "amount",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Code of currency to convert from",
                        "name": "fromCurrency",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Code of currency to convert to",
                        "name": "amount",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ConvertCurrencyResponse"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "get the status indicating if the server is operational.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Gets the status of the server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.StatusResponse"
                        }
                    }
                }
            }
        },
        "/project": {
            "post": {
                "description": "register a new project.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Registers a new project.",
                "parameters": [
                    {
                        "description": "Contact email details",
                        "name": "contactEmail",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ConvertCurrencyResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ConvertCurrencyResponse": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "models.StatusResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}