package main

type Repository[T any] interface {
	Add(item T)
	Get(uniqueId string) T
	Update(item T)
	Delete(item T)
}

type GameRegistry struct {
	Name          string `json:"name"`
	GameId        string `json:"GameId"`
	GameSecretKey string `json:"GameSecretKey"`
}

type Authentication struct {
	Email     string `json:"email"`
	SecretKey string `json:"secretKey"`
}
