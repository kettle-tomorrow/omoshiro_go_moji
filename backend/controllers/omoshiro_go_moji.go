package controller

import (
	// "backend/model"
	// "backend/service"

	"github.com/gin-gonic/gin"
)

func OmoshiroGoMojiAdd(c *gin.Context) {
	// 	omoshiroGoMoji := model.OmoshiroGoMoji{}
	// 	err := c.Bind(&omoshiroGoMoji)
	// 	if err != nil {
	// 		c.String(http.StatusBadRequest, "Bad request")
	// 		return
	// 	}
	// 	omoshiroGoMojiService := service.OmoshiroGoMojiService{}
	// 	err = omoshiroGoMojiService.SetOmoshiroGoMoji(&omoshiroGoMoji)
	// 	if err != nil {
	// 		c.String(http.StatusInternalServerError, "Server Error")
	// 		return
	// 	}
	// 	c.JSON(http.StatusCreated, gin.H{
	// 		"status": "ok",
	// 	})
	// }

	// func OmoshiroGoMojiList(c *gin.Context) {
	// 	omoshiroGoMojiService := service.OmoshiroGoMojiService{}
	// 	omoshiroGoMojiLists := omoshiroGoMojiService.GetOmoshiroGoMojiList()
	// 	c.JSONP(http.StatusOK, gin.H{
	// 		"message": "ok",
	// 		"data":    omoshiroGoMojiLists,
	// 	})
	// }

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
}
