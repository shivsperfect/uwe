package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/shivsperfect/uwe/types"
)

func HandleGetCustomer(w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return ShowAPIError(http.StatusBadRequest, err)
	}
	customer := types.Customer{
		ID: id,
	}

	writeJSON(w, http.StatusOK, customer)
	return nil
}
