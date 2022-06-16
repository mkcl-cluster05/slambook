package auth

type Auth struct {
	AuthId    string `bson:"authId" json:"authId"`
	Username  string `bson:"username" json:"username"`
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password" json:"password"`
	Role      string `bson:"role" json:"role"`
	CreatedAt int64  `bson:"createdAt" json:"createdAt"`
	UpdatedAt int64  `bson:"updatedAt" json:"updatedAt"`
}

type RegisterDTO struct {
	Username string `json:"username" binding:"required,min=2,max=25"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type ChangePasswordDTO struct {
	OldPassword string `json:"oldPassword" binding:"required,min=8,max=20"`
	NewPassword string `json:"newPassword" binding:"required,min=8,max=20"`
}

type ForgotPasswordDTO struct {
	Email string `json:"email" binding:"required,email"`
}

type AuthResponse struct {
	AccessToken string `json:"accessToken"`
}
