// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package swagger

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
        "/authCode": {
            "post": {
                "description": "发送手机验证码",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "手机验证码"
                ],
                "summary": "发送手机验证码",
                "parameters": [
                    {
                        "type": "string",
                        "default": "13881887710",
                        "description": "手机号码",
                        "name": "mobile",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "1",
                        "description": "发送类型:1登陆,2注册,3忘记密码",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.JsonResult"
                        }
                    }
                }
            }
        },
        "/forgetPwd": {
            "post": {
                "description": "忘记密码",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户登陆"
                ],
                "summary": "忘记密码",
                "parameters": [
                    {
                        "type": "string",
                        "default": "13881887710",
                        "description": "手机号码",
                        "name": "mobile",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "123456",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "432343",
                        "description": "短信验证码",
                        "name": "sms_code",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "1614479454",
                        "description": "时间戳",
                        "name": "timestamp",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "5fae1a0e4156e1ad4e742f4b9b7c1416",
                        "description": "签名",
                        "name": "sign",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.JsonResult"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "用户登陆",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户登陆"
                ],
                "summary": "用户登陆",
                "parameters": [
                    {
                        "type": "string",
                        "default": "13881887710",
                        "description": "手机号码",
                        "name": "mobile",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "123456",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "432343",
                        "description": "短信验证码",
                        "name": "sms_code",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "1614479454",
                        "description": "时间戳",
                        "name": "timestamp",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "5fae1a0e4156e1ad4e742f4b9b7c1416",
                        "description": "签名",
                        "name": "sign",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.JsonResult"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户登陆"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "type": "string",
                        "default": "13881887710",
                        "description": "手机号码",
                        "name": "mobile",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "123456",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "432343",
                        "description": "短信验证码",
                        "name": "sms_code",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "张三",
                        "description": "用户名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "1",
                        "description": "性别:1男,2女",
                        "name": "sex",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "1614479454",
                        "description": "时间戳",
                        "name": "timestamp",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "5fae1a0e4156e1ad4e742f4b9b7c1416",
                        "description": "签名",
                        "name": "sign",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.JsonResult"
                        }
                    }
                }
            }
        },
        "/user/addAddr": {
            "post": {
                "description": "添加收货地址",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "添加收货地址",
                "parameters": [
                    {
                        "type": "string",
                        "default": "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm",
                        "description": "用户token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "张三",
                        "description": "姓名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "13881887710",
                        "description": "联系电话",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "四川/成都/温江区",
                        "description": "地址",
                        "name": "addr",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "南薰大道塞纳湖畔7栋503",
                        "description": "详细地址",
                        "name": "detail",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "是否默认，1默认",
                        "name": "is_default",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.JsonResult"
                        }
                    }
                }
            }
        },
        "/user/addrList": {
            "post": {
                "description": "获取收货地址列表",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "获取收货地址列表",
                "parameters": [
                    {
                        "type": "string",
                        "default": "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm",
                        "description": "用户token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.JsonResult"
                        }
                    }
                }
            }
        },
        "/user/bandingAlipay": {
            "post": {
                "description": "绑定alipay",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "绑定alipay",
                "parameters": [
                    {
                        "type": "string",
                        "default": "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm",
                        "description": "用户token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "张三",
                        "description": "用户名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "320553500@qq.com",
                        "description": "账号",
                        "name": "account",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.JsonResult"
                        }
                    }
                }
            }
        },
        "/user/deleteAddr": {
            "post": {
                "description": "删除收货地址",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "删除收货地址",
                "parameters": [
                    {
                        "type": "string",
                        "default": "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm",
                        "description": "用户token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "1",
                        "description": "用户收货地址ID",
                        "name": "shop_addr_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.JsonResult"
                        }
                    }
                }
            }
        },
        "/user/editAddr": {
            "post": {
                "description": "修改收货地址",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "修改收货地址",
                "parameters": [
                    {
                        "type": "string",
                        "default": "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm",
                        "description": "用户token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "1",
                        "description": "用户收货地址ID",
                        "name": "shop_addr_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "张三",
                        "description": "姓名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "13881887710",
                        "description": "联系电话",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "四川/成都/温江区",
                        "description": "地址",
                        "name": "addr",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "南薰大道塞纳湖畔7栋503",
                        "description": "详细地址",
                        "name": "detail",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "是否默认，1默认",
                        "name": "is_default",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.JsonResult"
                        }
                    }
                }
            }
        },
        "/user/editPwd": {
            "post": {
                "description": "修改用户密码",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "修改用户密码",
                "parameters": [
                    {
                        "type": "string",
                        "default": "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm",
                        "description": "用户token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "123456",
                        "description": "就密码",
                        "name": "old_pwd",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "3221",
                        "description": "新密码",
                        "name": "new_pwd",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.JsonResult"
                        }
                    }
                }
            }
        },
        "/user/getAlipay": {
            "post": {
                "description": "获取alipay",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "获取alipay",
                "parameters": [
                    {
                        "type": "string",
                        "default": "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm",
                        "description": "用户token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.JsonResult"
                        }
                    }
                }
            }
        },
        "/user/headImg": {
            "post": {
                "description": "更新用户头像",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "更新用户头像",
                "parameters": [
                    {
                        "type": "string",
                        "default": "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm",
                        "description": "用户token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "用户头像",
                        "name": "img",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.JsonResult"
                        }
                    }
                }
            }
        },
        "/user/integral": {
            "post": {
                "description": "用户积分",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户积分",
                "parameters": [
                    {
                        "type": "string",
                        "default": "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm",
                        "description": "用户token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.JsonResult"
                        }
                    }
                }
            }
        },
        "/user/updateUser": {
            "post": {
                "description": "更新用户信息",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "更新用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "default": "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm",
                        "description": "用户token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "张三",
                        "description": "昵称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "1男/2女",
                        "description": "性别",
                        "name": "sex",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "default": "320553500@qq.com",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "default": "四川省/成都市/成华区",
                        "description": "用户地址",
                        "name": "user_addr",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.JsonResult"
                        }
                    }
                }
            }
        },
        "/user/userInfo": {
            "post": {
                "description": "获取用户信息",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "default": "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm",
                        "description": "用户token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.JsonResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.JsonResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
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
