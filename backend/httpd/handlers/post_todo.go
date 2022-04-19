package handlers

import (
	"todoList/platform"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type itemRequest struct {
	ID string `json:"id"`
	Title string `json:"title"`	
	Completed bool `json:"completed"`
}

func Post_todo (itemList *platform.ItemList) gin.HandlerFunc{
	return func (c *gin.Context) {
		// requestBody := itemRequest{}
		// c.Bind(&requestBody)

		// item := platform.Item{
		// 	ID: requestBody.ID,
		// 	Title: requestBody.Title,			
		// 	Done: requestBody.Done,
		// }

		title := c.Param("title")

		id := uuid.New().String()

		item := platform.Item{
			ID: id,
			Title: title,
			Completed: false,
		}

		itemList.Add(item)

		c.JSON(200, item)
	}
}