package input

func NewCreateUser(UID string, name string) CreateUser {
	return CreateUser{
		UID:      UID,
		Nickname: name,
	}
}

type CreateUser struct {
	UID      string `validate:"required"`
	Nickname string `validate:"required"`
}

func NewCurrentUser(UID string) CurrentUser {
	return CurrentUser{
		UID: UID,
	}
}

type CurrentUser struct {
	UID string `validate:"required"`
}

func NewUpdateUser(UID string, name string) UpdateUser {
	return UpdateUser{
		Nickname: name,
	}
}

type UpdateUser struct {
	ULID     string `validate:"required"`
	Nickname string `validate:"required"`
}

func NewDeleteUser(ULID string) DeleteUser {
	return DeleteUser{
		ULID: ULID,
	}
}

type DeleteUser struct {
	ULID string `validate:"required"`
}
