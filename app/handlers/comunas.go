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

// SetUserRoutes creates user routes
func SetComunasRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.ComunasService) {
	//	r.Use(middlewares.Recover())

	r.POST("/v2.0/comunas", createComunas(ctx, cfg, p))
	r.GET("/v2.0/comunas", getAllComunas(ctx, cfg, p))
	r.GET("/v2.0/comunas/:id", getComunasByID(ctx, cfg, p))
	r.PATCH("/v2.0/comunas/:id", updateComunas(ctx, cfg, p))
	r.DELETE("/v2.0/comunas/:id", deleteComunas(ctx, cfg, p))
}

// @Summary Create comunas
// @Description Creates a new comunas
// @Tags comunas
// @Param user body models.CreatecomunasReq true "New comunas to be created"
// @Success 201 {object} models.comunasResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/comunas [post]
func createComunas(ctx context.Context, cfg config.Config, p ports.ComunasService) gin.HandlerFunc {
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
		var Comunas models.CreateComunasReq
		err = json.Unmarshal(body, &Comunas)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Llamar al servicio para crear los comunas
		result, err := p.Create(ctx, Comunas)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Devolver la respuesta con el código de estado adecuado
		utils.ResponseJSON(c.Writer, c.Request, body, http.StatusCreated, result)
	}
}

// @Summary Get all comunas
// @Description Gets all the comunas
// @Tags comunas
// @Success 200 {array} models.comunasResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/comunas [get]
func getAllComunas(ctx context.Context, cfg config.Config, p ports.ComunasService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		comunas, err := p.GetAll(ctx)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, comunas)
	}
}

// @Summary Get comunas by ID
// @Description Gets a comunas by ID
// @Tags comunas
// @Param id path string true "ID"
// @Success 200 {object} models.comunasResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/comunas{id} [get]
func getComunasByID(ctx context.Context, cfg config.Config, p ports.ComunasService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		//	var params = mux.Vars(c.Request)
		comunas, err := p.GetByID(ctx, c.Param("id"))
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, comunas)
	}
}

// @Summary Update comunas
// @Description Updates a comunas
// @Tags comunas
// @Param id path string true "ID"
// @Param User body models.UpdatecomunasReq true "comunas"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/comunas/{id} [patch]
func updateComunas(ctx context.Context, cfg config.Config, p ports.ComunasService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		params := c.Params
		var Comunas models.UpdateComunasReq
		err = json.Unmarshal(body, &Comunas)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), Comunas)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, "Updated id: "+id)
	}
}

// @Summary Delete comunas
// @Description Delete a comunas
// @Tags comunas
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/comunas/{id} [delete]
func deleteComunas(ctx context.Context, cfg config.Config, p ports.ComunasService) gin.HandlerFunc {
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
