package collectiondto

type CreateCollection struct {
	ID          int `json:"id"`
	UserID      int `json:"user_id"`
	LiteraturID int `json:"literatur_id"`
}

type UpdateCollection struct {
	ID int `json:"id"`
}

type CollectionResponse struct {
	ID int `json:"id"`
}

