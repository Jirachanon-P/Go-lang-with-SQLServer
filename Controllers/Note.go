package Controllers

import (
	"first-api/Models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllNotes godoc
// @summary GET All Notes
// @description Get all notes
// @tags Note
// @id GetAllNotes
// @accept json
// @produce json
// @response 200 {object} Models.Response "OK"
// @response 400 {object} Models.Response "Bad Request"
// @response 401 {object} Models.Response "Unauthorized"
// @response 500 {object} Models.Response "Internal Server Error"
// @Router /note-api/note [get]
func GetAllNotes(c *gin.Context) {
	var notes []Models.Note
	err := Models.GetAllNote(&notes)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		c.JSON(http.StatusNotFound, "No Note Found")
	} else {
		c.JSON(http.StatusOK, notes)
	}
}

// CreateNote godoc
// @summary Create Note
// @description Create Note
// @tags Note
// @id CreateNote
// @accept json
// @produce json
// @param Note body Models.Note true "Note data to be create"
// @response 200 {object} Models.Response "OK"
// @response 400 {object} Models.Response "Bad Request"
// @response 401 {object} Models.Response "Unauthorized"
// @response 500 {object} Models.Response "Internal Server Error"
// @Router /note-api/note [post]
func CreateNote(c *gin.Context) {
	var notes Models.Note
	c.BindJSON(&notes)
	err := Models.CreateNote(&notes)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, notes)
	}
}

// GetNoteById godoc
// @summary Get NoteBy Id
// @description Get NoteBy Id
// @tags Note
// @id GetNoteById
// @accept json
// @produce json
// @param id path string true "id of note"
// @response 200 {object} Models.Response "OK"
// @response 400 {object} Models.Response "Bad Request"
// @response 401 {object} Models.Response "Unauthorized"
// @response 500 {object} Models.Response "Internal Server Error"
// @Router /note-api/note/{id} [get]
func GetNoteById(c *gin.Context) {
	idStr := c.Params.ByName("id")
	var note Models.Note

	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Handle error if conversion fails
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	err = Models.GetNoteById(&note, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
	} else {
		c.JSON(http.StatusOK, note)
	}
}

// UpdateNote godoc
// @summary Update Note
// @description Update Note
// @tags Note
// @id UpdateNote
// @accept json
// @produce json
// @param id path string true "id of note to be update"
// @param Note body Models.Note true "Note data to update"
// @response 200 {object} Models.Response "OK"
// @response 400 {object} Models.Response "Bad Request"
// @response 401 {object} Models.Response "Unauthorized"
// @response 500 {object} Models.Response "Internal Server Error"
// @Router /note-api/note/{id} [put]
func UpdateNote(c *gin.Context) {
	var note Models.Note
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	err = Models.GetNoteById(&note, id)
	if err != nil {
		// Handle error if conversion fails
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	err = Models.UpdateNote(&note, id)
	if err != nil {
		c.JSON(http.StatusNotFound, note)
	}
	c.BindJSON(&note)
	err = Models.UpdateNote(&note, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, note)
	}
}

// DeleteNote godoc
// @summary Delete Note
// @description Delete Note
// @tags Note
// @id DeleteNote
// @accept json
// @produce json
// @param id path string true "id of note to be delete"
// @response 200 {object} Models.Response "OK"
// @response 400 {object} Models.Response "Bad Request"
// @response 401 {object} Models.Response "Unauthorized"
// @response 500 {object} Models.Response "Internal Server Error"
// @Router /note-api/note/{id} [delete]
func DeleteNote(c *gin.Context) {
	var user Models.Note
	idStr := c.Params.ByName("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Handle error if conversion fails
		c.JSON(400, gin.H{"error": err})
		return
	}

	err = Models.DeleteNote(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + idStr: "is deleted"})
	}
}
