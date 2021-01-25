package server

import (
	"context"
	"github.com/abylq/learning-management-system/components/admin/quiz"
	quhttp "github.com/abylq/learning-management-system/components/admin/quiz/delivery/http"
	qumysql "github.com/abylq/learning-management-system/components/admin/quiz/repository/mysql"
	quusecase "github.com/abylq/learning-management-system/components/admin/quiz/usecase"
	"github.com/abylq/learning-management-system/components/auth"
	authhttp "github.com/abylq/learning-management-system/components/auth/delivery/http"
	authmysql "github.com/abylq/learning-management-system/components/auth/repository/mysql"
	authusecase "github.com/abylq/learning-management-system/components/auth/usecase"
	"github.com/abylq/learning-management-system/components/common/orders"
	orhttp "github.com/abylq/learning-management-system/components/common/orders/delivery/http"
	ormysql "github.com/abylq/learning-management-system/components/common/orders/repository/mysql"
	orusecase "github.com/abylq/learning-management-system/components/common/orders/usecase"
	"github.com/abylq/learning-management-system/mysql/connection"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/itsjamie/gin-cors"
)



type App struct {
	httpServer *http.Server
	authUC     auth.UseCase
	orUC       orders.UseCase
	quizUC     quiz.UseCase
}

func NewApp() *App {
	db := connection.Initialize()
	userRepo := authmysql.NewUserRepository(db)
	orderRepo := ormysql.NewOrderRepository(db)
	quizRepo := qumysql.NewQuizCategoryRepository(db)

	return &App{

		authUC: authusecase.NewAuthUseCase(
			userRepo,
			viper.GetString("auth.hash_salt"),
			[]byte(viper.GetString("auth.signing_key")),
			viper.GetDuration("auth.token_ttl"),
		),
		orUC:   orusecase.NewOrderUseCase(orderRepo),
		quizUC: quusecase.NewOrderUseCase(quizRepo),
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		cors.Middleware(cors.Config{
			Origins:        "*",
			Methods:        "GET, PUT, POST, DELETE",
			RequestHeaders: "Origin, Authorization, Content-Type",
			ExposedHeaders: "",
			MaxAge: 50 * time.Second,
			Credentials: true,
			ValidateHeaders: false,
		}),
	)

	// Set up http handlers
	// SignUp/SignIn endpoints
	authhttp.RegisterHTTPEndpoints(router, a.authUC)

	// API endpoints
	authMiddleware := authhttp.NewAuthMiddleware(a.authUC)
	api := router.Group("/api", authMiddleware)
	orhttp.RegisterHTTPEndpoints(api,a.orUC)
	quhttp.RegisterHTTPEndpoints(api,a.quizUC)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}


