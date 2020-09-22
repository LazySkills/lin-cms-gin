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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cms/admin/all": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员"
                ],
                "summary": "查询所有权限组",
                "parameters": [
                    {
                        "type": "string",
                        "description": "授权token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "[{\"id\":2,\"name\":\"guest\",\"info\":\"游客组\"}]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cms/admin/group": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员"
                ],
                "summary": "查询所有权限组及其权限",
                "parameters": [
                    {
                        "type": "string",
                        "description": "授权token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{\"access_token\":\"...\",\"refresh_token\":\"...\"},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员"
                ],
                "summary": "新建权限组",
                "parameters": [
                    {
                        "type": "string",
                        "description": "授权token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "组名",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "备注",
                        "name": "info",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "array",
                        "description": "权限ID集合",
                        "name": "permission_ids",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "[{\"id\":2,\"name\":\"guest\",\"info\":\"游客组\"}]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cms/admin/group/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员"
                ],
                "summary": "查询一个权限组及其权限",
                "parameters": [
                    {
                        "type": "string",
                        "description": "授权token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "[{\"id\":2,\"name\":\"guest\",\"info\":\"游客组\"}]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cms/admin/permission": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员"
                ],
                "summary": "查询所有可分配的权限",
                "parameters": [
                    {
                        "type": "string",
                        "description": "授权token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{....}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cms/admin/user/:id": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员"
                ],
                "summary": "管理员更新用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "授权token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "ID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "昵称",
                        "name": "nickname",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "头像",
                        "name": "avatar",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{\"access_token\":\"...\",\"refresh_token\":\"...\"},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "授权token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cms/admin/user/:id/password": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "修改用户密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "授权token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "新密码",
                        "name": "new_password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "确认密码",
                        "name": "confirm_password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cms/admin/users": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员"
                ],
                "summary": "查询所有用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "授权token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{....}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cms/user": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户更新信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "ID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "昵称",
                        "name": "nickname",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "头像",
                        "name": "avatar",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{\"access_token\":\"...\",\"refresh_token\":\"...\"},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cms/user/information": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "查询用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "授权token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{....}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cms/user/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"access_token\":\"...\",\"refresh_token\":\"...\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cms/user/password": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "修改密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "旧密码",
                        "name": "old_password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "新密码",
                        "name": "new_password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cms/user/permissions": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "查询自己拥有的权限",
                "parameters": [
                    {
                        "type": "string",
                        "description": "授权token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{....}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cms/user/refresh": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "刷新令牌",
                "parameters": [
                    {
                        "type": "string",
                        "description": "刷新token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"access_token\":\"...\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cms/user/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{\"access_token\":\"...\",\"refresh_token\":\"...\"},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/upload": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件管理"
                ],
                "summary": "上传文件",
                "parameters": [
                    {
                        "type": "file",
                        "description": "文件资源,可以写多个",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "文件资源1,可以写多个",
                        "name": "file1",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/book": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图书"
                ],
                "summary": "获取图书列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/book/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图书"
                ],
                "summary": "获取图书信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/books": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图书"
                ],
                "summary": "新增图书",
                "parameters": [
                    {
                        "type": "string",
                        "description": "图书名称",
                        "name": "title",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "图书作者",
                        "name": "author",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "图书简介",
                        "name": "summary",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "图书图片",
                        "name": "image",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/books/{id}": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图书"
                ],
                "summary": "修改图书",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "图书名称",
                        "name": "title",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "图书作者",
                        "name": "author",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "图书简介",
                        "name": "summary",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "图书图片",
                        "name": "image",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图书"
                ],
                "summary": "删除图书",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
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