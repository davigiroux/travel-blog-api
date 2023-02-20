package api

import (
	"net/http"

	db "github.com/davigiroux/travel-blog-api/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createArticleRequest struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	PosterID int64  `json:"posterId" binding:"required"`
}

func (s *Server) createArticle(ctx *gin.Context) {
	var req createArticleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var arg = db.CreateArticleParams{Content: req.Content, Title: req.Title, PosterID: req.PosterID}

	article, err := s.store.CreateArticle(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, article)
}
