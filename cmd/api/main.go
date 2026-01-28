package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"
    "smarttodo-api/internal/config"
    "smarttodo-api/internal/infrastructure/database"
)

func main() {
	fmt.Println("🚀 SmartTodo+ API")

	// Carrega as configurações
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	// Conectar ao banco de dados
	db, err := database.NewPostgresConnection(&cfg.Database)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer database.Close(db)

	// Configurar o router
	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(120 * time.Second))

	// Configurar CORS
	router.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"http://localhost:*", "http://127.0.0.1:*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300,
    }))

	// Rotas de health check
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "SmartTodo+ API está funcionando corretamente"}`))
	})

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        // Testa conexão com banco
        sqlDB, _ := db.DB()
        err := sqlDB.Ping()

        status := "healthy"
        statusCode := http.StatusOK

        if err != nil {
            status = "unhealthy"
            statusCode = http.StatusServiceUnavailable
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(statusCode)
        w.Write([]byte(fmt.Sprintf(`{"status":"%s","database":"connected"}`, status)))
    })

	// Iniciar servidor
	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	server := &http.Server{
		Addr: serverAddr,
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout: 120 * time.Second,
	}

	// Canal para ouvir sinais do sistema operacional
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Goroutine para iniciar o servidor
	go func() {
        log.Printf("Servidor rodando em http://localhost%s", serverAddr)
        log.Println("Endpoints disponíveis:")
        log.Println("   GET  /        - Root endpoint")
        log.Println("   GET  /health  - Health check")
        log.Println("\n✨ Pressione CTRL+C para parar o servidor")

        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("❌ Erro ao iniciar servidor: %v", err)
        }
    }()
	
	 // Aguarda sinal de interrupção
	 <-quit
	 log.Println("Encerrando servidor...")
 
	 // Graceful shutdown
	 log.Println("Servidor encerrado com sucesso!")
 }
