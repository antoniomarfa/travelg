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

// SetUserRoutes creates sale routes
func SetSaleRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.SaleService) {
	//	r.Use(middlewares.Recover())

	r.POST("/v2.0/sale", createSale(ctx, cfg, p))
	r.GET("/v2.0/sale", getAllSale(ctx, cfg, p))
	r.GET("/v2.0/sale/:id", getSaleByID(ctx, cfg, p))
	r.PATCH("/v2.0/sale/:id", updateSale(ctx, cfg, p))
	r.DELETE("/v2.0/sale/:id", deleteSale(ctx, cfg, p))
}

// @Summary Create sale
// @Description Creates a new sale
// @Tags sale
// @Param user body models.CreatesaleReq true "New sale to be created"
// @Success 201 {object} models.saleResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/sale [post]
func createSale(ctx context.Context, cfg config.Config, p ports.SaleService) gin.HandlerFunc {
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
		var sale models.CreateSaleReq
		err = json.Unmarshal(body, &sale)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Llamar al servicio para crear los sale
		result, err := p.Create(ctx, sale)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Devolver la respuesta con el código de estado adecuado
		utils.ResponseJSON(c.Writer, c.Request, body, http.StatusCreated, result)
	}
}

// @Summary Get all sale
// @Description Gets all the sale
// @Tags sale
// @Success 200 {array} models.saleResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/sale [get]
func getAllSale(ctx context.Context, cfg config.Config, p ports.SaleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		sale, err := p.GetAll(ctx)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, sale)
	}
}

// @Summary Get sale by ID
// @Description Gets a sale by ID
// @Tags sale
// @Param id path string true "ID"
// @Success 200 {object} models.saleResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/sale/{id} [get]
func getSaleByID(ctx context.Context, cfg config.Config, p ports.SaleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		//	var params = mux.Vars(c.Request)
		sale, err := p.GetByID(ctx, c.Param("id"))
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, sale)
	}
}

// @Summary Update sale
// @Description Updates a sale
// @Tags sale
// @Param id path string true "ID"
// @Param User body models.UpdatesaleReq true "sale"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/users/{id} [patch]
func updateSale(ctx context.Context, cfg config.Config, p ports.SaleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		params := c.Params
		var sale models.UpdateSaleReq
		err = json.Unmarshal(body, &sale)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), sale)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, "Updated id: "+id)
	}
}

// @Summary Delete sale
// @Description Delete a sale
// @Tags sale
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/sale/{id} [delete]
func deleteSale(ctx context.Context, cfg config.Config, p ports.SaleService) gin.HandlerFunc {
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
