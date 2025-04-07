package customer

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/josevitorrodriguess/client-manager/internal/utils"
)

type CustomerPFRequest struct {
	Type        string      `json:"type"`
	Email       string      `json:"email"`
	Phone       string      `json:"phone"`
	Cpf         string      `json:"cpf"`
	Name        string      `json:"name"`
	BirthDate   pgtype.Date `json:"birth_date"`
	AddressType string      `json:"address_type"`
	Street      string      `json:"street"`
	Number      string      `json:"number"`
	Complement  pgtype.Text `json:"complement"`
	State       string      `json:"state"`
	City        string      `json:"city"`
	Cep         string      `json:"cep"`
}

func (cPFr *CustomerPFRequest) IsValid() (bool, error) {
	if !(utils.MinChars(cPFr.Name, 5) && utils.MaxChars(cPFr.Name, 100)) {
		return false, fmt.Errorf("name must have between 5 and 100 characters")
	}
	if !utils.Matches(cPFr.Email, utils.EmailRegex) {
		return false, fmt.Errorf("invalid email")
	}
	if !utils.Matches(cPFr.Phone, utils.PhoneRegex) {
		return false, fmt.Errorf("invalid phone")
	}
	if !utils.Matches(cPFr.Cpf, utils.CPFRegex) {
		return false, fmt.Errorf("invalid cpf")
	}
	if cPFr.BirthDate.Valid {
		dateStr := fmt.Sprintf("%04d-%02d-%02d", cPFr.BirthDate.Time.Year(), cPFr.BirthDate.Time.Month(), cPFr.BirthDate.Time.Day())
		if !utils.Matches(dateStr, utils.BirthDateRegex) {
			return false, fmt.Errorf("invalid birth date format")
		}
	}
	if !utils.Matches(cPFr.Cep, utils.CEPRegex) {
		return false, fmt.Errorf("invalid cep")
	}

	return true, nil
}

type CustomerPJRequest struct {
	Type        string      `json:"type"`
	Email       string      `json:"email"`
	Phone       string      `json:"phone"`
	Cnpj        string      `json:"cnpj"`
	CompanyName string      `json:"company_name"`
	AddressType string      `json:"address_type"`
	Street      string      `json:"street"`
	Number      string      `json:"number"`
	Complement  pgtype.Text `json:"complement"`
	State       string      `json:"state"`
	City        string      `json:"city"`
	Cep         string      `json:"cep"`
}

func (cPJr *CustomerPJRequest) IsValid() (bool, error) {
	if !(utils.MinChars(cPJr.CompanyName, 5) && utils.MaxChars(cPJr.CompanyName, 100)) {
		return false, fmt.Errorf("company name must have between 5 and 100 characters")
	}
	if !utils.Matches(cPJr.Email, utils.EmailRegex) {
		return false, fmt.Errorf("invalid email")
	}
	if !utils.Matches(cPJr.Phone, utils.PhoneRegex) {
		return false, fmt.Errorf("invalid phone")
	}
	if !utils.Matches(cPJr.Cep, utils.CEPRegex) {
		return false, fmt.Errorf("invalid cep")
	}
	if !utils.Matches(cPJr.Cnpj, utils.CNPJRegex) {
		return false, fmt.Errorf("invalid cnpj")
	}

	return true, nil
}

type AddAddressRequest struct {
	CustomerID  uuid.UUID   `json:"customer_id"`
	AddressType string      `json:"address_type"`
	Street      string      `json:"street"`
	Number      string      `json:"number"`
	Complement  pgtype.Text `json:"complement"`
	State       string      `json:"state"`
	City        string      `json:"city"`
	Cep         string      `json:"cep"`
}
