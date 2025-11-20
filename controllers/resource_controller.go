package controllers

import (
	"net/http"
	"shikkhok_shohayok/db/sqlc"

	"github.com/gin-gonic/gin"
)

type ResourceController struct {
	store *db.Queries
}

func NewResourceController(store *db.Queries) *ResourceController {
	return &ResourceController{store: store}
}

type createResourceRequest struct {
	Title    string `json:"title" binding:"required"`
	FilePath string `json:"file_path" binding:"required"`
	Type     string `json:"type" binding:"required"`
}

func (ctrl *ResourceController) CreateResource(ctx *gin.Context) {
	var req createResourceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateResourceParams{
		Title:    req.Title,
		FilePath: req.FilePath,
		Type:     req.Type,
	}

	resource, err := ctrl.store.CreateResource(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resource)
}

func (ctrl *ResourceController) ListResources(ctx *gin.Context) {
	resources, err := ctrl.store.ListResources(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resources)
}
