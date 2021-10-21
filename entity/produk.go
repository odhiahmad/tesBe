package entity

import "time"

type Produk struct {
	Id        uint64 `gorm:"primary_key:auto_increament" json:"id"`
	Nama      string `gorm:"uniqueIndex;type:varchar(255)" json:"Nama"`
	Jenis     string `gorm:"->;<-;not null" json:"Jenis"`
	Deskripsi string `gorm:"->;<-;not null" json:"Deskripsi"`
	Foto      string `gorm:"->;<-;not null" json:"Foto"`
	Status    uint8  `gorm:"not null" json:"Status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
