
# Grep like API

The service has the following responsibilities

1. Enables client to search the log files stored at s3


## Run Locally

Clone the project

```bash
  git clone git@github.com:Sant470/grep-api.git
```

Go to the project directory

```bash
  cd grep-api
```

```bash
  add credential in config.yaml
```


Install dependencies

```bash
  go mod tidy
```

Start the server

```bash
  go run main.go
```

Use the below curl to test it

```bash
  curl --location 'http://localhost:8000/api/v1/search' \
--header 'Content-Type: application/json' \
--data '{
    "search_keyword": "hello, world",
    "from": 1704516560,
    "to": 1708636658
}'
```
## Dependency && Installation

It requires go version 1.21.5, you can download it following the guide mentioned below 

```bash
   https://go.dev/dl/
```
    
