## Context

`SendEmail` in `pkg/security/email_service.go:106-116` calls `smtp.SendMail` but always returns `nil` regardless of the result. When SMTP fails, the error is only printed via `fmt.Printf`, and the caller proceeds as if the email was sent. See `proposal.md` for motivation.

## Goals / Non-Goals

**Goals:**
- Propagate SMTP errors to the caller
- Fail user registration when email delivery fails

**Non-Goals:**
- Implementing email retries or queuing
- Refactoring the email template system (the TODO about learning `html/template`)
- Adding email delivery tracking or analytics

## Decisions

### Decision 1: Return SMTP error directly

Replace the `fmt.Printf` + `return nil` pattern with `return err`. This is the simplest fix and makes the error visible to the caller.

**Alternative considered:** Wrapping the error with context (e.g., `fmt.Errorf("email send failed: %w", err)`). Rejected for now to keep the change minimal; error wrapping can be added later if needed.

### Decision 2: No changes to the use case

The user creation use case already checks the error from `SendEmail`. Once `SendEmail` returns the error properly, the use case will automatically fail user creation. No changes needed to `application/user/usecase.go`.

## Risks / Trade-offs

- **User creation now fails on email issues** → Previously, users were created even if email failed. Now they won't be. This is the correct behavior, but it's a behavior change. Mitigation: This is the intended fix; silent failures are worse than explicit failures.
