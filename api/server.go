package api

import (
	"encoding/json"
	"log"
	"net/http"

	"game-backend/engine"
	"game-backend/models"
)

type Server struct {
	engine *engine.GameEngine
}

func NewServer(engine *engine.GameEngine) *Server {
	return &Server{engine: engine}
}

func (s *Server) Start(port string) {
	http.HandleFunc("/submit", s.handleSubmit)

	log.Println("🚀 Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func (s *Server) handleSubmit(w http.ResponseWriter, r *http.Request) {
	var resp models.UserResponse

	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Non-blocking send
	select {
	case s.engine.ResponseChan <- resp:
	default:
		log.Println("⚠️ Dropped response (channel full)")
	}

	w.WriteHeader(http.StatusOK)
}