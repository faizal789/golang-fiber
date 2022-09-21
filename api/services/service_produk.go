package services

import (
	"fmt"
	"gobe/api/model"
	"gobe/api/utils"
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


func (s services)BanyakKategoriX(page model.Halaman) (data int64)  {
	var where string
	if page.Param != nil && page.Param["nama"] != "" && page.Param["deptId"] != ""{
		where = "and name ilike '%' ||'" + page.Param["nama"] + "'||'%' and department_id = '" + page.Param["deptId"] + "'"
	}
	err := s.db.Get(&data, `SELECT count(category_id) FROM category WHERE 1 = 1`+where+ ``)
	if err != nil {
		fmt.Print(err.Error())
		return 0
	}
	return data
}


func (s services) PanggilkategoriX(page model.Halaman)[]model.Kategori {
	var kategoriSlice []model.Kategori
	var where string
	if page.Param != nil && page.Param["nama"] != "" && page.Param["deptId"] != ""{
		where = "and name ilike '%' ||'" + page.Param["nama"] + "'||'%' and department_id = '" + page.Param["deptId"] + "'"
	}
		err := s.db.Select(&kategoriSlice,
		`SELECT category_id as KatId, department_id as DeptId,
		name as Name, description as Desc
		FROM public.category WHERE 1 = 1`+where+` limit $1 offset $2;`,page.Size, utils.GetEnd(page.Pagenumber, page.Size))
		if err != nil {
			fmt.Print(err.Error())
		}
		return kategoriSlice
}

func (s services) PanggilDeptdbx() []model.Department{
	var deptSlice []model.Department

	err := s.db.Select(&deptSlice, `SELECT department_id as DeptId,
	coalesce("name",'') as Name, coalesce(description,'-') as Ket from department`)
	if err != nil {
		fmt.Print(err.Error())
	}
	return deptSlice
}

func (s services) InsertDeptdbx(dept model.DepartmentRequest) model.DepartmentRequest{
	tx := s.db.MustBegin()
	tx.MustExec(`INSERT INTO public.department(
		name, description)
		VALUES ($1,$2);`,dept.Param["name"],dept.Param["ket"])
	tx.Commit()
	return dept
	}
