package handlers

import (
	"context"
	"net/http"
	"strconv"

	"travel/config"

	"travel/tools/api/utils"

	"github.com/gin-gonic/gin"
)

// SetHealthRoutes creates health routes
func SetHealthRoutes(ctx context.Context, cfg config.Config, r *gin.Engine) {
	//	r.Handle("/health", healthCheck(ctx, cfg)).Methods(http.MethodGet)
	//	r.Handle("/health", healthCheck(cfg)).Methods(http.MethodGet)
	r.GET("/health", healthCheck(cfg))
}

// @Summary Health Check
// @Description Runs a Health Check
// @Tags Health
// @Success 200 "OK"
// @Failure 500 {object} object
// @Failure 503 {object} object
// @Router /health [get]
// func healthCheck(ctx context.Context, cfg config.Config) http.Handler {
func healthCheck(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Añadiendo headers con la información de configuración
		c.Header("Version", cfg.Version)
		c.Header("Environment", cfg.Environment)
		c.Header("Port", strconv.Itoa(cfg.Port))
		c.Header("Database", cfg.Database)
		c.Header("DSN", cfg.DSN)

		// Llamada a la utilidad para responder en JSON
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusOK, nil)
	}
}
