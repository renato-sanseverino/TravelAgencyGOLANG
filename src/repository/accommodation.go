package repository

import (
	"context"
	"travelagency/src/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)


type AccommodationRepository struct {
    pool *pgxpool.Pool
}

func NewAccommodationRepository(connPool *pgxpool.Pool) IRepository[domain.Accommodation] {
	return &AccommodationRepository{
		pool: connPool,
	}
}

func (ar *AccommodationRepository) Insert(ctx context.Context, a domain.Accommodation) error {

	_, err := ar.pool.Exec(ctx, "INSERT INTO accommodations (id, hotel, guests, checkin, checkout, room) VALUES ($1, $2, $3, $4, $5, $6)",
		a.ID,
		a.Hotel,
		a.Guests,
		a.Checkin,
		a.Checkout,
		a.Room,
	)

	if err != nil {
		return err
	}
	return nil // TODO: retornar o registro inserido, mudar na interface tambem
}

func (ar *AccommodationRepository) GetByID(ctx context.Context, id int) (*domain.Accommodation, error) {
	// TODO: implementar usando pgx

	obj := &domain.Accommodation{}
	return obj, nil
}

func (ar *AccommodationRepository) Delete(ctx context.Context, id int) error {
	_, err := ar.pool.Exec(ctx, "DELETE FROM accommodations WHERE id = $1", id)

	if err != nil {
		return err
	}
	return nil
}

func (ar *AccommodationRepository) Patch(ctx context.Context, id int, a domain.Accommodation) error {
	// TODO: implementar usando pgx
	return nil
}
