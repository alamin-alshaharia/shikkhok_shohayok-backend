package controllers

import (
	"database/sql"
	"net/http"
	"shikkhok_shohayok/db/sqlc"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClassController struct {
	store *db.Queries
}

func NewClassController(store *db.Queries) *ClassController {
	return &ClassController{store: store}
}

type createClassRequest struct {
	Name      string `json:"name" binding:"required"`
	Section   string `json:"section" binding:"required"`
	TeacherID int64  `json:"teacher_id" binding:"required"` // In real app, get from token
}

func (ctrl *ClassController) CreateClass(ctx *gin.Context) {
	var req createClassRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateClassParams{
		Name:      req.Name,
		Section:   req.Section,
		TeacherID: req.TeacherID,
	}

	class, err := ctrl.store.CreateClass(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, class)
}

func (ctrl *ClassController) ListClasses(ctx *gin.Context) {
	teacherIDStr := ctx.Query("teacher_id")
	teacherID, err := strconv.ParseInt(teacherIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher_id"})
		return
	}

	classes, err := ctrl.store.ListClasses(ctx, teacherID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, classes)
}

func (ctrl *ClassController) GetClass(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	class, err := ctrl.store.GetClass(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, class)
}
