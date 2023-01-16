package models

type Wallet struct {
	Model
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(100)"`
	PubKey      string `json:"pub_key" gorm:"type:varchar(100)"`
	PrivKey     string `json:"priv_key" gorm:"type:varchar(100)"`
	Mnemonic    string `json:"mnemonic" gorm:"type:varchar(250)"`
}
