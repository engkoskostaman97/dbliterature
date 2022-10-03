package models

import "time"

type Collection struct {
	ID          int                  `json:"id" gorm:"primary_key:auto_increment"`
	LiteraturID int                  `json:"literatur_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Literatur   LiteraturResponse    `json:"literatur" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID      int                  `json:"user_id"`
	User        UsersProfileResponse `json:"user"`
	CreatedAt   time.Time            `json:"-"`
	UpdatedAt   time.Time            `json:"-"`
}

type CollectionResponse struct {
	ID          int                  `json:"id"`
	LiteraturID int                  `json:"literatur_id"`
	Literatur   LiteraturResponse    `json:"literatur"`
	UserID      int                  `json:"user_id"`
	User        UsersProfileResponse `json:"user"`
}

func (CollectionResponse) TableName() string {
	return "collections"
}
