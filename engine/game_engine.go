package engine

import (
	"fmt"
	"sync"
	"time"

	"game-backend/models"
)

type GameEngine struct {
	ResponseChan chan models.UserResponse

	mu          sync.Mutex
	winnerFound bool
	winnerID    int

	totalCorrect   int
	totalIncorrect int

	startTime time.Time
}

func NewGameEngine(bufferSize int) *GameEngine {
	return &GameEngine{
		ResponseChan: make(chan models.UserResponse, bufferSize),
		startTime:    time.Now(),
	}
}

func (g *GameEngine) Start(wg *sync.WaitGroup) {
	go func() {
		for resp := range g.ResponseChan {
			g.processResponse(resp)
			wg.Done() // Mark response as processed
		}
	}()
}

func (g *GameEngine) processResponse(resp models.UserResponse) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if resp.IsCorrect {
		g.totalCorrect++
	} else {
		g.totalIncorrect++
	}

	if g.winnerFound {
		return
	}

	if resp.IsCorrect {
		g.winnerFound = true
		g.winnerID = resp.UserID
		fmt.Printf("Winner: User %d\n", g.winnerID)
		fmt.Printf("Time taken to find winner: %v\n", time.Since(g.startTime))
		fmt.Printf("Total Correct Answers By Winner: %d\n", g.totalCorrect)
		fmt.Printf("Total Incorrect Answers By Winner: %d\n", g.totalIncorrect)
	}
}

func (g *GameEngine) PrintMetrics() {
	fmt.Printf("Total Correct Answers: %d\n", g.totalCorrect)
	fmt.Printf("Total Incorrect Answers: %d\n", g.totalIncorrect)
}

func (g *GameEngine) WaitForCompletion(wg *sync.WaitGroup) {
	wg.Wait() // Wait for all responses to be processed
	fmt.Printf("Final Metrics:\n")
	fmt.Printf("Total Correct Answers: %d\n", g.totalCorrect)
	fmt.Printf("Total Incorrect Answers: %d\n", g.totalIncorrect)
}
