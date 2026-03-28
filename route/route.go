package route

import (
	"net/http"
	"solid/solid"
)

type Response struct {
	Message string `xml:"message"`
}

func Routes() {
	r := solid.NewRouter()

	r.Get("/", indexHandler)
}

func indexHandler(c *solid.Context) {
	var response = Response{
		Message: "Hello World",
	}
	
	solid.XmlResponse(c, response, http.StatusOK)
}
