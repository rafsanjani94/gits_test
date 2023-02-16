package user

//interface untuk overide table name
type Tabler interface {
	TableName() string
}

// overide table: default table name GO is 'tableS'
func (User) TableName() string {
	return "users"
}

type User struct {
	Id       int    `json:"id"`
	AuthorId int    `json:"author_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
