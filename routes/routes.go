package routes

import (
	"database/sql"
	"shikkhok_shohayok/controllers"
	db "shikkhok_shohayok/db/sqlc"
	"time"

	"shikkhok_shohayok/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(conn *sql.DB, config util.Config) *gin.Engine {
	router := gin.Default()

	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	store := db.New(conn)
	authController := controllers.NewAuthController(store, config)
	classController := controllers.NewClassController(store)
	studentController := controllers.NewStudentController(store)
	attendanceController := controllers.NewAttendanceController(store)
	routineController := controllers.NewRoutineController(store)
	examController := controllers.NewExamController(store)
	resourceController := controllers.NewResourceController(store)
	dashboardController := controllers.NewDashboardController(store)

	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)

	router.GET("/dashboard/stats", dashboardController.GetDashboardStats)

	router.POST("/classes", classController.CreateClass)
	router.GET("/classes", classController.ListClasses)
	router.GET("/classes/:id", classController.GetClass)

	router.POST("/students", studentController.CreateStudent)
	router.GET("/students", studentController.ListStudents)
	router.GET("/students/:id", studentController.GetStudent)
	router.PUT("/students/:id", studentController.UpdateStudent)

	router.POST("/attendance", attendanceController.CreateAttendance)
	router.GET("/attendance", attendanceController.GetAttendanceByDate)

	router.POST("/routines", routineController.CreateRoutine)
	router.GET("/routines", routineController.ListRoutines)

	router.POST("/lesson-plans", routineController.CreateLessonPlan)
	router.GET("/lesson-plans", routineController.ListLessonPlans)

	router.POST("/exams", examController.CreateExam)
	router.GET("/exams", examController.ListExams)
	router.POST("/marks", examController.SubmitMarks)
	router.GET("/results/:exam_id", examController.GetResults)

	router.POST("/resources", resourceController.CreateResource)
	router.GET("/resources", resourceController.ListResources)

	return router
}
