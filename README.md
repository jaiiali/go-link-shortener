# Link Shortener

Link Shortener with Hexagonal Architecture

## Up and Running

### Export environment variables
```shell
export $(grep -v '^#' ./.env | xargs -d '\n')
```

### Run database migrations
```shell
make migrate-up
```

### Build the project

```shell
make build
```

### Run the project

Run the project on `ip` and `port` that are sets in `.env` file:
```shell
make run
```
or
```shell
./bin/go-linke-shotener
```

## APIs

### Create Short Link

Request:
```shell
curl -X POST http://0.0.0.0:8000/api/links --data '{"url": "https://google.com/"}'
```

Response:
```json
{
    "id": "ZDI3Z",
    "original": "https://google.com/",
    "created_at": "2023-01-30 00:00:00"
}
```

### Links List

Request:
```shell
curl -X GET http://0.0.0.0:8000/api/links
```

Response:
```json
[
    {
        "id": "ZDI3Z",
        "original": "https://google.com/",
        "created_at": "2023-01-30 00:00:00"
    }
]
```

### Get Link

Request:
```shell
curl -X GET http://0.0.0.0:8000/api/links/{id}
```

Response:
```json
{
    "id": "ZDI3Z",
    "original": "https://google.com/",
    "created_at": "2023-01-30 00:00:00"
}
```

### Healthcheck

Request:
```shell
curl -X POST http://0.0.0.0:8000/api/health
```