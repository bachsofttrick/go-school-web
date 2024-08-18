package api

import (
	"fmt"
	"net/http"
	"os"
	"schoolweb/model"
	"schoolweb/mytime"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

// In-memory web storage
var webs = map[int]model.SchoolWeb{}
var fileName = "web.csv"

func AttachApi(r *gin.Engine) {
	loadWebs(webs)

	// Routes
	r.GET("/webs", getWebs)
	r.GET("/webs/:id", getWeb)
	r.POST("/webs", createWeb)
	r.DELETE("/webs/:id", deleteWeb)
}

func loadWebs(webs map[int]model.SchoolWeb) {
	// Read from csv
	webFile, err := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer webFile.Close()

	// Load webs from file
	websLoad := []model.SchoolWeb{}
	if err := gocsv.UnmarshalFile(webFile, &websLoad); err != nil {
		panic(err)
	}
	for _, web := range websLoad {
		fmt.Println(web)
		webs[web.ID] = web
	}

	// Save to csv
	// webFile, err = os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	// if err != nil {
	// 	panic(err)
	// }
	// gocsv.MarshalFile(websLoad, webFile)
}

func getWebs(c *gin.Context) {
	websList := make([]model.SchoolWeb, 0, len(webs))
	for _, web := range webs {
		websList = append(websList, web)
	}

	c.JSON(http.StatusOK, websList)
}

func getWeb(c *gin.Context) {
	id := c.Param("id")
	var idN int
	fmt.Sscan(id, &idN)
	web, exists := webs[idN]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "Web not found"})
		return
	}

	c.JSON(http.StatusOK, web)
}

func createWeb(c *gin.Context) {
	var web model.SchoolWeb
	if err := c.ShouldBindJSON(&web); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	if _, exists := webs[web.ID]; exists {
		c.JSON(http.StatusConflict, gin.H{"message": "Web already exists"})
		return
	}

	web.CreatedDate = mytime.GetUnixTime()
	webs[web.ID] = web
	c.JSON(http.StatusCreated, web)
}

func deleteWeb(c *gin.Context) {
	id := c.Param("id")
	var idN int
	fmt.Sscan(id, &idN)
	if _, exists := webs[idN]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "Web not found"})
		return
	}

	delete(webs, idN)
	c.JSON(http.StatusOK, gin.H{"message": "Web deleted"})
}
