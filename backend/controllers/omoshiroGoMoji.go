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
	if omoshiroGoMoji.ID == 0 {
		c.String(http.StatusNotFound, "not found")
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    omoshiroGoMoji,
	})
}

func OmoshiroGoMojiUpdate(c *gin.Context) {
	omoshiroGoMojiService := services.OmoshiroGoMojiService{}
	omoshiroGoMoji := omoshiroGoMojiService.GetOmoshiroGoMoji(c.Param("id"))
	err := c.Bind(&omoshiroGoMoji)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	omoshiroGoMojiService.UpdateOmoshiroGoMoji(&omoshiroGoMoji)
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"data":   omoshiroGoMoji,
	})
}

func OmoshiroGoMojiDelete(c *gin.Context) {
	omoshiroGoMojiService := services.OmoshiroGoMojiService{}
	omoshiroGoMojiService.DeleteOmoshiroGoMoji(c.Param("id"))
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
