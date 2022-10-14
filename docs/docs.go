// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Sergey Nekrasov",
            "url": "https://vk.com/serega_nekrasov",
            "email": "79508031750@yandex.ru"
        },
        "license": {
            "name": "AS IS (NO WARRANTY)"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/promos/create": {
            "post": {
                "description": "Adding a new promo to database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Add"
                ],
                "summary": "Add a new promo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Магазин",
                        "name": "Store",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Скидка",
                        "name": "Discount",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Цена",
                        "name": "Price",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Количество",
                        "name": "Quantity",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Промокоды(запись в виде массива)",
                        "name": "Promo",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/promos.PromoCreated"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/promos.PromoError"
                        }
                    }
                }
            }
        },
        "/promos/get": {
            "get": {
                "description": "Get a list of all promos",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Info"
                ],
                "summary": "Get all records",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ds.Promos"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/promos.PromoError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ds.Promos": {
            "type": "object",
            "properties": {
                "discount": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "promo": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "quantity": {
                    "type": "integer"
                },
                "store": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "promos.PromoCreated": {
            "type": "object",
            "properties": {
                "success": {
                    "description": "success",
                    "type": "boolean"
                }
            }
        },
        "promos.PromoError": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "description",
                    "type": "string"
                },
                "error": {
                    "description": "error",
                    "type": "string"
                },
                "type": {
                    "description": "type",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "/",
	Schemes:          []string{"https", "http"},
	Title:            "Freebie-shop",
	Description:      "Store with promo codes for various stores",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
