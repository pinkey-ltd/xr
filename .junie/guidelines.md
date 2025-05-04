# XR Project Development Guidelines

This document provides essential information for developers working on the XR project.

## Build/Configuration Instructions

### Prerequisites

- Go 1.24.2 or later
- Git

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/xr.git
   cd xr
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Build the project:
   ```bash
   go build ./...
   ```

## Testing Information

### Running Tests

To run all tests in the project:

```bash
go test ./...
```

To run tests with verbose output:

```bash
go test -v ./...
```

To run tests in a specific package:

```bash
go test -v ./go3d/vec2
```

To run a specific test:

```bash
go test -v ./go3d/vec2 -run TestNormal
```

### Adding New Tests

1. Create a test file with the naming convention `*_test.go` in the same package as the code you're testing.
2. Import the testing package and any assertion libraries (this project uses `github.com/stretchr/testify/assert`).
3. Write test functions with the naming convention `TestXxx` where `Xxx` is the function you're testing.
4. Use table-driven tests for comprehensive test coverage.

### Example Test

Here's a simple example of a test file structure:

```go
package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEven(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Zero", 0, true},
		{"Positive Even", 2, true},
		{"Positive Odd", 3, false},
		{"Negative Even", -4, true},
		{"Negative Odd", -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEven(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
```

## Code Style and Development Guidelines

### Code Organization

The project is organized into several packages:

- `go3d`: Contains 3D math utilities (vectors, matrices, etc.)
  - `vec2`: 2D vector operations
  - `vec3`: 3D vector operations
  - `mat2`, `mat3`, `mat4`: Matrix operations
  - Other geometric utilities
- `mst`: Mesh and 3D model handling
- `utils`: General utility functions

### Coding Conventions

1. **Naming**:
   - Use camelCase for variable and function names
   - Use PascalCase for exported functions, types, and variables
   - Use snake_case for file names

2. **Documentation**:
   - Add comments for all exported functions, types, and variables
   - Use godoc-compatible comments (starting with the name of the element)

3. **Error Handling**:
   - Return errors rather than using panic
   - Check all errors and handle them appropriately

4. **Testing**:
   - Write tests for all exported functions
   - Use table-driven tests for comprehensive coverage
   - Aim for high test coverage, especially for critical components

### Common Development Tasks

1. **Adding a new feature**:
   - Create a new branch for your feature
   - Implement the feature with appropriate tests
   - Submit a pull request

2. **Fixing a bug**:
   - Create a test that reproduces the bug
   - Fix the bug
   - Verify the test passes

3. **Debugging**:
   - Use `go test -v` for verbose test output
   - Use `t.Logf()` in tests for debugging information
   - For more complex debugging, use the standard library's `log` package

## Known Issues and Limitations

- The Cross function in the vec2 package has an implementation issue that causes test failures
- The project is still in early development (version 0.0.1)