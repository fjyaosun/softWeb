package main

import (
	"SoftWeb/jhttp/api"
	"SoftWeb/lib"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	for key, value := range api.GetRouterMap() {
		r.HandleFunc(key, value)
	}

	n := negroni.New(lib.NewRecoveryNotice(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")))
	n.UseHandler(r)
	n.Run(":3000")
}
