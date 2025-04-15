package service

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/josevitorrodriguess/client-manager/internal/utils"
	"github.com/josevitorrodriguess/client-manager/internal/validators"
)

type ServiceRequest struct {
	CustomerID  uuid.UUID      `json:"customer_id"`
	TypeProduct string         `json:"type_product"`
	Description string         `json:"description"`
	TotalValue  pgtype.Numeric `json:"total_value"`
	DownPayment pgtype.Numeric `json:"down_payment"`
	IsPaid      bool           `json:"is_paid"`
	IsFinished  bool           `json:"is_finished"`
}

func (pr *ServiceRequest) IsValid() (bool, error) {
	validationErrs := validators.ValidationErrors{
		Errors: make(map[string]string),
	}
	if !utils.NotBlank(pr.CustomerID.String()) {
		validationErrs.Errors["customer_id"] = "customer id cannot be empty"
	}

	if !utils.NotBlank(pr.TypeProduct) {
		validationErrs.Errors["type_product"] = "type product cannot be empty"
	}

	if !(utils.MinChars(pr.Description, 5) && utils.MaxChars(pr.Description, 255)) {
		validationErrs.Errors["description"] = "description must have between 5 and 255 characters"
	}

	if !validationErrs.HasErrors() {
		return true, nil
	}

	return false, validationErrs
}

type UpdateServiceStatusRequest struct {
	ID     int32 `json:"id"`
	Status bool  `json:"status"`
}
