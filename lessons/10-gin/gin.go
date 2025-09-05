package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var certificates = []Certificate{
	{Id: 36, Name: "Abigail", Course: "CED", Grade: "A", Date: "20-1-2044"},
	{Id: 39, Name: "Leopold", Course: "CBA", Grade: "S", Date: "20-1-2044"},
	{Id: 42, Name: "Zachary", Course: "CHF", Grade: "B", Date: "20-1-2044"},
}

func create(c *gin.Context) {
	var newCertificate Certificate
	if err := c.BindJSON(&newCertificate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	certificates = append(certificates, newCertificate)
	c.IndentedJSON(http.StatusCreated, newCertificate)
}

func readAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, certificates)
}

func main() {
	router := gin.Default()
	router.POST("/create", create)
	router.GET("/read", readAll)
	router.Run("localhost:8080")
}
