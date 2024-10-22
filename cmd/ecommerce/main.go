package main

import "github.com/devpenguin/go-ecommerce/internal/initialize"

func main() {
	r := initialize.Run()

	r.Run(":8002")
}