package solid

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gogf/gf/v2/encoding/gjson"
)

func StringResponse(c *Context, s string, status int) {
	c.w.Header().Set("Content-Type", "text/plain")
	c.w.WriteHeader(status)

	fmt.Fprintf(c.w, "%s", s)
}

func JsonResponse(c *Context, data any, status int) {
	var jsonData = gjson.New(data).MustToJsonString()

	c.w.Header().Set("Content-Type", "application/json")
	c.w.WriteHeader(status)

	fmt.Fprintf(c.w, "%s", jsonData)
}

func HtmlResponse(c *Context, html string, status int) {
	c.w.Header().Set("Content-Type", "text/html")
	c.w.WriteHeader(status)

	fmt.Fprintf(c.w, "%s", html)
}

func HtmlFileResponse(c *Context, file string, status int) {
	var html, err = os.ReadFile(filepath.Join(".", "resource", "view", file))
	if err != nil {
		c.w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(c.w, "Failed to read html file: %s", err)
		return
	}
	c.w.Header().Set("Content-Type", "text/html")
	c.w.WriteHeader(status)

	fmt.Fprintf(c.w, "%s", html)
}

func XmlResponse(c *Context, data any, status int) {
	var xmlData = gjson.New(data).MustToXmlString()

	c.w.Header().Set("Content-Type", "application/xml")
	c.w.WriteHeader(status)

	fmt.Fprintf(c.w, "%s", xmlData)
}