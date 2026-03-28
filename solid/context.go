package solid

import "net/http"

type Context struct {
	w http.ResponseWriter
	r *http.Request
}