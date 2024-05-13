package course

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qnify/api-server/utils/db"
	"github.com/redis/go-redis/v9"
	"github.com/zerodha/logf"
)

type courseModule struct {
	redis redis.Cmdable
	db    *db.DbX
	log   *logf.Logger
}

func Routes(app *fiber.App, redis redis.Cmdable, db *db.DbX, logger *logf.Logger) {
	controller := courseModule{
		redis: redis,
		db:    db,
		log:   logger,
	}

	/*
		Get one course
		@api GET /public/courses/{id}
		@description Get details about a course
		@operationId getCourse
		@param id path integer true "Course ID"
		@response responses/Ok
		@tags courses
	*/
	app.Get("/public/courses/:id", controller.getCourse)

	/*
		Get all courses
		@api GET /public/courses
		@description List all courses on the platform
		@operationId getAllCourses
		@response responses/AuthResponse
		@tags courses
	*/
	app.Get("/public/courses", controller.getAllCourses)

	/*
		Create a course
		@api POST /admin/courses
		@description Create a new course on the platform
		@operationId createCourse
		@request schemas/User
		@response responses/AuthResponse
		@tags courses
	*/
	app.Post("/admin/courses", controller.createCourse)

	/*
		Update a course
		@api PUT /admin/courses/{id}
		@description Update details about an existing course on the platform
		@operationId updateCourse
		@param id path integer true "Course ID"
		@response responses/Ok
		@tags courses
	*/
	app.Put("/admin/courses/:id", controller.updateCourse)

	/*
		Delete a course
		@api DELETE /admin/courses/{id}
		@description Delete details about an existing course on the platform
		@operationId deleteCourse
		@param id path integer true "Course ID"
		@response responses/Ok
		@tags courses
	*/
	app.Delete("/admin/courses/:id", controller.deleteCourse)
}
