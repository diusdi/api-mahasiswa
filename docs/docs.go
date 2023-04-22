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
        "license": {
            "name": "Licensi MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/mhs": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "menampilkan semua data mahasiswa",
                "operationId": "read-mahasiswa",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "menambahkan data mahasiswa baru",
                "operationId": "create-mahasiswa",
                "parameters": [
                    {
                        "description": "Menambahkan data mahasiswa",
                        "name": "mahasiswa",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Mahasiswa"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        },
        "/mhs/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "menampilkan data mahasiswa berdasarkan id",
                "operationId": "read-mahasiswa-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id mahasiswa",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "update data mahasiswa",
                "operationId": "update-mahasiswa",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id mahasiswa",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Data yang bisa diupdate : nama, umur, gender ('0' untuk perempuan dan '1' untuk laki-laki)",
                        "name": "mahasiswa",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Mahasiswa"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "menghapus data mahasiswa",
                "operationId": "delete-mahasiswa",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id mahasiswa",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Mahasiswa": {
            "type": "object",
            "properties": {
                "gender": {
                    "type": "string",
                    "example": "1"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "is_active": {
                    "type": "string",
                    "example": "1"
                },
                "nama": {
                    "type": "string",
                    "example": "Dion"
                },
                "tanggal_registrasi": {
                    "type": "string",
                    "example": "2020-01-02T15:04:05Z"
                },
                "usia": {
                    "type": "integer",
                    "example": 21
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API Mahasiswa",
	Description:      "API untuk mengatur data mahasiswa Jobhun. Untuk source code dapat dilihat di https://github.com/diusdi/api-mahasiswa",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
