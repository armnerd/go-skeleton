package mail

// Record 留言表
type Record struct {
	ID      int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name    string `gorm:"type:char(255)"`
	Email   string `gorm:"type:char(255)"`
	Message string `gorm:"type:text"`
	Ctime   int    `gorm:"type:int(11)"`
}

// TableName 声明表名
func (Record) TableName() string {
	return "mail"
}
