package models

type Orders struct {
	ID int64
	UserID int
	DateStart string
	Items int
	DisciplineID int64
	Status int
}
