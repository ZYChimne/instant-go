# instant-go

This is the Back-end Project of Instant, and you can visit the Front-end Project at [instant-next](https://github.com/ZYChimne/instant-next).

## Features

* High Performance, High Availability and Scalable
* Access: RESTful
* Storage: Redis & PostgreSQL
* Fan out on Write (Both Posts and Messages)
* [Efficient Flow Counting Algorithm](https://www.usenix.org/system/files/conference/atc18/atc18-gong.pdf) and Hot Spot Detection

## Project setup

```bash
gofmt -w ./
go run cmd/main.go
```
```
$env:GOPROXY = "https://proxy.golang.com.cn,direct"
set http_proxy=socks5://127.0.0.1:7890
set https_proxy=%http_proxy%
```