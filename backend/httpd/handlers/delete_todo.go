package handlers

import (
	"net/http"
	"todoList/platform"

	"github.com/gin-gonic/gin"
)

func Delete_todo (itemList *platform.ItemList) gin.HandlerFunc{
	return func (c *gin.Context) {
		id := c.Param("id")
		result, err := itemList.DeleteItem(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		
		c.JSON(http.StatusOK, result)

	}
}