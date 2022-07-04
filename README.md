# instant-go

This is the Back-end Project of Instant, and you can visit the Front-end Project at [instant-vue](https://github.com/ZYChimne/instant-vue).

## Features

* High Performance and Scalable
* Access: RESTful
* Logical: grpc
* Storage: Redis, Mongodb (https://www.mongodb.com/developer/products/mongodb/mongodb-schema-design-best-practices/)

## Project setup

```bash
git clone https://github.com/redis/redis.git # install redis 7
./redis-7.0.2/src/redis-server
brew services start mongodb-community@5.0
go run cmd/main.go
```

## TODO
  
* Read Escape Analysis