package data

import (
	"VK-Pilot-Project/internal/models/data"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

// Запись данных
// @Summary Wite
// @Description Запись данных
// @Security ApiKeyAuth
// @ID data-write
// @Tags data
// @Accept json
// @Produce json
// @Param input body data.Write true "write data"
// @Success 201 {object} data.StatusResponse статус операции
// @Failure 401 {object} nil Пользователь не авторизован
// @Failure 400 {object} string Некорректные данные
// @Failure 500 {object} string Внутренняя ошибка сервера
// @Router /api/write [post]
func (handler *handler) write(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		handler.logger.Error("read body", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var request data.Write

	if err := json.Unmarshal(body, &request); err != nil {
		handler.logger.Error("unmarshal body", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.service.Write(r.Context(), request); err != nil {
		handler.logger.Error("service error", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handler.writeResult(w)
}

func (handler *handler) writeResult(w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("content-type", "application/json")

	resp, err := json.Marshal(data.StatusResponse{
		Status: "success",
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
