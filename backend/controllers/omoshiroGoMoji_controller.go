package controllers

import (
	"net/http"
	"omoshiroGoMoji/backend/models"

	"github.com/gin-gonic/gin"
)

type OmoshiroGoMojiController struct{}

func (OmoshiroGoMojiController) Create(c *gin.Context) {
	omoshiroGoMoji := models.OmoshiroGoMoji{UserID: c.GetUint("currentUserID")}
	err := c.Bind(&omoshiroGoMoji)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	omoshiroGoMoji.CreateOmoshiroGoMoji(&omoshiroGoMoji)
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"data":   omoshiroGoMoji,
	})
}

func (OmoshiroGoMojiController) Index(c *gin.Context) {
	omoshiroGoMoji := models.OmoshiroGoMoji{}
	omoshiroGoMojiList := omoshiroGoMoji.GetOmoshiroGoMojiList()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    omoshiroGoMojiList,
	})
}

func (OmoshiroGoMojiController) Show(c *gin.Context) {
	omoshiroGomoji := models.OmoshiroGoMoji{}
	omoshiroGoMoji := omoshiroGomoji.GetOmoshiroGoMoji(c.Param("id"))
	if omoshiroGoMoji.ID == 0 {
		c.String(http.StatusNotFound, "not found")
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    omoshiroGoMoji,
	})
}

func (OmoshiroGoMojiController) Update(c *gin.Context) {
	omoshiroGoMoji := models.OmoshiroGoMoji{}
	omoshiroGoMoji = omoshiroGoMoji.GetOmoshiroGoMoji(c.Param("id"))
	err := c.Bind(&omoshiroGoMoji)
	if omoshiroGoMoji.UserID != c.GetUint("currentUserID") {
		c.String(http.StatusNotFound, "not found")
		return
	}
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	omoshiroGoMoji.UpdateOmoshiroGoMoji(&omoshiroGoMoji)
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"data":   omoshiroGoMoji,
	})
}

func (OmoshiroGoMojiController) Delete(c *gin.Context) {
	omoshiroGoMoji := models.OmoshiroGoMoji{}
	omoshiroGoMoji = omoshiroGoMoji.GetOmoshiroGoMoji(c.Param("id"))
	if omoshiroGoMoji.UserID != c.GetUint("currentUserID") {
		c.String(http.StatusNotFound, "not found")
		return
	}
	omoshiroGoMoji.DeleteOmoshiroGoMoji(c.Param("id"))
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
