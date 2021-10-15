package invoiceRepo

import (
	"TiBO_API/businesses/invoiceEntity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type InvoicesRepository struct {
	db *gorm.DB
}

func NewInvoicesRepository(db *gorm.DB) invoiceEntity.Repository {
	return &InvoicesRepository{
		db: db,
	}
}

func (r *InvoicesRepository) Create(ctx context.Context, invoiceData *invoiceEntity.Domain) (invoiceEntity.Domain, error) {
	rec := fromDomain(*invoiceData)

	err := r.db.Create(&rec).Error
	if err != nil {
		return invoiceEntity.Domain{}, err
	}
	return rec.toDomain(), err
}

func (r *InvoicesRepository) GetByUserID(ctx context.Context, userId string) ([]invoiceEntity.Domain, error) {
	rec := []Invoices{}

	err := r.db.Find(&rec, "user_id = ?", userId).Error
	if len(rec) == 0 {
		err = errors.New("Not Found")
		return nil, err
	}
	invoces := toDomainArray(rec)
	return invoces, nil
}

func (r *InvoicesRepository) GetID(ctx context.Context, ID uint) ([]invoiceEntity.Domain, error) {
	rec := []Invoices{}

	err := r.db.Find(&rec, "user_id = ?", ID).Error
	if len(rec) == 0 {
		err = errors.New("Not Found")
		return nil, err
	}
	invoces := toDomainArray(rec)
	return invoces, nil
}
