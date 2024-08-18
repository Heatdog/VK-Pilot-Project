package login

import (
	"VK-Pilot-Project/internal/models/auth"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

// Вход в систему
// @Summary Login
// @Description Вход в систему
// @ID login
// @Tags login
// @Accept json
// @Produce json
// @Param input body auth.ModelRequest true "auth info"
// @Success 200 {object} auth.ModelResponse токен аутентификации
// @Failure 400 {object} auth.ErrorResponse Некорректные данные
// @Failure 500 {object} auth.ErrorResponse Внутренняя ошибка сервера
// @Router /api/login [post]
func (handler *handler) login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		handler.logger.Error("read body", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var user auth.ModelRequest

	if err := json.Unmarshal(body, &user); err != nil {
		handler.logger.Error("unmarshal body", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := handler.loginService.Login(r.Context(), user)
	if err != nil {
		handler.logger.Error("user login", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := handler.tokenService.Generate(r.Context(), id)
	if err != nil {
		handler.logger.Error("token generate", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handler.writeToken(w, token)
}

func (handler *handler) writeToken(w http.ResponseWriter, token string) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")

	resp, err := json.Marshal(auth.ModelResponse{
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
