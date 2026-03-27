package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/shivsperfect/uwe/types"
)

func (db DB) CreateCustomer(c *types.Customer) error {
	_, err := db.NewInsert().Model(c).Exec(context.Background())
	return err
}

func (db DB) CreateFileUpload(f *types.FileUpload) error {
	_, err := db.NewInsert().Model(f).Exec(context.Background())
	return err
}

func (db DB) GetFileUploadByID(id uuid.UUID) (types.FileUpload, error) {
	var fileUpload types.FileUpload
	err := db.NewSelect().
		Model(&fileUpload).
		Where("id = ?", id).
		Scan(context.TODO(), &fileUpload)

	return fileUpload, err
}
