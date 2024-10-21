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

// SetUsersRoutes creates users routes
func SetUsersRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.UsersService) {
	//	r.Use(middlewares.Recover())

	r.POST("/v2.0/users", createUsers(ctx, cfg, p))
	r.GET("/v2.0/users", getAllUsers(ctx, cfg, p))
	r.GET("/v2.0/users/:id", getUsersByID(ctx, cfg, p))
	r.PATCH("/v2.0/users/:id", updateUsers(ctx, cfg, p))
	r.DELETE("/v2.0/users/:id", deleteUsers(ctx, cfg, p))
}

// @Summary Create users
// @Description Creates a new users
// @Tags users
// @Param user body models.CreateusersReq true "New users to be created"
// @Success 201 {object} models.usersResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/users [post]
func createUsers(ctx context.Context, cfg config.Config, p ports.UsersService) gin.HandlerFunc {
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
		var users models.CreateUsersReq
		err = json.Unmarshal(body, &users)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Llamar al servicio para crear los users
		result, err := p.Create(ctx, users)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Devolver la respuesta con el código de estado adecuado
		utils.ResponseJSON(c.Writer, c.Request, body, http.StatusCreated, result)
	}
}

// @Summary Get all users
// @Description Gets all the users
// @Tags users
// @Success 200 {array} models.usersResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/users [get]
func getAllUsers(ctx context.Context, cfg config.Config, p ports.UsersService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		filter := make(map[string]interface{})
		users, err := p.GetAll(ctx, filter)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, users)
	}
}

// @Summary Get users by ID
// @Description Gets a users by ID
// @Tags users
// @Param id path string true "ID"
// @Success 200 {object} models.usersResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/users/{id} [get]
func getUsersByID(ctx context.Context, cfg config.Config, p ports.UsersService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		//	var params = mux.Vars(c.Request)
		users, err := p.GetByID(ctx, c.Param("id"))
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, users)
	}
}

// @Summary Update users
// @Description Updates a users
// @Tags users
// @Param id path string true "ID"
// @Param User body models.UpdateusersReq true "users"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/users/{id} [patch]
func updateUsers(ctx context.Context, cfg config.Config, p ports.UsersService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		params := c.Params
		var users models.UpdateUsersReq
		err = json.Unmarshal(body, &users)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), users)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, "Updated id: "+id)
	}
}

// @Summary Delete users
// @Description Delete a users
// @Tags users
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/users/{id} [delete]
func deleteUsers(ctx context.Context, cfg config.Config, p ports.UsersService) gin.HandlerFunc {
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
