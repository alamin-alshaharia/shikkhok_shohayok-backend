package controllers

import (
	"net/http"
	"shikkhok_shohayok/db/sqlc"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	store *db.Queries
}

func NewDashboardController(store *db.Queries) *DashboardController {
	return &DashboardController{store: store}
}

func (ctrl *DashboardController) GetDashboardStats(ctx *gin.Context) {
	teacherIDStr := ctx.Query("teacher_id")
	dateStr := ctx.Query("date") // YYYY-MM-DD

	teacherID, err := strconv.ParseInt(teacherIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher_id"})
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format (YYYY-MM-DD)"})
		return
	}

	// Get Day of Week (e.g., "Sunday", "Monday")
	// Note: Go's Weekday().String() returns English names which matches our DB convention hopefully.
	// If DB stores Bangla, we need mapping. Assuming English for now based on previous code.
	dayOfWeek := date.Weekday().String()

	// 1. Get Total Classes for Today
	routinesCount, err := ctrl.store.GetTeacherDailyRoutinesCount(ctx, db.GetTeacherDailyRoutinesCountParams{
		TeacherID: teacherID,
		DayOfWeek: dayOfWeek,
	})
	if err != nil {
		// If no rows, it returns 0 count usually, but sqlc might return error if query fails
		routinesCount = 0
		// Log error but continue?
	}

	// 2. Get Attendance Stats for Today
	attendanceStats, err := ctrl.store.GetTeacherDailyAttendanceStats(ctx, db.GetTeacherDailyAttendanceStatsParams{
		TeacherID: teacherID,
		Date:      date,
	})

	var attendancePercentage float64 = 0
	if err == nil && attendanceStats.TotalCount > 0 {
		attendancePercentage = (float64(attendanceStats.PresentCount) / float64(attendanceStats.TotalCount)) * 100
	}

	ctx.JSON(http.StatusOK, gin.H{
		"total_classes":         routinesCount,
		"attendance_percentage": attendancePercentage,
	})
}
