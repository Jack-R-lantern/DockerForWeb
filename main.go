package main

import (
	"golang_framework/routers"
)

func main() {
	r := routers.InitRouter()
	r.Run()
}
