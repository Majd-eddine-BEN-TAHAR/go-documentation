package handler

import (
	"encoding/json"
	"net/http"
	"validate/internal/model"
	"validate/internal/util"
	"validate/internal/validator"
)

// UserHandler handles the /user endpoint
func UserHandler(w http.ResponseWriter, r *http.Request) {
    var user model.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	if err := validator.Validate(user); err != nil {
        util.HandleValidationError(w, err, user)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}