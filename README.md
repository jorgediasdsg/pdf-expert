# 📄 PDF Analyzer — Phase 2 (from script to structured service)

This project started in **Phase 1** as a deliberately coupled and naive implementation:

- a single `main.go`,
- one `/analyze` endpoint,
- direct file handling,
- direct PDF parsing,
- naive word counting,
- no layers, no tests, no abstractions.

The goal of Phase 1 was to **validate the PDF extraction approach** using `github.com/ledongthuc/pdf` with the simplest possible code.  
This validation is documented in `ADR/ADR001.md`.

Now, in **Phase 2**, we begin to move from “just works” to “maintainable”.

---

## 🎯 Phase 2 goals

Phase 2 focuses on **incremental refactoring**, not full clean architecture.

The goals are:

1. **Introduce a basic package structure**  
   - Move logic out of `main.go`.
   - Create minimal internal packages (e.g. `internal/pdf`, `internal/http` or similar).
   - Reduce coupling without over-engineering.

2. **Extract PDF analysis into a dedicated component**  
   - Separate HTTP concerns from parsing logic.
   - Make PDF parsing reusable.
   - Prepare ground for future interfaces and adapters.

3. **Improve HTTP handlers and JSON responses**  
   - Centralize response formatting.
   - Use consistent JSON structure for errors and success.
   - Make the API easier to consume and debug.

4. **Add basic unit tests**  
   - Start testing core logic (e.g. word counting, simple PDF wrapper).
   - Build habits for testability in later phases.

We still **do not** introduce worker pools, job queues, storage layers, or advanced patterns in this phase. Those remain future work.

---

## 🚦 Project phases (high-level roadmap)

### Phase 1 — Naive and coupled (completed)

- Single `main.go`.
- Direct `pdf.ToText` usage.
- Temporary file on disk.
- Naive `countWords` function.
- No packages, no layers, no tests.
- Purpose: Validate the PDF extraction library and establish a minimal end-to-end flow.

Details and rationale are in:

- `ADR/ADR001.md` — *Initial simple and coupled architecture to validate PDF extraction*.

---

### Phase 2 — Basic structure and separation (current)

- Split logic into small packages.
- Extract PDF analysis into its own component.
- Cleaner `main.go` that wires things together.
- More consistent HTTP responses.
- Introduce unit tests for core logic.

Documented in:

- `ADR/ADR002.md` — *Introduce basic package structure*  
- `ADR/ADR003.md` — *Extract PDF analysis into a dedicated component*  
- `ADR/ADR004.md` — *Improve HTTP handler and JSON response structure*  
- `ADR/ADR005.md` — *Add unit tests for core logic*

---

### Phase 3 — Toward clean architecture

Planned (not implemented yet):

- Domain-layer models.
- Application services (use cases).
- Ports and adapters (HTTP, storage, PDF).
- Configuration via environment variables.
- Structured logging.
- More complete test coverage.

---

### Phase 4 — Asynchronous processing and worker pool

Planned (not implemented yet):

- In-memory job queue (channels).
- Worker pool with N goroutines.
- Retry with backoff and per-job timeouts.
- Persistence (SQLite or Redis).
- Health and metrics endpoints.
- Docker and docker-compose setup.
- Extended ADR set documenting trade-offs.

---

## 📂 Directory structure (Phase 2 – target)

Planned structure for Phase 2 (still intentionally simple):

```text
pdf-analyzer/
  cmd/
    api/
      main.go                # Wire HTTP server and handlers
  internal/
    httpapi/
      handler.go             # HTTP handlers and routing
      response.go            # JSON response helpers
    pdfanalyzer/
      analyzer.go            # PDF text extraction and word counting
  ADR/
    ADR001.md
    ADR002.md
    ADR003.md
    ADR004.md
    ADR005.md
  Makefile
  README.md
  DEVELOPMENT.md
  go.mod
  go.sum
```

## 🧭 Design principles for this phase

Phase 2 follows these principles:

- Change one dimension at a time:
  - Only introduce structure that solves a real problem seen in Phase 1.

- Avoid premature abstraction:
  - No complex interfaces or patterns before they are needed.

- Prepare for Phase 3:
  - The refactoring here should make it easier, not harder, to later adopt clean/hexagonal architecture.
  
- Document decisions explicitly:
  - Every meaningful structural change is backed by an ADR in ADR/.