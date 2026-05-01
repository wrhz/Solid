package solid

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type Session struct {
	session *sessions.Session
	w http.ResponseWriter
	r *http.Request
}

func (c *Context) Session(name string) (*Session, error) {
	session, err := settings.GetSessionStore().Get(c.Request, name)
	if err != nil {
		return nil, err
	}

	return &Session{ session: session, w: c.Writer, r: c.Request }, nil
}

func (s *Session) SetSession(name string, value any) (error) {
	s.session.Values[name] = value
	return s.session.Save(s.r, s.w)
}

func (s *Session) GetSession(name string) any {
	return s.session.Values[name]
}