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
			Message: utils.RECORD_NOT_FOUND,
			Error:   err.Error(),
		}
	}

	return &commondto.ErrorResponse{
		Code:    utils.StatusCode[utils.InternalServer],
		Message: utils.INTERNAL_SERVER_ERROR,
		Error:   err.Error(),
	}
}
