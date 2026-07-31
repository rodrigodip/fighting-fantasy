## Why

`CreateHero` in `application/hero/usecase.go` discards the error from `FindByOwner` and unconditionally dereferences the returned pointer. When no hero exists for the user, `foundHero` is nil and `*foundHero` causes a runtime panic, crashing the request.

## What Changes

- `CreateHero` properly handles the `FindByOwner` error: if no hero is found, proceed with creation; if a hero is found, reject the request with the existing "already has a hero" logic.

## Capabilities

### New Capabilities

### Modified Capabilities
- `hero-create-error-feedback`: `CreateHero` must not panic when no hero exists — it must safely handle the "not found" case from the repository.

## Impact

- `internal/application/hero/usecase.go`: Fix nil pointer dereference in `CreateHero`
