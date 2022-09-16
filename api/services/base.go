package services

import (
	"gobe/api/model"
	"github.com/jmoiron/sqlx"
)


type Service interface{
	PanggilKategoriX() []model.Kategori
	PanggilProdukdbx() []model.Produk

}


type services struct{
	db *sqlx.DB
}


func InitService(db *sqlx.DB) Service {
	return &services{db}
}