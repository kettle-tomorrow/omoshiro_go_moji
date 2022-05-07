package controllers

import (
	"net/http"
	"omoshiroGoMoji/backend/models"
	"omoshiroGoMoji/backend/services"

	"github.com/gin-gonic/gin"
)

// "backend/model"
// "backend/service"

func OmoshiroGoMojiCreate(c *gin.Context) {
	omoshiroGoMoji := models.OmoshiroGoMoji{}
	err := c.Bind(&omoshiroGoMoji)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	omoshiroGoMojiService := services.OmoshiroGoMojiService{}
	omoshiroGoMojiService.CreateOmoshiroGoMojiService(&omoshiroGoMoji)
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func OmoshiroGoMojiList(c *gin.Context) {
	omoshiroGoMojiService := services.OmoshiroGoMojiService{}
	omoshiroGoMojiLists := omoshiroGoMojiService.GetOmoshiroGoMojiList()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    omoshiroGoMojiLists,
	})
}

// func OmoshiroGoMojiUpdate(c *gin.Context) {
// 	omoshiroGoMoji := model.OmoshiroGoMoji{}
// 	err := c.Bind(&omoshiroGoMoji)
// 	if err != nil {
// 		c.String(http.StatusBadRequest, "Bad request")
// 		return
// 	}
// 	omoshiroGoMojiService := service.OmoshiroGoMojiService{}
// 	err = omoshiroGoMojiService.UpdateOmoshiroGoMoji(&omoshiroGoMoji)
// 	if err != nil {
// 		c.String(http.StatusInternalServerError, "Server Error")
// 		return
// 	}
// 	c.JSON(http.StatusCreated, gin.H{
// 		"status": "ok",
// 	})
// }
// func OmoshiroGoMojiDelete(c *gin.Context) {
// 	id := c.PostForm("id")
// 	intId, err := strconv.ParseInt(id, 10, 0)
// 	if err != nil {
// 		c.String(http.StatusBadRequest, "Bad request")
// 		return
// 	}
// 	omoshiroGoMojiService := service.OmoshiroGoMojiService{}
// 	err = omoshiroGoMojiService.DeleteOmoshiroGoMoji(int(intId))
// 	if err != nil {
// 		c.String(http.StatusInternalServerError, "Server Error")
// 		return
// 	}
// 	c.JSON(http.StatusCreated, gin.H{
// 		"status": "ok",
// 	})
