# Simple Rest Api Blog 

## Golang Rest api with Sqlite

### require 
- golang v.1.18
- sqlite

## How to install

- Clone Repository 
  ```$ git clone https://github.com/```
- ```$ go mod tidy ```
- ```$ go run . ```


## Features
- Create Post ( Blog content)
- FindAll
- FindOne
- Update
- Delete


## Api Specifications

- Create Post

```
Request:

- Method : POST
- Endpoint : `baseURL/`
- Header :
    - Content-Type : application/json
    - Accept : application/json
```
```json
{
    "title": "bogor saja",
    "content": "Hello world"
}
```

- GET ALL Post

```
Request:

- Method : GET
- Endpoint : `baseURL/`
- Header :
    - Content-Type : application/json
    - Accept : application/json
```
Response:
```json
[
    {
        "id": 1,
        "title": "Hello saja",
        "content": "Hello world test",
        "published_at": "0001-01-01T00:00:00Z",
        "created_at": "2022-05-21T13:51:12.620949+07:00",
        "updated_at": "2022-05-21T14:30:40.54439+07:00"
    },
    {
        "id": 2,
        "title": "Hello saja yes",
        "content": "Hello world test",
        "published_at": "0001-01-01T00:00:00Z",
        "created_at": "2022-05-21T14:24:55.130547+07:00",
        "updated_at": "2022-05-21T14:30:52.909087+07:00"
    }
]
```

- GET Post By ID

```
Request:

- Method : GET
- Endpoint : `baseURL/{id}`
- Header :
    - Content-Type : application/json
    - Accept : application/json
```
Response:
```json
 {
    "id": 1,
    "title": "Hello saja",
    "content": "Hello world test",
    "published_at": "0001-01-01T00:00:00Z",
    "created_at": "2022-05-21T13:51:12.620949+07:00",
    "updated_at": "2022-05-21T14:30:40.54439+07:00"
}
```


- Update Post

```
Request:

- Method : PUT
- Endpoint : `baseURL/{id}`
- Header :
    - Content-Type : application/json
    - Accept : application/json
```
```json
{
    "title": "bogor saja",
    "content": "Hello world"
}
```

Response:

```json
{
    "id":4,
    "title": "bogor saja",
    "content": "Hello world",
    "published_at": "0001-01-01T00:00:00Z",
    "created_at": "2022-05-21T13:51:12.620949+07:00",
    "updated_at": "2022-05-21T14:30:40.54439+07:00"
}
```


- Delete Post

```
Request:

- Method : DELETE
- Endpoint : `baseURL/{id}`
- Header :
    - Content-Type : application/json
    - Accept : application/json
```

Response:

```json
{
    "id":4,
    "title": "bogor saja",
    "content": "Hello world",
    "published_at": "0001-01-01T00:00:00Z",
    "created_at": "2022-05-21T13:51:12.620949+07:00",
    "updated_at": "2022-05-21T14:30:40.54439+07:00"
}
```