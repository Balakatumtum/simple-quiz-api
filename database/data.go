package database

type Question struct {
	Question        string    `json:"question"`
	Answers         [3]string `json:"answers"`
	CorrectAnswerID int       `json:"-"`
}

var ExamPaper = []Question{
	{Question: "Berapa Jumlah fakultas di ITS?",
		Answers:         [3]string{"13", "7", "9"},
		CorrectAnswerID: 1},
	{Question: "Berapa jumlah departemen di ITS?",
		Answers:         [3]string{"39", "42", "27"},
		CorrectAnswerID: 0},
	{Question: "Apa salam yang mempersatukan mahasiswa ITS?",
		Answers:         [3]string{"VERITAS!", "HYGGE!", "VIVAT!"},
		CorrectAnswerID: 2},
	{Question: "Apa motto dari ITS?",
		Answers:         [3]string{"Excellence with morality", "Advancing Humanity", "Building up noble future"},
		CorrectAnswerID: 1},
}
