package presenters

type SignUpPresenter struct{}

func NewSignUpPresenter() *SignUpPresenter {
	return &SignUpPresenter{}
}

func (p *SignUpPresenter) ToResponse() struct{} {
	return struct{}{}
}
