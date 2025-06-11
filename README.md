
# 📝 CLI Task Manager in Go

A lightweight and efficient command-line task manager built using Go. This tool helps users maintain a to-do list directly from the terminal.

## 🔧 Features

- Add new tasks
- List all tasks
- Mark tasks as completed
- Delete tasks

## 🖥️ Commands

```bash
# Add a new task
go run main.go add "Your task here"

# List all tasks
go run main.go list

# Mark a task as completed
go run main.go do 1

# Delete a task
go run main.go del 1
```

## 💾 Data Storage

Tasks are stored locally in a JSON file.

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Storage:** JSON file
- **Interface:** Command Line Interface (CLI)

## 🚀 Getting Started

1. Clone the repository:
```bash
git clone https://github.com/Divyanth2468/CLI-TaskManager-Go.git
cd CLI-TaskManager-Go
```

2. Run the task manager with Go commands as shown above.
