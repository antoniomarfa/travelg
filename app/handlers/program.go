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

// SetProgramRoutes creates program routes
func SetProgramRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.ProgramService) {
	//	r.Use(middlewares.Recover())

	r.POST("/v2.0/programa", createProgram(ctx, cfg, p))
	r.GET("/v2.0/programa", getAllProgram(ctx, cfg, p))
	r.GET("/v2.0/programa/:id", getProgramByID(ctx, cfg, p))
	r.PATCH("/v2.0/programa/:id", updateProgram(ctx, cfg, p))
	r.DELETE("/v2.0/programa/:id", deleteProgram(ctx, cfg, p))
}

// @Summary Create program
// @Description Creates a new program
// @Tags program
// @Param user body models.CreateprogramReq true "New program to be created"
// @Success 201 {object} models.programResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/program [post]
func createProgram(ctx context.Context, cfg config.Config, p ports.ProgramService) gin.HandlerFunc {
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
		var program models.CreateProgramReq
		err = json.Unmarshal(body, &program)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Llamar al servicio para crear los program
		result, err := p.Create(ctx, program)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Devolver la respuesta con el código de estado adecuado
		utils.ResponseJSON(c.Writer, c.Request, body, http.StatusCreated, result)
	}
}

// @Summary Get all program
// @Description Gets all the program
// @Tags program
// @Success 200 {array} models.programResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/program [get]
func getAllProgram(ctx context.Context, cfg config.Config, p ports.ProgramService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		program, err := p.GetAll(ctx)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, program)
	}
}

// @Summary Get program by ID
// @Description Gets a program by ID
// @Tags program
// @Param id path string true "ID"
// @Success 200 {object} models.programResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/program/{id} [get]
func getProgramByID(ctx context.Context, cfg config.Config, p ports.ProgramService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		//	var params = mux.Vars(c.Request)
		program, err := p.GetByID(ctx, c.Param("id"))
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, program)
	}
}

// @Summary Update program
// @Description Updates a program
// @Tags program
// @Param id path string true "ID"
// @Param User body models.UpdateprogramReq true "program"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/program/{id} [patch]
func updateProgram(ctx context.Context, cfg config.Config, p ports.ProgramService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		params := c.Params
		var program models.UpdateProgramReq
		err = json.Unmarshal(body, &program)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), program)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, "Updated id: "+id)
	}
}

// @Summary Delete program
// @Description Delete a program
// @Tags program
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/program/{id} [delete]
func deleteProgram(ctx context.Context, cfg config.Config, p ports.ProgramService) gin.HandlerFunc {
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
