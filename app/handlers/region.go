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

// SetRegionRoutes creates region routes
func SetRegionRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.RegionService) {
	//	r.Use(middlewares.Recover())

	r.POST("/v2.0/region", createRegion(ctx, cfg, p))
	r.GET("/v2.0/region", getAllRegion(ctx, cfg, p))
	r.GET("/v2.0/region/:id", getRegionByID(ctx, cfg, p))
	r.PATCH("/v2.0/region/:id", updateRegion(ctx, cfg, p))
	r.DELETE("/v2.0/region/:id", deleteRegion(ctx, cfg, p))
}

// @Summary Create region
// @Description Creates a new region
// @Tags region
// @Param user body models.CreateregionReq true "New region to be created"
// @Success 201 {object} models.regionResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/region [post]
func createRegion(ctx context.Context, cfg config.Config, p ports.RegionService) gin.HandlerFunc {
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
		var region models.CreateRegionReq
		err = json.Unmarshal(body, &region)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Llamar al servicio para crear los region
		result, err := p.Create(ctx, region)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Devolver la respuesta con el código de estado adecuado
		utils.ResponseJSON(c.Writer, c.Request, body, http.StatusCreated, result)
	}
}

// @Summary Get all region
// @Description Gets all the region
// @Tags region
// @Success 200 {array} models.regionResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/region [get]
func getAllRegion(ctx context.Context, cfg config.Config, p ports.RegionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		filter := make(map[string]interface{})
		region, err := p.GetAll(ctx, filter)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, region)
	}
}

// @Summary Get region by ID
// @Description Gets a region by ID
// @Tags region
// @Param id path string true "ID"
// @Success 200 {object} models.regionResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/region/{id} [get]
func getRegionByID(ctx context.Context, cfg config.Config, p ports.RegionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		//	var params = mux.Vars(c.Request)
		region, err := p.GetByID(ctx, c.Param("id"))
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, region)
	}
}

// @Summary Update region
// @Description Updates a region
// @Tags region
// @Param id path string true "ID"
// @Param User body models.UpdateregionReq true "region"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/region/{id} [patch]
func updateRegion(ctx context.Context, cfg config.Config, p ports.RegionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		params := c.Params
		var region models.UpdateRegionReq
		err = json.Unmarshal(body, &region)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), region)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, "Updated id: "+id)
	}
}

// @Summary Delete region
// @Description Delete a region
// @Tags region
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/region/{id} [delete]
func deleteRegion(ctx context.Context, cfg config.Config, p ports.RegionService) gin.HandlerFunc {
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
