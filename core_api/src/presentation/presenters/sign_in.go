package presenters

type SignInPresenter struct{}

func NewSignInPresenter() *SignInPresenter {
	return &SignInPresenter{}
}

func (p *SignInPresenter) ToResponse() struct{} {
	return struct{}{}
}
