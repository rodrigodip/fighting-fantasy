## Why

The project has zero test coverage. Domain services contain pure business logic that's ideal for unit testing — no dependencies, deterministic behavior, and clear validation rules. Adding tests now establishes a testing foundation before building the adventure engine.

## What Changes

- Add unit tests for `hero.Service`: `ValidateInput`, `HasHero`, `SelectPotion`
- Add unit tests for `usr.Service`: `ValidadeUserInput`, `ValidatePassword`, `IsUserVerified`
- Use table-driven tests (Go idiomatic pattern) for comprehensive coverage

## Capabilities

### New Capabilities

### Modified Capabilities

## Impact

- `internal/domain/hero/service_test.go`: New test file
- `internal/domain/user/service_test.go`: New test file
