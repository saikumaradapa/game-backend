package simulator

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"game-backend/models"
)

func SimulateUsers(n int, endpoint string) {
	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		wg.Add(1)

		go func(userID int) {
			defer wg.Done()

			// Create a thread-safe random number generator
			randSource := rand.New(rand.NewSource(time.Now().UnixNano() + int64(userID)))

			// Random delay (10–1000 ms)
			delay := time.Duration(randSource.Intn(991)+10) * time.Millisecond
			time.Sleep(delay)

			resp := models.UserResponse{
				UserID:    userID,
				IsCorrect: randSource.Intn(2) == 1,
			}

			body, _ := json.Marshal(resp)

			http.Post(endpoint, "application/json", bytes.NewBuffer(body))

		}(i)
	}

	wg.Wait()
}
