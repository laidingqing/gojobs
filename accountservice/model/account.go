package model

//Account defines ...
type Account struct {
	ID       string `bson:"_id" json:"id,omitempty"`
	Name     string `bson:"name" json:"name,omitempty"`
	ServedBy string `bson:"-" json:"servedBy,omitempty"`
}
