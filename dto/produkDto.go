package dto

type ProdukUpdateDTO struct {
	Id        uint64 `json:"id" form:"id" binding:"required"`
	Nama      string `json:"nama" form:"nama" binding:"required"`
	Jenis     string `json:"jenis" form:"jenis" binding:"required"`
	Deskripsi string `json:"deskripsi" form:"deskripsi" binding:"required"`
	Foto      string `json:"foto" form:"foto" binding:"required"`
	Status    uint8  `json:"status,string,omitempty" form:"status,omitempty"`
}

type ProdukCreateDTO struct {
	Nama      string `json:"nama" form:"nama" binding:"required"`
	Jenis     string `json:"jenis" form:"jenis" binding:"required"`
	Deskripsi string `json:"deskripsi" form:"deskripsi" binding:"required"`
	Foto      string `json:"foto" form:"foto" binding:"required"`
	Status    uint8  `json:"status,string,omitempty" form:"status,omitempty"`
}

type ProdukDeleteDTO struct {
	Id uint64 `json:"id" form:"id" binding:"required"`
}
