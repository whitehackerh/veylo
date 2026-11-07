package dtos

type RegisterUserInput struct {
	Email    string
	Password string
}

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	User        interface{}
	AccessToken string
}
