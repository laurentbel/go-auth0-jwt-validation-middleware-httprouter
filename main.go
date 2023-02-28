package main

import (
	"fmt"
	"go-auth0-jwt-validation-middleware-httprouter/middlewares"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HelloRoute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s\n", ps.ByName("name"))
}

func HelloRouteWithParameters(append string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "hello, %s %s\n", ps.ByName("name"), append)
	}
}

func main() {

	issuer := "https://<your-auth0-issuer-url>/"
	scope := "<your-api-scope>"

	router := httprouter.New()
	router.GET("/helloUnsecure/:name", HelloRoute)
	router.GET("/helloUnsecureWrapped/:name", HelloRouteWithParameters(" !!!"))
	router.GET("/helloSecure/:name", middlewares.JwtValidationMiddleware(HelloRoute, issuer, scope))
	router.GET("/helloSecureWrapped/:name", middlewares.JwtValidationMiddleware(HelloRouteWithParameters(" !!!"), issuer, scope))

	log.Fatal(http.ListenAndServe(":8080", router))
}
