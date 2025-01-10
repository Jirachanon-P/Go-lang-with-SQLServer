package Routes

import (
	"first-api/Controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	note_service := r.Group("/note-api")
	{
		note_service.GET("note", Controllers.GetAllNotes)
		note_service.POST("note", Controllers.CreateNote)
		note_service.GET("note/:id", Controllers.GetNoteById)
		note_service.PUT("note/:id", Controllers.UpdateNote)
		note_service.DELETE("note/:id", Controllers.DeleteNote)
	}
	return r
}
