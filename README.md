# Go-Do: Idiomatic Golang CLI Task Manager

[![Go Version](https://img.shields.io/github/go-mod/go-version/aliozdemir13/Go-Do)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go CI](https://github.com/aliozdemir13/Lumina-CLI/actions/workflows/ci.yaml/badge.svg)](https://github.com/aliozdemir13/Go-Do/actions)

**Go-Do** is a polished Command Line Interface (CLI) productivity tool built entirely in **Go (Golang)**. This project serves as a practical deep-dive into building modern terminal applications with **JSON persistence**, local storage, and high-contrast ANSI styling.

Designed with an "Educational First" mindset, this repository documents the transition from basic syntax to idiomatic Go project structures.

## Key Features

- **Local JSON Persistence:** Automatically saves and loads tasks from `tasks.json` using standard library encoding.
- **Modern Terminal UI:** High-contrast ANSI colors and progress bars for a premium CLI experience.
- **Automated Task Id Management:** Sequential Id assignment for easy completion and deletion.
- **Progress Tracking:** Real-time calculation of completion percentages and visual status bars.
- **Idiomatic Project Layout:** Follows the `internal/` package pattern for clean code encapsulation.

## Learning Journey & Technical Highlights

As a learning project, I used Go-Do to master several core Golang concepts. Here is how they are implemented:

### 1. Persistence via JSON Marshalling
Instead of using a heavy database, I utilized the `encoding/json` package. I learned how to use **struct tags** (e.g., ``json:"title"``) to control how Go data structures are converted into human-readable JSON files.

### 2. Pointer vs. Value Receivers
I explored Go's efficiency by using **Pointer Receivers** (`*TodoList`) for methods that modify data (like `Add` or `Complete`) and understanding when to use value receivers for read-only operations.

### 3. Slice Manipulation (The "Delete" Pattern)
Since Go doesn't have a built-in `remove` method for slices, I implemented the idiomatic way to delete items by reslicing and joining:
`l.Tasks = append(l.Tasks[:i], l.Tasks[i+1:]...)`

### 4. Robust Input Handling
I implemented `bufio.Scanner` for handling multi-word task titles and `strconv.Atoi` for type-safe integer conversion of user inputs.

## Architectural Decisions
- It is a conscious decision not to introduce a database at the current stage and use simple json file due to repo is dedicated for learning and personal use only. It can be introduced in the future for cloud-based multi device usage.

## Installation & Usage

### Prerequisites
- [Go 1.25+](https://go.dev/dl/)

### Setup for Windows
1. Clone the repository:
   ```bash
   git clone https://github.com/aliozdemir13/Go-Do.git
   cd Go-Do
   ```
2. Build the app:
   ```bash 
   go build
   ```
3. Run the app
    ```bash
   ./Go-Do
   ```
4. List of commands
    ```bash
    ./Go-Do --help
    ```

### Setup for Linux/MacOS
1. Clone the repository:
   ```bash
   git clone https://github.com/aliozdemir13/Go-Do.git
   cd Go-Do
   ```
2. Build the app:
   ```bash 
   make build
   ```
3. Run the app
    ```bash
   ./go-do
   ```
4. List of commands
    ```bash
    ./go-do --help
    ```

### Menu Commands
```bash
    Usage:
    go-do [flags]
    go-do [command]
    go-do [command] [flags]

    Available Commands:
    add               Add a new task
    complete          Mark a task as complete
    delete            Delete a task
    help              Help about any command
    list              Show tasks (open by default)
    list --done       Show closed tasks
```

### Project Structure
```bash
.
├── main.go            # Application entry point
├── main_test.go
├── tasks.json         # Local data storage (auto-generated)
├── Makefile
├── cmd/
│   ├── root.go        # Root cobra command + header/progress helpers
│   ├── root_test.go
│   ├── add.go
│   ├── complete.go
│   ├── delete.go
│   └── list.go
└── internal/
    └── todo/
        ├── todo.go       # Core logic and Task/TodoList structs
        ├── todo_test.go
        ├── io.go         # JSON read/write operations
        ├── io_test.go
        ├── style.go      # ANSI styling and progress bar logic
        └── style_test.go
```

### License

Distributed under the MIT License. See LICENSE for more information.

Follow my learning journey as I explore the power of Go!

## Contribution
- Every PR requires at least 90% code coverage for entire code base.
- Linter rules can be found inside .golangci.yaml
