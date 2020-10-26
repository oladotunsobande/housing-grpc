package main

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/oladotunsobande/housing-grpc/config"
	"github.com/oladotunsobande/housing-grpc/gateway/routes"
)

func main() {
	n := negroni.Classic()
	n.UseHandler(routes.RegisterRoutes())

	err := http.ListenAndServe(":"+config.GetSecrets().ApplicationPort, n)
	if err != nil {
		log.Fatal(err)

		if r := recover(); r != nil {
			err = r.(error)
		}
	}
}
