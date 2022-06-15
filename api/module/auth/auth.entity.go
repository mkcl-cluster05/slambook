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
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangePasswordDTO struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type ForgotPasswordDTO struct {
	Email string `json:"email"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}
