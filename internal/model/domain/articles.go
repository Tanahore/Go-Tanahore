package domain

import "time"

type Articles struct {
	ArticleID uint      `gorm:"type:int;autoIncrement;primaryKey" json:"articleID"`
	SoilType  SoilTypes `gorm:"type:enum('Aluvial','Andosol','Latosol','Humus');not null" json:"soilType" validate:"required"`
	Title     string    `gorm:"type:varchar(50);not null" json:"title" validate:"required"`
	Content   string    `gorm:"type:text;not null" json:"content" validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
