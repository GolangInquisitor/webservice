package user

type User struct {
	Uuid       string `json:"uuid"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	MiddleName string `json:"middlename"`
	Fio        string `json:"fio"`
	Gender     string `json:"gender"`
	Age        string `json:"age"`
}
