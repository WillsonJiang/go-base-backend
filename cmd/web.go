package cmd

import (
	"backend/internal/database"
	"backend/internal/web/route"
	"backend/pkg/env"
	"backend/pkg/logger"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/coreos/go-systemd/daemon"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "backend",
	Short: "Api Server",
	Long:  `Api Server`,
	Run: func(cmd *cobra.Command, args []string) {
		WebServerMain()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func WebServerMain() {
	env.Init()
	logger.Init(env.Get("LOG_LEVEL", "info"))
	database.Init()

	router := route.Init()
	serviceHTTP := initHTTPListen(router)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logger.Info("Shutdown Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if serviceHTTP != nil {
		if err := serviceHTTP.Shutdown(ctx); err != nil {
			logger.Fatal("HTTP Service Shutdown:", err)
		}
	}
	logger.Info("Server Exist")
}

func initHTTPListen(router *gin.Engine) *http.Server {
	port := env.Get("HTTP_PORT", "8080")
	service := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	go func() {
		if err := service.ListenAndServe(); err != nil {
			logger.Errorf("listen http error, port: %s, err %s", port, err)
		}
		daemon.SdNotify(false, "READY=1")
	}()

	logger.Infof("listen http port: %s", port)

	return service
}
