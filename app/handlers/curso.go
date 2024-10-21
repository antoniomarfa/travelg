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

// SetCursoRoutes creates curso routes
func SetCursoRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.CursoService) {
	//	r.Use(middlewares.Recover())

	r.POST("/v2.0/curso", createCurso(ctx, cfg, p))
	r.GET("/v2.0/curso", getAllCurso(ctx, cfg, p))
	r.GET("/v2.0/curso/:id", getCursoByID(ctx, cfg, p))
	r.PATCH("/v2.0/curso/:id", updateCurso(ctx, cfg, p))
	r.DELETE("/v2.0/curso/:id", deleteCurso(ctx, cfg, p))
}

// @Summary Create curso
// @Description Creates a new curso
// @Tags colegion
// @Param user body models.CreatecursoReq true "New colegio to be created"
// @Success 201 {object} models.cursoResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/curso [post]
func createCurso(ctx context.Context, cfg config.Config, p ports.CursoService) gin.HandlerFunc {
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
		var curso models.CreateCursoReq
		err = json.Unmarshal(body, &curso)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Llamar al servicio para crear los curso
		result, err := p.Create(ctx, curso)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		// Devolver la respuesta con el código de estado adecuado
		utils.ResponseJSON(c.Writer, c.Request, body, http.StatusCreated, result)
	}
}

// @Summary Get all curso
// @Description Gets all the curso
// @Tags curso
// @Success 200 {array} models.cursoResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/curso [get]
func getAllCurso(ctx context.Context, cfg config.Config, p ports.CursoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		filter := make(map[string]interface{})
		curso, err := p.GetAll(ctx, filter)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, curso)
	}
}

// @Summary Get curso by ID
// @Description Gets a curso by ID
// @Tags curso
// @Param id path string true "ID"
// @Success 200 {object} models.cursoResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/curso/{id} [get]
func getCursoByID(ctx context.Context, cfg config.Config, p ports.CursoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		//	var params = mux.Vars(c.Request)
		curso, err := p.GetByID(ctx, c.Param("id"))
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, nil, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, curso)
	}
}

// @Summary Update curso
// @Description Updates a curso
// @Tags curso
// @Param id path string true "ID"
// @Param User body models.UpdatecursoReq true "curso"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/users/{id} [patch]
func updateCurso(ctx context.Context, cfg config.Config, p ports.CursoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout.Duration)
		defer cancel()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}

		params := c.Params
		var curso models.UpdateCursoReq
		err = json.Unmarshal(body, &curso)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), curso)
		if err != nil {
			utils.ResponseError(c.Writer, c.Request, body, err)
			return
		}
		utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, "Updated id: "+id)
	}
}

// @Summary Delete curso
// @Description Delete a curso
// @Tags curso
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/curso/{id} [delete]
func deleteCurso(ctx context.Context, cfg config.Config, p ports.CursoService) gin.HandlerFunc {
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
