package main

import (
	"log"
	"net/http"

	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-2/4-apis-advanced/configs"
	_ "github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-2/4-apis-advanced/docs"
	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-2/4-apis-advanced/internal/entity"
	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-2/4-apis-advanced/internal/infra/database"
	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-2/4-apis-advanced/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ===== DOC PARA O SWAG =====

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with authentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Guilherme Cadilhe
// @contact.url    github.com/Guilherme-G-Cadilhe
// @contact.email  no-email

// @license.name  Lincensa exemplo
// @license.url   Sem Link

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})
	productDB := database.NewProductDB(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	// Middleware para passar Context, poderia ser utilizado para JWT
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", configs.JWTExpiresIn))

	// Logger avançado
	r.Use(middleware.Logger)

	// Logger criado para exemplo
	// r.Use(LogRequest)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth)) // Middleware para verificar o token JWT no cabeçalho Authorization automaticamente
		r.Use(jwtauth.Authenticator)               // Middleware para autenticar o token JWT de forma automatica

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", productHandler.GetProduct)
			r.Put("/", productHandler.UpdateProduct)
			r.Delete("/", productHandler.DeleteProduct)
		})
		r.Get("/", productHandler.GetAllProducts)
		r.Post("/", productHandler.CreateProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	// swag init -g cmd/server/main.go
	//http://localhost:8000/docs/index.html
	http.ListenAndServe(":8000", r)
}

// Custom Middleware log
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
