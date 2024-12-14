package models

import (
	"time"
)

type CounterpartyModel struct {
	ID                          int64     `json:"id" gorm:"primaryKey"`
	FirstName                   *string   `json:"first_name"`
	MiddleName                  *string   `json:"middle_name"`
	LastName                    *string   `json:"last_name"`
	BirthDate                   *time.Time `json:"birth_date"`
	PhoneNumber                 *string   `json:"phone_number"`
	PassportCode                *string   `json:"passport_code"`
	Email                       *string   `json:"email"`
	PassportIssuingDate         *time.Time `json:"passport_issuing_date"`
	IdentificationCode          *string   `json:"identification_code"`
	PassportIssuedBy            *string   `json:"passport_issued_by"`
	PassportIssuingDatePresence bool      `json:"passport_issuing_date_presence"`
	ExternalID                  *string   `json:"external_id"`
	CounterpartyType            string    `json:"counterparty_type"`
	Street                      *string   `json:"street"`
	Region                      *string   `json:"region"`
	Locality                    *string   `json:"locality"`
	District                    *string   `json:"district"`
	HouseNumber                 *string   `json:"house_number"`
	Postcode                    *string   `json:"postcode"`
	CreatedAt                   time.Time `json:"created_at"`
	UpdatedAt                   time.Time `json:"updated_at"`
	LegalName                   *string   `json:"legal_name"`
	LegalAddress                *string   `json:"legal_address"`
	Index                       *string   `json:"index"`
	CompanyNumber               *string   `json:"company_number"`
	Signatory                   *string   `json:"signatory"`
	ContactPerson               *string   `json:"contact_person"`
	VATPayer                    bool      `json:"vat_payer"`
	CounterpartyGroupID         *int64    `json:"counterparty_group_id"`
	UsageScope                  []string  `json:"usage_scope" gorm:"-"`
}
