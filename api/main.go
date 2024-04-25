package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// pet represents data about a pet.
type pet struct {
	ID         int    `json:"id"`
	UserName   string `json:"userName"`
	PetName    string `json:"petName"`
	PetImage   string `json:"petImage"`
	IsFriendly bool   `json:"isFriendly"`
}

// petsResponse is the structure used for the JSON response
type petsResponse struct {
	Data []pet `json:"data"`
}

// pets slice to seed pet data.
var pets = []pet{
	{ID: 1, UserName: "Leanne Graham", PetName: "Bret", PetImage: "https://cdn.shibe.online/shibes/04be82add971f2b8490b5ec2308d3426e8494ad9.jpg", IsFriendly: true},
	{ID: 2, UserName: "Ervin Howell", PetName: "Anto", PetImage: "https://cdn.shibe.online/shibes/655be60d14f2a0dfb3060dae20c4aad0194b3393.jpg", IsFriendly: false},
	{ID: 3, UserName: "Clementine Bauch", PetName: "Samaurai", PetImage: "https://cdn.shibe.online/shibes/2d2e282fe2af42373600b819f80b869e1e7b18e9.jpg", IsFriendly: false},
}

func main() {
	router := gin.Default()
	router.GET("/pets", getPets)
	router.Run("localhost:8088")
}

// getPets responds with the list of all pets as JSON.
func getPets(c *gin.Context) {
	response := petsResponse{
		Data: pets,
	}

	c.IndentedJSON(http.StatusOK, response)
}
