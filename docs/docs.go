// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/Even": {
            "get": {
                "description": "Проверка чётности числа",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "numbers"
                ],
                "summary": "Четность",
                "operationId": "Even",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Число",
                        "name": "num",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/service.EvenRecponce"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/service.errorResponce"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/service.errorResponce"
                        }
                    }
                }
            }
        },
        "/Ping": {
            "get": {
                "description": "Проверка доступности сервера",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "Ping",
                "operationId": "Ping",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/service.EvenRecponce"
                        }
                    }
                }
            }
        },
        "/Sum": {
            "post": {
                "description": "Сумма массива чисел переданого в структуре SumData",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "numbers"
                ],
                "summary": "Сумма массива чисел",
                "operationId": "SumSlice",
                "parameters": [
                    {
                        "description": "Данные",
                        "name": "query",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.SumData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/service.SumResponce"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/service.errorResponce"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/service.errorResponce"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "service.EvenRecponce": {
            "type": "object",
            "properties": {
                "even": {
                    "type": "boolean"
                }
            }
        },
        "service.SumData": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                }
            }
        },
        "service.SumResponce": {
            "type": "object",
            "properties": {
                "sum": {
                    "type": "number"
                }
            }
        },
        "service.errorResponce": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}