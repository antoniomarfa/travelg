package handlers

/*
import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"travel/config"
	"travel/core/models"
	"travel/core/ports"

	"travel/tools/api/middlewares"
	"travel/tools/api/utils"

	"github.com/gin-gonic/gin"
)

// SetUserRoutes creates colegios routes
func SetColegiosRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.ColegiosService) {
	r.Use(middlewares.Recover())

	r.POST("/v2.0/colegio", createColegios(ctx, cfg, p))
	r.GET("/v2.0/colegio", getAllColegios(ctx, cfg, p))
	r.GET("/v2.0/colegio/:id", getColegiosByID(ctx, cfg, p))
	r.PATCH("/v2.0/colegio/:id", updateColegios(ctx, cfg, p))
	r.DELETE("/v2.0/colegio/:id", deleteColegios(ctx, cfg, p))
}

// @Summary Create colegios
// @Description Creates a new colegios
// @Tags colegion
// @Param user body models.CreatecolegiosReq true "New colegio to be created"
// @Success 201 {object} models.colegiosResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/colegios [post]
func createColegios(ctx context.Context, cfg config.Config, p ports.ColegiosService) gin.HandlerFunc {
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
		var Colegios models.CreateColegiosReq
		err = json.Unmarshal(body, &Colegios)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Llamar al servicio para crear los Colegios
		result, err := p.Create(ctx, Colegios)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Devolver la respuesta con el código de estado adecuado
		utils.ResponseJSON(c.Writer, c.Request, body, http.StatusCreated, result)
	}
}

// @Summary Get all Colegios
// @Description Gets all the Colegios
// @Tags Colegios
// @Success 200 {array} models.ColegiosResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/Colegios [get]
func getAllColegios(ctx context.Context, cfg config.Config, p ports.ColegiosService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		colegios, err := p.GetAll(ctx)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, colegios)
	}
}

// @Summary Get Colegios by ID
// @Description Gets a Colegios by ID
// @Tags Colegios
// @Param id path string true "ID"
// @Success 200 {object} models.ColegiosResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/Colegios/{id} [get]
func getColegiosByID(ctx context.Context, cfg config.Config, p ports.ColegiosService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		//	var params = mux.Vars(c.Request)
		colegios, err := p.GetByID(ctx, c.Param("id"))
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, colegios)
	}
}

// @Summary Update Colegios
// @Description Updates a Colegios
// @Tags Colegios
// @Param id path string true "ID"
// @Param User body models.UpdateColegiosReq true "Colegios"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/users/{id} [patch]
func updateColegios(ctx context.Context, cfg config.Config, p ports.ColegiosService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		params := c.Params
		var colegios models.UpdateColegiosReq
		err = json.Unmarshal(body, &colegios)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), colegios)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, "Updated id: "+id)
	}
}

// @Summary Delete Colegios
// @Description Delete a Colegios
// @Tags Colegios
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/Colegios/{id} [delete]
func deleteColegios(ctx context.Context, cfg config.Config, p ports.ColegiosService) gin.HandlerFunc {
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
*/
