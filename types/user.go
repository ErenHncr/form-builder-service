package types

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func validateUser(u *User) bool { return true }
