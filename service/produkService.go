package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/odhiahmad/apiuser/dto"
	"github.com/odhiahmad/apiuser/entity"
	"github.com/odhiahmad/apiuser/repository"
)

type ProdukService interface {
	CreateProduk(produk dto.ProdukCreateDTO) entity.Produk
	UpdateProduk(produk dto.ProdukUpdateDTO) entity.Produk
	FindAll(pagination dto.Pagination) []entity.Produk
	Delete(id uint64) entity.Produk
	IsDuplicateProdukName(nama string) bool
}

type produkService struct {
	produkRepository repository.ProdukRepository
}

func NewProdukService(produkRepo repository.ProdukRepository) ProdukService {
	return &produkService{
		produkRepository: produkRepo,
	}
}

func (service *produkService) CreateProduk(produk dto.ProdukCreateDTO) entity.Produk {
	produkToCreate := entity.Produk{}
	err := smapping.FillStruct(&produkToCreate, smapping.MapFields(&produk))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.produkRepository.InsertProduk((produkToCreate))
	return res
}

func (service *produkService) UpdateProduk(produk dto.ProdukUpdateDTO) entity.Produk {
	produkToUpdate := entity.Produk{}
	err := smapping.FillStruct(&produkToUpdate, smapping.MapFields(&produk))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.produkRepository.UpdateProduk((produkToUpdate))
	return res
}

func (service *produkService) Delete(id uint64) entity.Produk {

	res := service.produkRepository.Delete(id)
	return res
}

func (service *produkService) IsDuplicateProdukName(nama string) bool {
	res := service.produkRepository.IsDuplicateProdukname(nama)
	return !(res.Error == nil)
}

func (service *produkService) FindAll(pagination dto.Pagination) []entity.Produk {
	return service.produkRepository.FindAll(pagination)
}
