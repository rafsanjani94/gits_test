package author

//interface untuk overide table name
type Tabler interface {
	TableName() string
}

// overide table: default table name GO is 'tableS'
func (Author) TableName() string {
	return "author"
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
