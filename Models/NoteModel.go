package Models

type Note struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	IsDelete bool   `json:"isDelete"`
}

func (b *Note) TableName() string {
	return "dbo.note"
}
