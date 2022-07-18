package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "errors"
)

type car struct{
	ID    string  `json:"id"`
	Model string	  `json:"model"`
	Make string  `json:"make"`
	OnHand int	  `json:"onhand"`
}

var cars = []car{
	{ID: "1", Model: "TL", Make: "Acura", OnHand: 1},
	{ID: "2", Model: "Accord", Make: "Honda", OnHand: 5},
	{ID: "3", Model: "A3", Make: "Audi", OnHand: 2},
	{ID: "4", Model: "Raptor", Make: "Ford", OnHand: 2},
	{ID: "5", Model: "Stinger", Make: "KIA", OnHand: 3},
}

func getCars(c *gin.Context){
	c.IndentedJSON(http.StatusOK, cars)

}

func createCar(c *gin.Context) {
	var newCar car
	
	if err := c.BindJSON(&newCar); err != nil {
		return
	}
	
	cars = append(cars, newCar)
	c.IndentedJSON(http.StatusCreated, newCar)
}


func main() {
	router := gin.Default()
	router.GET("/cars", getCars)
	router.POST("/cars", createCar)
	router.Run("localhost:8000")
	

}