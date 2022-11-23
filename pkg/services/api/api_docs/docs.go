// Package api_docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package api_docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ping": {
            "get": {
                "description": "Ping the server",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "Ping the server",
                "operationId": "ping",
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/url": {
            "post": {
                "description": "Create a new url redirect",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "url"
                ],
                "summary": "Create a new url redirect",
                "operationId": "create-url",
                "parameters": [
                    {
                        "description": "url",
                        "name": "url",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CreateUrlRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.UrlDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/url/{urlId}": {
            "get": {
                "description": "Get url",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "url"
                ],
                "summary": "Get url",
                "operationId": "get-url",
                "parameters": [
                    {
                        "type": "string",
                        "description": "url id",
                        "name": "urlId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.UrlDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/{urlId}": {
            "get": {
                "description": "Redirect to url",
                "tags": [
                    "url"
                ],
                "summary": "Redirect to url",
                "operationId": "redirect-url",
                "parameters": [
                    {
                        "type": "string",
                        "description": "url id",
                        "name": "urlId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Found"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.CreateUrlRequest": {
            "type": "object",
            "required": [
                "redirectUrl"
            ],
            "properties": {
                "redirectUrl": {
                    "type": "string"
                }
            }
        },
        "api.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "api.UrlDTO": {
            "type": "object",
            "properties": {
                "newUrl": {
                    "type": "string"
                },
                "redirectUrl": {
                    "type": "string"
                },
                "urlId": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT_AUTH": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "URL Shortener API",
	Description:      "This is a URL shortener service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
