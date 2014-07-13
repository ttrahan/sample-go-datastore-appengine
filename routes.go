package routes

import (
	"net/http"

	"github.com/Shippable/sample-go-datastore-appengine/hello"
)

func init() {
	http.HandleFunc("/", hello.Handler)
}
