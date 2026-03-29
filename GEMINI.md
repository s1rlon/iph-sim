# GEMINI.md

## Project Overview

**iph-sim** (Idle Planet Hunter Simulator) is a simulation tool built in Go for the game "Idle Planet Miner." It allows users to model and calculate various game mechanics, including planet mining, crafting, manager assignments, and station upgrades.

The project is structured as a monolithic web application using the **Gin** web framework for routing and **SQLite3** for persistent storage of game states and configurations.

### Main Technologies
- **Language:** Go 1.23
- **Web Framework:** [Gin Gonic](https://github.com/gin-gonic/gin)
- **Database:** SQLite3 (`github.com/mattn/go-sqlite3`)
- **Template Engine:** Go `html/template`
- **Development Tool:** [Air](https://github.com/air-verse/air) for live reloading

## Project Structure

- `main.go`: The entry point of the application. It initializes the game state, sets up the Gin engine, loads templates, and registers routes.
- `game/`: Contains the core simulation logic.
    - `game.go`: Defines the primary `Game` struct and its initialization.
    - `planet.go`: Logic related to planet upgrades (Mining, Speed, Cargo).
    - `manager.go`: Manager assignment and bonus calculations.
    - `crafting.go` & `alloys.go` & `items.go`: Logic for resource processing and recipe management.
    - `db.go`: Database schema initialization and helper functions.
    - `calcer.go`: Core calculation engine for ROI and production optimization.
- `routes/`: Gin route handlers.
    - `planets.go`: Routes for managing and upgrading planets.
    - `managers.go`: Routes for manager management.
    - `market.go`: Routes for setting market trends and stars.
    - `misc.go`: Routes for ships, rooms, beacon, and station upgrades.
- `templates/`: HTML templates for the web interface.
- `ipm2.sql`: The SQLite database file used for persistence (created on first run).

## Building and Running

### Prerequisites
- Go 1.23 or higher
- [Air](https://github.com/air-verse/air) (optional, for development)

### Development Mode (with Live Reload)
```bash
air
```
This will automatically rebuild and restart the server whenever code or template changes are detected.

### Normal Run
```bash
go run main.go
```
The application will be available at `http://localhost:8080`.

### Building the Binary
```bash
go build -o iph-sim .
./iph-sim
```

## Development Conventions

- **Separation of Concerns:** Keep core simulation logic in the `game/` package and web-related logic in the `routes/` package.
- **Persistence:** All game state changes should be persisted to the SQLite database. Use the `SyncDB` or specific `saveToDB` methods defined in the `game/` structs.
- **Templates:** Templates are located in the `templates/` directory and use standard Go `html/template` syntax. Custom template functions (like `formatNumber`) are defined in `main.go`.
- **Game Data:** The database `ipm2.sql` is automatically initialized with necessary tables if it doesn't exist.
