## 1. Rewrite HeroCreateHandler error paths

- [x] 1.1 Replace missing user context path: remove `log.Println`, add `c.Redirect(303, "/")`
- [x] 1.2 Replace bind error path: remove `log.Println`, re-render dashboard content template with `FeedBack.Error`
- [x] 1.3 Replace creation error path: remove `log.Printf`, re-render dashboard content template with parsed error via `weberrors.ParseErrorForWeb`
- [x] 1.4 Remove `log.Printf("context keys: %v", c.Keys)` debug line
- [x] 1.5 Remove unused `"log"` import, add `"html/template"` and `weberrors` imports

## 2. Verify

- [x] 2.1 Run `go build ./...` to confirm the project compiles
