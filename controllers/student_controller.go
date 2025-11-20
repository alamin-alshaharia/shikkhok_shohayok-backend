package controllers

import (
	"database/sql"
	"net/http"
	"shikkhok_shohayok/db/sqlc"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	store *db.Queries
}

func NewStudentController(store *db.Queries) *StudentController {
	return &StudentController{store: store}
}

type createStudentRequest struct {
	Name       string `json:"name" binding:"required"`
	Roll       string `json:"roll" binding:"required"`
	FatherName string `json:"father_name" binding:"required"`
	MobileNo   string `json:"mobile_no" binding:"required"`
	ClassID    int64  `json:"class_id" binding:"required"`
}

func (ctrl *StudentController) CreateStudent(ctx *gin.Context) {
	var req createStudentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateStudentParams{
		Name:       req.Name,
		Roll:       req.Roll,
		FatherName: req.FatherName,
		MobileNo:   req.MobileNo,
		ClassID:    req.ClassID,
	}

	student, err := ctrl.store.CreateStudent(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func (ctrl *StudentController) ListStudents(ctx *gin.Context) {
	classIDStr := ctx.Query("class_id")
	classID, err := strconv.ParseInt(classIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class_id"})
		return
	}

	students, err := ctrl.store.ListStudents(ctx, classID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, students)
}

func (ctrl *StudentController) GetStudent(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	student, err := ctrl.store.GetStudent(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

type updateStudentRequest struct {
	Name       string `json:"name" binding:"required"`
	Roll       string `json:"roll" binding:"required"`
	FatherName string `json:"father_name" binding:"required"`
	MobileNo   string `json:"mobile_no" binding:"required"`
}

func (ctrl *StudentController) UpdateStudent(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req updateStudentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.UpdateStudentParams{
		ID:         id,
		Name:       req.Name,
		Roll:       req.Roll,
		FatherName: req.FatherName,
		MobileNo:   req.MobileNo,
	}

	student, err := ctrl.store.UpdateStudent(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, student)
}
