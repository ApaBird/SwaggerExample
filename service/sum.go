package service

import (
	"encoding/json"
	"net/http"
)

type SumData struct {
	Data []float64
}

type SumResponce struct {
	Sum float64
}

// @Summary		Сумма массива чисел
// @Tags		numbers
// @Description	Сумма массива чисел переданого в структуре SumData
// @ID			SumSlice
// @Accept		json
// @Produce		json
// @Param		query		body		SumData		true	"Данные"
// @Success		200			{object}	SumResponce		"success"
// @Failure		400,500		{object}	errorResponce	"error"
// @Router		/Sum [post]
func SumSlice(w http.ResponseWriter, r *http.Request) {
	var data SumData

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		ErrorResponce(err, "invalid data", http.StatusBadRequest, w)
		return
	}

	var responce SumResponce

	for _, v := range data.Data {
		responce.Sum += v
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responce); err != nil {
		ErrorResponce(err, "error encode", http.StatusInternalServerError, w)
	}
}
