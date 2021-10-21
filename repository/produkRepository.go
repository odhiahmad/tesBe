package repository

import (
	"github.com/odhiahmad/apiuser/dto"
	"github.com/odhiahmad/apiuser/entity"
	"gorm.io/gorm"
)

type ProdukRepository interface {
	InsertProduk(produk entity.Produk) entity.Produk
	UpdateProduk(produk entity.Produk) entity.Produk
	IsDuplicateProdukname(nama string) (tx *gorm.DB)
	FindAll(pagination dto.Pagination) []entity.Produk
	Delete(id uint64) entity.Produk
}

var count int64

type produkConnection struct {
	connection *gorm.DB
}

func NewProdukRepository(db *gorm.DB) ProdukRepository {
	return &produkConnection{
		connection: db,
	}
}

func (db *produkConnection) InsertProduk(produk entity.Produk) entity.Produk {
	db.connection.Save(&produk)

	return produk
}

func (db *produkConnection) UpdateProduk(produk entity.Produk) entity.Produk {

	var tempProduk entity.Produk
	db.connection.Find(&tempProduk, produk.Id)

	db.connection.Save(&produk)

	return produk
}

func (db *produkConnection) IsDuplicateProdukname(nama string) (tx *gorm.DB) {
	var produk entity.Produk
	return db.connection.Where("nama = ?", nama).Take(&produk)
}

func (db *produkConnection) FindByProdukName(nama string) entity.Produk {
	var produk entity.Produk
	db.connection.Where("nama = ?", nama).Take(&produk)

	return produk
}

func (db *produkConnection) Delete(id uint64) entity.Produk {
	var produk entity.Produk
	db.connection.Where("id = ?", id).Delete(&produk)

	return produk
}

func (db *produkConnection) FindAll(pagination dto.Pagination) []entity.Produk {
	var produk []entity.Produk

	offset := (pagination.Page - 1) * pagination.Limit
	db.connection.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&produk).Count(&count)
	return produk
}
