package web

import (
	"tanahore/internal/model/domain"
)

type CreateArticle struct {
	ArticleID uint             `gorm:"type:int;autoIncrement;primaryKey" json:"articleID"`
	SoilType  domain.SoilTypes `gorm:"type:enum('Aluvial','Andosol','Latosol','Humus');not null" json:"soilType" validate:"required"`
	Title     string           `gorm:"type:varchar(50);not null" json:"title" validae:"required"`
	Content   string           `gorm:"type:text;not null" json:"content" validate:"required"`
}
