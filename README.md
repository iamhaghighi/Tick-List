<div align="center">

# 📝 Todo on Terminal

**A fast, lightweight terminal-based todo manager built with Go.**

[![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=flat&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20macOS%20%7C%20Linux-lightgrey.svg)]()

TickList is a minimalist CLI todo application that lives in your terminal. Built with the [Bubble Tea](https://github.com/charmbracelet/bubbletea) TUI framework and powered by SQLite for persistence, it provides a clean, keyboard-driven interface for managing your tasks without ever leaving the command line.

</div>

---

## ✨ Features

- 📋 **Full CRUD** -- Create, read, update, and delete todos from a single screen
- ✅ **Toggle completion** -- Mark tasks as done with a single keystroke
- 💾 **Persistent storage** -- Todos are saved to a local SQLite database (no setup required)
- ⌨️ **Vim-friendly keybindings** -- Navigate with `j`/`k` or arrow keys
- 🖥️ **Clean TUI** -- Full-screen terminal interface with a responsive header and inline help
- 🕐 **Timestamps** -- Track when each todo was created and last updated
- ⚡ **Zero dependencies to install** -- Pure Go SQLite driver (no CGO required)

---

## 👀 Demo

### 📋 Todo List Screen

```
┌──────────────────────────────────────────────────────────────────────────────┐
│ Todo Management  A Fast, lightweight Todo        Go 1.26.4   v1.0.0        │
└──────────────────────────────────────────────────────────────────────────────┘

 ▶ [✓]  Buy groceries                              created_at: 14:30:00  updated_at: 14:30:00
   [ ]  Finish project documentation               created_at: 13:00:00  updated_at: 13:00:00
   [ ]  Review pull requests                       created_at: 12:00:00  updated_at: 12:00:00

 ↑/k move up   ↓/j move down   enter select   a add   e edit   delete delete   'q' quit
```

## 🚀 Getting Started

### 📋 Prerequisites

- **Go 1.26** or later

### 📥 Installation

Clone the repository and build:

```bash
git clone https://github.com/iamhaghighi/Tick-List.git
cd Todo_CLi
go mod download
```

### ▶️ Running

From the project root, run:

```bash
go run cmd/main.go
```

Or build a binary:

```bash
go build -o TickList.exe cmd/main.go
```

---

## 🎮 Usage

### 🧭 Navigation

| 🔑 Key | ⚡ Action |
|---|---|
| `↑` / `k` | ⬆️ Move cursor up |
| `↓` / `j` | ⬇️ Move cursor down |
| `enter` | 🔄 Toggle todo completion |
| `a` | ➕ Add a new todo |
| `e` | ✏️ Edit the selected todo |
| `delete` | 🗑️ Delete the selected todo |
| `y` | ✅ Confirm delete |
| `esc` | ⬅️ Cancel / go back |
| `q` / `ctrl+c` | 🚪 Quit |

### ➕ Creating a Todo

1. Press `a` to open the editor
2. Type your todo title
3. Press `enter` to save

### ✏️ Editing a Todo

1. Navigate to the todo you want to edit
2. Press `e` to open the editor
3. Update the title
4. Press `enter` to save

### 🗑️ Deleting a Todo

1. Navigate to the todo you want to delete
2. Press `delete`
3. Confirm by pressing `y` (or cancel with `esc`)

### 🔄 Toggling Completion

Navigate to a todo and press `enter` to flip its done state.

---

## 📂 Project Structure

```
todo_cli/
├── cmd/
│   └── main.go                          # 🚀 Entry point & CLI flags
├── internal/
│   ├── app/
│   │   ├── model.go                     # 🧩 Root Bubble Tea model
│   │   ├── state.go                     # 🖥️ Screen state (Todo / Editor)
│   │   ├── update.go                    # ⌨️ Global key handling & routing
│   │   └── view.go                      # 🎨 Screen renderer
│   ├── components/
│   │   ├── editor/
│   │   │   ├── model.go                 # ✏️ Editor state (Create / Edit / Delete)
│   │   │   ├── update.go                # ⌨️ Editor input handling
│   │   │   └── view.go                  # 🖼️ Editor UI
│   │   ├── header/
│   │   │   ├── header.go                # 📌 Header model & initialization
│   │   │   └── view.go                  # 📊 Responsive full-width header bar
│   │   ├── help/
│   │   │   └── help.go                  # ❓ Key bindings & help panel
│   │   └── todo/
│   │       ├── model.go                 # 📋 Todo list state & cursor
│   │       ├── update.go                # 🧭 Cursor movement
│   │       └── view.go                  # 🖼️ Todo list rendering
│   ├── constants/
│   │   └── constant.go                  # 🏷️ App name, version, description
│   ├── domain/
│   │   └── todo.go                      # 📦 Todo entity definition
│   ├── messages/
│   │   └── todo.go                      # 📨 Bubble Tea message types
│   ├── repository/
│   │   ├── todo_repository.go           # 🔌 Repository interface
│   │   ├── memory/
│   │   │   └── todo_repository.go       # 💭 In-memory implementation
│   │   └── sqlite/
│   │       └── sqlite_repository.go     # 🗄️ SQLite implementation
│   ├── service/
│   │   └── todo_service.go              # ⚙️ Business logic & validation
│   └── styles/
│       ├── color.go                     # 🎨 Color palette
│       └── style.go                     # 🖌️ Lipgloss style definitions
├── .gitignore
├── go.mod
└── go.sum
```

---

## 🏗️ Architecture

TickList follows a **layered architecture** inspired by clean architecture principles:

```
┌─────────────────────────────────────────────────┐
│              🖥️  UI Layer                       │
│  app/ ─► components/editor, todo, header, help  │
├─────────────────────────────────────────────────┤
│              ⚙️  Service Layer                   │
│          service/todo_service.go                 │
├─────────────────────────────────────────────────┤
│              🔌 Repository Layer                 │
│     repository/sqlite/ (SQLite implementation)   │
├─────────────────────────────────────────────────┤
│              📦 Domain Layer                     │
│              domain/todo.go                      │
└─────────────────────────────────────────────────┘
```

- 🖥️ **UI Layer** -- Built with Bubble Tea (Elm architecture). The root `app` model owns all sub-components and routes messages based on the active screen.
- ⚙️ **Service Layer** -- Handles validation and business logic. Wraps the repository with input checks before persisting.
- 🔌 **Repository Layer** -- Interface-based (`TodoRepository`). The active implementation uses SQLite via `modernc.org/sqlite` (pure Go, no CGO).
- 📦 **Domain Layer** -- Defines the `Todo` entity shared across all layers.

---

## 🛠️ Tech Stack

| 🔧 Component | 💡 Technology |
|---|---|
| 📝 Language | [Go](https://go.dev) 1.26 |
| 🖥️ TUI Framework | [Bubble Tea v2](https://github.com/charmbracelet/bubbletea) |
| 🧩 UI Components | [Bubbles v2](https://github.com/charmbracelet/bubbles) |
| 🖌️ Styling | [Lipgloss v2](https://github.com/charmbracelet/lipgloss) |
| 🗄️ Database | [SQLite](https://github.com/nicholasgasior/goinernsql) via [modernc.org/sqlite](https://modernc.org/sqlite) (pure Go) |
| 🏗️ Architecture | Elm architecture (MVU) with layered service/repository pattern |

---

## 🗄️ Database

TickList uses a local SQLite database stored at `todos.db` in the project root. The database is created automatically on first run.

### 📐 Schema

```sql
CREATE TABLE IF NOT EXISTS todos (
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    title      TEXT    NOT NULL,
    done       INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);
```

No migrations or manual setup needed -- the schema is applied automatically by the SQLite repository on startup.

## 🤝 Contributing

1. 🍴 Fork the repository
2. 🌿 Create a feature branch (`git checkout -b feature/my-feature`)
3. 💾 Commit your changes (`git commit -m 'feat: add my feature'`)
4. 🚀 Push to the branch (`git push origin feature/my-feature`)
5. 📬 Open a Pull Request

Please follow [Conventional Commits](https://conventional-emoji-commits.site/) for commit messages.

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---