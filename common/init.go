package common

import "github.com/gorilla/sessions"

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func init() {
	store.Options = &sessions.Options{
		//Domain:   "localhost",
		Path:   "/",
		MaxAge: 0, // 8 hours
		//HttpOnly: true,
	}

}
