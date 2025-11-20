package controllers

import (
	"net/http"
	"shikkhok_shohayok/db/sqlc"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExamController struct {
	store *db.Queries
}

func NewExamController(store *db.Queries) *ExamController {
	return &ExamController{store: store}
}

type createExamRequest struct {
	ClassID    int64  `json:"class_id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	TotalMarks int32  `json:"total_marks" binding:"required"`
}

func (ctrl *ExamController) CreateExam(ctx *gin.Context) {
	var req createExamRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateExamParams{
		ClassID:    req.ClassID,
		Name:       req.Name,
		TotalMarks: req.TotalMarks,
	}

	exam, err := ctrl.store.CreateExam(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exam)
}

func (ctrl *ExamController) ListExams(ctx *gin.Context) {
	classIDStr := ctx.Query("class_id")
	classID, err := strconv.ParseInt(classIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class_id"})
		return
	}

	exams, err := ctrl.store.ListExams(ctx, classID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exams)
}

type resultItem struct {
	StudentID     int64 `json:"student_id" binding:"required"`
	ObtainedMarks int32 `json:"obtained_marks" binding:"required"`
}

type submitMarksRequest struct {
	ExamID  int64        `json:"exam_id" binding:"required"`
	Results []resultItem `json:"results" binding:"required"`
}

func (ctrl *ExamController) SubmitMarks(ctx *gin.Context) {
	var req submitMarksRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var results []db.Result
	for _, item := range req.Results {
		arg := db.UpsertResultParams{
			ExamID:        req.ExamID,
			StudentID:     item.StudentID,
			ObtainedMarks: item.ObtainedMarks,
		}
		res, err := ctrl.store.UpsertResult(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		results = append(results, res)
	}

	ctx.JSON(http.StatusOK, results)
}

func (ctrl *ExamController) GetResults(ctx *gin.Context) {
	examIDStr := ctx.Param("exam_id")
	examID, err := strconv.ParseInt(examIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam_id"})
		return
	}

	results, err := ctrl.store.GetResultsByExam(ctx, examID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, results)
}
