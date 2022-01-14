package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest/src/config"
	"rest/src/infrastructure/gateways/rest"
	"rest/src/interfaces/api/docs"
	"rest/src/pkg/cors"
	"rest/src/pkg/logger"
	"rest/src/storage/sqlstorage"
	"syscall"
	"time"
)

// @Title Jasur's Swagger
// @description CRUD
// @contact.name API Support
// @contact.url https://translate.google.com/?sl=en&tl=ru&text=scammer&op=translate
// @contact.email https://translate.google.com/?sl=en&tl=ru&text=scammer&op=translate
// @license.name Scam
// @license.url https://translate.google.com/?sl=en&tl=ru&text=scammer&op=translate
func main() {
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, os.Interrupt,syscall.SIGTERM)

	cfg := config.Load()
	l := logger.New(cfg.LogLevel, "postgres")

	docs.SwaggerInfo.Host = cfg.HTTPHost + cfg.HTTPPort
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	db, err := sqlx.Connect("postgres", cfg.PostgresURL())
	if err != nil {
		panic(err)
	}
	storage := sqlstorage.New(db, l)

	r := gin.New()
	r.Use(cors.CORSMiddleware())
	r.Use(gin.Logger(), gin.Recovery())

	handler := rest.NewAPI(cfg, l, r, storage)

	srv := &http.Server{
		Addr:    cfg.HTTPPort,
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Fatal(fmt.Sprintf("Failed To Start REST Server: %s\n", err.Error()))
		}
	}()
	l.Info("REST Server started at port" + cfg.HTTPPort)


	OSCall := <-quitSignal


	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	l.Info(fmt.Sprintf("\nSystem Call:%+v", OSCall))
	fmt.Printf("system call:%+v", OSCall)

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(fmt.Sprintf("REST Server Graceful Shutdown Failed: %s\n", err))
	}
	cancel()
}