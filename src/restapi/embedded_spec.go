// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Filmoteka",
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "basePath": "/Filmoteka/",
  "paths": {
    "/actors": {
      "get": {
        "security": [
          {
            "isUser": []
          },
          {
            "isAdmin": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Actor"
        ],
        "summary": "Получить список всех актеров",
        "operationId": "getAllActors",
        "responses": {
          "200": {
            "description": "Список актеров",
            "schema": {
              "type": "array",
              "items": {
                "type": "object",
                "properties": {
                  "date_of_birthday": {
                    "type": "string",
                    "format": "date"
                  },
                  "films": {
                    "type": "array",
                    "items": {
                      "type": "string"
                    }
                  },
                  "gender": {
                    "enum": [
                      "M",
                      "F"
                    ]
                  },
                  "name": {
                    "type": "string"
                  }
                }
              }
            }
          },
          "400": {
            "description": "Ошибка",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Actor"
        ],
        "summary": "Изменить информацию об актере",
        "operationId": "updateActor",
        "parameters": [
          {
            "type": "string",
            "name": "name",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "format": "date",
            "name": "b_date",
            "in": "query",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "date_of_birthday": {
                  "type": "string",
                  "format": "date"
                },
                "name": {
                  "type": "string"
                },
                "sex": {
                  "type": "string",
                  "enum": [
                    "M",
                    "F"
                  ]
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Обновлено удачно"
          },
          "400": {
            "description": "Ошибка",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Actor"
        ],
        "summary": "Добавить актера",
        "operationId": "addActor",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Actor"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Успешно"
          },
          "400": {
            "description": "Ошибка",
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
          "Actor"
        ],
        "summary": "Удалить запись об актере",
        "operationId": "deleteActor",
        "parameters": [
          {
            "type": "string",
            "name": "name",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "format": "date",
            "name": "b_date",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Успешное удаление"
          },
          "400": {
            "description": "Некорректные входные данные",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/films": {
      "get": {
        "security": [
          {
            "isUser": []
          },
          {
            "isAdmin": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Film"
        ],
        "summary": "Получить список всех фильмов с возможностью сортировки",
        "operationId": "getAllFilmsSorted",
        "parameters": [
          {
            "enum": [
              "title",
              "rating",
              "datepremiere"
            ],
            "type": "string",
            "default": "rating",
            "name": "sortBy",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Список фильмов",
            "schema": {
              "type": "array",
              "items": {
                "type": "object",
                "required": [
                  "name",
                  "year",
                  "description",
                  "rating"
                ],
                "properties": {
                  "description": {
                    "type": "string",
                    "maxLength": 1000
                  },
                  "name": {
                    "type": "string"
                  },
                  "rating": {
                    "type": "number",
                    "maximum": 10
                  },
                  "year": {
                    "type": "integer"
                  }
                }
              }
            }
          },
          "400": {
            "description": "Ошибка",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Film"
        ],
        "summary": "Изменить информацию о фильме",
        "operationId": "updateFilm",
        "parameters": [
          {
            "type": "string",
            "name": "name",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "name": "year",
            "in": "query",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "description": {
                  "type": "string",
                  "maxLength": 1000
                },
                "name": {
                  "type": "string",
                  "maxLength": 150
                },
                "rate": {
                  "type": "number",
                  "maximum": 10
                },
                "year": {
                  "type": "integer"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Обновлено удачно"
          },
          "400": {
            "description": "Ошибка",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Film"
        ],
        "summary": "Добавить фильм",
        "operationId": "addFilm",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Film"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Успешно"
          },
          "400": {
            "description": "Ошибка",
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
          "Film"
        ],
        "summary": "Удалить запись о фильме",
        "operationId": "deleteFilm",
        "parameters": [
          {
            "type": "string",
            "name": "name",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "name": "year",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Успешное удаление"
          },
          "400": {
            "description": "Некорректные входные данные",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/films/find": {
      "get": {
        "security": [
          {
            "isUser": []
          },
          {
            "isAdmin": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Film"
        ],
        "summary": "Поиск фильма по фрагменту описания или имени актера",
        "operationId": "findFilm",
        "parameters": [
          {
            "maxLength": 1000,
            "type": "string",
            "name": "desc",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "actor",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Найдено",
            "schema": {
              "type": "array",
              "items": {
                "type": "object",
                "required": [
                  "name",
                  "year",
                  "description"
                ],
                "properties": {
                  "description": {
                    "type": "string",
                    "maxLength": 1000
                  },
                  "name": {
                    "type": "string"
                  },
                  "year": {
                    "type": "integer"
                  }
                }
              }
            }
          },
          "400": {
            "description": "Не найдено",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/user/login": {
      "get": {
        "security": [],
        "tags": [
          "User"
        ],
        "summary": "Вход пользователя в систему",
        "operationId": "userLogin",
        "parameters": [
          {
            "minLength": 5,
            "type": "string",
            "name": "username",
            "in": "query",
            "required": true
          },
          {
            "minLength": 5,
            "type": "string",
            "name": "password",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успешнный вход в систему",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Неверное имя пользователя / пароль"
          }
        }
      }
    },
    "/user/logout": {
      "get": {
        "security": [
          {
            "isUser": []
          },
          {
            "isAdmin": []
          }
        ],
        "tags": [
          "User"
        ],
        "summary": "Выход пользователя из системы",
        "operationId": "userLogout",
        "parameters": [
          {
            "type": "string",
            "name": "token",
            "in": "header",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успешно"
          }
        }
      }
    }
  },
  "definitions": {
    "Actor": {
      "type": "object",
      "required": [
        "name",
        "sex",
        "date_of_birthday"
      ],
      "properties": {
        "date_of_birthday": {
          "type": "string",
          "format": "date"
        },
        "name": {
          "type": "string"
        },
        "sex": {
          "type": "string",
          "enum": [
            "M",
            "F"
          ]
        }
      }
    },
    "ActorFilm": {
      "type": "object",
      "required": [
        "name",
        "sex",
        "date_of_birthday",
        "films"
      ],
      "properties": {
        "date_of_birthday": {
          "type": "string",
          "format": "date"
        },
        "films": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "name": {
          "type": "string"
        },
        "sex": {
          "type": "string",
          "enum": [
            "M",
            "F"
          ]
        }
      }
    },
    "Actors": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Actor"
      }
    },
    "Film": {
      "type": "object",
      "required": [
        "name",
        "description",
        "year",
        "rate",
        "actors"
      ],
      "properties": {
        "actors": {
          "$ref": "#/definitions/Actors"
        },
        "description": {
          "type": "string",
          "maxLength": 1000
        },
        "name": {
          "type": "string",
          "maxLength": 150
        },
        "rate": {
          "type": "number",
          "maximum": 10
        },
        "year": {
          "type": "integer"
        }
      }
    },
    "Films": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Film"
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string",
          "minLength": 5
        },
        "userStatus": {
          "type": "string",
          "enum": [
            "admin",
            "user"
          ]
        },
        "username": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "roles": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    }
  },
  "securityDefinitions": {
    "isAdmin": {
      "type": "apiKey",
      "name": "admin-key",
      "in": "header"
    },
    "isUser": {
      "type": "apiKey",
      "name": "user-key",
      "in": "header"
    }
  },
  "security": [
    {
      "isAdmin": []
    }
  ],
  "tags": [
    {
      "description": "Everything about actors",
      "name": "Actor"
    },
    {
      "description": "Everything about films",
      "name": "Film"
    },
    {
      "description": "Login / logout",
      "name": "User"
    }
  ],
  "x-components": {}
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Filmoteka",
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "basePath": "/Filmoteka/",
  "paths": {
    "/actors": {
      "get": {
        "security": [
          {
            "isUser": []
          },
          {
            "isAdmin": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Actor"
        ],
        "summary": "Получить список всех актеров",
        "operationId": "getAllActors",
        "responses": {
          "200": {
            "description": "Список актеров",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/GetAllActorsOKBodyItems0"
              }
            }
          },
          "400": {
            "description": "Ошибка",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Actor"
        ],
        "summary": "Изменить информацию об актере",
        "operationId": "updateActor",
        "parameters": [
          {
            "type": "string",
            "name": "name",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "format": "date",
            "name": "b_date",
            "in": "query",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "date_of_birthday": {
                  "type": "string",
                  "format": "date"
                },
                "name": {
                  "type": "string"
                },
                "sex": {
                  "type": "string",
                  "enum": [
                    "M",
                    "F"
                  ]
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Обновлено удачно"
          },
          "400": {
            "description": "Ошибка",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Actor"
        ],
        "summary": "Добавить актера",
        "operationId": "addActor",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Actor"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Успешно"
          },
          "400": {
            "description": "Ошибка",
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
          "Actor"
        ],
        "summary": "Удалить запись об актере",
        "operationId": "deleteActor",
        "parameters": [
          {
            "type": "string",
            "name": "name",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "format": "date",
            "name": "b_date",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Успешное удаление"
          },
          "400": {
            "description": "Некорректные входные данные",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/films": {
      "get": {
        "security": [
          {
            "isUser": []
          },
          {
            "isAdmin": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Film"
        ],
        "summary": "Получить список всех фильмов с возможностью сортировки",
        "operationId": "getAllFilmsSorted",
        "parameters": [
          {
            "enum": [
              "title",
              "rating",
              "datepremiere"
            ],
            "type": "string",
            "default": "rating",
            "name": "sortBy",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Список фильмов",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/GetAllFilmsSortedOKBodyItems0"
              }
            }
          },
          "400": {
            "description": "Ошибка",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Film"
        ],
        "summary": "Изменить информацию о фильме",
        "operationId": "updateFilm",
        "parameters": [
          {
            "type": "string",
            "name": "name",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "name": "year",
            "in": "query",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "description": {
                  "type": "string",
                  "maxLength": 1000
                },
                "name": {
                  "type": "string",
                  "maxLength": 150
                },
                "rate": {
                  "type": "number",
                  "maximum": 10,
                  "minimum": 0
                },
                "year": {
                  "type": "integer"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Обновлено удачно"
          },
          "400": {
            "description": "Ошибка",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Film"
        ],
        "summary": "Добавить фильм",
        "operationId": "addFilm",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Film"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Успешно"
          },
          "400": {
            "description": "Ошибка",
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
          "Film"
        ],
        "summary": "Удалить запись о фильме",
        "operationId": "deleteFilm",
        "parameters": [
          {
            "type": "string",
            "name": "name",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "name": "year",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Успешное удаление"
          },
          "400": {
            "description": "Некорректные входные данные",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/films/find": {
      "get": {
        "security": [
          {
            "isUser": []
          },
          {
            "isAdmin": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Film"
        ],
        "summary": "Поиск фильма по фрагменту описания или имени актера",
        "operationId": "findFilm",
        "parameters": [
          {
            "maxLength": 1000,
            "type": "string",
            "name": "desc",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "actor",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Найдено",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/FindFilmOKBodyItems0"
              }
            }
          },
          "400": {
            "description": "Не найдено",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/user/login": {
      "get": {
        "security": [],
        "tags": [
          "User"
        ],
        "summary": "Вход пользователя в систему",
        "operationId": "userLogin",
        "parameters": [
          {
            "minLength": 5,
            "type": "string",
            "name": "username",
            "in": "query",
            "required": true
          },
          {
            "minLength": 5,
            "type": "string",
            "name": "password",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успешнный вход в систему",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Неверное имя пользователя / пароль"
          }
        }
      }
    },
    "/user/logout": {
      "get": {
        "security": [
          {
            "isUser": []
          },
          {
            "isAdmin": []
          }
        ],
        "tags": [
          "User"
        ],
        "summary": "Выход пользователя из системы",
        "operationId": "userLogout",
        "parameters": [
          {
            "type": "string",
            "name": "token",
            "in": "header",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успешно"
          }
        }
      }
    }
  },
  "definitions": {
    "Actor": {
      "type": "object",
      "required": [
        "name",
        "sex",
        "date_of_birthday"
      ],
      "properties": {
        "date_of_birthday": {
          "type": "string",
          "format": "date"
        },
        "name": {
          "type": "string"
        },
        "sex": {
          "type": "string",
          "enum": [
            "M",
            "F"
          ]
        }
      }
    },
    "ActorFilm": {
      "type": "object",
      "required": [
        "name",
        "sex",
        "date_of_birthday",
        "films"
      ],
      "properties": {
        "date_of_birthday": {
          "type": "string",
          "format": "date"
        },
        "films": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "name": {
          "type": "string"
        },
        "sex": {
          "type": "string",
          "enum": [
            "M",
            "F"
          ]
        }
      }
    },
    "Actors": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Actor"
      }
    },
    "Film": {
      "type": "object",
      "required": [
        "name",
        "description",
        "year",
        "rate",
        "actors"
      ],
      "properties": {
        "actors": {
          "$ref": "#/definitions/Actors"
        },
        "description": {
          "type": "string",
          "maxLength": 1000
        },
        "name": {
          "type": "string",
          "maxLength": 150
        },
        "rate": {
          "type": "number",
          "maximum": 10,
          "minimum": 0
        },
        "year": {
          "type": "integer"
        }
      }
    },
    "Films": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Film"
      }
    },
    "FindFilmOKBodyItems0": {
      "type": "object",
      "required": [
        "name",
        "year",
        "description"
      ],
      "properties": {
        "description": {
          "type": "string",
          "maxLength": 1000
        },
        "name": {
          "type": "string"
        },
        "year": {
          "type": "integer"
        }
      }
    },
    "GetAllActorsOKBodyItems0": {
      "type": "object",
      "properties": {
        "date_of_birthday": {
          "type": "string",
          "format": "date"
        },
        "films": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "gender": {
          "enum": [
            "M",
            "F"
          ]
        },
        "name": {
          "type": "string"
        }
      }
    },
    "GetAllFilmsSortedOKBodyItems0": {
      "type": "object",
      "required": [
        "name",
        "year",
        "description",
        "rating"
      ],
      "properties": {
        "description": {
          "type": "string",
          "maxLength": 1000
        },
        "name": {
          "type": "string"
        },
        "rating": {
          "type": "number",
          "maximum": 10,
          "minimum": 0
        },
        "year": {
          "type": "integer"
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string",
          "minLength": 5
        },
        "userStatus": {
          "type": "string",
          "enum": [
            "admin",
            "user"
          ]
        },
        "username": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "roles": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    }
  },
  "securityDefinitions": {
    "isAdmin": {
      "type": "apiKey",
      "name": "admin-key",
      "in": "header"
    },
    "isUser": {
      "type": "apiKey",
      "name": "user-key",
      "in": "header"
    }
  },
  "security": [
    {
      "isAdmin": []
    }
  ],
  "tags": [
    {
      "description": "Everything about actors",
      "name": "Actor"
    },
    {
      "description": "Everything about films",
      "name": "Film"
    },
    {
      "description": "Login / logout",
      "name": "User"
    }
  ],
  "x-components": {}
}`))
}
