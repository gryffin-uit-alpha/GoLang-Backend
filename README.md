# BlogSpot Backend

The core backend system for the Blogspot platform.

## 1. Introduction
BlogSpot is a core backend development project, built with Go (Golang) for high performance and optimal processing speed. The system uses PostgreSQL as its database and follows the Go Standard Layout (categorized into `cmd`, `internal/db`, `internal/handler`, `internal/utils`), keeping the codebase clean and maintainable.

**Tech Stack:**
- **Language:** Go (Golang 1.22+)
- **Database:** PostgreSQL
- **Routing:** `github.com/go-chi/chi`
- **Tooling:** 
  - `sqlc` (for generating type-safe Go from SQL)
  - `goose` (for running database migrations)
- **Auth:** Bcrypt (Password hashing)