package handlers

import (
	"net/http"

	"github.com/demola234/defifundr/internal/adapters/dto/request"
	"github.com/demola234/defifundr/internal/adapters/dto/response"
	"github.com/demola234/defifundr/internal/core/domain"
	"github.com/demola234/defifundr/internal/core/ports"
	"github.com/demola234/defifundr/pkg/app_errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService ports.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService ports.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetProfile godoc
// @Summary Get user profile
// @Description Retrieve authenticated user's profile
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.UserResponse "User profile"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Router /users/profile [get]
func (h *UserHandler) GetProfile(ctx *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Error: "Unauthorized",
		})
		return
	}

	// Convert user ID to UUID
	userUUID, ok := userID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Error: "Invalid user ID",
		})
		return
	}

	// Get user profile
	user, err := h.userService.GetUserByID(ctx, userUUID)
	if err != nil {
		errResponse := response.ErrorResponse{
			Error: app_errors.ErrInternalServer.Error(),
		}

		if app_errors.IsAppError(err) {
			appErr := err.(*app_errors.AppError)
			errResponse.Error = appErr.Error()

			if appErr.ErrorType == app_errors.ErrorTypeNotFound {
				ctx.JSON(http.StatusNotFound, errResponse)
				return
			}

			ctx.JSON(http.StatusBadRequest, errResponse)
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// Create response DTO
	userResponse := response.UserResponse{
		ID:                  user.ID,
		Email:               user.Email,
		FirstName:           user.FirstName,
		LastName:            user.LastName,
		AccountType:         user.AccountType,
		PersonalAccountType: user.PersonalAccountType,
		Nationality:         user.Nationality,
		CreatedAt:           user.CreatedAt,
	}

	if user.Gender != nil {
		userResponse.Gender = *user.Gender
	}

	if user.ResidentialCountry != nil {
		userResponse.ResidentialCountry = *user.ResidentialCountry
	}

	if user.JobRole != nil {
		userResponse.JobRole = *user.JobRole
	}

	if user.CompanyWebsite != nil {
		userResponse.CompanyWebsite = *user.CompanyWebsite
	}

	if user.EmploymentType != nil {
		userResponse.EmploymentType = *user.EmploymentType
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "User profile retrieved",
		Data:    userResponse,
	})
}

// UpdateProfile godoc
// @Summary Update user profile
// @Description Update authenticated user's profile information
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param profile body request.UpdateProfileRequest true "Profile data to update"
// @Success 200 {object} response.UserResponse "Updated user profile"
// @Failure 400 {object} response.ErrorResponse "Invalid request"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Router /users/profile [put]
func (h *UserHandler) UpdateProfile(ctx *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Error: "Unauthorized",
		})
		return
	}

	// Convert user ID to UUID
	userUUID, ok := userID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Error: "Invalid user ID",
		})
		return
	}

	var req request.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error:   app_errors.ErrInvalidRequest.Error(),
			Details: err.Error(),
		})
		return
	}

	// Validate request data
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error:   app_errors.ErrInvalidRequest.Error(),
			Details: err.Error(),
		})
		return
	}

	// Get existing user
	currentUser, err := h.userService.GetUserByID(ctx, userUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Error: "Failed to retrieve user profile",
		})
		return
	}

	// Update user object with new values
	updatedUser := domain.User{
		ID:                  userUUID,
		Email:               currentUser.Email, // Email cannot be changed
		FirstName:           req.FirstName,
		LastName:            req.LastName,
		AccountType:         currentUser.AccountType, // Account type cannot be changed
		PersonalAccountType: currentUser.PersonalAccountType,
		Nationality:         req.Nationality,
		Gender:              &req.Gender,
		ResidentialCountry:  &req.ResidentialCountry,
		JobRole:             &req.JobRole,
		CompanyWebsite:      &req.CompanyWebsite,
		EmploymentType:      &req.EmploymentType,
	}

	// Update user profile
	user, err := h.userService.UpdateUser(ctx, updatedUser)
	if err != nil {
		errResponse := response.ErrorResponse{
			Error: app_errors.ErrInternalServer.Error(),
		}

		if app_errors.IsAppError(err) {
			appErr := err.(*app_errors.AppError)
			errResponse.Error = appErr.Error()
			ctx.JSON(http.StatusBadRequest, errResponse)
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// Create response DTO
	userResponse := response.UserResponse{
		ID:                  user.ID,
		Email:               user.Email,
		FirstName:           user.FirstName,
		LastName:            user.LastName,
		AccountType:         user.AccountType,
		PersonalAccountType: user.PersonalAccountType,
		Nationality:         user.Nationality,
		CreatedAt:           user.CreatedAt,
		UpdatedAt:           user.UpdatedAt,
	}

	if user.Gender != nil {
		userResponse.Gender = *user.Gender
	}

	if user.ResidentialCountry != nil {
		userResponse.ResidentialCountry = *user.ResidentialCountry
	}

	if user.JobRole != nil {
		userResponse.JobRole = *user.JobRole
	}

	if user.CompanyWebsite != nil {
		userResponse.CompanyWebsite = *user.CompanyWebsite
	}

	if user.EmploymentType != nil {
		userResponse.EmploymentType = *user.EmploymentType
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "User profile updated",
		Data:    userResponse,
	})
}

// ChangePassword godoc
// @Summary Change user password
// @Description Change authenticated user's password
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param password body request.ChangePasswordRequest true "Password change data"
// @Success 200 {object} response.SuccessResponse "Password changed successfully"
// @Failure 400 {object} response.ErrorResponse "Invalid request"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Router /users/change-password [post]
func (h *UserHandler) ChangePassword(ctx *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Error: "Unauthorized",
		})
		return
	}

	// Convert user ID to UUID
	userUUID, ok := userID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Error: "Invalid user ID",
		})
		return
	}

	var req request.ChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error:   app_errors.ErrInvalidRequest.Error(),
			Details: err.Error(),
		})
		return
	}

	// Validate request data
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error:   app_errors.ErrInvalidRequest.Error(),
			Details: err.Error(),
		})
		return
	}

	// Change password
	err := h.userService.UpdatePassword(ctx, userUUID, req.OldPassword, req.NewPassword)
	if err != nil {
		errResponse := response.ErrorResponse{
			Error: app_errors.ErrInternalServer.Error(),
		}

		if app_errors.IsAppError(err) {
			appErr := err.(*app_errors.AppError)
			errResponse.Error = appErr.Error()

			if appErr.ErrorType == app_errors.ErrorTypeUnauthorized {
				ctx.JSON(http.StatusUnauthorized, errResponse)
				return
			}

			ctx.JSON(http.StatusBadRequest, errResponse)
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Password changed successfully",
	})
}
