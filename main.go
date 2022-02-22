package main

import (
	"gin_demo/Routers"
	"gin_demo/app/global/variable"
	_ "gin_demo/bootstrap"
)

func main() {
    router := Routers.InitApiRouter()
    _ = router.Run(variable.ConfigYml.GetString("HttpServer.Api.Port"))
}
