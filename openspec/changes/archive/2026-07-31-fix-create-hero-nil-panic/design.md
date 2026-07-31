## Context

`CreateHero` in `application/hero/usecase.go:23` calls `FindByOwner` and discards the error, then unconditionally dereferences the returned pointer on line 24. When no hero exists, the repository returns a nil pointer and an error, causing a panic. See `proposal.md` for motivation.

## Goals / Non-Goals

**Goals:**
- Eliminate the nil pointer panic when no hero exists
- Preserve the existing "user already has a hero" rejection logic

**Non-Goals:**
- Refactoring the `HasHero` domain service (its design is questionable but works for now)
- Changing the repository interface or error types

## Decisions

### Decision 1: Use the repository error to determine hero existence

Instead of dereferencing the pointer and checking `HasHero`, use the error from `FindByOwner` as the signal:
- **Error returned** → no hero exists → proceed with creation
- **No error** → hero found → call `HasHero` to reject the duplicate

**Alternative considered:** Check `foundHero == nil` before dereferencing. Rejected because the error is the idiomatic Go signal and should not be ignored.

### Decision 2: Keep `HasHero` call for the "hero exists" path

When `FindByOwner` succeeds (no error), the existing `HasHero` check is called with the dereferenced hero. This preserves current behavior for the duplicate case.

## Risks / Trade-offs

- **Repository error ambiguity** → If `FindByOwner` returns an error for reasons other than "not found" (e.g., DB connection failure), we'd incorrectly proceed with creation. Mitigation: the `RegisterHero` call will also fail on DB issues, so no corrupt data is persisted. A proper "not found" error type would be a future improvement.
