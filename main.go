package main

import (
	"game-backend/api"
	"game-backend/engine"
	"game-backend/simulator"
	"sync"
	"time"
)

func main() {
	// Initialize Game Engine
	gameEngine := engine.NewGameEngine(1001) // keep buffer size slightly larger than expected to avoid blocking

	// Create a WaitGroup to track responses
	var wg sync.WaitGroup
	wg.Add(1000) // Add 1000 responses to track

	gameEngine.Start(&wg)

	// Start API Server
	server := api.NewServer(gameEngine)
	go server.Start("8080")

	// Give server time to start
	time.Sleep(1 * time.Second)

	// Simulate 1000 users
	simulator.SimulateUsers(1000, "http://localhost:8080/submit")

	// Wait for all responses and print final metrics
	gameEngine.WaitForCompletion(&wg)
}
