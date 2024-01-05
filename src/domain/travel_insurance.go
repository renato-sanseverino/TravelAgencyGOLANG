package domain

import (
	"travelagency/src/utils"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Decimal points to github.com/shopspring/decimal.Decimal, as Go does not have a native decimal type
type Decimal = decimal.Decimal

type Insurance struct {
	ID            uuid.UUID  // Anotations para o GIN FRAMEWORK
	ClientID      uuid.UUID  `json:"client_id" binding:"required"`
	PurposeOfTrip string     `json:"purpose_of_trip" binding:"required"` // campo está errado no banco, fora do padrão dos outros campos
	Luggage       *Decimal   `json:"luggage"`
	MedicalCover  *Decimal   `json:"medical_cover"`
	PriceTotal    Decimal    `json:"price_total" binding:"required"`
}

func NewInsurance(clientID uuid.UUID, purposeOfTrip string, luggage Decimal, medicalCover Decimal, priceTotal Decimal) (*Insurance, error) {
	obj := &Insurance{
		ID:            uuid.New(),  // evitar a duplicação de registros com UUID
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
	if obj.PurposeOfTrip == "" || obj.PriceTotal.IsZero() {
		return utils.ValidationError
	}
	return nil
}
