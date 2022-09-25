#### go test ####
```shell
go test -v ./... -coverprofile=c.out
```

#### test coverage ####
````shell
go tool cover -func=c.out
````
