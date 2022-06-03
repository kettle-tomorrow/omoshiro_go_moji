package controllers

import (
	"net/http"
	"omoshiroGoMoji/backend/models"
	"omoshiroGoMoji/backend/services"

	"github.com/gin-gonic/gin"
)

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
		"data":   omoshiroGoMoji,
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

func OmoshiroGoMojiShow(c *gin.Context) {
	omoshiroGomojiService := services.OmoshiroGoMojiService{}
	omoshiroGoMoji := omoshiroGomojiService.GetOmoshiroGoMoji(c.Param("id"))
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    omoshiroGoMoji,
	})
}

// func OmoshiroGoMojiUpdate(c *gin.Context) {
// 	omoshiroGoMoji := models.OmoshiroGoMoji{}
// 	err := c.Bind(&omoshiroGoMoji)
// 	if err != nil {
// 		c.String(http.StatusBadRequest, "Bad request")
// 		return
// 	}
// 	omoshiroGoMojiService := services.OmoshiroGoMojiService{}
// 	omoshiroGoMojiService.UpdateOmoshiroGoMoji(&omoshiroGoMoji)
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
