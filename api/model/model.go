package model

type Kategori struct {
	KatId  int  `json:"katId"`
	DeptId int	`json:"deptId"`
	Name   string `json:"name"`
	Desc   string `json:"ket"`
}

type Produk struct {
	Kategori     Kategori `json:"kategori"`
	ProdukId     string `json:"produkId"`
	Nama         string `json:"nama"`
	KatId        any `json:"katId"`
	NameKategori any `json:"nameKategori"`
}
