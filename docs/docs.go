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
        "/events": {
            "get": {
                "description": "get all events",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Event"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "parameters": [
                    {
                        "description": "Event",
                        "name": "event",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateEvent"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    }
                }
            }
        },
        "/events/{event_id}": {
            "get": {
                "description": "get event by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "event_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "event_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/faculties": {
            "get": {
                "description": "get faculty by name, country, city, domain, budget",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "faculties"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Faculty Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Faculty Country",
                        "name": "country",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Faculty City",
                        "name": "city",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Faculty Domain",
                        "name": "domain",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Faculty Budget",
                        "name": "budget",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Faculty"
                            }
                        }
                    }
                }
            }
        },
        "/faculties/{faculty_id}": {
            "get": {
                "description": "get faculty by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "faculties"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Faculty ID",
                        "name": "faculty_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Faculty"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update faculty",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "faculties"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Faculty ID",
                        "name": "faculty_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Faculty",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateFaculty"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Faculty"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete faculty",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "faculties"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Faculty ID",
                        "name": "faculty_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Faculty deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/faculties/{university_id}": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "create faculty",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "faculties"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "University ID",
                        "name": "university_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Faculty",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateFaculty"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Faculty"
                        }
                    }
                }
            }
        },
        "/universities": {
            "get": {
                "description": "get university by name, country, city",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "universities"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "University Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "University Country",
                        "name": "country",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "University City",
                        "name": "city",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.University"
                            }
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
                "description": "create university",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "universities"
                ],
                "parameters": [
                    {
                        "description": "University",
                        "name": "university",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUniversity"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.University"
                        }
                    }
                }
            }
        },
        "/universities/{university_id}": {
            "get": {
                "description": "get university by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "universities"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "University ID",
                        "name": "university_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.University"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update university",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "universities"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "University ID",
                        "name": "university_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "University",
                        "name": "university",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUniversity"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.University"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete university",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "universities"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "University ID",
                        "name": "university_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "University deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "get all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    }
                }
            }
        },
        "/users/current": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get current user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Token"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "register user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RegisterUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Token"
                        }
                    }
                }
            }
        },
        "/users/{user_id}": {
            "get": {
                "description": "get user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreateEvent": {
            "type": "object",
            "required": [
                "name",
                "url",
                "visitor_id"
            ],
            "properties": {
                "campaign_id": {
                    "type": "string"
                },
                "meta": {
                    "$ref": "#/definitions/models.JSONB"
                },
                "name": {
                    "type": "string"
                },
                "payload": {
                    "$ref": "#/definitions/models.JSONB"
                },
                "url": {
                    "type": "string"
                },
                "visitor_id": {
                    "type": "string"
                }
            }
        },
        "models.CreateFaculty": {
            "type": "object",
            "required": [
                "budget",
                "domains",
                "name"
            ],
            "properties": {
                "about": {
                    "type": "string"
                },
                "academic_requirements": {
                    "type": "string"
                },
                "apply_date": {
                    "type": "string"
                },
                "budget": {
                    "type": "integer"
                },
                "domains": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "duration": {
                    "type": "number"
                },
                "language_requirements": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "other_requirements": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "models.CreateUniversity": {
            "type": "object",
            "required": [
                "country",
                "img_link",
                "name"
            ],
            "properties": {
                "about": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "img_link": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "ranking": {
                    "type": "string"
                }
            }
        },
        "models.Event": {
            "type": "object",
            "properties": {
                "campaign_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "meta": {
                    "$ref": "#/definitions/models.JSONB"
                },
                "name": {
                    "type": "string"
                },
                "payload": {
                    "$ref": "#/definitions/models.JSONB"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "visitor_id": {
                    "type": "string"
                }
            }
        },
        "models.Faculty": {
            "type": "object",
            "properties": {
                "about": {
                    "type": "string"
                },
                "academic_requirements": {
                    "type": "string"
                },
                "apply_date": {
                    "type": "string"
                },
                "budget": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "domains": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "duration": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                },
                "language_requirements": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "other_requirements": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                },
                "university": {
                    "$ref": "#/definitions/models.University"
                },
                "university_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.JSONB": {
            "type": "object",
            "additionalProperties": true
        },
        "models.LoginUser": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 8
                }
            }
        },
        "models.RegisterUser": {
            "type": "object",
            "required": [
                "email",
                "name"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3
                },
                "password": {
                    "type": "string"
                },
                "visitor_id": {
                    "type": "string"
                }
            }
        },
        "models.Token": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "models.University": {
            "type": "object",
            "properties": {
                "about": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "faculties": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Faculty"
                    }
                },
                "id": {
                    "type": "string"
                },
                "img_link": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "ranking": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "visitor_id": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Type \"bearer\" followed by a space and JWT token",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{"http", "https"},
	Title:            "Easy-Uni API",
	Description:      "This is the API for the Easy-Uni application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
