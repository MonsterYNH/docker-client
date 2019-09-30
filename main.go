package main

import "test/router"

func main() {
	router.GetRouter().Run("0.0.0.0:8999")
}