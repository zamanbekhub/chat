// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/api/v1/chat": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "Получить всех пользователей",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Chat ID",
                        "name": "chat_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Response-model_Chat"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Response-schema_Empty"
                        }
                    }
                }
            }
        },
        "/api/v1/chat/all": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "Получить всех пользователей",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Response-model_Chat"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Response-schema_Empty"
                        }
                    }
                }
            }
        },
        "/api/v1/message/push": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "message"
                ],
                "summary": "Новое сообщение",
                "parameters": [
                    {
                        "description": "Новое сообщение",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.MessagePush"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Response-schema_Empty"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Response-schema_Empty"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Chat": {
            "type": "object",
            "properties": {
                "chat_id": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type_code": {
                    "$ref": "#/definitions/model.ChatTypeCode"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.ChatTypeCode": {
            "type": "string",
            "enum": [
                "personal",
                "group",
                "channel"
            ],
            "x-enum-varnames": [
                "CHAT_TYPE_CODE_PERSONAL",
                "CHAT_TYPE_CODE_GROUP",
                "CHAT_TYPE_CODE_CHANNEL"
            ]
        },
        "schema.Empty": {
            "type": "object"
        },
        "schema.MessagePush": {
            "type": "object",
            "properties": {
                "channel": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "schema.Response-model_Chat": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "result": {
                    "$ref": "#/definitions/model.Chat"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "schema.Response-schema_Empty": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "result": {
                    "$ref": "#/definitions/schema.Empty"
                },
                "status": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "127.0.0.1:8000",
	BasePath:         "/chat",
	Schemes:          []string{"http"},
	Title:            "chat",
	Description:      "chat",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	//LeftDelim:        "{{",
	//RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
