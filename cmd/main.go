package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/phanletrunghieu/notifier/endpoints"

	"github.com/go-kit/kit/log"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

	"github.com/phanletrunghieu/notifier/database/pg"
	serviceHttp "github.com/phanletrunghieu/notifier/http"
	"github.com/phanletrunghieu/notifier/service"
	deviceSvc "github.com/phanletrunghieu/notifier/service/device"
	userSvc "github.com/phanletrunghieu/notifier/service/user"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("failed to load .env by errors: %v", err))
	}

	// setup addr
	httpAddr := ":" + os.Getenv("PORT")

	// setup log
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// setup locale
	{
		loc, err := time.LoadLocation("Asia/Bangkok")
		if err != nil {
			logger.Log("error", err)
			os.Exit(1)
		}
		time.Local = loc
	}

	pgDB, closeDB := pg.New(os.Getenv("PG_DATASOURCE"))
	defer closeDB()

	err = pg.MigrateTables(pgDB)
	if err != nil {
		panic(fmt.Sprintf("failed to migrate tables by errors: %v", err))
	}

	s := service.Service{
		UserService:   userSvc.NewPGService(pgDB),
		DeviceService: deviceSvc.NewPGService(pgDB),
	}

	var h http.Handler
	{
		h = serviceHttp.NewHTTPHandler(logger, endpoints.NewEndpoints(s))
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", httpAddr)
		errs <- http.ListenAndServe(httpAddr, h)
	}()

	logger.Log("exit", <-errs)
}
