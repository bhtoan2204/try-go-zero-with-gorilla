package http

import (
	"context"
	"go-socket/config"
	"go-socket/constant"
	"go-socket/core/acl/idempotency"
	appCtx "go-socket/core/context"
	"go-socket/core/delivery/http/middleware"
	infraacl "go-socket/core/infra/acl"
	"go-socket/core/pkg/server"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var _ App = (*Server)(nil)

type App interface {
	Routes(ctx context.Context, appCtx *appCtx.AppContext) *gin.Engine
	Start(ctx context.Context, appCtx *appCtx.AppContext) error
}

type Server struct {
	cfg        *config.Config
	router     *gin.Engine
	httpServer *http.Server
	handler    RoutingHandler
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Routes(ctx context.Context, appCtx *appCtx.AppContext) *gin.Engine {
	r := gin.New()
	r.MaxMultipartMemory = 50 << 20
	r.RedirectTrailingSlash = false
	r.Use(middleware.SetRequestID())
	idemStore := infraacl.NewRedisIdempotencyStore(appCtx.GetCache())
	idemManager := idempotency.NewManager(
		idemStore,
		constant.DEFAULT_IDEMPOTENCY_LOCK_TTL,
		constant.DEFAULT_IDEMPOTENCY_DONE_TTL,
	)
	r.Use(middleware.IdempotencyMiddleware(idemManager))
	r.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": gin.H{"error": "something went wrong"}})
	}))

	// cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{
		"*",
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
		"X-Inside-Token",
	}
	r.Use(cors.New(corsConfig))

	pingHandler := func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"clientIP": ctx.ClientIP(),
			},
		})
	}
	r.GET("/health-check", pingHandler)
	r.HEAD("/health-check", pingHandler)

	s.router = r
	s.handler = NewRoutingHandler(s.cfg, appCtx.GetRedisClient())

	// public api
	s.registerPublicAPI()

	return r
}

func (s *Server) Start(ctx context.Context, appCtx *appCtx.AppContext) error {
	srv, err := server.New(s.cfg.HttpConfig.Port)
	if err != nil {
		return err
	}

	return srv.ServeHTTPHandler(ctx, s.Routes(ctx, appCtx))
}

func (s *Server) registerPublicAPI() {
	h := s.handler
	router := s.router.Group("")
	apiV1 := router.Group("/api/v1")
	h.RegisterHandlers(apiV1)
}
