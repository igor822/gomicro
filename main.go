package main

import (
    "log"
    "net/http"

    "github.com/igor822/gomicro/controller"
)

func main() {
	router, err := controller.MapHandlers()
	if err != nil {
		panic(err)
	}

	log.Println("shipping-api started listening in :")
	http.ListenAndServe(":80", router)
}
