## Context

The `HeroCreateHandler` is an HTMX endpoint. The hero creation form in `dashboard.html` uses `hx-post="/hero/create" hx-target="#dashboard-content" hx-swap="innerHTML"`, meaning error responses must return an HTML fragment (not a full page) that replaces the dashboard content. See `proposal.md` for motivation.

## Goals / Non-Goals

**Goals:**
- Every error path returns a meaningful HTML response or redirect
- Error messages use the existing `auth-feedback.html` partial rendered inside the dashboard layout
- Consistent with patterns established in `auth.go` and `pages.go`

**Non-Goals:**
- Adding a generic error container system (the TODO in the code mentions this, but it's a larger refactor)
- Changing the success path behavior
- Adding client-side validation

## Decisions

### Decision 1: Re-render dashboard content on error

On bind or creation errors, parse `dashboard.html` (without `base.html` layout) and execute the `"content"` template with `FeedBack` populated. This returns just the HTML fragment HTMX expects for `hx-target="#dashboard-content"`.

**Alternative considered:** Return only the `auth-feedback.html` partial. Rejected because the form lives inside the dashboard and HTMX would replace the dashboard content with just a toast, losing the form.

### Decision 2: Redirect on missing user context

Use `c.Redirect(303, "/")` for missing `user_id`, matching the pattern in `DashboardPageHandler` and `AdventurePageHandler`.

### Decision 3: Use `weberrors.ParseErrorForWeb` for use-case errors

The `CreateHero` use case may return domain errors. Use the existing `weberrors.ParseErrorForWeb()` to format them for display, consistent with `auth.go`.

## Risks / Trade-offs

- **Dashboard re-render needs user/hero data** → The error paths for bind/creation failures still need to fetch user and hero data to render the dashboard template. This means an extra DB call on error, but it's the cost of a consistent UX.
