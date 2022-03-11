package user_dto

type CreateUserDto struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UpdateUserDto struct {
	ID       uint64 `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
