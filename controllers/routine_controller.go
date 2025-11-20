package controllers

import (
	"database/sql"
	"net/http"
	"shikkhok_shohayok/db/sqlc"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type RoutineController struct {
	store *db.Queries
}

func NewRoutineController(store *db.Queries) *RoutineController {
	return &RoutineController{store: store}
}

type createRoutineRequest struct {
	ClassID     int64  `json:"class_id" binding:"required"`
	DayOfWeek   string `json:"day_of_week" binding:"required"`
	StartTime   string `json:"start_time" binding:"required"` // HH:MM
	EndTime     string `json:"end_time" binding:"required"`   // HH:MM
	Subject     string `json:"subject" binding:"required"`
	TeacherName string `json:"teacher_name" binding:"required"`
}

func (ctrl *RoutineController) CreateRoutine(ctx *gin.Context) {
	var req createRoutineRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse times. Assuming input is "15:04"
	startTime, err := time.Parse("15:04", req.StartTime)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_time format (HH:MM)"})
		return
	}
	endTime, err := time.Parse("15:04", req.EndTime)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_time format (HH:MM)"})
		return
	}

	arg := db.CreateRoutineParams{
		ClassID:     req.ClassID,
		DayOfWeek:   req.DayOfWeek,
		StartTime:   startTime,
		EndTime:     endTime,
		Subject:     req.Subject,
		TeacherName: req.TeacherName,
	}

	routine, err := ctrl.store.CreateRoutine(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, routine)
}

func (ctrl *RoutineController) ListRoutines(ctx *gin.Context) {
	classIDStr := ctx.Query("class_id")
	classID, err := strconv.ParseInt(classIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class_id"})
		return
	}

	routines, err := ctrl.store.ListRoutinesByClass(ctx, classID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, routines)
}

type createLessonPlanRequest struct {
	ClassID int64  `json:"class_id" binding:"required"`
	Date    string `json:"date" binding:"required"` // YYYY-MM-DD
	Topic   string `json:"topic" binding:"required"`
	Note    string `json:"note"`
}

func (ctrl *RoutineController) CreateLessonPlan(ctx *gin.Context) {
	var req createLessonPlanRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format (YYYY-MM-DD)"})
		return
	}

	arg := db.CreateLessonPlanParams{
		ClassID: req.ClassID,
		Date:    date,
		Topic:   req.Topic,
		Note:    sql.NullString{String: req.Note, Valid: req.Note != ""},
	}

	plan, err := ctrl.store.CreateLessonPlan(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, plan)
}

func (ctrl *RoutineController) ListLessonPlans(ctx *gin.Context) {
	classIDStr := ctx.Query("class_id")
	classID, err := strconv.ParseInt(classIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class_id"})
		return
	}

	plans, err := ctrl.store.ListLessonPlans(ctx, classID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, plans)
}
