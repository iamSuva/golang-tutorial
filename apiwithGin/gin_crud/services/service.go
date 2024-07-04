package services


type Noteservice struct{

}
type Note struct{
	Title string
	Id int
}
// func (n* Noteservice) GetNoteserive() string{
// 	return "get notes service"
// }
func (n* Noteservice) GetNoteserive() []Note{
	notes:=[]Note{
		{Id:1,Title: "note 1"},
		{Id:2,Title: "note 2"},
		{Id:3,Title: "note 3"},
	}
	return notes
}

// func (n* Noteservice) CreateNoteService() string{
//     return "create note service"
// }
func (n* Noteservice) CreateNoteService() Note{
    note:=Note{
		Title: "new note",
		Id: 4,
	}
	return note
}