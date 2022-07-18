package main

import (
	"net/http"
	"gitub.com/gin-gonic/gin"
	"errors"
)

type car struct{
	ID    string  `json:"id:`
	Model string	  `json:"title"`
	Make string  `json:"brand"`
	OnHand int	  `json:"onhand"`
}

var cars = []car{
	{ID: "1", Model: "TL", Make: "Acura", OnHand: 1},
	{ID: "2", Model: "Accord", Make: "Honda", OnHand: 5},
	{ID: "3", Model: "A3", Make: "Audi", OnHand: 2},
	{ID: "4", Model: "Raptor", Make: "Ford", OnHand: 2},
	{ID: "5", Model: "Stinger", Make: "KIA", OnHand: 3},
}