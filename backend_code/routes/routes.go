package routes

import (
	"culturehouse/handlers"
	"culturehouse/middleware"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/api/groups", middleware.WithCORS(handlers.GetGroupsHandler))
	http.HandleFunc("/api/groups/", middleware.WithCORS(handlers.GetGroupWithMediaHandler))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
	http.HandleFunc("/api/news", middleware.WithCORS(handlers.GetNewsHandler))
	http.HandleFunc("/api/event", middleware.WithCORS(handlers.GetEventsHandler))
	http.HandleFunc("/api/festival", middleware.WithCORS(handlers.GetFestivalsHandler))
	http.HandleFunc("/api/news/", middleware.WithCORS(handlers.GetContentItemWithMediaHandler))
	http.HandleFunc("/api/event/", middleware.WithCORS(handlers.GetContentItemWithMediaHandler))
	http.HandleFunc("/api/festival/", middleware.WithCORS(handlers.GetContentItemWithMediaHandler))
	http.HandleFunc("/api/vacancies", middleware.WithCORS(handlers.GetActiveVacanciesHandler))
	http.HandleFunc("/api/contacts", middleware.WithCORS(handlers.GetContactPersonsHandler))
	http.HandleFunc("/api/media/pdf", middleware.WithCORS(handlers.GetPDFMediaHandler))

	http.HandleFunc("/", middleware.WithCORS(handlers.HomeHandler))
}
