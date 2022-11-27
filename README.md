# Test Go Kit

This repository is for testing [Go kit](https://github.com/go-kit/kit) repository by extending some of the [gokit.io/examples](https://gokit.io/examples). 

## Dev
```shell
# Start service
docker-compose up
```

After start run:
```shell
# hit count endpoint
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
# hit to uppercase endpoint
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase

```

## TODO
* Add linter
* Test runner
* Code coverage

## DONE
* Dockerfile
* Makefile
