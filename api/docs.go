// Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://pangud.org",
        "contact": {
            "name": "服务支持",
            "url": "https://pangud.org",
            "email": "dev_support@gail.com"
        },
        "license": {
            "name": "AGPL-3.0",
            "url": "https://www.gnu.org/licenses/agpl-3.0.en.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ssl_certs/dns_providers": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "查询DNS提供商列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SSL证书"
                ],
                "summary": "查询DNS提供商列表",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "分页查询用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "分页查询用户",
                "parameters": [
                    {
                        "type": "string",
                        "name": "keywords",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "example": 10,
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "example": 0,
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ResponseEntity-types_Page-biz_User"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据用户名、密码、昵称、姓名创建新用户, 其中用户名、密码为必填、其他为可选项",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "新增用户",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "CreateUserCommand",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/biz.CreateUserCommand"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "更新用户信息 姓名、昵称、头像",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "更新用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "use id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "用户信息",
                        "name": "UpdateUserCommand",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/biz.UpdateUserCommand"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "biz.CreateUserCommand": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "nickname": {
                    "type": "string",
                    "example": "Gaga"
                },
                "password": {
                    "type": "string",
                    "example": "your password"
                },
                "realname": {
                    "type": "string",
                    "example": "张三"
                },
                "username": {
                    "type": "string",
                    "example": "pduser"
                }
            }
        },
        "biz.UpdateUserCommand": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string",
                    "example": "https://xxxx.com/xb.jpg"
                },
                "locked": {
                    "type": "boolean",
                    "example": false
                },
                "nickname": {
                    "type": "string",
                    "example": "Gaga"
                },
                "realname": {
                    "type": "string",
                    "example": "张三"
                }
            }
        },
        "biz.User": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "create_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_login_time": {
                    "type": "string"
                },
                "locked": {
                    "type": "boolean"
                },
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "realname": {
                    "type": "string"
                },
                "update_time": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "errors.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/errors.ErrorCode"
                        }
                    ],
                    "example": 0
                },
                "msg": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "errors.ErrorCode": {
            "type": "integer",
            "enum": [
                0,
                99999,
                100000,
                100010,
                100011,
                200000,
                200001,
                300000
            ],
            "x-enum-comments": {
                "ErrCodeInvalidParam": "参数错误"
            },
            "x-enum-varnames": [
                "ErrCodeOK",
                "ErrCodeUnknownError",
                "ErrCodeSystemError",
                "ErrCodeDBError",
                "ErrCodeDBConnectionError",
                "ErrCodeUserError",
                "ErrCodeInvalidParam",
                "ErrCodeExternalError"
            ]
        },
        "types.Page-biz_User": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/biz.User"
                    }
                },
                "offset": {
                    "type": "integer"
                },
                "total": {
                    "description": "Total 总数",
                    "type": "integer"
                }
            }
        },
        "types.ResponseEntity-types_Page-biz_User": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/types.Page-biz_User"
                },
                "msg": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:2345",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "PangudOS API",
	Description:      "PANGUD OS API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
