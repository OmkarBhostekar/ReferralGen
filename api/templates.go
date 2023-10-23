package api

import (
	"fmt"
	"net/http"
	"strings"

	db "example.com/referralgen/db/sqlc"
	"example.com/referralgen/token"
	"github.com/gin-gonic/gin"
)

type CreateTemplateRequest struct {
	Name     string   `json:"name" binding:"required"`
	Template string   `json:"template" binding:"required"`
	Params   []string `json:"params" binding:"required"`
}

func (server *Server) CreateTemplate(ctx *gin.Context) {
	var req CreateTemplateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Printf(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateTemplateParams{
		UserID:   authPayload.UserID,
		Name:     req.Name,
		Template: req.Template,
		Params:   req.Params,
	}
	template, err := server.store.CreateTemplate(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, template)
}

func (server *Server) GetTemplatesByUser(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	templates, err := server.store.GetTemplatesByUser(ctx, authPayload.UserID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, templates)
}

type GetTemplateByNameRequest struct {
	Name string `form:"name" binding:"required"`
}

func (server *Server) GetTemplateByName(ctx *gin.Context) {
	var req GetTemplateByNameRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.GetTemplatesByNameParams{
		Name:   strings.ToLower(req.Name) + "%",
		UserID: authPayload.UserID,
	}
	templates, err := server.store.GetTemplatesByName(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, templates)
}

type DeleteTemplateByIdRequest struct {
	TemplateID int64 `uri:"id" binding:"required"`
}

type DeleteTemplateByIdResponse struct {
	Result string `json:"result" binding:"required"`
}

func (server *Server) DeleteTemplateById(ctx *gin.Context) {
	var req DeleteTemplateByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	template, err := server.store.GetTemplateById(ctx, req.TemplateID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	if template.UserID != authPayload.UserID {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("User can only delete own templates")))
		return
	}
	template, err = server.store.DeleteTemplateById(ctx, req.TemplateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := DeleteTemplateByIdResponse{
		Result: "Template deleted successfully!",
	}
	ctx.JSON(http.StatusOK, rsp)
}

type UpdateTemplateRequest struct {
	ID       int64    `json:"id" binding:"required"`
	Name     string   `json:"name" binding:"required"`
	Template string   `json:"template" binding:"required"`
	Params   []string `json:"params" binding:"required"`
}

func (server *Server) UpdateTemplate(ctx *gin.Context) {
	var req UpdateTemplateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	template, err := server.store.GetTemplateById(ctx, req.ID)
	if template.UserID != authPayload.UserID {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	arg := db.UpdateTemplateByIdParams{
		ID:       req.ID,
		Name:     req.Name,
		Template: req.Template,
		Params:   req.Params,
	}
	template, err = server.store.UpdateTemplateById(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, template)
}
