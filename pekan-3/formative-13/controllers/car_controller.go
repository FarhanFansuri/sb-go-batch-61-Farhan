package controllers

import (
	"fmt"
	. "formative-13/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCar(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"data": Cars_data,
	})
}

func CreateCar(ctx *gin.Context) {
	var item Car
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	item.CarID = fmt.Sprintf("c%d", len(Cars_data)+1)
	Cars_data = append(Cars_data, item)
	ctx.JSON(http.StatusCreated, gin.H{
		"data": Cars_data,
	})
}

func UpdateCar(ctx *gin.Context) {
	var item Car
	var CarID = ctx.Param("carId")
	var condition = false
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, d := range Cars_data {
		if CarID == d.CarID {
			condition = true
			Cars_data[i] = item
			Cars_data[i].CarID = CarID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error status":  "Data Not Found",
			"error message": fmt.Sprintf("Data with id %v is not found", CarID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func DeleteCar(ctx *gin.Context) {

	CarID := ctx.Param("carId")
	var condition bool

	for i, d := range Cars_data {
		if CarID == d.CarID {

			Cars_data = append(Cars_data[:i], Cars_data[i+1:]...)
			condition = true
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error status":  "Data Not Found",
			"error message": fmt.Sprintf("Data with id %v is not found", CarID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": fmt.Sprintf("Car with ID %v successfully deleted", CarID),
	})
}
