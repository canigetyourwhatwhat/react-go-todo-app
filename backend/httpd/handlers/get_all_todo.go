package handlers

import (
	"net/http"
	"todoList/platform"

	"github.com/gin-gonic/gin"
)

func Get_all_todo(itemList *platform.ItemList) gin.HandlerFunc {
	return func (c *gin.Context){	
		result := itemList.GetAll()
		c.JSON(http.StatusOK, result)
	}
}