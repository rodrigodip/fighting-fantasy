## Purpose

Ensures email delivery failures are properly reported to callers, preventing silent failures where users are created but never receive verification emails.

## ADDED Requirements

### Requirement: Email delivery errors are reported

The email sending function SHALL return an error when SMTP delivery fails.

#### Scenario: SMTP send fails
- **WHEN** the SMTP server rejects the message or is unreachable
- **THEN** the function SHALL return an error to the caller
- **AND** the function SHALL NOT return `nil`

#### Scenario: SMTP send succeeds
- **WHEN** the SMTP server accepts the message
- **THEN** the function SHALL return `nil`

### Requirement: User creation fails when email delivery fails

The user registration use case SHALL fail with an error when email delivery fails.

#### Scenario: Email delivery fails during user creation
- **WHEN** a new user is created and the verification email fails to send
- **THEN** the use case SHALL return an error to the caller
- **AND** the user SHALL be informed that registration failed
