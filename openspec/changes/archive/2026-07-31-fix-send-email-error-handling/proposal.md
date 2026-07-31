## Why

`SendEmail` in `pkg/security/email_service.go` always returns `nil` regardless of whether the SMTP send succeeds or fails. When email delivery fails, the error is only printed to stdout, and the caller (user registration use case) proceeds as if the email was sent successfully. Users never receive their verification email but the system believes they did.

## What Changes

- `SendEmail` returns the SMTP error to the caller instead of swallowing it
- User creation will now fail with an error if email delivery fails
- Replace `fmt.Printf` debug statement with proper error return

## Capabilities

### New Capabilities

### Modified Capabilities
- `email-verification`: `SendEmail` must return an error when email delivery fails instead of silently returning `nil`

## Impact

- `internal/pkg/security/email_service.go`: Fix `SendEmail` to return SMTP errors
