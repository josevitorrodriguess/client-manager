package product

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/josevitorrodriguess/client-manager/internal/utils"
)

type ProductRequest struct {
	CustomerID  uuid.UUID      `json:"customer_id"`
	TypeProduct string         `json:"type_product"`
	Description string         `json:"description"`
	TotalValue  pgtype.Numeric `json:"total_value"`
	DownPayment pgtype.Numeric `json:"down_payment"`
	IsPaid      pgtype.Bool    `json:"is_paid"`
	IsFinished  pgtype.Bool    `json:"is_finished"`
}

func (pr *ProductRequest) IsValid() (bool, error) {
	if uuid.Nil == pr.CustomerID {
		return false, fmt.Errorf("customer id cannot be empty")
	}

	if pr.TypeProduct == "" {
		return false, fmt.Errorf("type product cannot be empty")
	}

	if !(utils.MinChars(pr.Description, 5) && utils.MaxChars(pr.Description, 255)) {
		return false, fmt.Errorf("description must have between 5 and 255 characters")
	}


	return true, nil
}





