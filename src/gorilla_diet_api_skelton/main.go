package main

import (
	"fmt"
	"net/http"

	"gorilla_diet_api_skelton/web"
)

var port = 8080

func main() {
	router := web.Router()
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
