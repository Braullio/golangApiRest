
# GOLANG API REST

![So](https://img.shields.io/badge/S.O.-Linux-informational.svg)
![Go Version](https://img.shields.io/badge/Go-v.1.18-informational.svg)
![Database](https://img.shields.io/badge/Database-BigQuery-informational.svg)
<br>
![Repo Size](https://img.shields.io/badge/Repo_size-xxx_MB-informational.svg)
![Exec Size](https://img.shields.io/badge/Exec_size-xxx_MB-informational.svg)
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
  PROJECT_ID = "project-123456789"
  FILE_GOOGLE_CLOUD_KEY = "./security/bigquery_key.json"
  GOOGLE_CHAT_WEBHOOK = "https://chat.googleapis.com/v1/spaces/AAAAgOxrNhU/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=zk3322S04VnWdHW9y_4Zsg62CUZP9HvIvy0000uJ1qo%3D"
```

## Tables
Locate: Biquery <br>
Name: users
<br>

| Column     | Type     | Compilate |
|------------|----------|-----------|
| id         | STRING   | NULLABLE  |
| status     | STRING   | NULLABLE  |
| name       | STRING   | NULLABLE  |
| phone      | STRING   | NULLABLE  |
| created_at | DATETIME | NULLABLE  |
| updated_at | DATETIME | NULLABLE  |

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
<br>

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
<br>

#### -- *Search User*
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

## License

[MIT](https://choosealicense.com/licenses/mit/)
