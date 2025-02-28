### Kafka Installetion

```js
docker run -d --name kafka -p 9092:9092 -p 2181:2181 wurstmeister/kafka:latest
```

### Run Service

- Run main service

```js
    go run api/main.go
```

- Run user service

```js
    cd cmd/userService && go run .
```
