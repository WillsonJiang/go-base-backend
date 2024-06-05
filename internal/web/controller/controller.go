package controller

import (
	"backend/internal/web/helper/api"
	"time"

	"github.com/gin-gonic/gin"
)

var serverStarTime = time.Now()

func ServerCheck(ctx *gin.Context) {
	api.ApiResponseOK(ctx, api.SELECT_SUCCESS, struct {
		Test            string    `json:"test"`
		ServerStartTime time.Time `json:"server_start_time"`
	}{
		ServerStartTime: serverStarTime,
		Test:            "welcome to backend",
	})
}
