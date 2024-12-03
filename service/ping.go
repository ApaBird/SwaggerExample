package service

import "net/http"

// @Summary		Ping
// @Tags		ping
// @Description	Проверка доступности сервера
// @ID			Ping
// @Produce		text/plain
// @Success		200			{object}	EvenRecponce		"success"
// @Router		/Ping [get]
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
