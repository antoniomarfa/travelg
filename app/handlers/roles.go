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

// SetRolesRoutes creates roles routes
func SetRolesRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.RolesService) {
	//	r.Use(middlewares.Recover())

	r.POST("/v2.0/roles", createRoles(ctx, cfg, p))
	r.GET("/v2.0/roles", getAllRoles(ctx, cfg, p))
	r.GET("/v2.0/roles/:id", getRolesByID(ctx, cfg, p))
	r.PATCH("/v2.0/roles/:id", updateRoles(ctx, cfg, p))
	r.DELETE("/v2.0/roles/:id", deleteRoles(ctx, cfg, p))
}

// @Summary Create roles
// @Description Creates a new roles
// @Tags roles
// @Param user body models.CreaterolesReq true "New roles to be created"
// @Success 201 {object} models.rolesResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/roles [post]
func createRoles(ctx context.Context, cfg config.Config, p ports.RolesService) gin.HandlerFunc {
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
		var roles models.CreateRolesReq
		err = json.Unmarshal(body, &roles)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Llamar al servicio para crear los roles
		result, err := p.Create(ctx, roles)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Devolver la respuesta con el código de estado adecuado
		utils.ResponseJSON(c.Writer, c.Request, body, http.StatusCreated, result)
	}
}

// @Summary Get all roles
// @Description Gets all the roles
// @Tags roles
// @Success 200 {array} models.rolesResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/roles [get]
func getAllRoles(ctx context.Context, cfg config.Config, p ports.RolesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		roles, err := p.GetAll(ctx)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, roles)
	}
}

// @Summary Get roles by ID
// @Description Gets a roles by ID
// @Tags roles
// @Param id path string true "ID"
// @Success 200 {object} models.rolesResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/roles/{id} [get]
func getRolesByID(ctx context.Context, cfg config.Config, p ports.RolesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		//	var params = mux.Vars(c.Request)
		roles, err := p.GetByID(ctx, c.Param("id"))
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, roles)
	}
}

// @Summary Update roles
// @Description Updates a roles
// @Tags roles
// @Param id path string true "ID"
// @Param User body models.UpdaterolesReq true "roles"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/roles/{id} [patch]
func updateRoles(ctx context.Context, cfg config.Config, p ports.RolesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		params := c.Params
		var roles models.UpdateRolesReq
		err = json.Unmarshal(body, &roles)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), roles)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, "Updated id: "+id)
	}
}

// @Summary Delete roles
// @Description Delete a roles
// @Tags roles
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/roles/{id} [delete]
func deleteRoles(ctx context.Context, cfg config.Config, p ports.RolesService) gin.HandlerFunc {
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
