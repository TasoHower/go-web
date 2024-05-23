# Odin-validator

## Getting started

### Prerequisites
- **[Go](https://go.dev/)**: Must use go 1.21.6 or higher version **[releases](https://go.dev/doc/devel/release)**.

### Get odin-validator
Clone Odin-validator source code from github with:
```s
$ git clone https://github.com/odinbtc/validator.git
```
### Config file
First you must create your configuration file. The configuration file of Odin-validator is a JSON file, and the path to the configuration file can be arbitrary. The default configuration file path is `./conf.json` under the source code path of Odin-validator.

#### Default configuration file
```json
{
    "app": {
        "runtime_rootPath": "./runtime/",
        "log_save_path": "logs/",
        "log_save_name": "web",
        "log_file_ext": "log",
        "log_level": "debug",
        "time_format": 2006010203,
        "expire_time": 1
    },
    "server": {
        "run_mode": "debug",
        "http_port": 8081,
        "read_timeout": 30,
        "write_timeout": 30,
        "shut_down_timeout": 30
    },
    "merkle": {
        "remote_path": "./data/merkle/remote/",
        "file_path": "./data/merkle/local/",
        "file_ext": ".json"
    }
}
```
- `app` specifies the path where the log file is stored
- `server` specifies the port and timeout time occupied by the local network server integrated by Odin-validator.
- `merkle` is used to configure the data storage path of Odin-validator (WARNING: the current data is stored in the form of json files, please do not modify `merkle-file_ext`)

### Load third-party libraries
Use the following code to load third-party libraries:
```s
$ go mod tidy
```

#### Third-party components used by Odin-validator
- **[gin-gonic/gin](https://github.com/gin-gonic/gin)**: Odin-validator integrates Gin to provide local web services.
- **[uber-go/zap](https://github.com/uber-go/zap)**: zap provides logging capabilities for Odin-validator.

### Run code
Use the following code in the Odin-validator root directory to start:
```sh
$ go run main.go
```
At the same time, you can use the following command to specify the configuration file to use:
```sh
$ go run main.go -c your_config_file_path
```
If you see the following output it means Odin-validator ran successfully:
```text
2024-02-21T15:59:53.548+0800    debug   config/config.go:42     Init config running.file path:[./conf.json]     {"pid": 6013, "process": "main"}
2024-02-21T15:59:53.549+0800    info    config/config.go:64     Init config success     {"pid": 6013, "process": "main"}
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /api/ping                 --> validator/web/router.InitRouter.TRPathParamHandler[...].func2 (4 handlers)
2024-02-21T15:59:53.550+0800    info    validator/main.go:73    Start http server listening :8081       {"pid": 6013, "process": "main"}
```

## Web api
### Web server test
#### Ping
- **Url**: /api/ping
- **Method**: GET
- **Request** :
- **Response**: pong

### Date
#### Date injection
- **Url**: /api/merkle/build
- **Method**: PUT
- **Request** :
```json
{
    "block": 779833,
    "ins":[
        "{\"tick\":\"\u4e2d\u6587\",\"op\":\"mint\",\"from\":\"bc1pt87kqa72x0v2qq3xlxuw0muz94umgqmcmk3eqq06hr8tcasjgppqd5r04w\",\"to\":\"bc1pxaneaf3w4d27hl2y93fuft2xk6m4u3wc4rafevc6slgd7f5tq2dqyfgy06\",\"amt\":\"1000\",\"balance\":\"2000\",\"available\":\"1000\",\"transferable\":\"2000\"}"
    ],
    "trx":[
        "{\"tick\":\"\u4e2d\u6587\",\"op\":\"mint\",\"from\":\"bc1pt87kqa72x0v2qq3xlxuw0muz94umgqmcmk3eqq06hr8tcasjgppqd5r04w\",\"to\":\"bc1pxaneaf3w4d27hl2y93fuft2xk6m4u3wc4rafevc6slgd7f5tq2dqyfgy06\",\"amt\":\"1000\",\"balance\":\"2000\",\"available\":\"1000\",\"transferable\":\"2000\"}"
    ]
}
```
- **Response**: 
```json
{
    "data": {},
    "code": 200,
    "request_id": "",
    "msg": "Success"
}
```
#### Date file get
- **Url**: /api/merkle/file/get
- **Method**: GET
- **Request** : block
- **Response**: 
```json
{
    "data": {
        "path": "/Users/******/go/src/validator/data/merkle/local/32/779832.json"
    },
    "code": 200,
    "msg": "Success"
}
```
