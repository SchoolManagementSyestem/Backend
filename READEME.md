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

### Authorization and Authentication and TenantId

- Authorization

```js
	err = helpers.Authorization(&user, "principal")
	if err != nil {
		return nil, err
	}
```

- Authorizations

```js
	err = helpers.Authorizations(&user, []string{"teacher", "principal"})
	if err != nil {
		return nil, err
	}
```

- Authentication

```js
	user, err := helpers.Authentication(&p)

	if err != nil {
		return nil, err
	}

	fmt.Println(user)
```

- TenantId

```js
	tenantId, err := helpers.GetTenantId(&p)
	if err != nil {
		return nil, err
	}
    fmt.Println(tenantId)
```
