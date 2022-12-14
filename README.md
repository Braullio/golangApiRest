
# GOLANG API REST

![So](https://img.shields.io/badge/S.O.-Linux-informational.svg)
![Go Version](https://img.shields.io/badge/Go-v.1.18-informational.svg)
<br>
![Database](https://img.shields.io/badge/Database-BigQuery-informational.svg)
![Database](https://img.shields.io/badge/Database-Postgresql-informational.svg)
<br>
![Report](https://img.shields.io/badge/Report-Google_Chat-informational.svg)
<br>
![Repo Size](https://img.shields.io/badge/Repo_size-1,6_MB-informational.svg)
![Exec Size](https://img.shields.io/badge/Exec_size-26_MB-informational.svg)
<br>
[![MIT License](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://choosealicense.com/licenses/mit/)

## Installation GoLang
```bash
  sudo apt install snapd
  sudo snap install go
```

## Run Locally
```bash
  # Install dependencies
  go build
  # Start the server linux
  ./golangApiRest
```

## ENV (.env)
```
  GOOGLE_IAM_KEY = "./security/bigquery_key.json"
  BIGQUERY_PROJECT_ID = "project-123456789"
  GOOGLE_CHAT_WEBHOOK = "https://chat.googleapis.com/v1/spaces/AAAAgOxrNhU/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=zk3322S04V0WdHW9y_4Zsg62CUZP9HvIvy0000uJ1qo%3D"

  POSTGRESQL_HOST = "localhost"
  POSTGRESQL_USER = "postgres"
  POSTGRESQL_PASSWORD = "postgres"
  POSTGRESQL_DBNAME = "golangApiRest"
  POSTGRESQL_PORT = "5432"
```

## Tables
Locate: Biquery <br>
Name: users <br>
AutoMigrate: false <br>

| Column     | Type     | Compilate |
|------------|----------|-----------|
| id         | STRING   | NULLABLE  |
| status     | STRING   | NULLABLE  |
| name       | STRING   | NULLABLE  |
| phone      | STRING   | NULLABLE  |
| created_at | DATETIME | NULLABLE  |
| updated_at | DATETIME | NULLABLE  |

Locate: Postgresql <br>
Name: products <br>
AutoMigrate: true <br>
Soft Delete: deleted_at <br>

| Column      | Type      |
|-------------|-----------|
| id          | BIGINT    |
| name        | TEXT      |
| value       | NUMERIC   |
| created_at  | TIMESTAMP |
| updated_at  | TIMESTAMP |
| deleted_at  | TIMESTAMP |

## API Reference

### # Monitoring point
*Request*
```http request
  GET http://localhost:8001/ping
```
*Response*
```http
  ### Status code: 200
  {
    "msg": "pong"
  }
```
![So](https://img.shields.io/badge/Report-Google_Chat_Webhook-red.svg)
<br><br>

### # User point's
#### -- *Create User*
*Request*
```http
  POST http://localhost:8001/user
  Content-Type: application/json

  {
    "name": "Testenildos",
    "phone" : "5562981919191",
    "status": "Registred"
  }
```
*Response*
```http
  ### Status code: 201
  {
    "id": "b11b1522-cb7c-48b8-9b35-27ec3a343e34",
    "name": "Testenildos",
    "phone": "5562981919191",
    "status": "Registred",
    "created_at": "2022-09-24T23:01:15.748503488",
    "updated_at": "2022-09-24T23:01:15.748503488"
  }
```
![So](https://img.shields.io/badge/Report-Google_Chat_Webhook-red.svg)
<br>

#### -- *List User*
*Request*
```http
  GET http://localhost:8001/user/:id
```
*Response*
```http
  ### Status code: 200
  [
    {
      "id": "c8b948cd-4172-40a1-924a-fe815214f659",
      "name": "Testenildfffffo",
      "phone": "5562981921119191",
      "status": "Registr11ed",
      "created_at": "2022-09-22T18:31:17",
      "updated_at": "2022-09-22T18:31:17"
    }
  ]
```
![So](https://img.shields.io/badge/Report-Google_Chat_Webhook-red.svg)

*Request*
```http
  GET http://localhost:8001/user
```
*Response*
```http
  ### Status code: 200
  [
    {
      "id": "c8b948cd-4172-40a1-924a-fe815214f659",
      "name": "Testenildfffffo",
      "phone": "5562981921119191",
      "status": "Registr11ed",
      "created_at": "2022-09-22T18:31:17",
      "updated_at": "2022-09-22T18:31:17"
    },
    {
      "id": "ac668339-beee-46e6-87c5-8a482e5d9d26",
      "name": "Testenildo",
      "phone": "5562981919191",
      "status": "Registred",
      "created_at": "2022-09-22T18:31:22",
      "updated_at": "2022-09-22T18:31:22"
    },
    {
      "id": "b11b1522-cb7c-48b8-9b35-27ec3a343e34",
      "name": "Testenildos",
      "phone": "5562981919191",
      "status": "Registred",
      "created_at": "2022-09-24T23:01:15",
      "updated_at": "2022-09-24T23:01:15"
    }
  ]
```
![So](https://img.shields.io/badge/Report-Google_Chat_Webhook-red.svg)
<br>

#### -- *Update User*
*Request*
```http 
  PUT http://localhost:8001/user/:id
  Content-Type: application/json
  
  {
    "status": "Processed"
  }
```
*Response*
```http
  ### Status code: 200
```
![So](https://img.shields.io/badge/Report-Google_Chat_Webhook-red.svg)
<br>

#### -- *Delete User*
*Request*
```http request
  DELETE http://localhost:8001/user/:id
```
*Response*

```http
  ### Status code: 200
```
![So](https://img.shields.io/badge/Report-Google_Chat_Webhook-success.svg)

<br>

### # Products point's
#### -- *Create Product*
*Request*
```http
  POST http://localhost:8001/product
  Content-Type: application/json

  {
    "name": "balinha",
    "value": 0.11
  }
```
*Response*
```http
  ### Status code: 201
  {
    "id": 1,
    "name": "balinha",
    "value": 0.11,
    "created_at": "2022-09-25T19:49:19.758414735-03:00",
    "updated_at": "2022-09-25T19:49:19.758414735-03:00"
  }
```
![So](https://img.shields.io/badge/Report-Google_Chat_Webhook-red.svg)
<br>

#### -- *List Product*
*Request*
```http
  GET http://localhost:8001/product/:id
```
*Response*
```http
  ### Status code: 200
  [
    {
      "id": 1,
      "name": "balinha",
      "value": 0.11,
      "created_at": "2022-09-25T19:49:19.758414-03:00",
      "updated_at": "2022-09-25T19:49:19.758414-03:00"
    }
  ]
```
![So](https://img.shields.io/badge/Report-Google_Chat_Webhook-red.svg)

*Request*
```http
  GET http://localhost:8001/product
```
*Response*
```http
  ### Status code: 200
  [
    {
      "id": 1,
      "name": "balinha",
      "value": 0.11,
      "created_at": "2022-09-25T19:49:19.758414-03:00",
      "updated_at": "2022-09-25T19:49:19.758414-03:00"
    },
    {
      "id": 2,
      "name": "Bola",
      "value": 10.5,
      "created_at": "2022-09-25T19:52:01.388815-03:00",
      "updated_at": "2022-09-25T19:52:01.388815-03:00"
    }
  ]
```
![So](https://img.shields.io/badge/Report-Google_Chat_Webhook-red.svg)
<br>

#### -- *Update Product*
*Request*
```http 
  PUT http://localhost:8001/product/:id
  Content-Type: application/json
  
  {
    "value": 11.00
  }
```
*Response*
```http
  ### Status code: 200
  {
    "id": 2,
    "name": "Bola",
    "value": 11,
    "created_at": "2022-09-25T19:52:01.388815-03:00",
    "updated_at": "2022-09-25T19:53:54.586532334-03:00"
  }
```
![So](https://img.shields.io/badge/Report-Google_Chat_Webhook-red.svg)
<br>

#### -- *Delete Product [ Soft Delete ]*
*Request*
```http request
  DELETE http://localhost:8001/product/:id
```
*Response*

```http
  ### Status code: 200
```
![So](https://img.shields.io/badge/Report-Google_Chat_Webhook-red.svg)


## License

[MIT](https://choosealicense.com/licenses/mit/)
