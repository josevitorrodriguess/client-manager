package customer

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/josevitorrodriguess/client-manager/internal/db/sqlc"
	"github.com/josevitorrodriguess/client-manager/internal/utils"
	"github.com/josevitorrodriguess/client-manager/internal/validators"
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
	validationErrs := validators.ValidationErrors{
		Errors: make(map[string]string),
	}

	if !(utils.MinChars(cPFr.Name, 5) && utils.MaxChars(cPFr.Name, 100)) {
		validationErrs.Errors["Name"] = "name must have between 5 and 100 characters"
	}

	if !utils.Matches(cPFr.Email, utils.EmailRegex) {
		validationErrs.Errors["Email"] = "invalid email"
	}

	if !utils.Matches(cPFr.Phone, utils.PhoneRegex) {
		validationErrs.Errors["Phone"] = "invalid phone"
	}

	if !utils.Matches(cPFr.Cpf, utils.CPFRegex) {
		validationErrs.Errors["Cpf"] = "invalid cpf"
	}

	if cPFr.BirthDate.Valid {
		dateStr := fmt.Sprintf("%04d-%02d-%02d", cPFr.BirthDate.Time.Year(), cPFr.BirthDate.Time.Month(), cPFr.BirthDate.Time.Day())
		if !utils.Matches(dateStr, utils.BirthDateRegex) {
			validationErrs.Errors["BirthDate"] = "invalid birth date format"
		}
	}

	if !utils.Matches(cPFr.Cep, utils.CEPRegex) {
		validationErrs.Errors["Cep"] = "invalid cep"
	}

	if !validationErrs.HasErrors() {
		return true, nil
	}

	return false, validationErrs
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
	validationErrs := validators.ValidationErrors{
		Errors: make(map[string]string),
	}

	if !(utils.MinChars(cPJr.CompanyName, 5) && utils.MaxChars(cPJr.CompanyName, 100)) {
		validationErrs.Errors["CompanyName"] = "company name must have between 5 and 100 characters"
	}

	if !utils.Matches(cPJr.Email, utils.EmailRegex) {
		validationErrs.Errors["Email"] = "invalid email"
	}

	if !utils.Matches(cPJr.Phone, utils.PhoneRegex) {
		validationErrs.Errors["Phone"] = "invalid phone"
	}

	if !utils.Matches(cPJr.Cep, utils.CEPRegex) {
		validationErrs.Errors["Cep"] = "invalid cep"
	}

	if !utils.Matches(cPJr.Cnpj, utils.CNPJRegex) {
		validationErrs.Errors["Cnpj"] = "invalid cnpj"
	}

	if !validationErrs.HasErrors() {
		return true, nil
	}

	return false, validationErrs
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

type CustomerResponse struct {
	ID          uuid.UUID          `json:"id"`
	Type        sqlc.CustomerType  `json:"type"`
	Email       string             `json:"email"`
	Phone       string             `json:"phone"`
	IsActive    bool               `json:"is_active"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
	Cpf         interface{}        `json:"cpf,omitempty"`
	PfName      interface{}        `json:"pf_name,omitempty"`
	BirthDate   interface{}        `json:"birth_date,omitempty"`
	Cnpj        interface{}        `json:"cnpj,omitempty"`
	CompanyName interface{}        `json:"company_name,omitempty"`
	Addresses   []AddressResponse  `json:"addresses"`
}

type AddressResponse struct {
	ID          int32       `json:"id"`
	AddressType string      `json:"address_type"`
	Street      string      `json:"street"`
	Number      string      `json:"number"`
	Complement  pgtype.Text `json:"complement"`
	State       string      `json:"state"`
	City        string      `json:"city"`
	Cep         string      `json:"cep"`
}
