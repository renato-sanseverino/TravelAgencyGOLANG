package repository

import (
	"context"
	"travelagency/src/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)


type InsuranceRepository struct {
    pool *pgxpool.Pool
}

func NewInsuranceRepository(connPool *pgxpool.Pool) IRepository[domain.Insurance] {
	return &InsuranceRepository{
		pool: connPool,
	}
}

func (ir *InsuranceRepository) Insert(ctx context.Context, i domain.Insurance) error {

	_, err := ir.pool.Exec(ctx, "INSERT INTO travelinsurance (client_id, \"purposeOfTrip\", luggage, medical_cover, price_total) VALUES ($1, $2, $3, $4, $5)",
		i.ClientID,
		i.PurposeOfTrip,
		i.Luggage,
		i.MedicalCover,
		i.PriceTotal,
	)

	if err != nil {
		return err
	}
	return nil
}

func (ir *InsuranceRepository) GetByID(ctx context.Context, id int) (*domain.Insurance, error) {
	// TODO: implementar usando pgx	
	obj := &domain.Insurance{}
	return obj, nil
}

func (ir *InsuranceRepository) Delete(ctx context.Context, id int) error {
	_, err := ir.pool.Exec(ctx, "DELETE FROM travelinsurance WHERE id = $1", id)

	if err != nil {
		return err
	}
	return nil
}

func (ir *InsuranceRepository) Patch(ctx context.Context, id int, i domain.Insurance) error {
	// TODO: implementar usando pgx
	return nil
}
