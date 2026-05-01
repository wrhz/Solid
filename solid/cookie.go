package solid

import (
	"net/http"
	"time"
)

type Cookie struct {
	Name   string
	Value  string
	Quoted bool
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string
	MaxAge      int
	Secure      bool
	HttpOnly    bool
	SameSite    http.SameSite
	Partitioned bool
	Raw         string
	Unparsed    []string
}

func (c *Context) SetCookie(cookie *Cookie) {
	cookies := &http.Cookie{
		Name:   cookie.Name,
		Value:  cookie.Value,
		Quoted: cookie.Quoted,
		Path:   cookie.Path,
		Domain: cookie.Domain,
		Expires:  cookie.Expires,
		RawExpires: cookie.RawExpires,
		MaxAge:   cookie.MaxAge,
		Secure:   cookie.Secure,
		HttpOnly: cookie.HttpOnly,
		SameSite: cookie.SameSite,
		Partitioned: cookie.Partitioned,
		Raw:    cookie.Raw,
		Unparsed: cookie.Unparsed,
	}

	http.SetCookie(c.Writer, cookies)
}

func (c *Context) GetCookie(name string) (*Cookie, error) {
	cookie, err := c.Request.Cookie(name)
	if err != nil {
		return nil, err
	}

	cookies := &Cookie{
		Name:   cookie.Name,
		Value:  cookie.Value,
		Quoted: cookie.Quoted,
		Path:   cookie.Path,
		Domain: cookie.Domain,
		Expires:  cookie.Expires,
		RawExpires: cookie.RawExpires,
		MaxAge:   cookie.MaxAge,
		Secure:   cookie.Secure,
		HttpOnly: cookie.HttpOnly,
		SameSite: cookie.SameSite,
		Partitioned: cookie.Partitioned,
		Raw:    cookie.Raw,
		Unparsed: cookie.Unparsed,
	}

	return cookies, nil
}