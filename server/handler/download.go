package handler

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"autograph-backend-search/logging"
	"autograph-backend-search/repository/filesave"
	"autograph-backend-search/rpc"
	"autograph-backend-search/server/common"
	"autograph-backend-search/utils"

	"github.com/gin-gonic/gin"
)

func Download(ctx *gin.Context) {
	handler := downloadHandler{
		ctx: ctx,
	}

	if err := handler.checkParam(); err != nil {
		logging.Default().WithError(err).Errorf("parse req error: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, common.MakeUnknownErrorResp())
		return
	}

	err := handler.produce()
	if err != nil {
		logging.Default().WithError(err).Errorf("produce error: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, common.MakeUnknownErrorResp())
		return
	}
}

type downloadHandler struct {
	ctx *gin.Context

	id uint
}

func (h *downloadHandler) checkParam() error {
	id := h.ctx.Query("id")

	if len(id) == 0 {
		return utils.WrapError(common.ErrRequestParamEmpty, "query 'id' is empty")
	}

	idInteger, err := strconv.Atoi(id)
	if err != nil {
		return utils.WrapErrorf(err, "atoi(%s) fail", id)
	}

	if idInteger < 0 {
		return utils.WrapErrorf(common.ErrRequestParamInvalid, "id(%d) cannot be negative", idInteger)
	}

	h.id = uint(idInteger)

	return nil
}

func (h *downloadHandler) produce() error {
	fileInfo, err := rpc.ControllerGetFileInfo(context.TODO(), h.id)
	if err != nil {
		return utils.WrapErrorf(err, "call ControllerGetFileInfo(%d) fail", h.id)
	}

	fileContent, err := filesave.ReadFile(fileInfo.URL)
	if err != nil {
		return utils.WrapErrorf(err, "call ReadFile(%#v) fail", fileInfo.URL)
	}

	// Content-Type, see
	// https://www.iana.org/assignments/media-types/application/octet-stream
	contentType := "application/octet-stream"

	// Content-Disposition, see:
	// https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Content-Disposition
	filename := fmt.Sprintf("%s.%s", fileInfo.Name, fileInfo.Type)
	filename = url.PathEscape(filename)
	h.ctx.Header("Content-Disposition", "attachment; filename="+filename)

	// Access-Control-Expose-Headers, see:
	// https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Access-Control-Expose-Headers
	h.ctx.Header("Access-Control-Expose-Headers", "Content-Disposition")

	h.ctx.Data(http.StatusOK, contentType, fileContent)

	return nil
}
