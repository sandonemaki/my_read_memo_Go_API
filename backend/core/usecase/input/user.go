package input

type CreateUser struct {
	UID      string `validate:"required"`
	Nickname string `validate:"required"`
}

type CurrentUser struct {
	UID string `validate:"required"`
}
