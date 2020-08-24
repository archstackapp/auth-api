package oauth2

import (
	"net/http"

	"github.com/go-session/session"
)

func (o *OAuth2) passwordAuthorizationHandler(mail string, password string) (string, error) {
	user, err := o.users.Repository.GetByMail(mail)
	if err != nil {
		return "", err
	}

	err = o.verifyUser(user, password)
	if err != nil {
		return "", err
	}

	return user.ID.String(), err
}

func (o *OAuth2) userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		return
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}

		store.Set("ReturnUri", r.Form)
		store.Save()

		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	userID = uid.(string)
	store.Delete("LoggedInUserID")
	store.Save()
	return
}
