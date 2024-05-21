package web

import (
	"tanahore/internal/model/domain"
)

type CreateArticle struct {
	ArticleID uint             `gorm:"type:int;autoIncrement;primaryKey" json:"articleID" form:"articleID"`
	SoilType  domain.SoilTypes `gorm:"type:enum('Aluvial','Andosol','Latosol','Humus');not null" json:"soilType" validate:"required" form:"soilType"`
	Title     string           `gorm:"type:varchar(50);not null" json:"title" validae:"required" form:"title"`
	Content   string           `gorm:"type:text;not null" json:"content" validate:"required" form:"content"`
	Image     string           `gorm:"type:varchar(255)" json:"image_url" form:"image"`
}

type GetArticlesByID struct {
	ArticleID uint `json:"articleID" validate:"required"`
}

type GetArticlesByTitle struct {
	Title string `json:"title" validate:"required"`
}
