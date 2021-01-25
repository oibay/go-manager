package models

type Question struct {
	ID 	  		int64
	QuizID 		int64
	Title 		string
	Image 		string
	CorrectID 	int64
	CreatedAT  	string
	UpdatedAT  	string

}
