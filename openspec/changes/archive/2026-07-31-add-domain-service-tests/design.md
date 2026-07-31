## Context

The project has zero tests. Domain services (`hero.Service` and `usr.Service`) contain pure validation logic with no external dependencies, making them ideal candidates for the first test suite. See `proposal.md` for motivation.

## Goals / Non-Goals

**Goals:**
- Establish a testing pattern for the project
- Cover all public methods in domain services
- Use table-driven tests for clarity and maintainability

**Non-Goals:**
- Testing application use cases (requires mocking repositories)
- Testing infrastructure layer (requires database setup)
- Achieving 100% coverage (focus on critical paths)

## Decisions

### Decision 1: Table-driven tests

Use Go's idiomatic table-driven test pattern with `t.Run` for subtests. Each test case has a name, inputs, and expected output.

**Alternative considered:** Individual test functions per case. Rejected — table-driven tests reduce boilerplate and make it easy to add new cases.

### Decision 2: Standard library only

Use only `testing` package. No external assertion libraries (testify, etc.).

**Rationale:** Keeps dependencies minimal. The validation logic is simple enough that `if got != want` is clear.

### Decision 3: Test file placement

Place `service_test.go` alongside `service.go` in the same package. This allows testing unexported behavior if needed later.

## Risks / Trade-offs

- **No mocking framework** → Use cases will need mocks later. Mitigation: Can add a simple mock library when needed, or use interface-based fakes.
