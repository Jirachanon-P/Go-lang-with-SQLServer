package Routes

import (
	"first-api/Controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		note_service := v1.Group("/note")
		{
			note_service.GET("", Controllers.GetAllNotes)
			note_service.POST("", Controllers.CreateNote)
			note_service.GET(":id", Controllers.GetNoteById)
			note_service.PUT(":id", Controllers.UpdateNote)
			note_service.DELETE(":id", Controllers.DeleteNote)
		}

		user_service := v1.Group("/user", authorizationMiddleware)
		{
			user_service.GET("", Controllers.GetAllUsers)
			user_service.POST("", Controllers.CreateUser)
			user_service.GET(":id", Controllers.GetUserById)
			user_service.PUT(":id", Controllers.UpdateUser)
			user_service.DELETE(":id", Controllers.DeleteUser)
		}

		auth_service := v1.Group("/auth")
		{
			auth_service.POST("login", Controllers.Login)
		}
	}
	return r
}
