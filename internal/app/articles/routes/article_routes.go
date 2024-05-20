package routes

import "github.com/labstack/echo/v4"

func (routes *ArticleRoutesImpl) MobileArticleRoutes(apiGroup echo.Group) {
	ArticleGroup := apiGroup.Group("/api/articles")

	ArticleGroup.GET("", routes.Handler.GetAllArticles)
	ArticleGroup.GET("/:id", routes.Handler.GetArticleByID)
	ArticleGroup.GET("/search/:title", routes.Handler.ArticleSearch)
}

func (routes *ArticleRoutesImpl) WebArticleRoutes(apiGroup echo.Group) {
	ArticleGroup := apiGroup.Group("/api/articles")

	ArticleGroup.POST("/create", routes.Handler.CreateArticle)
}
