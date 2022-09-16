package services

import (
	"fmt"
	"gobe/api/model"
)



func (s services) PanggilKategoriX() []model.Kategori {
	var kategoriSlice []model.Kategori
		err := s.db.Select(&kategoriSlice,
		`SELECT category_id as KatId, department_id as DeptId,
		name as Name, description as Desc
		FROM public.category;`)
		if err != nil {
			fmt.Print(err.Error())
		}
		return kategoriSlice
}


func (s services) PanggilProdukdbx() []model.Produk {
	var ProdukSlice []model.Produk
	err := s.db.Select(&ProdukSlice,`
	SELECT DISTINCT product.name as Nama,product_category.category_id as KatId,category.name as NameKategori from product
	inner join product_category
	on product_category.product_id = product.product_id
	inner join category
	on category.category_id = product_category.category_id`)
	if err != nil {
		fmt.Print(err.Error())
	}
	return ProdukSlice
}

