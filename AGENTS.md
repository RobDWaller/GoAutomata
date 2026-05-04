# Developer Guidelines

## Code Standards
- Keep changes small and explicit.
- Keep cyclomatic complexity low.
- Ensure tests are written for all new code.
- Use meaningful names for variables and methods.
- Preserve existing public behavior unless the change explicitly requires API updates.

## Testing
- Every behavior change must include or update tests.
- Prefer table-driven tests for multi-case logic.
- Before submitting code commits, run:
  - `go test -v ./...`
  - `go vet ./...`

## Linting
- Run `gofmt ./...` and keep code/lint output clean.
- Ensure spelling is correct in packages, modules, methods and comments.

## Development Process
- Ask the human user at least two relevant questions before you begin planning your changes to gain additional context.
- Read related code and tests before editing.
- Always confirm proposed changes with human user before implementation.
- Implement the smallest correct change first; avoid unrelated refactors.
- Always add tests for new code.
- If you are told to edit only a single file always check with human user before editing any other files.
- After code changes are made run tests and linting steps.
- If there is a CI pipeline, ensure it passes.
- Update inline code comments when new functionality is added.
- Update `README.md` when public behavior or usage changes.
- At the end of the process ask the human user if they want you to produce an Architecture Decision Record and Architecture Diagram.
  - ADRs and Architecture Diagrams should be stored in the `./.adr` directory.
  - ADRs should be concise, no more than 300 wors.
  - ADR files should have the timestamp in the file name.

## Pull Requests
- Keep commits/PRs focused: one logical change, tests included, checks passing.
- Commit messages should be descriptive but short.

