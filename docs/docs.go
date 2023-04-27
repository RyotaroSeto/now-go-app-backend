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
        "/api/v1/like": {
            "get": {
                "description": "自身にいいねをしたユーザー一覧を表示した時呼ばれる API",
                "summary": "いいね一覧参照 API",
                "parameters": [
                    {
                        "description": "UserID",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userinterface.GetLikeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/userinterface.GetLikedResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "ユーザーがいいねした時呼ばれる API",
                "summary": "いいね API",
                "parameters": [
                    {
                        "description": "UserID, LikedUserID, MessageBody",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userinterface.LikeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "description": "指定ユーザーのプロフィール確認時呼ばれる API",
                "summary": "ユーザープロフィール情報参照 API",
                "parameters": [
                    {
                        "description": "ID",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userinterface.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userinterface.UserResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "自身のプロフィール情報更新時呼ばれる API",
                "summary": "ユーザープロフィール情報更新 API",
                "parameters": [
                    {
                        "description": "ID, DateOfBirth, Gender, Residence, Occupation, Height, Weight",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userinterface.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userinterface.UserDetailResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "userinterface.GetLikeRequest": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "userinterface.GetLikedResponse": {
            "type": "object",
            "properties": {
                "liked_date": {
                    "type": "string"
                },
                "liked_user_id": {
                    "type": "integer"
                },
                "message_body": {
                    "type": "string"
                }
            }
        },
        "userinterface.LikeRequest": {
            "type": "object",
            "properties": {
                "liked_user_id": {
                    "type": "integer"
                },
                "message_body": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "userinterface.UserDetailResponse": {
            "type": "object",
            "properties": {
                "date_of_birth": {
                    "type": "integer"
                },
                "gender": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "occupation": {
                    "type": "string"
                },
                "residence": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "userinterface.UserRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "userinterface.UserResponse": {
            "type": "object",
            "properties": {
                "date_of_birth": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "occupation": {
                    "type": "string"
                },
                "residence": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "user_name": {
                    "type": "string"
                },
                "weight": {
                    "type": "integer"
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