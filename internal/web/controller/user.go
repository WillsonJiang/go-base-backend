package controller

import (
	"backend/internal/web/helper/api"
	"backend/internal/web/service"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	data, err := service.GetUsers()
	if err != nil {
		api.ApiResponseAbort(ctx, api.SELECT_FAILED, nil)
		return
	}
	api.ApiResponseOK(ctx, api.SELECT_SUCCESS, data)
}
