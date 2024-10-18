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
func SetCompanyRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.CompanyService) {
	//	r.Use(middlewares.Recover())

	r.POST("/v2.0/company", createCompany(ctx, cfg, p))
	r.GET("/v2.0/company", getAllCompany(ctx, cfg, p))
	r.GET("/v2.0/company/:id", getCompanyByID(ctx, cfg, p))
	r.PATCH("/v2.0/company/:id", updateCompany(ctx, cfg, p))
	r.DELETE("/v2.0/company/:id", deleteCompany(ctx, cfg, p))
}

// @Summary Create company
// @Description Creates a new company
// @Tags company
// @Param user body models.CreatecompanyReq true "New company to be created"
// @Success 201 {object} models.companyResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/company [post]
func createCompany(ctx context.Context, cfg config.Config, p ports.CompanyService) gin.HandlerFunc {
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
		var Company models.CreateCompanyReq
		err = json.Unmarshal(body, &Company)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Llamar al servicio para crear los company
		result, err := p.Create(ctx, Company)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Devolver la respuesta con el código de estado adecuado
		utils.ResponseJSON(c.Writer, c.Request, body, http.StatusCreated, result)
	}
}

// @Summary Get all company
// @Description Gets all the company
// @Tags company
// @Success 200 {array} models.companyResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/company [get]
func getAllCompany(ctx context.Context, cfg config.Config, p ports.CompanyService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		company, err := p.GetAll(ctx)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, company)
	}
}

// @Summary Get company by ID
// @Description Gets a company by ID
// @Tags company
// @Param id path string true "ID"
// @Success 200 {object} models.companyResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/company{id} [get]
func getCompanyByID(ctx context.Context, cfg config.Config, p ports.CompanyService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		//	var params = mux.Vars(c.Request)
		company, err := p.GetByID(ctx, c.Param("id"))
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, company)
	}
}

// @Summary Update company
// @Description Updates a company
// @Tags company
// @Param id path string true "ID"
// @Param User body models.UpdatecompanyReq true "company"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/company/{id} [patch]
func updateCompany(ctx context.Context, cfg config.Config, p ports.CompanyService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		params := c.Params
		var Company models.UpdateCompanyReq
		err = json.Unmarshal(body, &Company)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), Company)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, "Updated id: "+id)
	}
}

// @Summary Delete company
// @Description Delete a company
// @Tags company
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/company/{id} [delete]
func deleteCompany(ctx context.Context, cfg config.Config, p ports.CompanyService) gin.HandlerFunc {
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
