package model

import "time"

//Profile defines ...
type Profile struct {
	ID         string    `bson:"_id" json:"id,omitempty"`
	Name       string    `bson:"name" json:"name,omitempty"`
	Avatar     string    `bson:"avatar" json:"avatar,omitempty"`
	Experience int32     `bson:"experience" json:"experience,omitempty"`
	Degree     int32     `bson:"degree" json:"degree,omitempty"`
	Phone      int32     `bson:"phone" json:"phone,omitempty"`
	Gender     int32     `bson:"gender" json:"gender,omitempty"`
	Birth      time.Time `bson:"birth" json:"birth,omitempty"`
	WeiChat    string    `bson:"weichat" json:"weichat,omitempty"`
	JobStatus  int32     `bson:"jobstatus" json:"jobstatus,omitempty"`
	CreateAt   time.Time `bson:"createAt" json:"createAt,omitempty"`
	UpdateAt   time.Time `bson:"updateAt" json:"updateAt,omitempty"`
}
