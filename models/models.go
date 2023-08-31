package models
type Customer struct {
	ID      uint    `json:"id"`
	Name    string  `gorm:"type:varchar(255)" json:"name"`
	Address Address `gorm:"embedded" json:"address"`
}

	type Address struct {
		City    string `gorm:"type:varchar(255)" json:"city"`
		State   string `gorm:"type:varchar(255)" json:"state"`
		Zip     int    `gorm:"type:int" json:"zip"`
		Country string `gorm:"type:varchar(255)" json:"country"`
	}

//for authentication
type Details struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

