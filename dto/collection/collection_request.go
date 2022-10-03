package collectiondto

type CollectionRequest struct {
	UserID      int `json:"user_id" form:"user_id" gorm:"type: varchar(255)" validate:"required"`
	LiteraturID int `json:"literatur_id" form:"literatur_id" gorm:"type: varchar(255)" validate:"required"`
}
