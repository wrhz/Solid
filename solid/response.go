package solid

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gogf/gf/v2/encoding/gjson"
)

func StringResponse(c *Context, s string, status int) {
	c.Writer.Header().Set("Content-Type", "text/plain")
	c.Writer.WriteHeader(status)

	fmt.Fprintf(c.Writer, "%s", s)
}

func JsonResponse(c *Context, data any, status int) {
	var jsonData = gjson.New(data).MustToJsonString()

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)

	fmt.Fprintf(c.Writer, "%s", jsonData)
}

func HtmlResponse(c *Context, html string, status int) {
	c.Writer.Header().Set("Content-Type", "text/html")
	c.Writer.WriteHeader(status)

	fmt.Fprintf(c.Writer, "%s", html)
}

func HtmlFileResponse(c *Context, file string, status int) {
	var html, err = os.ReadFile(filepath.Join(".", "resource", "view", file))
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(c.Writer, "Failed to read html file: %s", err)
		return
	}
	c.Writer.Header().Set("Content-Type", "text/html")
	c.Writer.WriteHeader(status)

	fmt.Fprintf(c.Writer, "%s", html)
}

func XmlResponse(c *Context, data any, status int) {
	var xmlData = gjson.New(data).MustToXmlString()

	c.Writer.Header().Set("Content-Type", "application/xml")
	c.Writer.WriteHeader(status)

	fmt.Fprintf(c.Writer, "%s", xmlData)
}