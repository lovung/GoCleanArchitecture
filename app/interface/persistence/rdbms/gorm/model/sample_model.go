package model

type SampleModel struct {
	ID   string `gorm:"primarykey"`
	Name string
}
