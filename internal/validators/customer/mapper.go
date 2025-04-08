package customer

import "github.com/josevitorrodriguess/client-manager/internal/db/sqlc"

func MapAddresses(rows []sqlc.GetCustomerAddressesRow) []AddressResponse {
	addrs := make([]AddressResponse, 0, len(rows))
	for _, a := range rows {
		addrs = append(addrs, AddressResponse{
			ID:          a.ID,
			AddressType: a.AddressType,
			Street:      a.Street,
			Number:      a.Number,
			Complement:  a.Complement,
			State:       a.State,
			City:        a.City,
			Cep:         a.Cep,
		})
	}
	return addrs
}

func MapCustomer(data sqlc.GetCustomerByIDRow, addresses []sqlc.GetCustomerAddressesRow) CustomerResponse {
	return CustomerResponse{
		ID:          data.ID,
		Type:        data.Type,
		Email:       data.Email,
		Phone:       data.Phone,
		IsActive:    data.IsActive,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		Cpf:         data.Cpf,
		PfName:      data.PfName,
		BirthDate:   data.BirthDate,
		Cnpj:        data.Cnpj,
		CompanyName: data.CompanyName,
		Addresses:   MapAddresses(addresses),
	}
}
