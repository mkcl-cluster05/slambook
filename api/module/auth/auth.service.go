package auth

import (
	"encoding/json"
	"net/http"
	"slambook/api/middlewere"
	"slambook/utils/binding"
	r "slambook/utils/response"

	"github.com/gin-gonic/gin"
)

const (
	USER_FOUND     = "account exists"
	USER_NOT_FOUND = "account not exists"
)

type AuthService interface {
	registerHandler(*gin.Context)
	loginHandler(*gin.Context)
	changePasswordHandler(*gin.Context)
	forgotPasswordHandler(*gin.Context)
}

type authService struct {
	authRepository AuthRepository
}

func NewAuthService(authRepository AuthRepository) AuthService {
	return &authService{
		authRepository: authRepository,
	}
}

func (service *authService) registerHandler(ctx *gin.Context) {

	var registerDTO RegisterDTO
	if valid := binding.BindData(ctx, &registerDTO); !valid {
		return
	}

	isUserPresent := service.authRepository.checkUser(ctx, registerDTO.Email)

	if isUserPresent {
		ctx.JSON(http.StatusBadRequest, r.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Error:   USER_FOUND,
		})
		return
	}

	c := ctx.Request.Context()
	auth, err := service.authRepository.register(c, registerDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, r.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Error:   err.Error(),
		})
		return
	}

	token, err := generateJWT(auth)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, r.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, r.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Result: AuthResponse{
			AccessToken: token,
		},
	})
}

func (service *authService) loginHandler(ctx *gin.Context) {

	var loginDTO LoginDTO
	if valid := binding.BindData(ctx, &loginDTO); !valid {
		return
	}

	isUserPresent := service.authRepository.checkUser(ctx, loginDTO.Email)

	if !isUserPresent {
		ctx.JSON(http.StatusNotFound, r.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "error",
			Error:   USER_NOT_FOUND,
		})
		return
	}

	c := ctx.Request.Context()
	auth, err := service.authRepository.login(c, loginDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, r.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Error:   err.Error(),
		})
		return
	}

	token, err := generateJWT(auth)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, r.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, r.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Result: AuthResponse{
			AccessToken: token,
		},
	})

}
func (service *authService) changePasswordHandler(ctx *gin.Context) {

	var user middlewere.User
	reqUser := ctx.Request.Header.Get("user")
	json.Unmarshal([]byte(reqUser), &user)

	var changePasswordDTO ChangePasswordDTO
	if valid := binding.BindData(ctx, &changePasswordDTO); !valid {
		return
	}

	c := ctx.Request.Context()
	auth, err := service.authRepository.changePassword(c, user, changePasswordDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, r.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Error:   err.Error(),
		})
		return
	}

	token, err := generateJWT(auth)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, r.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, r.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Result: AuthResponse{
			AccessToken: token,
		},
	})

}
func (service *authService) forgotPasswordHandler(ctx *gin.Context) {

	var user middlewere.User
	reqUser := ctx.Request.Header.Get("user")
	json.Unmarshal([]byte(reqUser), &user)

	var forgotPasswordDTO ForgotPasswordDTO
	if valid := binding.BindData(ctx, &forgotPasswordDTO); !valid {
		return
	}

	isUserPresent := service.authRepository.checkUser(ctx, forgotPasswordDTO.Email)

	if !isUserPresent {
		ctx.JSON(http.StatusBadRequest, r.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Error:   USER_NOT_FOUND,
		})
		return
	}

	ctx.JSON(http.StatusOK, r.SuccessResponse{
		Status:  http.StatusOK,
		Message: "success",
		Result:  "check your email to reset passowrd",
	})

}
