package attendance

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qnify/api-server/utils/db"
	"github.com/redis/go-redis/v9"
	"github.com/zerodha/logf"
)

type module struct {
	redis redis.Cmdable
	db    *db.DbX
	log   *logf.Logger
}

func Routes(app *fiber.App, redis redis.Cmdable, db *db.DbX, logger *logf.Logger) {
	controller := module{
		redis: redis,
		db:    db,
		log:   logger,
	}

	// Get total count of absent, late etc. for a student
	app.Get("/attendance/total/:studentId", controller.getStudentTotal)

	// Get total count of absent, late etc. for all the student of class
	app.Get("/attendance/total/:studentId", controller.getSectionSummary)

}
