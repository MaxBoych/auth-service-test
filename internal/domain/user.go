package domain

type User struct {
	//mutable
	Firstname   string
	Lastname    string
	Email       string
	PhoneNumber string
	Role        string

	//immutable
	Username Username
	Password string
}

type UpdatedUserData struct {
	Username Username

	Role        *string
	Firstname   *string
	Lastname    *string
	Email       *string
	PhoneNumber *string
}

func (d *User) IsEmpty() bool {
	return d.Username == ""
}
