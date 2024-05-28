package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yourusername/flowerstore/internal/app/usecase"
	"github.com/yourusername/flowerstore/pkg/logger"
)

type UserHandler struct {
	userUsecase usecase.Usercase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	logger.InfoLogger.Println("Received GetUserByID request")

	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.userUsecase.GetUserByID(userID)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		logger.ErrorLogger.Printf("Failed to get user: %v", err)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		logger.ErrorLogger.Printf("Failed to marshal user to JSON: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
