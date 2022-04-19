package handlers

import (
	"net/http"
	"todoList/platform"

	"github.com/gin-gonic/gin"
)

func Put_todo (itemList *platform.ItemList) gin.HandlerFunc{
	return func (c *gin.Context){
		id := c.Param("id")

		item, err := itemList.PutItem(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err})
			return
		}		
		c.JSON(http.StatusOK, item)
	}
}