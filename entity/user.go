package entity

type User struct {
	Id ID
	Devices []Device
}

func NewUser(id string) *User {
	return &User{Id: StrToID(id)}
}