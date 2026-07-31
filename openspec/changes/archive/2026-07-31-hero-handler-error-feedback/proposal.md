## Why

The `HeroCreateHandler` uses debug `log` statements and silent `return`s for all error paths, leaving users with a blank HTTP response when something goes wrong. Every other handler in the project renders proper feedback via the `auth-feedback.html` partial—this one is the outlier.

## What Changes

- Replace all `log.Println`/`log.Printf` debug calls with proper error responses
- On missing user context: redirect to `/` (consistent with other handlers)
- On bind error: re-render dashboard with a user-friendly error message
- On hero creation failure: re-render dashboard with the parsed error message
- Remove the unused `"log"` import

## Capabilities

### New Capabilities
- `hero-create-error-feedback`: Proper error handling and user feedback for the hero creation flow

### Modified Capabilities

## Impact

- `internal/interface/web/handlers/hero.go`: Rewrite error paths to render feedback instead of logging
