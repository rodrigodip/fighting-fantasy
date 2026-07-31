## 1. Hero service tests

- [x] 1.1 Create `internal/domain/hero/service_test.go` with table-driven tests for `ValidateInput` (empty name, short name, long name, empty potion, valid input)
- [x] 1.2 Add tests for `HasHero` (user has hero, user doesn't have hero)
- [x] 1.3 Add tests for `SelectPotion` (dexterity, strength, fortune, invalid potion)

## 2. User service tests

- [x] 2.1 Create `internal/domain/user/service_test.go` with table-driven tests for `ValidadeUserInput` (empty name, short name, empty email, invalid email, empty password, valid input)
- [x] 2.2 Add tests for `ValidatePassword` (too short, no lowercase, no uppercase, no digit, no special char, valid password)
- [x] 2.3 Add tests for `IsUserVerified` (verified, unverified)

## 3. Verify

- [x] 3.1 Run `go test ./internal/domain/...` and confirm all tests pass
