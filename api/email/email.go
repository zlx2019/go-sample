package email

const name = "email"

type Email struct {
}

func NewEmail() *Email {
	return &Email{}
}

func (e *Email) Init() {
}

func (e *Email) Name() string {
	return name
}
