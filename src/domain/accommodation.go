package entity

import (
	"time"
	"travelagency/src/utils"
)

type Accommodation struct {
	ID       int
	Hotel    string
	Guests   int
	Checkin  time.Time
	Checkout *time.Time
	Room     *int
	Charges  []int
}

func NewAccommodation(hotel string, guests int, checkin time.Time, checkout time.Time, room int) (*Accommodation, error){
	obj := &Accommodation{
		ID:     9999,   // gerar uuid
		Guests: guests,
		Checkin: time.Now(),
		Checkout: &checkout,
		Room: &room,
	}
	err := obj.Validate()
	if err != nil {
		return nil, utils.ValidationError
	}
	return obj, nil
}

func (obj *Accommodation) AddCharge(id int) error {
	charge, _ := obj.GetCharge(id)
	if (charge != -1) { // verifica se já foi adicionado
		return utils.DuplicationError
	}
	obj.Charges = append(obj.Charges, id)
	return nil
}

func (obj *Accommodation) RemoveCharge(id int) error {
	for i, j := range obj.Charges {
		if j == id {
			obj.Charges = append(obj.Charges[:i], obj.Charges[i+1:]...)
			return nil
		}
	}
	return utils.NotFoundError
}

func (obj *Accommodation) GetCharge(id int) (int, error) {
	// não reinventar a roda, verificar  slices.IndexFunc()
	for _, v := range obj.Charges {
		if v == id {
			return id, nil
		}
	}
	return -1, utils.NotFoundError
}

func (obj *Accommodation) Validate() error {
	if obj.Hotel == "" || obj.Guests < 1 || obj.Checkin.IsZero() {
		return utils.ValidationError
	}
	return nil
}
