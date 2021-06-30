package entity

type User struct {
	Id ID
	Devices []Device
}

func NewUser(id string) *User {
	return &User{Id: StrToID(id)}
}

// Находит девайс в коллекции девайсов пользователя (по токену) и удаляет его
func (u *User) RemoveDevice(deviceToken string) {
    for i, d := range u.Devices {
        if d.Token == deviceToken {
            u.Devices = append(u.Devices[:i], u.Devices[i+1:]...)
			break
        }
    }
}