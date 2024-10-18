package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"travel/app/handlers"
	"travel/config"
	"travel/core/ports"
	"travel/core/services"
	"travel/infrastructure/mongo"
	"travel/infrastructure/postgres"

	"travel/tools/infrastructure"

	"github.com/gin-gonic/gin"
)

type api struct {
	config   config.Config
	services svs
}

type svs struct {
	company    ports.CompanyService
	colegios   ports.ColegiosService
	sale       ports.SaleService
	ingreso    ports.IngresoService
	pagos      ports.PagosService
	curso      ports.CursoService
	voucher    ports.VoucherService
	fmedica    ports.FmedicaService
	program    ports.ProgramService
	region     ports.RegionService
	permission ports.PermissionService
	roles      ports.RolesService
	users      ports.UsersService
	comunas    ports.ComunasService
}

// New creates a new API
func New(ctx context.Context, cfg config.Config) (a api) {
	a.config = cfg

	var companyRepo ports.CompanyRepository
	var colegiosRepo ports.ColegiosRepository
	var saleRepo ports.SaleRepository
	var ingresoRepo ports.IngresoRepository
	var pagosRepo ports.PagosRepository
	var cursoRepo ports.CursoRepository
	var voucherRepo ports.VoucherRepository
	var fmedicaRepo ports.FmedicaRepository
	var programRepo ports.ProgramRepository
	var regionRepo ports.RegionRepository
	var permissionRepo ports.PermissionRepository
	var rolesRepo ports.RolesRepository
	var usersRepo ports.UsersRepository
	var comunasRepo ports.ComunasRepository

	switch a.config.Database {
	case "mongo":
		db, err := infrastructure.ConnectMongoDB(ctx, a.config.DSN)
		if err != nil {
			log.Fatal(err)
		}

		companyRepo, err = mongo.NewCompanyRepository(ctx, db)
		if err != nil {
			log.Fatal(err)
		}
	case "postgres":
		db, err := infrastructure.ConnectPostgresOrm(ctx, a.config.DSN)
		if err != nil {
			log.Fatal(err)
		}
		/*
			_, filePath, _, _ := runtime.Caller(0)
			migrationsDir := filepath.Join(filePath, "../../..", cfg.PostgresMigrationsDir)
			err = infrastructure.MigratePostgresDB(db, migrationsDir)
			if err != nil {
				log.Fatal(err)
			}
		*/

		companyRepo = postgres.NewCompanyRepository(ctx, db)
		colegiosRepo = postgres.NewColegiosRepository(ctx, db)
		saleRepo = postgres.NewSaleRepository(ctx, db)
		ingresoRepo = postgres.NewIngresoRepository(ctx, db)
		pagosRepo = postgres.NewPagosRepository(ctx, db)
		cursoRepo = postgres.NewCursoRepository(ctx, db)
		voucherRepo = postgres.NewVoucherRepository(ctx, db)
		fmedicaRepo = postgres.NewFmedicaRepository(ctx, db)
		programRepo = postgres.NewProgramRepository(ctx, db)
		regionRepo = postgres.NewRegionRepository(ctx, db)
		permissionRepo = postgres.NewPermissionRepository(ctx, db)
		rolesRepo = postgres.NewRolesRepository(ctx, db)
		usersRepo = postgres.NewUsersRepository(ctx, db)
		comunasRepo = postgres.NewComunasRepository(ctx, db)

	default:
		log.Fatalf("database flag %s not valid", a.config.Database)
	}

	a.services.company = services.NewCompanyService(a.config, companyRepo)
	a.services.colegios = services.NewColegiosService(a.config, colegiosRepo)
	a.services.sale = services.NewSaleService(a.config, saleRepo)
	a.services.ingreso = services.NewIngresoService(a.config, ingresoRepo)
	a.services.pagos = services.NewPagosService(a.config, pagosRepo)
	a.services.curso = services.NewCursoService(a.config, cursoRepo)
	a.services.voucher = services.NewVoucherService(a.config, voucherRepo)
	a.services.fmedica = services.NewFmedicaService(a.config, fmedicaRepo)
	a.services.program = services.NewProgramService(a.config, programRepo)
	a.services.region = services.NewRegionService(a.config, regionRepo)
	a.services.permission = services.NewPermissionService(a.config, permissionRepo)
	a.services.roles = services.NewRolesService(a.config, rolesRepo)
	a.services.users = services.NewUsersService(a.config, usersRepo)
	a.services.comunas = services.NewComunasService(a.config, comunasRepo)

	return a
}

// Run API
func (a *api) Run(ctx context.Context, cancel context.CancelFunc) func() error {
	return func() error {
		defer cancel()
		gin.SetMode(gin.ReleaseMode)
		router := gin.Default()

		handlers.SetHealthRoutes(ctx, a.config, router)
		handlers.SetCompanyRoutes(ctx, a.config, router, a.services.company)
		handlers.SetColegiosRoutes(ctx, a.config, router, a.services.colegios)
		handlers.SetSaleRoutes(ctx, a.config, router, a.services.sale)
		handlers.SetIngresoRoutes(ctx, a.config, router, a.services.ingreso)
		handlers.SetPagosRoutes(ctx, a.config, router, a.services.pagos)
		handlers.SetCursoRoutes(ctx, a.config, router, a.services.curso)
		handlers.SetVoucherRoutes(ctx, a.config, router, a.services.voucher)
		handlers.SetFmedicaRoutes(ctx, a.config, router, a.services.fmedica)
		handlers.SetProgramRoutes(ctx, a.config, router, a.services.program)
		handlers.SetRegionRoutes(ctx, a.config, router, a.services.region)
		handlers.SetPermissionRoutes(ctx, a.config, router, a.services.permission)
		handlers.SetRolesRoutes(ctx, a.config, router, a.services.roles)
		handlers.SetUsersRoutes(ctx, a.config, router, a.services.users)
		handlers.SetComunasRoutes(ctx, a.config, router, a.services.comunas)

		log.Printf("Version: %s", a.config.Version)
		log.Printf("Environment: %s", a.config.Environment)
		log.Printf("Database: %s", a.config.Database)
		log.Printf("Listening on port %d", a.config.Port)

		server := &http.Server{
			Addr:    fmt.Sprintf(":%d", a.config.Port),
			Handler: router,
		}
		go shutdown(ctx, server)
		return server.ListenAndServe()
	}
}

func shutdown(ctx context.Context, server *http.Server) {
	<-ctx.Done()
	log.Printf("Shutting down API gracefully...")
	server.Shutdown(ctx)
}
