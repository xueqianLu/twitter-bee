package openapi

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/xueqianLu/twitter-bee/config"
	"github.com/xueqianLu/twitter-bee/models/apimodels"
	"github.com/xueqianLu/twitter-bee/types"
	"net/http"
)

type apiHandler struct {
	conf    *config.Config
	backend types.ServiceBackend
}

// @Summary Get Follower Count
// @Description Get the count of followers for a specific account
// @Tags Twitter
// @Accept json
// @Produce json
// @Param account body apimodels.FollowerCountRequest true "Account"
// @Success 200 {object} apimodels.FollowerCountResponse
// @Failure 500 {object} apimodels.BaseResponse
func (api apiHandler) FollowerCount(c *gin.Context) {
	req := apimodels.FollowerCountRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.WithError(err).Error("FollowerCount ctx.ShouldBindJSON error")
		api.errResponse(c, err)
		return
	}
	if data, err := api.backend.FollowerCount(req); err != nil {
		api.errResponse(c, err)
	} else {
		res := apimodels.BaseResponse{
			Code:    http.StatusOK,
			Message: "success",
			Data:    data,
		}
		api.response(c, http.StatusOK, res)
	}
}

// @Summary Get Follower List
// @Description Get the list of followers for a specific account
// @Tags Twitter
// @Accept json
// @Produce json
// @Param account body apimodels.FollowerListRequest true "Account"
// @Success 200 {object} apimodels.FollowerListResponse
// @Failure 500 {object} apimodels.BaseResponse
func (api apiHandler) FollowerList(c *gin.Context) {
	req := apimodels.FollowerListRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.WithError(err).Error("FollowerList ctx.ShouldBindJSON error")
		api.errResponse(c, err)
		return
	}
	if data, err := api.backend.FollowerList(req); err != nil {
		api.errResponse(c, err)
	} else {
		res := apimodels.BaseResponse{
			Code:    http.StatusOK,
			Message: "success",
			Data:    data,
		}
		api.response(c, http.StatusOK, res)
	}
}

func (api apiHandler) response(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}

func (api apiHandler) errResponse(c *gin.Context, err error) {
	res := apimodels.BaseResponse{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	}
	api.response(c, http.StatusInternalServerError, res)
}
