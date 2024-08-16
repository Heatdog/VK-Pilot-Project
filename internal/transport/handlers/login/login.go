package login

import (
	"VK-Pilot-Project/internal/models/auth"
	"encoding/json"
	"io"
	"net/http"
)

func (handler *handler) login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var user auth.Model

	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := handler.loginService.Login(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := handler.tokenService.Generate(r.Context(), id.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handler.writeToken(w, token)
}

func (handler *handler) writeToken(w http.ResponseWriter, token string) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")

	resp, err := json.Marshal(struct {
		Token string
	}{
		Token: token,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
