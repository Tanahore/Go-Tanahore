package web

import "tanahore/internal/model/domain"

type AllArticlesResponse struct {
	ArticleID uint             `gorm:"type:int;autoIncrement;primaryKey" json:"articleID"`
	SoilType  domain.SoilTypes `gorm:"type:enum('Aluvial','Andosol','Latosol','Humus');not null" json:"soilType"`
	Title     string           `gorm:"type:varchar(50);not null" json:"title"`
	Content   string           `gorm:"type:text;not null" json:"content"`
	ImageURL  string           `json:"imageURL"`
}

type GetAllArticlesResponse struct {
	ArticleID uint   `json:"articleID"`
	Title     string `json:"title"`
	ImageURL  string `json:"imageURL"`
}
