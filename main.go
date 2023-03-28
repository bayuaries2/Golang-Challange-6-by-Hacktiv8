package main

import (
	"Challanges-6/routers"
)

func main() {
	var PORT = ":8080"
	routers.StarServer().Run(PORT)
}
