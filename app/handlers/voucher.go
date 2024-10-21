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

// SetVoucherRoutes creates voucher routes
func SetVoucherRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.VoucherService) {
	//	r.Use(middlewares.Recover())

	r.POST("/v2.0/voucher", createVoucher(ctx, cfg, p))
	r.GET("/v2.0/voucher", getAllVoucher(ctx, cfg, p))
	r.GET("/v2.0/voucher/:id", getVoucherByID(ctx, cfg, p))
	r.PATCH("/v2.0/voucher/:id", updateVoucher(ctx, cfg, p))
	r.DELETE("/v2.0/voucher/:id", deleteVoucher(ctx, cfg, p))
}

// @Summary Create voucher
// @Description Creates a new voucher
// @Tags voucher
// @Param user body models.CreatevoucherReq true "New voucher to be created"
// @Success 201 {object} models.voucherResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/voucher [post]
func createVoucher(ctx context.Context, cfg config.Config, p ports.VoucherService) gin.HandlerFunc {
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
		var voucher models.CreateVoucherReq
		err = json.Unmarshal(body, &voucher)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Llamar al servicio para crear los voucher
		result, err := p.Create(ctx, voucher)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Devolver la respuesta con el código de estado adecuado
		utils.ResponseJSON(c.Writer, c.Request, body, http.StatusCreated, result)
	}
}

// @Summary Get all voucher
// @Description Gets all the voucher
// @Tags voucher
// @Success 200 {array} models.voucherResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/voucher [get]
func getAllVoucher(ctx context.Context, cfg config.Config, p ports.VoucherService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		filter := make(map[string]interface{})
		voucher, err := p.GetAll(ctx, filter)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, voucher)
	}
}

// @Summary Get voucher by ID
// @Description Gets a voucher by ID
// @Tags voucher
// @Param id path string true "ID"
// @Success 200 {object} models.voucherResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/voucher/{id} [get]
func getVoucherByID(ctx context.Context, cfg config.Config, p ports.VoucherService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		//	var params = mux.Vars(c.Request)
		voucher, err := p.GetByID(ctx, c.Param("id"))
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, voucher)
	}
}

// @Summary Update voucher
// @Description Updates a voucher
// @Tags voucher
// @Param id path string true "ID"
// @Param User body models.UpdatevoucherReq true "voucher"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/voucher/{id} [patch]
func updateVoucher(ctx context.Context, cfg config.Config, p ports.VoucherService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		params := c.Params
		var voucher models.UpdateVoucherReq
		err = json.Unmarshal(body, &voucher)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), voucher)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, "Updated id: "+id)
	}
}

// @Summary Delete voucher
// @Description Delete a voucher
// @Tags voucher
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/voucher/{id} [delete]
func deleteVoucher(ctx context.Context, cfg config.Config, p ports.VoucherService) gin.HandlerFunc {
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
