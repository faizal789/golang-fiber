package services

import (
	"gobe/api/model"
	"github.com/jmoiron/sqlx"
)


type Service interface{
	PanggilProdukdbx() []model.Produk
	PanggilkategoriX(page model.Halaman) []model.Kategori
	BanyakKategoriX(page model.Halaman) int64
	PanggilDeptdbx() []model.Department
	InsertDeptdbx(dept model.DepartmentRequest) model.DepartmentRequest

}


type services struct{
	db *sqlx.DB
}


func InitService(db *sqlx.DB) Service {
	return &services{db}
}