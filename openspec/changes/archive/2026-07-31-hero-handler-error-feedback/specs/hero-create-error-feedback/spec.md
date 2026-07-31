## Purpose

Provides proper error handling and user feedback for the hero creation flow, ensuring users always receive a meaningful response instead of blank pages.

## ADDED Requirements

### Requirement: Hero creation redirects unauthenticated users

The handler SHALL redirect to `/` when the user context is missing.

#### Scenario: Missing user context
- **WHEN** a request to create a hero has no `user_id` in the context
- **THEN** the handler SHALL respond with a 303 redirect to `/`

### Requirement: Hero creation shows error on invalid input

The handler SHALL re-render the dashboard with an error message when the request body cannot be parsed.

#### Scenario: Invalid request binding
- **WHEN** the form data cannot be bound to the hero creation request
- **THEN** the handler SHALL re-render the dashboard with a `FeedBack.Error` message indicating invalid input

### Requirement: Hero creation shows error on creation failure

The handler SHALL re-render the dashboard with a user-friendly error message when hero creation fails.

#### Scenario: Hero creation returns an error
- **WHEN** the use case returns an error during hero creation
- **THEN** the handler SHALL re-render the dashboard with a `FeedBack.Error` message containing the parsed error

### Requirement: No debug logging in handler

The handler SHALL NOT contain `log.Println`, `log.Printf`, or `fmt.Print` debug statements.

#### Scenario: Handler code is clean of debug artifacts
- **WHEN** the handler code is inspected
- **THEN** no `log` or `fmt` print calls SHALL be present
