package api

import (
	"fmt"
	"net/http"
	"time"

	db "example.com/referralgen/db/sqlc"
	"example.com/referralgen/token"
	"github.com/gin-gonic/gin"
)

type GenerateReferralTemplateRequest struct {
	Prompt string `json:"prompt" binding:"required"`
}

type GenerateReferralTemplateResponse struct {
	Template string `json:"template" binding:"required"`
}

func (server *Server) GenerateReferralTemplate(ctx *gin.Context) {
	// add rate limiting
	var req GenerateReferralTemplateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	// check if user has generated 5 templates today
	arg := db.GetTodayUserGenerationCountParams{
		UserID:      authPayload.UserID,
		CreatedDate: time.Now().Format("2006-01-02"),
	}
	generations, err := server.store.GetTodayUserGenerationCount(ctx, arg)
	if err != nil {
		// check if err is no rows
		if err.Error() == "sql: no rows in result set" {
			arg := db.CreateGenerationParams{
				UserID:      authPayload.UserID,
				CreatedDate: time.Now().Format("2006-01-02"),
			}
			generations, err = server.store.CreateGeneration(ctx, arg)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}
	if generations.Count.Int32 >= 5 {
		ctx.JSON(http.StatusTooManyRequests, errorResponse(fmt.Errorf("You have reached your daily generation limit")))
		return
	}
	// open ai call to generate template
	template := OpenAICompletion(req.Prompt)

	// increase generation count for user
	_, err = server.store.IncreaseGenerationCount(ctx, generations.ID)

	// send template to user
	rsp := GenerateReferralTemplateResponse{
		Template: template,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func OpenAICompletion(prompt string) string {
	// TODO: call open ai api to generate template
	return "Hello, this is sample template for " + prompt
}
