package main

import "gorm.io/gorm"

type SqliteGameRepository struct {
	db *gorm.DB
}

func NewSqliteGameRegistryRepository(db *gorm.DB) *SqliteGameRepository {
	return &SqliteGameRepository{
		db: db,
	}
}

func (r *SqliteGameRepository) Add(item GameRegistry) error {
	return r.db.Create(&item).Error
}

func (r *SqliteGameRepository) Get(uniqueId string) (GameRegistry, error) {
	var game GameRegistry
	err := r.db.Where(("game_id = ?"), uniqueId).First(&game).Error
	return game, err
}

func (r *SqliteGameRepository) Update(item GameRegistry) error {
	return r.db.Save(&item).Error
}

func (r *SqliteGameRepository) Delete(gameId string) error {
	return r.db.Where("game_id = ?", gameId).Delete(&GameRegistry{}).Error
}

type Repository[T any] interface {
	Add(item T) error
	Get(uniqueId string) (T, error)
	Update(item T) error
	Delete(item T) error
}

type GameRegistry struct {
	GameId        string `json:"gameId"`
	Name          string `json:"name"`
	GameSecretKey string `json:"gameSecretKey"`
}

type Authentication struct {
	Email     string `json:"email"`
	SecretKey string `json:"secretKey"`
}
