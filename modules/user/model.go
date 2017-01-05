package user

type User struct {
	Id       string `bson:"_id,omitempty"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type LoginResponse struct {
	User  User
	Token string
	Error error
}

type RegisterResponse struct {
	User    User
	Success bool
	Error   error
}
