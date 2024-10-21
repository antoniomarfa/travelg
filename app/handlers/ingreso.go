package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"travel/config"
	"travel/core/models"
	"travel/core/ports"

	//	"travel/tools/api/middlewares"
	"travel/tools/api/utils"

	"github.com/gin-gonic/gin"
)

// SetIngresoRoutes creates ingreso routes
func SetIngresoRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.IngresoService) {
	//	r.Use(middlewares.Recover())

	r.POST("/v2.0/ingreso", createIngreso(ctx, cfg, p))
	r.GET("/v2.0/ingreso", getAllIngreso(ctx, cfg, p))
	r.GET("/v2.0/ingreso/:id", getIngresoByID(ctx, cfg, p))
	r.PATCH("/v2.0/ingreso/:id", updateIngreso(ctx, cfg, p))
	r.DELETE("/v2.0/ingreso/:id", deleteIngreso(ctx, cfg, p))
}

// @Summary Create ingreso
// @Description Creates a new ingreso
// @Tags ingreso
// @Param user body models.CreateingresoReq true "New colegio to be created"
// @Success 201 {object} models.ingresoResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/ingreso [post]
func createIngreso(ctx context.Context, cfg config.Config, p ports.IngresoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Establecer un timeout en el contexto  c.Request.Context()
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		// Leer el cuerpo de la solicitud
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Parsear el JSON al modelo adecuado
		var ingreso models.CreateIngresoReq
		err = json.Unmarshal(body, &ingreso)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Llamar al servicio para crear los ingreso
		result, err := p.Create(ctx, ingreso)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Devolver la respuesta con el código de estado adecuado
		utils.ResponseJSON(c.Writer, c.Request, body, http.StatusCreated, result)
	}
}

// @Summary Get all ingreso
// @Description Gets all the ingreso
// @Tags ingreso
// @Success 200 {array} models.ingresoResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/ingreso [get]
func getAllIngreso(ctx context.Context, cfg config.Config, p ports.IngresoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		filter := make(map[string]interface{})
		ingreso, err := p.GetAll(ctx, filter)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, ingreso)
	}
}

// @Summary Get ingreso by ID
// @Description Gets a ingreso by ID
// @Tags ingreso
// @Param id path string true "ID"
// @Success 200 {object} models.ingresoResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/ingreso/{id} [get]
func getIngresoByID(ctx context.Context, cfg config.Config, p ports.IngresoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		//	var params = mux.Vars(c.Request)
		ingreso, err := p.GetByID(ctx, c.Param("id"))
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, ingreso)
	}
}

// @Summary Update ingreso
// @Description Updates a ingreso
// @Tags ingreso
// @Param id path string true "ID"
// @Param User body models.UpdateingresoReq true "ingreso"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/users/{id} [patch]
func updateIngreso(ctx context.Context, cfg config.Config, p ports.IngresoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		params := c.Params
		var ingreso models.UpdateIngresoReq
		err = json.Unmarshal(body, &ingreso)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), ingreso)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, "Updated id: "+id)
	}
}

// @Summary Delete ingreso
// @Description Delete a ingreso
// @Tags ingreso
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/ingreso/{id} [delete]
func deleteIngreso(ctx context.Context, cfg config.Config, p ports.IngresoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		//	var params = mux.Vars(c.Request)
		id := c.Param("id")
		err := p.Delete(ctx, c.Param("id"))
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, "deleted id: "+id)
	}
}
