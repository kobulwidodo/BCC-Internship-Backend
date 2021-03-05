package entity

type RegisterUser struct {
	Name     string `json:"name" binding:"required"`
	NoHp     string `json:"no_hp" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
