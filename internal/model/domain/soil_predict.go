package domain

type SoilTypes string
type Volcanics int

const (
	Aluvial SoilTypes = "Aluvial"
	Latosol SoilTypes = "Latosol"
	Humus   SoilTypes = "Humus"
	Andosol SoilTypes = "Andosol"
)

const (
	VolcanicTrue  Volcanics = 1
	VolcanicFalse Volcanics = 0
)

type Plants struct {
	PlantID   uint      `gorm:"type:int;primarykey" json:"plantID"`
	PlantName string    `gorm:"type:varchar(20);not null" json:"plantName" validate:"required"`
	SoilType  SoilTypes `gorm:"type:varchar(20);not null" json:"soilType" validate:"required"`
	Volcanic  Volcanics `gorm:"type:int; not null" json:"volcanic" validate:"required"`
}
