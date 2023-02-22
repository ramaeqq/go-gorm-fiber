package request

type AddPerson struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Address  string `json:"address" validate:"required"`
	Password string `json:"password" validate:"required,min=3"`
}

type UpdatePerson struct {
	Username string `json:"username"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type PersonEmail struct {
	Email string `json:"email" validate:"required"`
}
