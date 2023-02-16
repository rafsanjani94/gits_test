package publisher

//interface untuk overide table name
type Tabler interface {
	TableName() string
}

// overide table: default table name GO is 'tableS'
func (Publisher) TableName() string {
	return "publisher"
}

type Publisher struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
