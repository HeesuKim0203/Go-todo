package main

import (
	"net/http"

	"github.com/goWeb/app"
	"github.com/urfave/negroni"
)

func main() {
	m := app.MakeHandler()
	n := negroni.Classic()

	n.UseHandler(m)

	err := http.ListenAndServe(":3001", n)

	if err != nil {
		panic(err)
	}
}