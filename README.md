# Game Backend

This project is a Go-based backend system that simulates multiple users answering a game question, evaluates responses in real-time, and announces a winner.

## Project Structure

```
.
├── api/
│   └── server.go          # API server to handle user responses
├── engine/
│   └── game_engine.go     # Game engine to process responses and determine the winner
├── models/
│   └── models.go          # Data models for the application
├── simulator/
│   └── mock_users.go      # Mock user simulator to generate responses
├── main.go                # Entry point of the application
├── go.mod                 # Go module file
```

## Features

1. **Mock User Engine**
   - Simulates `N` users.
   - Randomly assigns a correct answer flag (yes/no).
   - Adds a random delay (10–1000ms) to simulate network lag.
   - Sends all responses concurrently to the API server.

2. **API Server**
   - Exposes an endpoint `/submit` to receive user responses in JSON format.
   - Forwards each response to the Game Engine for evaluation.

3. **Game Engine**
   - Determines the first user who sent a correct answer.
   - Prints the winner's user ID on the server console.
   - Tracks and prints metrics for correct and incorrect answers.

## How to Run

1. Clone the repository:
   ```bash
   git clone https://github.com/saikumaradapa/game-backend.git
   cd game-backend
   ```

2. Run the application:
   ```bash
   go run -race .
   ```

3. The server will start on `http://localhost:8080`. The simulator will send 1000 user responses to the `/submit` endpoint.

## Example Output

```
2026/04/03 17:27:49 🚀 Server running on port 8080
Winner: User 708
Time taken to find winner: 1.039843s
Total Correct Answers: 500
Total Incorrect Answers: 500
```

## Requirements

- Go 1.18 or higher

## Contributing

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes and push to your fork.
4. Create a pull request.

## License

This project is licensed under the MIT License.