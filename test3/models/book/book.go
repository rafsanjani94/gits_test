package book

//interface untuk overide table name
type Tabler interface {
	TableName() string
}

// overide table: default table name GO is 'tableS'
func (Book) TableName() string {
	return "book"
}

type Book struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	AuthorId    int    `json:"author_id"`
	PublisherId int    `json:"publisher_id"`
}
