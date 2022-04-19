package handlers

import (
	"net/http"
	"todoList/platform"

	"github.com/gin-gonic/gin"
)

func Get_one_todo (itemList *platform.ItemList) gin.HandlerFunc{
	return func (c *gin.Context) {
		id := c.Param("id")
		item, err := itemList.GetOneItem(id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No such item"})
			return
		}
		c.JSON(http.StatusOK, item)
	}
}