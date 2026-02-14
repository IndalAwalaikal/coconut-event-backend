package router

import (
	"database/sql"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"

	admcontroller "github.com/IndalAwalaikal/coconut-event-hub/backend/internal/controller/admin"
	pubcontroller "github.com/IndalAwalaikal/coconut-event-hub/backend/internal/controller/public"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/middleware"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/repository"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/service"
)

// NewRouter wires controllers and returns an http.Handler
func NewRouter(db *sql.DB) http.Handler {
	r := mux.NewRouter()

	// Resolve absolute storage directory so file server serves correct files
	absStorageDir, err := filepath.Abs("storage")
	if err != nil {
		panic("failed to resolve storage path: " + err.Error())
	}

	// Serve uploaded files from /storage/* -> local storage/ directory
	// This makes poster and documentation images accessible via URLs returned by SaveMultipartFile
	r.PathPrefix("/storage/").Handler(http.StripPrefix("/storage/", http.FileServer(http.Dir(absStorageDir))))

	r.Use(middleware.CORS)

	// repositories & services (use absolute subpaths for storage)
	regRepo := repository.NewRegistrationRepository(db)
	// pass relative storage paths to services so saved file paths remain under /storage/*
	regService := service.NewRegistrationService(db, regRepo, "storage/registrations")

	evtRepo := repository.NewEventRepository(db)
	evtService := service.NewEventService(evtRepo)

	docRepo := repository.NewDocumentationRepository(db)
	docService := service.NewDocumentationService(db, docRepo, "storage/documentations")

	// posters
	posterRepo := repository.NewPosterRepository(db)
	posterService := service.NewPosterService(posterRepo, "storage")

	// controllers
	pubRegCtrl := pubcontroller.NewRegistrationController(regService)
	adminRegCtrl := admcontroller.NewAdminRegistrationController(regService)
	pubEvtCtrl := pubcontroller.NewEventController(evtService)
	adminEvtCtrl := admcontroller.NewAdminEventController(evtService)
	pubDocCtrl := pubcontroller.NewDocumentationController(docService)
	adminDocCtrl := admcontroller.NewAdminDocumentationController(docService)
	adminDashCtrl := admcontroller.NewDashboardController(db)
	pubPosterCtrl := pubcontroller.NewPosterController(posterService)
	adminPosterCtrl := admcontroller.NewAdminPosterController(posterService)

	// auth
	adminRepo := repository.NewAdminRepository(db)
	authService := service.NewAuthService(db, adminRepo)
	authCtrl := admcontroller.NewAuthController(authService)

	// public routes
	r.HandleFunc("/api/registrations", pubRegCtrl.Create).Methods("POST")
	r.HandleFunc("/api/events", pubEvtCtrl.List).Methods("GET")
	r.HandleFunc("/api/events/detail", pubEvtCtrl.Get).Methods("GET")
	r.HandleFunc("/api/documentations", pubDocCtrl.List).Methods("GET")
	r.HandleFunc("/api/documentations/{id}", pubDocCtrl.Get).Methods("GET")
	// posters
	r.HandleFunc("/api/posters", pubPosterCtrl.List).Methods("GET")
	r.HandleFunc("/api/posters/{id}", pubPosterCtrl.Get).Methods("GET")

	// admin auth (login)
	r.HandleFunc("/api/admin/login", authCtrl.Login).Methods("POST")

	// protected admin routes
	adminSub := r.PathPrefix("/api/admin").Subrouter()
	adminSub.Use(func(next http.Handler) http.Handler { return middleware.AdminAuth(next) })
	adminSub.HandleFunc("/registrations", adminRegCtrl.List).Methods("GET")
	adminSub.HandleFunc("/registrations/{id}", adminRegCtrl.Get).Methods("GET")
	adminSub.HandleFunc("/registrations/export", adminRegCtrl.ExportCSV).Methods("GET")
	adminSub.HandleFunc("/dashboard", adminDashCtrl.Get).Methods("GET")
	// admin event routes
	adminSub.HandleFunc("/events", adminEvtCtrl.Create).Methods("POST")
	adminSub.HandleFunc("/events/{id}", adminEvtCtrl.Update).Methods("PUT")
	adminSub.HandleFunc("/events/{id}", adminEvtCtrl.Delete).Methods("DELETE")

	// admin documentation routes
	adminSub.HandleFunc("/documentations", adminDocCtrl.Create).Methods("POST")
	adminSub.HandleFunc("/documentations", adminDocCtrl.List).Methods("GET")
	adminSub.HandleFunc("/documentations/{id}", adminDocCtrl.Update).Methods("PUT")
	adminSub.HandleFunc("/documentations/{id}", adminDocCtrl.Delete).Methods("DELETE")

	// admin posters
	adminSub.HandleFunc("/posters", adminPosterCtrl.Create).Methods("POST")
	adminSub.HandleFunc("/posters", adminPosterCtrl.List).Methods("GET")
	adminSub.HandleFunc("/posters/{id}", adminPosterCtrl.Update).Methods("PUT")
	adminSub.HandleFunc("/posters/{id}", adminPosterCtrl.Delete).Methods("DELETE")

	// health
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200); w.Write([]byte("ok")) })

	return r
}
