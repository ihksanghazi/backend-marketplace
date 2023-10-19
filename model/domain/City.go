package domain

type City struct {
	Id         string `gorm:"primaryKey"`
	ProvinceId string `gorm:"foreignKey"`
	Type       string
	CityName   string
	PostalCode string
	// Association
	User  User  `gorm:"foreignKey:CityId"`
	Store Store `gorm:"foreignKey:CityId"`
}
