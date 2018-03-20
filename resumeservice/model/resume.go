package model

import "time"

//Resume defines ...
type Resume struct {
	ID         string      `bson:"_id" json:"id,omitempty"`
	Name       string      `bson:"name" json:"name,omitempty"`
	Bio        string      `bson:"bio" json:"bio,omitempty"`
	Intentions []string    `bson:"intentions" json:"intentions,omitempty"` //意向
	Works      []Work      `bson:"works" json:"works,omitempty"`           //工作经历
	Projects   []Project   `bson:"projects" json:"projects,omitempty"`     //项目经验
	Educations []Education `bson:"educations" json:"educations,omitempty"` //项目经验
	Attachment string      `bson:"attachment" json:"attachment,omitempty"`
	Socials    []string    `bson:"socials" json:"socials,omitempty"`
	CreateAt   time.Time   `bson:"createAt" json:"createAt,omitempty"`
	UpdateAt   time.Time   `bson:"updateAt" json:"updateAt,omitempty"`
}

//Work 工作经历
type Work struct {
	ID             string    `bson:"id" json:"id,omitempty"`
	Name           string    `bson:"name" json:"name,omitempty"`
	JobPosition    int32     `bson:"jobPosition" json:"jobPosition,omitempty"`
	JobName        int32     `bson:"jobName" json:"jobName,omitempty"`
	Industry       int32     `bson:"industry" json:"industry,omitempty"`
	Tags           []string  `bson:"tags" json:"tags,omitempty"`
	Department     string    `bson:"department" json:"department,omitempty"`
	Responsibility string    `bson:"responsibility" json:"responsibility,omitempty"`
	Performance    string    `bson:"performance" json:"performance,omitempty"`
	StartAt        time.Time `bson:"startAt" json:"startAt,omitempty"`
	EndAt          time.Time `bson:"endAt" json:"endAt,omitempty"`
}

//Education 教育
type Education struct {
	ID          string    `bson:"id" json:"id,omitempty"`
	Name        string    `bson:"name" json:"name,omitempty"`
	Major       int32     `bson:"major" json:"major,omitempty"`
	Degree      int32     `bson:"degree" json:"degree,omitempty"`
	Description int32     `bson:"description" json:"description,omitempty"`
	StartAt     time.Time `bson:"startAt" json:"startAt,omitempty"`
	EndAt       time.Time `bson:"endAt" json:"endAt,omitempty"`
}

//Project 工作经历
type Project struct {
	ID          string    `bson:"id" json:"id,omitempty"`
	Name        string    `bson:"name" json:"name,omitempty"`
	Role        string    `bson:"role" json:"role,omitempty"`
	Description string    `bson:"description" json:"description,omitempty"`
	Performance string    `bson:"performance" json:"performance,omitempty"`
	StartAt     time.Time `bson:"startAt" json:"startAt,omitempty"`
	EndAt       time.Time `bson:"endAt" json:"endAt,omitempty"`
}
