package handler

import (
	"encoding/json"
	"net/http"
	"validate/internal/model"
	"validate/internal/util"
	"validate/internal/validator"
)

// CarHandler handles the /car endpoint
func CarHandler(w http.ResponseWriter, r *http.Request) {
    var car model.Car
    if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := validator.Validate(car); err != nil {
        util.HandleValidationError(w, err, car)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(car)
}