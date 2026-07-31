## 1. Fix CreateHero nil pointer dereference

- [x] 1.1 Handle the error from `FindByOwner` — if error is returned, proceed with creation (no existing hero); if no error, call `HasHero` with the found hero to reject duplicates
- [x] 1.2 Remove the unconditional `*foundHero` dereference on line 24

## 2. Verify

- [x] 2.1 Run `go build ./...` to confirm the project compiles
