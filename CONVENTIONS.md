# Go (Golang) Development Conventions

These conventions ensure a scalable, maintainable, and clean Go codebase that prioritizes simplicity, modularity, and readability. They integrate modern Go best practices, including concurrency, error management, and project organization.

---
## 0. Tech Stack

- **Build**: Docker Container - Use `makefile`
- **Backend:** Go 1.25 + [Gin Web Framework](https://gin-gonic.com/en/docs/) with [`html/template`](https://pkg.go.dev/html/template)
- **Frontend:** HTML, [HTMX](https://htmx.org/docs/), Vanilla JavaScript
- **Styling:** Tailwind CSS v4
- **Design:** Follow front-end guide-lines served from `FRONT-END.md`

---

## 1. Project Structure

- **Project architecture:** This project is structured using Vaughn Vernon Domain-driven design principals.
- **Prioritize uncoupled arquitecture** Use interfaces to serve functions
- **Interface layer:** This layer is used as "presentation layer". The <web> layer is responsible to handle the application front end.

---

## 2. Modularity & Simplicity

- **Single Responsibility:** Every file, type, and function should do one thing.
- **Short Functions:** Keep functions under 30 lines when possible.
- **Descriptive Names:** Use meaningful file, type, and function names (follow [Google Go standards](https://google.github.io/styleguide/go/decisions)).
- **No Printing/Direct Error Handling:** Never log or print errors except via centralized logging and error handling modules.

---

## 3. Concurrency

- Use goroutines and channels where suitable (for parallelism and asynchronous tasks).
- Avoid concurrency when it makes code less readable or more complex.
- Prefer vectorized computations (use slices and helper methods in `/pkg/vector`) over manual loops for data processing.
- Always document concurrent code for clarity.

---

## 4. Error Management

- **Centralize Errors:** Define all error types and helpers in `/errors/errors.go`.
- **Propagate Errors:** Always return errors to a single handling point, never handle or print errors directly in business logic.
- **Error Wrapping:** Use Go’s error wrapping (`fmt.Errorf("context: %w", err)`) for stack traces.
- **No Silent Failures:** Always check and return errors, never ignore them.

---

## 5. Code Quality

- **DRY:** Avoid duplication—use helpers or utility packages for repeated logic.
- **Readability:** Prefer clarity over cleverness. Add comments for complex logic.
- **Scalability:** Organize code into modules and packages so new features can be added without major refactoring.

---

## 6. Front-end

- [Read Front-end guideline](./FRONT-END.md)
---

## 6. References

- [Read this to know  more about the project](./README.md)
