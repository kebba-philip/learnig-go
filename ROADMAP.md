# Go Backend Development — Complete Roadmap
**Duration:** 24 weeks
**Primary framework:** Gin (after learning `net/http` first)

---

## 📌 How to Use This Document

- **Reference it** at the start of every week to know what's coming.
- **Pair it** with `PROGRESS.md` (your weekly progress track
---

## 🗓️ Phase Overview

| Phase | Weeks | Focus | Outcome |
|---|---|---|---|
| **1. Go Fundamentals** | 1–4 | Language syntax, types, control flow, functions, structs, interfaces | Write idiomatic Go programs |
| **2. Software Engineering** | 5–7 | OOP in Go, error handling, testing, Git/GitHub | Write testable, maintainable code |
| **3. HTTP & Backend** | 8–10 | HTTP, `net/http`, concurrency | Build web servers from scratch |
| **4. Databases** | 11–13 | SQL, PostgreSQL, `database/sql`, migrations | Persist data reliably |
| **5. Gin & REST APIs** | 14–17 | Gin, production APIs, auth, security | Build secure REST APIs |
| **6. Advanced Engineering** | 18–20 | Testing, architecture, background jobs | Engineer larger systems |
| **7. Deployment & Portfolio** | 21–24 | Docker, deployment, portfolio projects | Ship a job-ready portfolio |

---

# 📚 Phase 1: Go Fundamentals (Weeks 1–4)

## Week 1 — Introduction to Go

**Learning objectives**
- Understand what Go is and why it's used for backend development.
- Install Go and set up a development environment (WSL, VS Code/Zed).
- Understand Go workspace, project structure, and modules.
- Write, build, and run Go programs.
- Understand packages, imports, and the `main` package.
- Use variables, constants, and basic data types.
- Use `gofmt` and `go vet` as part of the workflow.

**Topics**
- What Go is, history, and backend use cases
- Installing Go, `go version`, WSL setup
- `go mod init`, `go.mod`, module basics
- `go run` vs `go build` vs `./binary`
- `package main`, `import`, `func main()`
- Variables (`var`, `:=`), constants, basic types (`bool`, `string`, `int`, `int64`, `float64`, `byte`, `rune`)
- Zero values, type conversions, `string(int)` trap
- Scope: package, function, block
- `fmt.Println`, `fmt.Printf` with `%v`, `%T`, `%d`, `%s`, `%q`, `%.2f`
- `gofmt -w`, `go vet`

**Exercises**
1. Hello program — print three lines about yourself.
2. Build and run a binary; test it from `/tmp` (proves standalone binaries).
3. `gofmt` the code; verify formatting.
4. Stripped binary experiment (`-ldflags="-s -w"`).
5. Types program — declare variables of each basic type, print value and type.
6. Scope program — demonstrate zero values, float truncation, string conversion trap, and block scope.
7. **Mini-Project: CLI Unit Converter** (details below).

**Mini-Project — CLI Unit Converter**
- Converts 0°C, 100°C, 37°C to Fahrenheit (hardcoded values).
- Uses `const` for the banner and conversion ratio.
- Uses `float64` for temperatures.
- Prints with `%.2f`.
- Passes `gofmt` and `go vet` clean.
- *Stretch (optional):* add km → miles conversion.

**Assessment**
- 10-question quiz on variables, types, zero values, `go run` vs `go build`, `gofmt`, `go vet`.
- Code review of the unit converter.
- One-sentence explanations of key concepts (binary size, `go vet` warning).

**Deliverables**
- `hello/`, `types/`, `scope/`, `unitconv/` folders.
- All code gofmt'd and vet-clean.

---

## Week 2 — Control Structures and Problem-Solving

**Learning objectives**
- Master operators and expressions.
- Use `if`, `else`, `switch`.
- Use `for` loops, `break`, `continue`.
- Write algorithms to solve simple problems.
- Understand time and space complexity basics.

**Topics**
- Arithmetic, comparison, logical operators
- `if` / `else if` / `else`, initialization in `if`
- `switch` (including no-condition switch), `fallthrough`
- `for` — the only loop in Go (3 forms: `for {}`, `for cond {}`, `for init; cond; post {}`)
- `range` over slices/maps/strings
- `break`, `continue`, labeled loops
- Basic algorithms: sum, max, reverse, count, filter
- Big-O intuition (O(1), O(n), O(n²))

**Exercises**
1. FizzBuzz.
2. Sum of numbers 1–N.
3. FizzBuzz with `switch`.
4. Reverse a slice of strings (loop, no built-ins).
5. Multiplication table with nested loops.
6. **Mini-Project: Number Guessing CLI** — the program picks a number, user guesses with hints (higher/lower). Uses `for`, `if`, `fmt.Scan`.
7. **Stretch:** prime checker, factorial.

**Assessment**
- Quiz on operators, `switch`, loop forms, `break`/`continue`.
- Code review of the guessing game.
- **This is where your Python experience helps — draw parallels.**

---

## Week 3 — Functions, Collections, and Data Structures

**Learning objectives**
- Write functions with parameters and return values.
- Use multiple return values and `error`.
- Use arrays, slices, and maps idiomatically.
- Use variadic functions and closures.

**Topics**
- Function signatures, named returns
- Multiple return values
- `error` and `nil` (first exposure; deep dive in Week 6)
- Arrays vs slices; `len`, `cap`, `append`, `copy`, `make`
- Slicing syntax and slice aliasing
- Maps: `make`, `delete`, iteration, comma-ok idiom
- Variadic functions (`...`)
- Closures and anonymous functions
- `strconv` package (`Atoi`, `Itoa`, `ParseFloat`)

**Exercises**
1. Function that returns (int, error) — division with divide-by-zero check.
2. Slice exercises: append, remove, filter, sum.
3. Map exercises: word count, frequency counter, deduplicate.
4. Variadic function: `sum(nums ...int) int`.
5. Closure: counter generator.
6. **Mini-Project: CLI Todo List (in-memory)** — add, list, delete todos using slices and maps. No database yet. Uses `fmt.Scan` or `bufio.Scanner`.

**Assessment**
- Quiz on slices vs arrays, map operations, closures, multiple returns.
- Code review of the todo CLI.

---

## Week 4 — Go's Core Language Features

**Learning objectives**
- Use pointers safely.
- Define structs and methods.
- Compose types via embedding.
- Define and implement interfaces.
- Organize code into packages.
- Use Go modules properly.

**Topics**
- Pointers (`*`, `&`), nil pointers, why pointers matter
- Structs, methods (value vs pointer receivers)
- Embedding and composition
- Interfaces, implicit implementation, empty interface (`any`)
- Type assertions and type switches
- Packages: `internal/`, naming, exports
- `go.mod`, `go.sum`, semantic imports, `go get`, `go mod tidy`
- Building a multi-package project

**Exercises**
1. Struct: `Book` with fields, methods, printing.
2. Pointer receiver vs value receiver demonstration.
3. Interface: `Shape` with `Area()`; `Circle`, `Rectangle`.
4. Type switch over `any`.
5. Multi-package project: `internal/models`, `internal/services`, `cmd/app`.
6. **Mini-Project: CLI Library Manager** — books, members, borrow/return, using structs, interfaces, and multiple packages. No database (in-memory).

**Assessment**
- Quiz on pointers, interfaces, embedding, modules.
- Code review of the library manager.
- **Phase 1 review:** 20-question cumulative quiz covering Weeks 1–4.

**Phase 1 deliverable:** a working CLI library manager with multi-package architecture.

---

# 📚 Phase 2: Intermediate Go and Software Engineering (Weeks 5–7)

## Week 5 — Object-Oriented Design in Go

**Learning objectives**
- Understand Go's take on OOP.
- Design reusable, composable components.
- Apply SOLID principles idiomatically.

**Topics**
- Structs + methods + interfaces as Go's OOP
- Composition over inheritance
- Encapsulation via exported/unexported identifiers
- Interfaces as contracts at boundaries
- SOLID in Go (with pragmatic caveats)
- Small interfaces (`io.Reader`, `io.Writer` as exemplars)

**Exercises**
1. Refactor Week 4 library manager to use interfaces and injection.
2. Define a `Storage` interface; implement in-memory and file-based versions.
3. Write a `Stringer` interface implementation (`String() string`).
4. **Mini-Project: Enhanced Library Manager** with pluggable storage backends.

**Assessment**
- Quiz on composition, interfaces, encapsulation.
- Code review of the refactor.
- Design discussion: when *not* to use an interface.

---

## Week 6 — Error Handling, Testing, and Debugging

**Learning objectives**
- Handle errors idiomatically.
- Create, wrap, and inspect errors.
- Use `panic`/`recover` appropriately.
- Write unit tests with table-driven style.
- Debug Go programs.

**Topics**
- `error` as a value; `errors.New`, `fmt.Errorf` with `%w`
- `errors.Is`, `errors.As`, `errors.Unwrap`
- Sentinel errors and custom error types
- `panic` and `recover` — when (rarely) appropriate
- `testing` package: `TestXxx`, `t.Run`, `t.Helper`
- Table-driven tests
- `go test`, `-v`, `-cover`, `-run`, `-race`
- Debugging: `log`, `fmt.Printf`, `dlv` basics

**Exercises**
1. Wrap errors through a call chain; inspect with `errors.Is`.
2. Custom error type with metadata.
3. Write table-driven tests for the unit converter from Week 1.
4. Test the todo CLI logic (Week 3) with mocks.
5. **Mini-Project: Tested Unit Converter** — refactor `unitconv` into a testable function + `main`, with table-driven tests and >80% coverage.

**Assessment**
- Quiz on error wrapping, `errors.Is/As`, table-driven tests.
- Code review of the tested converter.
- Run `go test -cover` and report coverage.

---

## Week 7 — Git, GitHub, and Professional Workflows

**Learning objectives**
- Use Git confidently.
- Publish projects on GitHub.
- Maintain a professional repo.

**Topics**
- Git fundamentals: `init`, `add`, `commit`, `status`, `log`, `diff`
- Branches, merges, conflicts
- GitHub: remotes, push/pull, PRs
- `.gitignore` for Go projects (binaries, `vendor/` if used)
- README structure
- `gofmt`, `go vet`, and CI-ready habits
- Conventional commits (light introduction)

**Exercises**
1. `git init` all Week 1–6 projects; make first commits.
2. Create a GitHub repo; push one project.
3. Create a branch, make a change, merge it.
4. Resolve a deliberate merge conflict.
5. Write a README for your best project.
6. **Mini-Project: Portfolio Repo** — one GitHub repository `go-academy` with all Week 1–7 work organized in folders and a top-level README.

**Assessment**
- Walkthrough: explain your Git history and README.
- Verify `git log` shows clean commits.
- **Phase 2 review:** 20-question cumulative quiz.

**Phase 2 deliverable:** GitHub repo `go-academy` with all work committed.

---

# 📚 Phase 3: HTTP and Backend Foundations (Weeks 8–10)

## Week 8 — HTTP and Client-Server Architecture

**Learning objectives**
- Understand how HTTP works.
- Understand REST fundamentals.
- Handle JSON.

**Topics**
- Internet basics: DNS, IP, TCP, TLS, HTTP
- Request/response lifecycle
- HTTP methods: GET, POST, PUT, PATCH, DELETE
- Status codes (1xx–5xx, common ones)
- Headers, bodies, content types
- JSON: `encoding/json`, `Marshal`, `Unmarshal`, struct tags
- REST design principles

**Exercises**
1. Use `curl` to explore HTTP: headers, methods, status codes.
2. Parse a JSON string into a struct.
3. Marshal a struct to JSON with tags.
4. Design a REST API for a todo app on paper (endpoints, methods, status codes).
5. **Mini-Project: JSON CLI Client** — a Go program that calls a public API (`https://api.github.com/users/...` or similar), decodes JSON, prints results. *Introduces `net/http` client.*

**Assessment**
- Quiz on HTTP methods, status codes, JSON.
- Code review of the JSON client.

---

## Week 9 — Building Web Servers with `net/http`

**Learning objectives**
- Build HTTP servers with the standard library.
- Understand handlers and middleware.
- Parse requests, return responses.

**Topics**
- `http.ListenAndServe`
- `http.HandleFunc`, `http.Handler`, `ServeMux`
- `http.ResponseWriter`, `*http.Request`
- Path routing (`/users/{id}` in Go 1.22+)
- Query parameters, form values
- Parsing JSON request bodies
- Writing JSON responses
- Middleware (function composition)
- Request validation and error responses
- Graceful shutdown (`http.Server.Shutdown`)

**Exercises**
1. "Hello, World" web server.
2. Echo server — returns the request path and query.
3. JSON echo — parse body, return it with a field added.
4. Middleware for logging requests.
5. Routing with Go 1.22+ path parameters.
6. **Mini-Project: Todo REST API (`net/http`)** — CRUD endpoints with in-memory storage. No database yet. Includes request validation and proper status codes.

**Assessment**
- Quiz on handlers, middleware, request parsing.
- Code review of the todo API.
- `curl` walkthrough of every endpoint.

---

## Week 10 — Concurrency and Backend Fundamentals

**Learning objectives**
- Use goroutines and channels safely.
- Prevent race conditions.
- Use `context` for cancellation.

**Topics**
- Goroutines (`go func()`)
- Channels (unbuffered, buffered), `close`
- `select` statement
- `sync.WaitGroup`, `sync.Mutex`
- Race conditions; `go test -race`
- `context.Background`, `WithCancel`, `WithTimeout`, `WithDeadline`
- Request contexts (`r.Context()`)
- Goroutine leaks
- Worker pools

**Exercises**
1. Spawn goroutines with `WaitGroup`.
2. Producer/consumer with channels.
3. `select` across two channels.
4. Mutex-protected counter; run with `-race`.
5. Timeout with `context.WithTimeout`.
6. **Mini-Project: Concurrent URL Checker** — given N URLs, fetch them concurrently with a worker pool, timeout, and aggregated results.

**Assessment**
- Quiz on goroutines, channels, `context`, `sync`.
- Code review of the URL checker.
- **Phase 3 review:** 20-question cumulative quiz.

**Phase 3 deliverable:** working `net/http` todo API + concurrent URL checker.

---

# 📚 Phase 4: Databases and Data Persistence (Weeks 11–13)

## Week 11 — SQL and Relational Database Fundamentals

**Learning objectives**
- Understand relational databases.
- Write SQL confidently.

**Topics**
- Tables, rows, columns, primary keys, foreign keys
- Constraints (`NOT NULL`, `UNIQUE`, `CHECK`)
- Relationships (1-1, 1-N, N-N)
- SQL data types (PostgreSQL)
- CRUD (`INSERT`, `SELECT`, `UPDATE`, `DELETE`)
- Filtering (`WHERE`, `LIKE`, `IN`, `BETWEEN`)
- Sorting (`ORDER BY`), pagination (`LIMIT`/`OFFSET`)
- Joins (`INNER`, `LEFT`, `RIGHT`)
- Normalization (1NF–3NF)
- Transactions (`BEGIN`, `COMMIT`, `ROLLBACK`)
- PostgreSQL setup locally (Docker recommended)

**Exercises**
1. Install PostgreSQL (Docker) and connect with `psql`.
2. Create a schema: `users`, `posts`, `comments`.
3. Insert sample data; run basic queries.
4. Write joins to fetch posts with author info and comment counts.
5. Use a transaction for a multi-step operation.
6. **Mini-Project: SQL Schema + Queries** — design and populate a blog schema; write 15 queries of increasing complexity.

**Assessment**
- Quiz on SQL syntax, joins, normalization, transactions.
- Review of the schema and queries.

---

## Week 12 — PostgreSQL Integration with Go

**Learning objectives**
- Connect Go to PostgreSQL.
- Use `database/sql` idiomatically.

**Topics**
- `database/sql` and drivers (`github.com/jackc/pgx/v5/stdlib` or `lib/pq`)
- `sql.Open`, `db.Ping`, `db.Close`
- Connection pooling (`SetMaxOpenConns`, `SetMaxIdleConns`, `SetConnMaxLifetime`)
- Executing queries: `QueryRow`, `Query`, `Exec`
- Scanning rows (`rows.Scan`, `rows.Next`, `rows.Err`)
- Parameterized queries (`$1`, `$2`) — SQL injection prevention
- Prepared statements
- Handling `sql.ErrNoRows`
- Struct scanning (manual, pre-ORM)
- Environment variables for config

**Exercises**
1. Connect to Postgres; verify with `db.Ping`.
2. Create a `users` table; insert with parameterized queries.
3. Query a user by ID; handle `ErrNoRows`.
4. Query all users into a slice of structs.
5. Update and delete users safely.
6. **Mini-Project: User CRUD CLI with Postgres** — refactor Week 3 todo or the user CLI to persist to Postgres.

**Assessment**
- Quiz on `database/sql`, connection pooling, parameterized queries.
- Code review of the CLI.
- Security check: confirm no string concatenation in SQL.

---

## Week 13 — Database Design, Migrations, and Data Access

**Learning objectives**
- Design schemas for real applications.
- Use migrations.
- Understand repository/service patterns.

**Topics**
- Schema design patterns for real apps
- Foreign keys and cascade behavior
- Migrations: `golang-migrate/migrate` (or `goose`)
- Versioned migration files (`up`/`down`)
- Transactions in Go (`db.Begin`, `tx.Commit`, `tx.Rollback`)
- Repository pattern (`UserRepository` interface + Postgres impl)
- Service layer (business logic)
- Pagination and filtering (server-side)
- Indexes and query performance (`EXPLAIN ANALYZE`)
- ORMs: GORM introduction (when to use vs raw SQL)

**Exercises**
1. Set up `golang-migrate` for a new project.
2. Write `up`/`down` migrations for a blog schema.
3. Implement a `UserRepository` with raw SQL.
4. Add pagination to a list query.
5. Add an index and measure performance with `EXPLAIN ANALYZE`.
6. **Mini-Project: Blog Backend (DB layer only)** — Postgres schema with migrations, repository + service layers for `users` and `posts`. No HTTP yet — test with CLI or Go tests.

**Assessment**
- Quiz on migrations, transactions, repository pattern, indexes.
- Code review of the blog backend.
- **Phase 4 review:** 20-question cumulative quiz.

**Phase 4 deliverable:** working blog schema + migrations + repository/service layers.

---

# 📚 Phase 5: Gin and REST API Development (Weeks 14–17)

## Week 14 — Introduction to Gin

**Learning objectives**
- Install and use Gin.
- Build a simple REST API with Gin.

**Topics**
- `go get github.com/gin-gonic/gin`
- Gin project structure (`cmd/`, `internal/handlers/`, `internal/services/`, `internal/repositories/`)
- Routes and route groups (`r.Group("/api/v1")`)
- Path parameters (`c.Param`), query params (`c.Query`)
- JSON binding (`c.ShouldBindJSON`)
- JSON responses (`c.JSON`)
- HTTP status codes with Gin
- Middleware (`r.Use`, `gin.Logger`, custom)
- Centralized error handling (`c.Error`, middleware to format errors)

**Exercises**
1. Gin "Hello, World".
2. CRUD with in-memory storage.
3. Route groups for `/api/v1/users`.
4. JSON binding + validation (`binding:"required"`).
5. Custom middleware (request ID or auth stub).
6. **Mini-Project: Refactor Todo API to Gin** — take Week 9's `net/http` todo API and rewrite it in Gin. Compare code size and clarity. Write a short note: what did Gin give you, and what did it hide?

**Assessment**
- Quiz on Gin basics, binding, middleware.
- Code review of the refactor.
- Reflective note: `net/http` vs Gin.

---

## Week 15 — Building Production-Style REST APIs

**Learning objectives**
- Design clean, consistent REST APIs.
- Structure larger codebases.

**Topics**
- REST design revisited (nouns, verbs, idempotency)
- Request validation (`go-playground/validator` via Gin)
- Consistent response envelope (`{data, error, meta}`)
- Pagination, filtering, sorting in APIs
- API versioning (`/api/v1`, `/api/v2`)
- Config management (env vars via `os.Getenv`, `godotenv`, or `envconfig`)
- Structured logging (`log/slog` — standard library)
- Graceful shutdown with Gin (`http.Server` + context)
- Layered architecture: handler → service → repository
- Error handling: domain errors mapped to HTTP status

**Exercises**
1. Add request validation to the todo API.
2. Implement a consistent response envelope.
3. Add pagination/filtering to a list endpoint.
4. Add `/api/v1` prefix and plan for `/v2`.
5. Add structured logging with `log/slog`.
6. Add graceful shutdown.
7. **Mini-Project: Production Todo API (Gin + Postgres)** — full stack: Gin handlers, service layer, Postgres repository, validation, pagination, structured logging, graceful shutdown. This is a *serious* project and will be a portfolio centerpiece.

**Assessment**
- Quiz on validation, response design, logging, config.
- Code review of the production todo API.
- Load testing with `hey` or `wrk` (briefly).

---

## Week 16 — Authentication and Authorization

**Learning objectives**
- Implement secure auth.
- Protect endpoints.
- Understand JWT and sessions.

**Topics**
- Authentication vs authorization
- User registration (email uniqueness, validation)
- Password hashing (`golang.org/x/crypto/bcrypt`, or `argon2`)
- Session-based auth (cookies) — overview
- JWT: structure, signing (`golang-jwt/jwt/v5`), claims, expiry
- Auth middleware (JWT validation)
- Role-based access control (RBAC)
- Permissions design
- Protecting routes
- Common vulnerabilities (broken auth, weak hashing, JWT misconfig)

**Exercises**
1. Implement `/register` and `/login`.
2. Hash passwords with bcrypt; verify on login.
3. Issue a JWT on login; middleware to validate.
4. Add RBAC (`admin`, `user`) to routes.
5. Write tests for auth middleware.
6. **Mini-Project: Auth-Enabled Todo API** — add users, JWT auth, per-user todos. Update the production todo API from Week 15.

**Assessment**
- Quiz on auth vs authz, JWT, bcrypt, RBAC.
- Code review with security focus.
- Security checklist walkthrough.

---

## Week 17 — API Security and Reliability

**Learning objectives**
- Harden the API.
- Document it.
- Test security.

**Topics**
- Input validation and sanitization
- SQL injection prevention (recap)
- CORS (`gin-contrib/cors`)
- Rate limiting (`golang.org/x/time/rate` or middleware)
- Secure config and secrets management
- Request timeouts (`http.Server` config + context)
- Context propagation
- Secure error responses (no stack traces to clients)
- Dependency vulnerabilities (`govulncheck`)
- OpenAPI/Swagger (`swaggo/swag` for Gin)
- Security testing (manual checklist + basic penetration)

**Exercises**
1. Add CORS.
2. Add rate limiting (e.g., 10 req/s per IP).
3. Add request timeouts.
4. Run `govulncheck ./...` and fix findings.
5. Generate Swagger docs with `swag init`.
6. Security-test the auth flow (bad tokens, expired tokens, missing auth).
7. **Mini-Project: Secured + Documented Todo API** — add CORS, rate limiting, timeouts, Swagger. Final version of the todo API before it becomes portfolio project 1.

**Assessment**
- Quiz on security hardening, rate limiting, OpenAPI.
- Code review of the final todo API.
- **Phase 5 review:** 20-question cumulative quiz.

**Phase 5 deliverable:** production-ready, secured, documented todo API with JWT auth, Postgres, Swagger.

---

# 📚 Phase 6: Advanced Backend Engineering (Weeks 18–20)

## Week 18 — Advanced Testing and Code Quality

**Learning objectives**
- Write unit, integration, and HTTP handler tests.
- Use `httptest`.
- Run benchmarks and fuzz tests.
- Set up CI.

**Topics**
- Unit testing (recap)
- Integration tests with Postgres (testcontainers-go or Dockerized test DB)
- `httptest.NewRecorder`, `httptest.NewServer`
- Testing Gin endpoints
- Mocking with interfaces (avoid heavy mocking frameworks)
- Database testing (fixtures, cleanup, transactional tests)
- Test coverage (`go test -cover`, `-coverprofile`)
- Benchmarks (`BenchmarkXxx`, `go test -bench=.`)
- Fuzz testing (`FuzzXxx`)
- GitHub Actions: Go test workflow (`.github/workflows/test.yml`)

**Exercises**
1. Test a Gin handler with `httptest.NewRecorder`.
2. Integration test against a real Postgres (Docker).
3. Benchmark a critical function.
4. Write a fuzz test for a parser.
5. Set up GitHub Actions to run `go test ./...`.
6. **Mini-Project: Fully Tested Todo API** — bring test coverage above 80%, add integration tests, benchmarks, and CI workflow.

**Assessment**
- Quiz on `httptest`, integration testing, CI.
- Coverage report walkthrough.
- Review GitHub Actions output.

---

## Week 19 — Application Architecture and Maintainability

**Learning objectives**
- Structure larger Go applications.
- Apply clean architecture carefully.
- Refactor responsibly.

**Topics**
- Organizing larger apps (`cmd/`, `internal/`, `pkg/`, `api/`)
- Clean architecture fundamentals (domain, use cases, adapters)
- Dependency injection via constructors (no DI framework)
- Interfaces at boundaries
- Service/repository separation (recap + refinement)
- Configuration via structs (`envconfig`)
- Logging and observability (`log/slog` + request IDs)
- Avoiding premature abstraction (YAGNI)
- Refactoring and technical debt

**Exercises**
1. Refactor the todo API into a clean architecture layout.
2. Introduce a domain layer with entity methods.
3. Add request IDs throughout (middleware + logger).
4. Write a "decisions" document (`DECISIONS.md`) explaining architectural choices.
5. **Mini-Project: Refactored Todo API** with clean architecture and a written architecture doc.

**Assessment**
- Quiz on clean architecture, DI, observability.
- Architecture walkthrough: explain your folder layout.
- Review of `DECISIONS.md`.

---

## Week 20 — Background Jobs and External Integrations

**Learning objectives**
- Send HTTP requests from Go reliably.
- Integrate third-party APIs.
- Run background jobs.
- Handle uploads.

**Topics**
- `net/http` client (recap), timeouts, retries, exponential backoff
- Third-party API integration patterns
- Background workers (`goroutine` + queue)
- Scheduled jobs (`time.Ticker`, or `robfig/cron`)
- Email notifications (SMTP via `net/smtp`, or a provider SDK)
- File uploads (multipart forms) and storage (local or S3-compatible)
- Message queues (intro: RabbitMQ / Kafka via overview only)
- Reliable task processing (idempotency, retries)

**Exercises**
1. HTTP client with timeout and retry.
2. Exponential backoff helper.
3. Background worker to process a queue.
4. Scheduled daily job (log something).
5. Send an email (using a test SMTP or provider).
6. File upload endpoint with validation.
7. **Mini-Project: Notifier Service** — a small service with: an HTTP endpoint that queues a "send email" job; a background worker that processes the queue; retries with backoff; logging.

**Assessment**
- Quiz on HTTP client patterns, retries, background jobs.
- Code review of the notifier service.
- **Phase 6 review:** 20-question cumulative quiz.

**Phase 6 deliverable:** notifier service + refactored clean-architecture todo API.

---

# 📚 Phase 7: Deployment and Portfolio Projects (Weeks 21–24)

## Week 21 — Deployment and Production Operations

**Learning objectives**
- Build production Go binaries.
- Containerize with Docker.
- Deploy and operate.

**Topics**
- Building Go binaries for Linux (`GOOS`, `GOARCH`, `-ldflags`)
- Static binaries; CGO off
- Linux server basics (SSH, users, systemd)
- Environment-based config (production vs dev)
- Docker fundamentals (`Dockerfile`, multi-stage builds)
- Docker Compose (app + Postgres)
- Deploying to a VPS (DigitalOcean, Hetzner, Fly.io, Render — pick one)
- Connecting production Postgres
- HTTPS (Caddy or Let's Encrypt)
- Health checks (`/healthz`)
- Graceful shutdown (recap)
- Logging and monitoring (slog to stdout; basic uptime monitoring)

**Exercises**
1. Cross-compile the todo API for Linux.
2. Write a multi-stage `Dockerfile` (small final image).
3. `docker-compose.yml` with app + Postgres.
4. Deploy to a VPS.
5. Add `/healthz`.
6. Set up a systemd service (or use Docker restart policies).
7. **Mini-Project: Deployed Todo API** — the production todo API, running publicly with HTTPS.

**Assessment**
- Quiz on Docker, deployment, health checks.
- Walkthrough of the live deployment.
- Review of Dockerfile and Compose file.

---

## Week 22 — Portfolio Project 1: Task Management API

**Goal:** Ship a portfolio-quality task management API.

**Requirements**
- User registration and login
- Secure password storage (bcrypt/argon2)
- JWT authentication + authorization
- CRUD for tasks (create, list, update, delete)
- Task fields: title, description, status, due date, priority
- PostgreSQL with migrations
- Request validation
- Pagination, filtering, sorting
- Automated tests (unit + integration)
- OpenAPI/Swagger docs
- Docker deployment
- README (setup, endpoints, architecture, deployment)
- GitHub repo with clean history

**Exercises**
- Day 1: Project scaffold + DB schema + migrations.
- Day 2: Auth endpoints (register, login, JWT).
- Day 3: Task CRUD.
- Day 4: Filtering/pagination.
- Day 5: Tests.
- Day 6: Docker + deployment.
- Day 7: Documentation and polish.

**Assessment**
- Full project review.
- Deployment walkthrough.
- README quality check.

---

## Week 23 — Portfolio Project 2: Social Media Backend

**Goal:** Ship a second portfolio project showing different skills.

**Requirements**
- User profiles
- Posts and comments
- Likes
- Follow relationships (N-N)
- Authentication and permissions
- Complex DB relationships
- Pagination and filtering (feed)
- File upload (avatars or post images)
- Database migrations
- Robust error handling
- Automated tests
- API documentation (Swagger)
- Docker deployment
- README

**Exercises**
- Day 1: Schema design + migrations (users, posts, comments, likes, follows).
- Day 2: Auth + profiles.
- Day 3: Posts + comments.
- Day 4: Likes + follows.
- Day 5: Feed with pagination.
- Day 6: File uploads.
- Day 7: Tests, docs, deploy.

**Assessment**
- Full project review.
- Compare architecture to Project 1 — what improved?
- README quality check.

---

## Week 24 — Final Capstone, Deployment, and Career Preparation

**Goal:** Ship a production-style capstone and prepare for the job market.

**Capstone requirements**
- Idiomatic Go
- Well-organized architecture
- Gin or `net/http`
- PostgreSQL integration
- Database migrations
- Authentication and authorization
- Secure API design
- Automated tests
- API documentation
- Environment-based configuration
- Docker-based deployment
- Logging and error handling
- Professional GitHub repo
- Clear README (setup, architecture, endpoints, deployment)

**Suggested capstone ideas**
- Booking system (rooms, reservations)
- URL shortener (with analytics)
- Job board (companies, jobs, applications)
- Habit tracker (habits, check-ins, streaks)

**Career preparation topics**
- Resume: how to present Go projects
- GitHub profile polish
- Explaining your projects in interviews
- Common Go interview questions
- Coding interview practice with Go
- What to learn next (see "After Week 24" below)

**Exercises**
- Day 1–2: Final capstone scaffold + schema.
- Day 3–4: Core features.
- Day 5: Tests and docs.
- Day 6: Deploy.
- Day 7: Resume + portfolio polish.

**Assessment**
- **Final skills assessment:** 40-question cumulative quiz.
- Capstone project review.
- Career readiness review: strengths, weaknesses, next steps.
- **Program complete:** you have 3 deployed, documented, tested Go projects on GitHub.

---

# 🎯 After Week 24 — What to Learn Next

**To be job-ready, continue with:**
- **System design** (scaling, caching, message queues, load balancing)
- **Advanced concurrency** (patterns: worker pools, pipelines, fan-in/out)
- **gRPC and Protocol Buffers**
- **Event-driven architecture** (Kafka, NATS)
- **Kubernetes** basics (after Docker)
- **Observability** (Prometheus, Grafana, OpenTelemetry)
- **Database tuning** (indexes, query plans, partitioning)
- **Go performance profiling** (`pprof`, flame graphs)
- **A second language or domain** — depending on target roles (TypeScript, Rust, or specialized backend domains)

---

# 📊 Week-by-Week Projects Summary

| Week | Project | Focus |
|---|---|---|
| 1 | CLI Unit Converter | Basics, types, fmt |
| 2 | Number Guessing CLI | Control flow |
| 3 | Todo CLI (in-memory) | Functions, slices, maps |
| 4 | Library Manager | Structs, interfaces, packages |
| 5 | Library Manager (enhanced) | OOP, composition, DI |
| 6 | Tested Unit Converter | Error handling, tests |
| 7 | Portfolio Repo | Git, GitHub |
| 8 | JSON CLI Client | HTTP client, JSON |
| 9 | Todo API (`net/http`) | HTTP servers |
| 10 | Concurrent URL Checker | Goroutines, channels |
| 11 | Blog Schema + Queries | SQL |
| 12 | User CRUD CLI + Postgres | `database/sql` |
| 13 | Blog Backend (
