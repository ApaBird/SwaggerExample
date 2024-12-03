package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type EvenRecponce struct {
	Even bool
}

// @Summary		Четность
// @Tags		numbers
// @Description	Проверка чётности числа
// @ID			Even
// @Accept		json
// @Produce		json
// @Param		num			query		integer		true	"Число"
// @Success		200			{object}	EvenRecponce		"success"
// @Failure		400,500		{object}	errorResponce		"error"
// @Router		/Even [get]
func Even(w http.ResponseWriter, r *http.Request) {
	num := r.URL.Query().Get("num")

	if num == "" {
		ErrorResponce(fmt.Errorf("num is empty"), "the num parameter requires an int", http.StatusBadRequest, w)
		return
	}

	var numInt int
	var err error
	if numInt, err = strconv.Atoi(num); err != nil {
		ErrorResponce(err, "the num parameter requires an int", http.StatusBadRequest, w)
		return
	}

	responce := EvenRecponce{Even: (num == strconv.Itoa(2*int(numInt)))}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responce); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
