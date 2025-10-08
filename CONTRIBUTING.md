# Contributing to go-svg-captcha

Thank you for considering contributing to this project! :tada:

We welcome contributions from everyone. By participating in this project, you agree to abide by our code of conduct (be respectful, constructive, and collaborative).

## Table of Contents

- [Ways to Contribute](#ways-to-contribute)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Making Changes](#making-changes)
- [Code Style Guidelines](#code-style-guidelines)
- [Testing Guidelines](#testing-guidelines)
- [Commit Message Guidelines](#commit-message-guidelines)
- [Submitting Pull Requests](#submitting-pull-requests)
- [Reporting Issues](#reporting-issues)
- [Community Guidelines](#community-guidelines)
- [License](#license)

## Ways to Contribute

There are many ways to contribute to go-svg-captcha:

### 🐛 Report Bugs

Found a bug? [Open an issue](https://github.com/reu98/go-svg-captcha/issues/new) with:

- Clear title and description
- Steps to reproduce
- Expected vs actual behavior
- Go version and OS
- Code sample if applicable

### 💡 Suggest Features

Have an idea? [Open an issue](https://github.com/reu98/go-svg-captcha/issues/new) with:

- Clear description of the feature
- Use cases and benefits
- Possible implementation approach (optional)

### 📝 Improve Documentation

- Fix typos or clarify existing docs
- Add examples or tutorials
- Improve code comments
- Translate documentation

### 🔧 Submit Code

- Fix bugs
- Implement new features
- Improve performance
- Refactor code

### 💬 Help Others

- Answer questions in issues
- Review pull requests
- Share your experience

## Getting Started

### Prerequisites

- **Go 1.21.4 or higher** - [Download Go](https://golang.org/dl/)
- **Git** - [Install Git](https://git-scm.com/downloads)
- A GitHub account

### Fork and Clone

1. **Fork the repository** on GitHub by clicking the "Fork" button

2. **Clone your fork** locally:

   ```bash
   git clone https://github.com/YOUR_USERNAME/go-svg-captcha.git
   cd go-svg-captcha
   ```

3. **Add upstream remote** to keep your fork in sync:

   ```bash
   git remote add upstream https://github.com/reu98/go-svg-captcha.git
   ```

4. **Verify remotes**:
   ```bash
   git remote -v
   # origin    https://github.com/YOUR_USERNAME/go-svg-captcha.git (fetch)
   # origin    https://github.com/YOUR_USERNAME/go-svg-captcha.git (push)
   # upstream  https://github.com/reu98/go-svg-captcha.git (fetch)
   # upstream  https://github.com/reu98/go-svg-captcha.git (push)
   ```

## Development Setup

### Install Dependencies

```bash
# Download all dependencies
go mod download

# Verify dependencies
go mod verify
```

### Project Structure

```
go-svg-captcha/
├── captcha.go          # Main API functions
├── option.go           # Configuration structs
├── draw.go             # SVG rendering logic
├── math.go             # Math CAPTCHA logic
├── random.go           # Random generation utilities
├── constants.go        # Package constants
├── enums.go            # Enumerations
├── fonts/              # Embedded fonts
│   ├── Comismsh.ttf
│   └── fonts.go
├── examples/           # Usage examples
│   └── basic/
└── *_test.go          # Test files
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Running Benchmarks

```bash
# Run all benchmarks
go test -bench=. -benchmem

# Run specific benchmark
go test -bench=BenchmarkCreateByText -benchmem
```

### Running Examples

```bash
# Run the basic example
cd examples/basic
go run main.go
```

## Making Changes

### Create a Branch

Always create a new branch for your changes:

```bash
# Update your local main branch
git checkout main
git pull upstream main

# Create a new branch
git checkout -b feature/your-feature-name
# or
git checkout -b fix/issue-description
```

Branch naming conventions:

- `feature/description` - New features
- `fix/description` - Bug fixes
- `docs/description` - Documentation changes
- `refactor/description` - Code refactoring
- `test/description` - Test improvements
- `perf/description` - Performance improvements

### Make Your Changes

1. Write your code
2. Add/update tests
3. Update documentation
4. Run tests and linting

## Code Style Guidelines

### Go Best Practices

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `go fmt` to format your code
- Use `go vet` to check for suspicious constructs
- Use `golint` for additional style checks (optional)

### Formatting

```bash
# Format all Go files
go fmt ./...

# Check for suspicious constructs
go vet ./...
```

### Code Quality

- **Clear naming**: Use descriptive variable and function names
- **Comments**: Add GoDoc comments for exported functions, types, and constants
- **Error handling**: Always handle errors appropriately
- **Keep it simple**: Prefer simple, readable code over clever solutions
- **No magic numbers**: Use named constants instead
- **DRY principle**: Don't repeat yourself

### Example GoDoc Comment

```go
// CreateByText generates a text-based CAPTCHA with the given options.
// It returns a Result containing the CAPTCHA text and SVG data.
//
// Example:
//   result, err := captcha.CreateByText(captcha.OptionText{
//       Size: 6,
//       IsColor: true,
//   })
func CreateByText(option OptionText) (*Result, error) {
    // implementation
}
```

## Testing Guidelines

### Writing Tests

- **Test all new features**: Add tests for any new functionality
- **Test edge cases**: Consider boundary conditions and error cases
- **Table-driven tests**: Use table-driven tests for multiple scenarios
- **Test naming**: Use descriptive test names: `TestFunctionName_Scenario`
- **Test coverage**: Aim for >80% code coverage

### Example Test

```go
func TestCreateByText_WithCustomSize(t *testing.T) {
    tests := []struct {
        name     string
        size     uint8
        wantLen  int
    }{
        {"default", 0, 4},
        {"custom 6", 6, 6},
        {"custom 8", 8, 8},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := CreateByText(OptionText{Size: tt.size})
            require.NoError(t, err)
            require.Len(t, result.Text, tt.wantLen)
        })
    }
}
```

### Test Requirements

Before submitting a PR:

- [ ] All tests pass: `go test ./...`
- [ ] No test coverage decrease
- [ ] Tests are deterministic (no random failures)
- [ ] Tests run quickly (< 1s per test if possible)

## Commit Message Guidelines

We follow [Conventional Commits](https://www.conventionalcommits.org/) for clear git history.

### Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types

- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, semicolons, etc.)
- `refactor`: Code refactoring without changing functionality
- `perf`: Performance improvements
- `test`: Adding or updating tests
- `chore`: Maintenance tasks, dependency updates

### Examples

```bash
feat(captcha): add support for multiplication in math CAPTCHA

Implement multiplication operation for math CAPTCHA generation.
This expands the math challenge options beyond addition and subtraction.

Closes #42

---

fix(color): correct HSL to RGB conversion

The previous implementation had an off-by-one error in the color
conversion algorithm, causing incorrect colors in certain hue ranges.

Fixes #38

---

docs(readme): add troubleshooting section

Added common issues and their solutions to help users debug problems.
```

### Best Practices

- Use imperative mood: "add feature" not "added feature"
- Keep subject line under 50 characters
- Capitalize subject line
- No period at the end of subject line
- Separate subject from body with a blank line
- Wrap body at 72 characters
- Reference issues and PRs in the footer

## Submitting Pull Requests

### Before Submitting

Complete this checklist:

- [ ] Code follows the style guidelines
- [ ] Code is properly formatted (`go fmt`)
- [ ] No suspicious constructs (`go vet`)
- [ ] All tests pass (`go test ./...`)
- [ ] New tests added for new features
- [ ] Documentation updated (if needed)
- [ ] CHANGELOG.md updated (for notable changes)
- [ ] Commit messages follow conventions
- [ ] Branch is up to date with main

### Creating the Pull Request

1. **Push your branch** to your fork:

   ```bash
   git push origin feature/your-feature-name
   ```

2. **Create Pull Request** on GitHub

3. **Fill out the PR template** with:
   - Clear description of changes
   - Related issue numbers (e.g., "Fixes #123")
   - Type of change (bug fix, feature, etc.)
   - Testing performed
   - Screenshots (if UI-related)

### PR Title Format

Follow the same format as commit messages:

```
feat: add multiplication support for math CAPTCHA
fix: correct color generation algorithm
docs: improve API documentation
```

### Review Process

- Maintainers will review your PR
- Address any feedback or requested changes
- Keep the PR focused and reasonably sized
- Be patient and respectful

### After Approval

- Maintainers will merge your PR
- Your contribution will be credited
- Delete your branch (optional)

```bash
git branch -d feature/your-feature-name
git push origin --delete feature/your-feature-name
```

## Reporting Issues

### Bug Reports

Use this template:

````markdown
**Description**
Clear description of the bug.

**To Reproduce**
Steps to reproduce the behavior:

1. Create CAPTCHA with options '...'
2. Call function '...'
3. See error

**Expected Behavior**
What you expected to happen.

**Actual Behavior**
What actually happened.

**Environment**

- Go version: [e.g., 1.21.4]
- OS: [e.g., macOS 14.0, Ubuntu 22.04]
- go-svg-captcha version: [e.g., v1.1.0]

**Code Sample**

```go
// Minimal code to reproduce
```
````

**Additional Context**
Any other relevant information.

````

### Feature Requests

Use this template:

```markdown
**Feature Description**
Clear description of the feature.

**Use Case**
Why is this feature needed? What problem does it solve?

**Proposed Solution**
Your ideas on how to implement it (optional).

**Alternatives Considered**
Other approaches you've thought about.

**Additional Context**
Any other relevant information.
````

## Community Guidelines

### Code of Conduct

- **Be respectful**: Treat everyone with respect and kindness
- **Be constructive**: Provide helpful feedback and suggestions
- **Be collaborative**: Work together towards the best solution
- **Be patient**: Everyone was a beginner once
- **Be inclusive**: Welcome contributors from all backgrounds

### Communication

- Use clear, professional language
- Assume good intentions
- Ask questions if something is unclear
- Provide context in discussions
- Be responsive to feedback

### Recognition

All contributors will be recognized in:

- Git commit history
- GitHub contributors page
- Release notes (for significant contributions)

## Questions?

- 📚 Check the [documentation](https://pkg.go.dev/github.com/reu98/go-svg-captcha)
- 💬 Open a [discussion](https://github.com/reu98/go-svg-captcha/discussions)
- 🐛 Search [existing issues](https://github.com/reu98/go-svg-captcha/issues)
- 📧 Contact maintainers via GitHub

## License

By contributing, you agree that your contributions will be licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

---

Thank you for contributing to go-svg-captcha! :tada: :sparkles:

Your contributions make this project better for everyone. :heart:
