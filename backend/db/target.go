package db

import (
	"context"
	"pingo/models"
)

func (db *PingoDB) CreateTarget(ctx context.Context, t *models.Target) (models.Target, error) {
	return *t, nil
}
func (db *PingoDB) ReadTarget(ctx context.Context, ID int) (string, error) {
	return "adam", nil
}
func (db *PingoDB) UpdateTarget(ctx context.Context, ID int, t models.Target) (models.Target, error) {
	return t, nil
}
func (db *PingoDB) DeleteTarget(ctx context.Context, ID int) (string, error) {
	return "adam", nil
}
