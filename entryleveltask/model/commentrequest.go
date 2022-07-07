package model

type CommentRequest struct {
	Comment   string `json:"message"`
	ProductID int64  `json:"productID"`
	ParentID int64 `json:"parentID"`
}
