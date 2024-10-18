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

// SetPagosRoutes creates pagos routes
func SetPagosRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.PagosService) {
	//	r.Use(middlewares.Recover())

	r.POST("/v2.0/pagos", createPagos(ctx, cfg, p))
	r.GET("/v2.0/pagos", getAllPagos(ctx, cfg, p))
	r.GET("/v2.0/pagos/:id", getPagosByID(ctx, cfg, p))
	r.PATCH("/v2.0/pagos/:id", updatePagos(ctx, cfg, p))
	r.DELETE("/v2.0/pagos/:id", deletePagos(ctx, cfg, p))
}

// @Summary Create pagos
// @Description Creates a new pagos
// @Tags pagos
// @Param user body models.CreatepagosReq true "New pago to be created"
// @Success 201 {object} models.pagosResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/pagos [post]
func createPagos(ctx context.Context, cfg config.Config, p ports.PagosService) gin.HandlerFunc {
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
		var pagos models.CreatePagosReq
		err = json.Unmarshal(body, &pagos)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Llamar al servicio para crear los pagos
		result, err := p.Create(ctx, pagos)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Devolver la respuesta con el código de estado adecuado
		utils.ResponseJSON(c.Writer, c.Request, body, http.StatusCreated, result)
	}
}

// @Summary Get all pagos
// @Description Gets all the pagos
// @Tags pagos
// @Success 200 {array} models.pagosResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/pagos [get]
func getAllPagos(ctx context.Context, cfg config.Config, p ports.PagosService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		pagos, err := p.GetAll(ctx)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, pagos)
	}
}

// @Summary Get pagos by ID
// @Description Gets a pagos by ID
// @Tags pagos
// @Param id path string true "ID"
// @Success 200 {object} models.pagosResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/pagos/{id} [get]
func getPagosByID(ctx context.Context, cfg config.Config, p ports.PagosService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		//	var params = mux.Vars(c.Request)
		pagos, err := p.GetByID(ctx, c.Param("id"))
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, pagos)
	}
}

// @Summary Update pagos
// @Description Updates a pagos
// @Tags pagos
// @Param id path string true "ID"
// @Param User body models.UpdatepagosReq true "pagos"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/users/{id} [patch]
func updatePagos(ctx context.Context, cfg config.Config, p ports.PagosService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		params := c.Params
		var pagos models.UpdatePagosReq
		err = json.Unmarshal(body, &pagos)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), pagos)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, "Updated id: "+id)
	}
}

// @Summary Delete pagos
// @Description Delete a pagos
// @Tags pagos
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/pagos/{id} [delete]
func deletePagos(ctx context.Context, cfg config.Config, p ports.PagosService) gin.HandlerFunc {
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
