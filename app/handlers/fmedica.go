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

// SetFmedicaRoutes creates fmedica routes
func SetFmedicaRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.FmedicaService) {
	//	r.Use(middlewares.Recover())

	r.POST("/v2.0/ficha", createFicha(ctx, cfg, p))
	r.GET("/v2.0/ficha", getAllFicha(ctx, cfg, p))
	r.GET("/v2.0/ficha/:id", getFichaByID(ctx, cfg, p))
	r.PATCH("/v2.0/ficha/:id", updateFicha(ctx, cfg, p))
	r.DELETE("/v2.0/ficha/:id", deleteFicha(ctx, cfg, p))
}

// @Summary Create fmedica
// @Description Creates a new fmedica
// @Tags colegion
// @Param user body models.CreatefmedicaReq true "New colegio to be created"
// @Success 201 {object} models.fmedicaResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/fmedica [post]
func createFicha(ctx context.Context, cfg config.Config, p ports.FmedicaService) gin.HandlerFunc {
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
		var ficha models.CreateFmedicaReq
		err = json.Unmarshal(body, &ficha)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Llamar al servicio para crear los ficha
		result, err := p.Create(ctx, ficha)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Devolver la respuesta con el código de estado adecuado
		utils.ResponseJSON(c.Writer, c.Request, body, http.StatusCreated, result)
	}
}

// @Summary Get all fmedica
// @Description Gets all the fmedica
// @Tags fmedica
// @Success 200 {array} models.fmedicaResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/fmedica [get]
func getAllFicha(ctx context.Context, cfg config.Config, p ports.FmedicaService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		ficha, err := p.GetAll(ctx)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, ficha)
	}
}

// @Summary Get fmedica by ID
// @Description Gets a fmedica by ID
// @Tags fmedica
// @Param id path string true "ID"
// @Success 200 {object} models.fmedicaResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/fmedica/{id} [get]
func getFichaByID(ctx context.Context, cfg config.Config, p ports.FmedicaService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		//	var params = mux.Vars(c.Request)
		ficha, err := p.GetByID(ctx, c.Param("id"))
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, ficha)
	}
}

// @Summary Update fmedica
// @Description Updates a fmedica
// @Tags fmedica
// @Param id path string true "ID"
// @Param User body models.UpdatefmedicaReq true "fmedica"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/fmedica/{id} [patch]
func updateFicha(ctx context.Context, cfg config.Config, p ports.FmedicaService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		params := c.Params
		var ficha models.UpdateFmedicaReq
		err = json.Unmarshal(body, &ficha)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), ficha)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, "Updated id: "+id)
	}
}

// @Summary Delete fmedica
// @Description Delete a fmedica
// @Tags fmedica
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/fmedica/{id} [delete]
func deleteFicha(ctx context.Context, cfg config.Config, p ports.FmedicaService) gin.HandlerFunc {
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
