package handler

import (
	"encoding/json"
	"net/http"

	"github.com/adullahkapadia/auth-service/internal/config"
	"github.com/adullahkapadia/auth-service/internal/middleware"
	"github.com/adullahkapadia/auth-service/internal/model"
	"github.com/adullahkapadia/auth-service/internal/service"
	"github.com/adullahkapadia/auth-service/internal/utils"
)

type AuthHandler struct {
	Service *service.AuthService
	Config  *config.Config
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {

	var req model.RegisterRequest

	// Decode JSON body
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		utils.JSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid JSON",
		})
		return
	}

	// Basic Validation
	if req.Name == "" || req.Email == "" || req.Password == "" {

		utils.JSON(w, http.StatusBadRequest, map[string]string{
			"error": "All fields are required",
		})
		return 
	}

	// Call Service
	err = h.Service.Register(&req)

	if err != nil {

		utils.JSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})

		return
	}

	utils.JSON(w, http.StatusCreated, map[string]string{
		"message": "User registered successfully",
	})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	var req model.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {

		utils.JSON(w,400,map[string]string{
			"error":"Invalid JSON",
		})

		return
	}

	token, err := h.Service.Login(
		&req,
		h.Config.JWTSecret,
	)

	if err != nil {

		utils.JSON(w,401,map[string]string{
			"error":err.Error(),
		})

		return
	}

	utils.JSON(w,200,map[string]string{
		"token":token,
	})
}

func (h *AuthHandler) Profile(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID := r.Context().Value(
		middleware.UserIDKey,
	).(string)

	user, err := h.Service.GetProfile(userID)

	if err != nil {

		utils.JSON(
			w,
			404,
			map[string]string{
				"error":"User not found",
			},
		)

		return
	}

	utils.JSON(
		w,
		200,
		user,
	)

}