# ADR-019 — Handler Tests via Gin Test Framework

## Status
Accepted

## Context
After Clean Architecture activation, handlers became thin adapters responsible for:
- parsing HTTP input,
- converting to DTOs,
- calling use cases,
- mapping output to responses.

To ensure correct HTTP behavior, the system must test:
- status codes,
- JSON structure,
- error mapping,
- middleware behavior (request ID, JSON envelope).

Gin provides a built-in test framework (`httptest` + `gin.Default()`)
which allows isolated HTTP handler tests without running a real server.

## Decision
Adopt Gin Test Framework for testing HTTP handlers.

Key elements:
- use `httptest.NewRecorder()` for capturing output,
- use `gin.CreateTestContext()` to build request + context,
- inject mock use cases or ports for full isolation.

## Consequences

### Positive
- Tests verify real HTTP behavior.
- Handlers remain thin and easy to test.
- Guarantees correct JSON envelope and status codes.

### Negative
- Requires mocks of use cases.
- Coupling to Gin at the HTTP layer (acceptable since it's an adapter).

## Alternatives
A) Test endpoints via `net/http/httptest` manually  
Rejected — more verbose, wastes abstractions Gin already provides.

B) Skip handler tests  
Rejected — leads to fragile API behavior.
