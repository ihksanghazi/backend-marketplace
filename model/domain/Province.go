package domain

type Province struct {
	Id       string `gorm:"primaryKey"`
	Province string
	Cities   []City `gorm:"foreignKey:ProvinceId"`
}
