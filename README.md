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

## Run Docker Images for background pass -d 
server porst must be same (e.g. 9080)
``` 
docker run -p 8080:9090 -e "PORT=9090" transaction  
```

## Run DB and Transaction docker all together 
```
docker compose up
```

## Bring all server Down
```
docker compose down
```

## Sample request
http://localhost:8080/transactions/0051c7df-5029-48ff-ad8b-0b5d8a4d3272
http://localhost:8080/transactions/account/GXZB08691413145620?status=complete

## Other Docker comand 
docker ps
docker logs -f 2032a1f8d008
docker stop 8e33bd1ffe8c


## Postgres deployment (Individual)
docker run -itd -e POSTGRES_USER=shail -e POSTGRES_PASSWORD=shail -p 5432:5432 -v data:/var/lib/postgresql/data --name postgresql postgres



