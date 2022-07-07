package model

type Comment struct {
	ID        int64  `json:"id"`
	Comment   string `json:"message"`
	ProductID int64  `json:"productID"`
	UserID    int64  `json:"userID"`
	ParentID int64 `json:"parentID"`
}
