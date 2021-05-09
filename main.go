package main

import (
	"echo-poc/handlers"
	"echo-poc/router"
)

func main() {
	r := router.New()
	v1 := r.Group("/api")
	h := handlers.NewHandler()
	h.Register(v1)
}
