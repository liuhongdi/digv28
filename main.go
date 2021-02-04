package main

import (
	"github.com/liuhongdi/digv28/router"
)

func main() {
	//引入路由
	r := router.Router()
	//run
	r.Run(":8080")
}
