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

// SetPermissionRoutes creates permission routes
func SetPermissionRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.PermissionService) {
	//	r.Use(middlewares.Recover())

	r.POST("/v2.0/permission", createPermission(ctx, cfg, p))
	r.GET("/v2.0/permission", getAllPermission(ctx, cfg, p))
	r.GET("/v2.0/permission/:id", getPermissionByID(ctx, cfg, p))
	r.PATCH("/v2.0/permission/:id", updatePermission(ctx, cfg, p))
	r.DELETE("/v2.0/permission/:id", deletePermission(ctx, cfg, p))
}

// @Summary Create permission
// @Description Creates a new permission
// @Tags permission
// @Param user body models.CreatepermissionReq true "New permission to be created"
// @Success 201 {object} models.permissionResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/permission [post]
func createPermission(ctx context.Context, cfg config.Config, p ports.PermissionService) gin.HandlerFunc {
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
		var permission models.CreateRolesPermissionsReq
		err = json.Unmarshal(body, &permission)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Llamar al servicio para crear los permission
		result, err := p.Create(ctx, permission)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Devolver la respuesta con el código de estado adecuado
		utils.ResponseJSON(c.Writer, c.Request, body, http.StatusCreated, result)
	}
}

// @Summary Get all permission
// @Description Gets all the permission
// @Tags permission
// @Success 200 {array} models.permissionResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/permission [get]
func getAllPermission(ctx context.Context, cfg config.Config, p ports.PermissionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		filter := make(map[string]interface{})
		permission, err := p.GetAll(ctx, filter)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, permission)
	}
}

// @Summary Get permission by ID
// @Description Gets a permission by ID
// @Tags permission
// @Param id path string true "ID"
// @Success 200 {object} models.permissionResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/permission/{id} [get]
func getPermissionByID(ctx context.Context, cfg config.Config, p ports.PermissionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		//	var params = mux.Vars(c.Request)
		permission, err := p.GetByID(ctx, c.Param("id"))
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, permission)
	}
}

// @Summary Update permission
// @Description Updates a permission
// @Tags permission
// @Param id path string true "ID"
// @Param User body models.UpdatepermissionReq true "permission"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/permission/{id} [patch]
func updatePermission(ctx context.Context, cfg config.Config, p ports.PermissionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		params := c.Params
		var permission models.UpdateRolesPermissionsReq
		err = json.Unmarshal(body, &permission)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), permission)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, "Updated id: "+id)
	}
}

// @Summary Delete permission
// @Description Delete a permission
// @Tags permission
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/permission/{id} [delete]
func deletePermission(ctx context.Context, cfg config.Config, p ports.PermissionService) gin.HandlerFunc {
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
