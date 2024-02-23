package models

type City struct {
	Generic

	CountryID uint    `gorm:"index;not null" json:"country_id"`
	Country   Country `json:"-"`

	Name             string `gorm:"not null" json:"name"`
	ServersAvailable int    `gorm:"<-:false;->;-:migration" json:"servers_available"`
}
