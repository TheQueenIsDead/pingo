package db

import (
	"context"
	"pingo/models"
)

func (pdb *PingoDB) CreateTarget(ctx context.Context, t *models.Target) (models.Target, error) {

	err := pdb.db.Insert(t)
	if err != nil {
		return models.Target{}, err
	}

	return *t, nil
}
func (pdb *PingoDB) ReadTarget(ctx context.Context, ID int) (string, error) {
	return "adam", nil
}
func (pdb *PingoDB) UpdateTarget(ctx context.Context, ID int, t models.Target) (models.Target, error) {
	return t, nil
}
func (pdb *PingoDB) DeleteTarget(ctx context.Context, ID int) (string, error) {
	return "adam", nil
}
