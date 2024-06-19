package main

import (
	"github.com/kd-o-guh/Library-API/routes"
)

func main() {
	router := routes.RouterInit()
	router.Run("127.0.0.1:8080")
}
