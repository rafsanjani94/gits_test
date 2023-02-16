package request

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type BookReq struct {
	Name        string `json:"name" binding:"required"`
	PublisherId int    `json:"publisher_id" binding:"required"`
}
