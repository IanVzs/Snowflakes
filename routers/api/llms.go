package api

import (
	"github.com/IanVzs/Snowflakes/models"
	"github.com/IanVzs/Snowflakes/pkgs/app"
	"github.com/IanVzs/Snowflakes/pkgs/e"
	"github.com/IanVzs/Snowflakes/services/llms"
	"github.com/gin-gonic/gin"
)

// @QA
// @Produce  json
// @Param name query content true "Content"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/qa [post]
func QA(c *gin.Context) {
	// 构建请求体
	var reqBody models.LLMRequest
	appG := app.Gin{C: c}
	err := appG.C.ShouldBindJSON(&reqBody)

	if err != nil {
		appG.Response(400, e.INVALID_PARAMS, "无效的请求参数")
	}
	rspContent := llms.QAOpenAI(reqBody.Content)
	appG.Response(200, e.SUCCESS, rspContent)
}
