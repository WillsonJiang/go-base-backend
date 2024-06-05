package api

import (
	"errors"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/xuri/excelize/v2"
)

type Res struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Timestamp  int64       `json:"timestamp"`
	DateTime   time.Time   `json:"datetime"`
	ReturnData interface{} `json:"return_data,omitempty"`
}

type Paginate struct {
	TotalPage       int `json:"total_page"`
	TotalCount      int `json:"total_count"`
	CurrentPage     int `json:"current_page"`
	PageSize        int `json:"page_size"`
	CurrentPageSize int `json:"current_page_size"`
	Data            any `json:"data"`
}

func ApiResponseOK(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(http.StatusOK, Res{
		StatusCode: statusCode,
		Message:    GetCodeMsg(statusCode),
		Timestamp:  time.Now().UTC().UnixNano(),
		DateTime:   time.Now(),
		ReturnData: data,
	})
}

func FileResponseOK(ctx *gin.Context, fileName string, file *excelize.File) {
	ctx.Writer.Header().Add("Content-Type", "application/octet-stream")
	ctx.Writer.Header().Add("Content-disposition", "attachment;filename="+fileName)
	ctx.Writer.Header().Add("Content-Transfer-Encoding", "binary")
	_ = file.Write(ctx.Writer)
}

func ApiResponseAbort(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, Res{
		StatusCode: statusCode,
		Message:    GetCodeMsg(statusCode),
		Timestamp:  time.Now().UTC().UnixNano(),
		DateTime:   time.Now(),
		ReturnData: data,
	})
}

func ApiAbortWithHttpCode(ctx *gin.Context, httpCode int, statusCode int, data interface{}) {
	ctx.AbortWithStatusJSON(httpCode, Res{
		StatusCode: statusCode,
		Message:    GetCodeMsg(statusCode),
		Timestamp:  time.Now().UTC().UnixNano(),
		DateTime:   time.Now(),
		ReturnData: data,
	})
}

func ValidateRequest(formData any) []string {
	validate := validator.New()
	err := validate.Struct(formData)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]string, len(ve))
			for i, fe := range ve {
				out[i] = strings.ToLower(fe.Field())
			}
			return out
		}
	}
	return []string{}
}

func GetPaginateData(totalCount, pageSize, currentPage, pageCount int, data any) *Paginate {
	var totalPage = 0
	if totalCount != 0 && pageSize != 0 && currentPage != 0 {
		totalPage = int(math.Ceil(float64(totalCount) / float64(pageSize)))
	}
	return &Paginate{
		TotalPage:       totalPage,
		TotalCount:      totalCount,
		PageSize:        pageSize,
		CurrentPage:     currentPage,
		CurrentPageSize: pageCount,
		Data:            data,
	}
}
