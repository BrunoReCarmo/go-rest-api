package controller 

import (
	"go-api/model",
	"github.com/gin-gonic/gin",
	"net/http"
)

type offerController struct {
	//useCase
}

//it will return offerController type
func newOfferController() offerController {
	return offerController()
}

func (p *offerController) GetOffers(ctx *gin.Context) {
	//Array in GO => []
	mockedOffers := []model.Offer{
		{
			ID: 1,
			Name: "Camisa malha fina para eventos sociais",
			Price: 49
		}
	}
	//http status from 'net/http'
	ctx.JSON(
		http.StatusOk, mockedOffers
	)
}