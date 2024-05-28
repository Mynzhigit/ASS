package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yourusername/flowerstore/internal/app/usecase"
	"github.com/yourusername/flowerstore/pkg/logger"
)

type AdminHandler struct {
	adminUsecase usecase.AdminUsecase
}

func NewAdminHandler(adminUsecase usecase.AdminUsecase) *AdminHandler {
	return &AdminHandler{
		adminUsecase: adminUsecase,
	}
}

func (h *AdminHandler) GetAdminByID(w http.ResponseWriter, r *http.Request) {
	logger.InfoLogger.Println("Received GetAdminByID request")

	adminIDStr := r.URL.Query().Get("admin_id")
	adminID, err := strconv.Atoi(adminIDStr)
	if err != nil {
		http.Error(w, "Invalid admin ID", http.StatusBadRequest)
		return
	}

	admin, err := h.adminUsecase.GetAdminByID(adminID)
	if err != nil {
		http.Error(w, "Failed to get admin", http.StatusInternalServerError)
		logger.ErrorLogger.Printf("Failed to get admin: %v", err)
		return
	}

	response, err := json.Marshal(admin)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		logger.ErrorLogger.Printf("Failed to marshal admin to JSON: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
