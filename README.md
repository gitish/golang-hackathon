# Hackathon
How to set in local
* Run below Command
```
    // Set environment variable as DATA_DIR for data store
    export DATA_DIR=/Users/shashail/data/go
    export STATIC_DIR=/Users/shashail/myWs/GoLang/go-credit/static
    //build and run go application
    go get
    go mod tidy
    go run .
```

## Create Docker Image
```
docker images ls
docker build --tag transaction .
```

## Run Docker Images in background
``` docker run -d -p 8080:8080 transaction  ```

## Other Docker comand 
docker ps
docker logs -f 2032a1f8d008
docker stop 8e33bd1ffe8c


## Postgres deployment 
docker run -itd -e POSTGRES_USER=shail -e POSTGRES_PASSWORD=shail -p 5432:5432 -v data:/var/lib/postgresql/data --name postgresql postgres



