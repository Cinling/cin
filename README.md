[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/cinling/cin)](https://github.com/cinling/cin/tags)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cinling/cin?color=red)
![GitHub last commit](https://img.shields.io/github/last-commit/cinling/cin)

# cin
[简体中文](https://github.com/Cinling/cin/blob/master/doc/README.zh-cn.md)

This repostory has been deprecated. new repostory:  [go-cam/cam](https://github.com/go-cam/cam)

The http and websocket server's framework.

Inspiration come form yii2.

# Dependencies
- github.com/go-sql-driver/mysql v1.4.1
- github.com/go-xorm/xorm v0.7.9
- github.com/golang-migrate/migrate/v4 v4.7.0
- github.com/gorilla/sessions v1.2.0
- github.com/gorilla/websocket v1.4.1
- github.com/tidwall/gjson v1.3.4

# Getting started

Edit file:  `main.go`
```go
package main

import (
    "github.com/cinling/cin"
    "github.com/cinling/cin/core/base"
)

func main() {
	config := cin.NewConfig()
    config.ComponentDict = map[string]base.ConfigComponentInterface{
        "ws":      cin.NewConfigWebsocketServer(24600),
        "http":    cin.NewConfigHttpServer(24601).SetSessionName("test"),
        "db":      cin.NewConfigDatabase("mysql", "127.0.0.1", "3306", "cin", "root", "root"),
        "console": cin.NewConfigConsole(),
    }
    cin.App.AddConfig(config)
    cin.App.Run()
}
```

Run go module to download dependency. 
> go mod tidy

Compiled code.
> go build main.go

Run bin file
> ./main.exe  `or` ./main


# Migrations
After Compiled code, run the following command to create migration's file.
> ./main.exe migrate/create [filename]
