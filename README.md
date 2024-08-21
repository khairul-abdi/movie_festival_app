# movie_festival_app

## Tools

- Go Programming Language
- Gin Gonic Framework
- Postgresql
- Orm Gorm

## Prepare

# Usage

First thing first we have to create db :

```
    create database postgresql with the same name like dicrectory package/database/config.go
```

and then install all of the dependencies by typing this command in terminal :

```
    go mod tidy
```

and then we are ready to execute the server by this following command :

```
    go run main.go
```

# Features

## User

### Create User Register

```
curl --location --request POST 'http://localhost:9001/users/register' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-raw '{
    "email": "nia@gmail.com",
    "username": "nia",
    "password":"123456",
    "age":9
}'
```

### Response

```
{
  "code": 201,
  "data": {
    "age": 9,
    "email": "nia@gmail.com",
    "id": 8,
    "username": "nia"
  },
  "message": "Success insert data"
}
```

### Create User Login

```
curl --location --request POST 'http://localhost:9001/users/login' \
--header 'Content-Type: application/json' \
--data-raw '{
  "email": "nia@gmail.com",
  "password":"123456"
}'
```

### Response

```
{
  "code": 200,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5pYUBnbWFpbC5jb20iLCJpZCI6OH0.bMWIYZRb7g78dgKr_0YfkWmYGLMFFfNMQcHMbfRJXww"
  },
  "message": "Success Create Token"
}
```

### Update User

```
curl --location --request PUT 'http://localhost:9001/users/8' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5pYUBnbWFpbC5jb20iLCJpZCI6OH0.bMWIYZRb7g78dgKr_0YfkWmYGLMFFfNMQcHMbfRJXww' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "niakurniati@gmail.com",
    "username": "nia kurniati"
}'
```

### Response

```
{
  "code": 200,
  "data": {
    "age": 9,
    "email": "niakurniati@gmail.com",
    "id": 8,
    "updated_at": "2022-03-29T08:51:36.031953569+07:00",
    "username": "nia kurniati"
  },
  "message": "Success Update user"
}
```

### Delete User

```
curl --location --request DELETE 'http://localhost:9001/users/8' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5pYUBnbWFpbC5jb20iLCJpZCI6N30.ZwNE9OAbPQCM3t18jkBf-jwLQxS4ppIWd0D1Wmt04IQ' \
--data-raw ''
```

### Response

```
{
  "code": 200,
  "message": "Your user has been successfully deleted"
}
```
