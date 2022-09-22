
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
    "name": "Testenildo",
    "phone" : "5562981919191",
    "status": "Registred"
  }
```
*Response*
```http
  ### Status code: 201
  {
      "id": "dc53b354-2658-4719-af2d-1185b7304df4",
      "name": "Testenildo",
      "phone": "5562981919191",
      "status": "Registred",
      "created_at": "2022-09-22T16:12:59.425329621-03:00",
      "updated_at": "2022-09-22T16:12:59.425329621-03:00"
  }
```
<br>

[//]: # (TODO)
[//]: # (#### -- *Search User*)
[//]: # (*Request*)
[//]: # (```http)
[//]: # (  GET http://localhost:8001/user/:id)
[//]: # (```)
[//]: # (*Response*)
[//]: # (```http)
[//]: # (  ### Status code: 200)
[//]: # (  {)
[//]: # (    "id": "61404d0a-0000-492a-ba0b-e82f4535adbe",)
[//]: # (    "name": "Testenildo",)
[//]: # (    "phone" : "5562981919191",)
[//]: # (    "status": "Registred",)
[//]: # (    "created_at": "",)
[//]: # (    "updated_at": "")
[//]: # (  })
[//]: # (```)
[//]: # (<br>)

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
