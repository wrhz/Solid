package solid

import "net/http"

type Context struct {
	Writer http.ResponseWriter
	Request *http.Request
}

func (c *Context) RequestID() string {
	if id, ok := c.Request.Context().Value("requestID").(string); ok {
		return id
	}

	return ""
}
