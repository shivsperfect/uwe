package handler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/shivsperfect/uwe/db"
	"github.com/shivsperfect/uwe/types"
)

type CreateFileUploadRequest struct {
	FileType types.FileType `json:"fileType"`
	Mapping  map[string]int `json:"mapping"`
}

type CreateFileUploadResponse struct {
	ID uuid.UUID `json:"id"`
}

type UploadHandler struct {
	db db.DB
}

func NewUploadHandler(db db.DB) *UploadHandler {
	return &UploadHandler{
		db: db,
	}
}

func (h *UploadHandler) HandleCreateFileUpload(w http.ResponseWriter, r *http.Request) error {
	var req CreateFileUploadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	fileUpload := types.FileUpload{
		ID:         uuid.New(),
		Type:       req.FileType,
		CustomerID: uuid.MustParse("25391335-1c6e-42be-9931-776179b4e8c1"),
		Mapping:    req.Mapping,
	}
	if err := h.db.CreateFileUpload(&fileUpload); err != nil {

		return err
	}
	resp := CreateFileUploadResponse{
		ID: fileUpload.ID,
	}
	return writeJSON(w, http.StatusCreated, resp)
}

func (h *UploadHandler) HandleFileUpload(w http.ResponseWriter, r *http.Request) error {
	fileID, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		return err
	}

	fileUpload, err := h.db.GetFileUploadByID(fileID)
	if err != nil {
		return err
	}

	fmt.Println(fileUpload)

	subs, err := processSubscriptionsUpload(r.Body, fileUpload.Mapping)
	if err != nil {
		return err
	}

	fmt.Println(subs)

	return nil
}

type Mapping struct {
	Amount     int
	Currency   int
	Period     int
	ExternalID int
	StartedAt  int
	CanceledAt int
}

func processSubscriptionsUpload(r io.Reader, mapping map[string]int) ([]types.Subscription, error) {
	if err := validateSubscriptionMapping(mapping); err != nil {
		return nil, fmt.Errorf("invalid mapping for subscriptions: %s", err)
	}
	reader := csv.NewReader(r)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading CSV: %s", err)
		}
		amount := record[mapping["amount"]]
		_ = amount
		fmt.Println("record", record)
	}
	return nil, nil
}

func validateSubscriptionMapping(m map[string]int) error {
	allowedKeys := []string{
		"amount",
		"currency",
		"period",
		"vat",
		"external_id",
		"started_at",
		"canceled_at",
	}
	for _, key := range allowedKeys {
		if _, ok := m[key]; !ok {
			return fmt.Errorf("invalid subscription mapping: %s", key)
		}
	}
	return nil
}
