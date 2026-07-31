## ADDED Requirements

### Requirement: Hero creation succeeds when user has no existing hero

The hero creation use case SHALL proceed without error when the repository reports that no hero exists for the user.

#### Scenario: No existing hero found
- **WHEN** `CreateHero` is called for a user who has no hero in the repository
- **THEN** the use case SHALL proceed with hero creation without panicking or returning an error
- **AND** the new hero SHALL be persisted successfully

#### Scenario: Repository error during hero lookup
- **WHEN** `CreateHero` is called and the repository returns an error other than "not found"
- **THEN** the use case SHALL return an error to the caller
- **AND** the use case SHALL NOT panic

### Requirement: Hero creation rejects duplicate heroes safely

The hero creation use case SHALL reject creation when the user already has an active hero, without dereferencing a nil pointer.

#### Scenario: User already has an active hero
- **WHEN** `CreateHero` is called for a user who already has a hero
- **THEN** the use case SHALL return an error indicating the user already has a hero
- **AND** no new hero SHALL be persisted
