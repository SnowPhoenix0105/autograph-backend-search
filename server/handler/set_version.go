package handler

import (
	"autograph-backend-search/logging"
	"autograph-backend-search/server/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var currentVersion uint = 29

func SetVersion(ctx *gin.Context) {
	v := ctx.Query("v")
	version, err := strconv.Atoi(v)
	if err != nil {
		logging.Default().Errorf("atoi(%#v) fail: %s", version, err.Error())
		ctx.JSON(http.StatusBadRequest, common.MakeUnknownErrorResp())
	}

	currentVersion = uint(version)
}

func GetVersion(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, common.MakeSuccessResp(currentVersion))
}
