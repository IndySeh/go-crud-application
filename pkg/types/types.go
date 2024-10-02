package types

type User struct {
	Id    int
	Name  string
	Email string
}

type DeleteResponse struct {
	Message string `json:"message"`
	User    *User  `json:"user"`
}

type PostResponse struct {
	Message string `json:"message"`
	User    *User  `json:"user,omitempty"`
}




