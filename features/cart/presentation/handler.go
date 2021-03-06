package presentation

import (
	"errors"
	"net/http"
	"strconv"

	"construct-week1/features/cart"
	"construct-week1/features/cart/presentation/request"
	"construct-week1/features/cart/presentation/response"
	"construct-week1/features/helper"
	"construct-week1/middlewares"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartBusiness cart.Business
}

func NewCartHandler(business cart.Business) *CartHandler {
	return &CartHandler{
		cartBusiness: business,
	}
}

func (handle *CartHandler) AddCart(c echo.Context) error {
	//	Merupakan validasi user id yaitu menggunakan token
	idUser := middlewares.ExtractToken(c)
	if idUser == 0 {
		return errors.New("failed")
	}

	var newCart request.Cart
	newCart.UserId = uint(idUser)
	newCart.Status = 1

	errBind := c.Bind(&newCart)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Failed to bind your data cart, check your input"))
	}

	dataCart := request.ToCore(newCart)
	err := handle.cartBusiness.InsertCart(dataCart)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to insert your cart"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("Success insert your cart"))
}

func (handle *CartHandler) GetCart(c echo.Context) error {
	//	Merupakan validasi user id yaitu menggunakan token
	idUser := middlewares.ExtractToken(c)
	if idUser == 0 {
		return errors.New("failed")
	}

	res, err := handle.cartBusiness.FindCart(uint(idUser))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get your data product cart"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success to get your data product cart", response.FromCoreList(res)))
}

func (handle *CartHandler) UpdateCart(c echo.Context) error {
	//	Merupakan validasi user id yaitu menggunakan token
	idUser := middlewares.ExtractToken(c)
	if idUser == 0 {
		return errors.New("failed")
	}

	id := c.Param("id")
	idCart, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("ID not recognized"))
	}

	var updateCart request.CartUpdate
	errBind := c.Bind(&updateCart)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Failed to bind your data quantity, check your input"))
	}

	dataUpdateCart := request.ToCoreUpdate(updateCart)
	err := handle.cartBusiness.UpdateCart(uint(idCart), uint(idUser), dataUpdateCart)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to update your quantity"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("Success update your quantity"))
}
