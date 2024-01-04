package domain

import (
	"travelagency/src/utils"

	"github.com/shopspring/decimal"
)

// Decimal points to github.com/shopspring/decimal.Decimal, as Go does not have a native decimal type
type Decimal = decimal.Decimal

type Insurance struct {
	ID            int      // Anotations para o GIN FRAMEWORK
	ClientID      int      `json:"client_id" binding:"required"`
	PurposeOfTrip string   `json:"purposeOfTrip" binding:"required"` // campo está errado no banco
	Luggage       *Decimal `json:"luggage"`
	MedicalCover  *Decimal `json:"medical_cover"`
	PriceTotal    Decimal  `json:"price_total" binding:"required"`
}

func NewInsurance(clientID int, purposeOfTrip string, luggage Decimal, medicalCover Decimal, priceTotal Decimal) (*Insurance, error) {
	obj := &Insurance{
		ID:            9999, // gerar uuid
		ClientID:      clientID,
		PurposeOfTrip: purposeOfTrip,
		Luggage:       &luggage,
		MedicalCover:  &medicalCover,
		PriceTotal:    priceTotal,
	}

	err := obj.Validate()
	if err != nil {
		return nil, utils.ValidationError
	}
	return obj, nil
}

func (obj *Insurance) Validate() error {
	if obj.ClientID == 0 || obj.PurposeOfTrip == "" || obj.PriceTotal.IsZero() {
		return utils.ValidationError
	}
	return nil
}
