# Network benchmark

Build the binary

```
$> go build cmd/server/server.go
```

Run net/http web server

```
$> ./server -port=8080
```

Run fasthttp web server

```
$> ./server -fast -port=8081
```

Benchmark

```
# Take a look on net/http performance result!
$> hey -n=500000 -c=100 http://localhost:8081
# Take a look on fasthttp performance result!
$> hey -n=500000 -c=100 http://localhost:8080 #
```
