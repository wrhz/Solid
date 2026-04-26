package solid

import (
	"encoding/xml"
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

func BytesResponse(c *Context, data []byte, status int) {
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.Writer.WriteHeader(status)

	c.Writer.Write(data)
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

func ViewResponse(c *Context, file string, status int) {
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
	var xmlData, err = xml.Marshal(data)
	
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(c.Writer, "Failed to marshal xml: %s", err)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/xml")
	c.Writer.WriteHeader(status)

	fmt.Fprintf(c.Writer, "%s", xmlData)
}