package data

import (
	"VK-Pilot-Project/internal/models/data"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

// Чтение данных
// @Summary Read
// @Description Чтение данных
// @Security ApiKeyAuth
// @ID data-read
// @Tags data
// @Accept json
// @Produce json
// @Param input body data.KeysStruct true "read keys"
// @Success 200 {object} data.DataStruct статус операции
// @Failure 401 {object} nil Пользователь не авторизован
// @Failure 400 {object} string Некорректные данные
// @Failure 500 {object} string Внутренняя ошибка сервера
// @Router /api/read [post]
func (handler *handler) read(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		handler.logger.Error("read body", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var request data.KeysStruct

	if err := json.Unmarshal(body, &request); err != nil {
		handler.logger.Error("unmarshal body", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := handler.service.Read(r.Context(), request)
	if err != nil {
		handler.logger.Error("service error", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handler.readResult(w, res)
}

func (handler *handler) readResult(w http.ResponseWriter, res data.DataStruct) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")

	resp, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
