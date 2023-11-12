package services

import (
	"errors"

	commondto "github.com/kushanD95/traval-diary/package/dto"
	"github.com/kushanD95/traval-diary/package/utils"
	"gorm.io/gorm"
)

type Context struct {
	_ struct{}
}

func (context *Context) BuildRepoErrRes(err error) *commondto.ErrorResponse {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &commondto.ErrorResponse{
			Code:    utils.StatusCode[utils.NotFound],
			Message: "Record not found",
			Error:   err.Error(),
		}
	}

	return &commondto.ErrorResponse{
		Code:    utils.StatusCode[utils.InternalServer],
		Message: "Internal server Error",
		Error:   err.Error(),
	}
}
