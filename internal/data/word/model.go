package word

// Record 词汇表
type Record struct {
	ID     int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name   string `gorm:"type:varchar(45)"`
	Mobile string `gorm:"type:varchar(45)"`
}

// TableName 声明表名
func (Record) TableName() string {
	return "word"
}
