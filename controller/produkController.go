package controller

import (
	"encoding/base64"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/odhiahmad/apiuser/dto"
	"github.com/odhiahmad/apiuser/helper"
	"github.com/odhiahmad/apiuser/service"
)

type ProdukController interface {
	CreateProduk(ctx *gin.Context)
	UpdateProduk(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type produkController struct {
	produkService service.ProdukService
	jwtService    service.JWTService
}

func NewProdukController(produkService service.ProdukService, jwtService service.JWTService) ProdukController {
	return &produkController{
		produkService: produkService,
		jwtService:    jwtService,
	}
}

func (c *produkController) CreateProduk(ctx *gin.Context) {
	var produkCreateDTO dto.ProdukCreateDTO
	errDTO := ctx.ShouldBind(&produkCreateDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.produkService.IsDuplicateProdukName(produkCreateDTO.Nama) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate Produk Name", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {

		dec, err := base64.StdEncoding.DecodeString(produkCreateDTO.Foto)
		if err != nil {
			panic(err)
		}

		f, err := os.Create("fotoProduk/" + produkCreateDTO.Nama + ".jpeg")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if _, err := f.Write(dec); err != nil {
			panic(err)
		}
		if err := f.Sync(); err != nil {
			panic(err)
		}
		produkCreateDTO.Foto = produkCreateDTO.Nama + ".jpeg"

		createdProduk := c.produkService.CreateProduk(produkCreateDTO)
		response := helper.BuildResponse(true, "!OK", createdProduk)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *produkController) UpdateProduk(ctx *gin.Context) {
	var produkUpdateDTO dto.ProdukUpdateDTO
	errDTO := ctx.ShouldBind(&produkUpdateDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.produkService.IsDuplicateProdukName(produkUpdateDTO.Nama) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate produkname", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		os.Remove("fotoProduk/" + produkUpdateDTO.Nama + ".jpeg")
		dec, err := base64.StdEncoding.DecodeString(produkUpdateDTO.Foto)
		if err != nil {
			panic(err)
		}

		f, err := os.Create("fotoProduk/" + produkUpdateDTO.Nama + ".jpeg")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if _, err := f.Write(dec); err != nil {
			panic(err)
		}
		if err := f.Sync(); err != nil {
			panic(err)
		}
		produkUpdateDTO.Foto = produkUpdateDTO.Nama + ".jpeg"

		updatedProduk := c.produkService.UpdateProduk(produkUpdateDTO)
		response := helper.BuildResponse(true, "!OK", updatedProduk)
		ctx.JSON(http.StatusCreated, response)

	}

}
func (c *produkController) FindAll(ctx *gin.Context) {
	var pagination dto.Pagination

	errDTO := ctx.ShouldBind(&pagination)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	// paginationSet := helper.GeneratePaginationFromRequest(ctx)
	getProduk := c.produkService.FindAll(pagination)
	response := helper.BuildResponsePagination(true, "!OK", getProduk, pagination)
	ctx.JSON(http.StatusOK, response)

}

func (c *produkController) Delete(ctx *gin.Context) {
	var produkDelete dto.ProdukDeleteDTO
	errDTO := ctx.ShouldBind(&produkDelete)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	log.Print("tes ", produkDelete.Id)
	deleteProduk := c.produkService.Delete(produkDelete.Id)
	response := helper.BuildResponse(true, "Data berhasil dihapus", deleteProduk)
	ctx.JSON(http.StatusOK, response)

}
