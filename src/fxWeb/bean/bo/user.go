package bo

type User struct {
	Id   int    `gorm:column:id;"primary_key;type:int4"`
	Name string `gorm:column:name;type:varchar(255)`
	Age  int    `gorm:column:age;type:int4`
	Sex  int    `gorm:column:sex;type:int4`
}

func (User) TableName() string {
	return "USER"
}
