package model

// Other Student fields are not needed from Db so no need to parse them
// Email Field Just has intials of email like yashps22
type Student struct {
	Id     string `json:"i"`
	Name   string `json:"n"`
	Gender string `json:"g"`
	Email  string `json:"u"`
	Hall   string `json:"h"`
	Room   string `json:"r"`
}

// ungroupedData
type Students []Student

// groupedData
type Room Students
type Wing map[string]Room
type Hall map[string]Wing

// Request json format for wingiesOrNot route
type WingiesOrNot struct {
	Id1 string `json:"id1"`
	Id2 string `json:"id2"`
}
