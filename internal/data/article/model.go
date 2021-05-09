package article

// Record 用户表
type Record struct {
	ID       int    `gorm:"primary_key;AUTO_INCREMENT"`
	Title    string `gorm:"type:varchar(255)"`
	Intro    string `gorm:"type:varchar(255)"`
	Cover    string `gorm:"type:varchar(255)"`
	Raw      string `gorm:"type:text"`
	Author   string `gorm:"type:varchar(32)"`
	Type     int    `gorm:"type:int(11)"`
	Timeline int    `gorm:"type:int(11)"`
	Ctime    string `gorm:"type:char(255)"`
}

// TableName 声明表名
func (Record) TableName() string {
	return "article"
}
