# redis-info-svc

A basic service proxy to `redis-cli info` based on go-kit.

## Installation

Install Go. https://golang.org/doc/install

```
git clone https://github.com/ericksond/redis-info-svc.git
cd redis-info-svc
go get
go build
./redis-info-svc -port=[default 8080]
```

Make a post to retrieve redis info.

```
curl -XPOST -d'{"addr":"localhost:6379", "passwd": ""}' http://localhost:8080/info 
```

Password is optional.

### Sample Output

```
{"Info":{"aof_current_rewrite_time_sec":"-1","aof_enabled":"0","aof_last_bgrewrite_status":"ok"...}}
```
