package main

import (
	"github.com/qnify/api-server/modules/attendance"
	"github.com/qnify/api-server/modules/auth"
	"github.com/qnify/api-server/modules/course"
	"github.com/qnify/api-server/utils/config"
	"github.com/qnify/api-server/utils/db"
	"github.com/qnify/api-server/utils/fiber"
	"github.com/qnify/api-server/utils/helper"
)

func main() {

	config := config.LoadConfig("config.yaml")

	dbSql := db.InitDB(config.Db)
	dbx := db.NewDBX(config.Db.Type, dbSql)
	redis := helper.GetRedis(config.RedisURL)
	logger := helper.InitLogger()

	fiberApp := fiber.InitFiber(logger)
	auth.Routes(fiberApp, redis, dbSql, logger, config.Auth)
	course.Routes(fiberApp, redis, dbx, logger)
	attendance.Routes(fiberApp, redis, dbx, logger)

	fiber.StartWithGracefulShutdown(fiberApp, config.Port)
}
