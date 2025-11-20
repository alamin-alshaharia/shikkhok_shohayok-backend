package controllers

import (
	"net/http"
	"shikkhok_shohayok/db/sqlc"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AttendanceController struct {
	store *db.Queries
}

func NewAttendanceController(store *db.Queries) *AttendanceController {
	return &AttendanceController{store: store}
}

type attendanceItem struct {
	StudentID int64  `json:"student_id" binding:"required"`
	Status    string `json:"status" binding:"required"` // present, absent, late
}

type createAttendanceRequest struct {
	Date       string           `json:"date" binding:"required"` // YYYY-MM-DD
	Attendances []attendanceItem `json:"attendances" binding:"required"`
}

func (ctrl *AttendanceController) CreateAttendance(ctx *gin.Context) {
	var req createAttendanceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format (YYYY-MM-DD)"})
		return
	}

	// In a real app, use a transaction here
	var results []db.Attendance
	for _, item := range req.Attendances {
		arg := db.UpsertAttendanceParams{
			StudentID: item.StudentID,
			Date:      date,
			Status:    item.Status,
		}
		res, err := ctrl.store.UpsertAttendance(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		results = append(results, res)
	}

	ctx.JSON(http.StatusOK, results)
}

func (ctrl *AttendanceController) GetAttendanceByDate(ctx *gin.Context) {
	classIDStr := ctx.Query("class_id")
	dateStr := ctx.Query("date")

	classID, err := strconv.ParseInt(classIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class_id"})
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format (YYYY-MM-DD)"})
		return
	}

	arg := db.GetAttendanceByDateParams{
		ClassID: classID,
		Date:    date,
	}

	attendance, err := ctrl.store.GetAttendanceByDate(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, attendance)
}
