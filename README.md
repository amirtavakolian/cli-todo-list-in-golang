
# To-Do List CLI App in Golang

A command-line To-Do list application built with Golang.  
Easily add, list, complete, and delete your tasks right from the terminal.

## ✨ Features

- Add new tasks
- List all tasks
- Mark tasks as completed
- Delete tasks
- Data persistence using a local file (e.g., `tasks.json`)

## 📦 Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/todo-cli-go.git
   cd todo-cli-go
   ```
   
Build the application:
```bash
go build -o todo
```

Run the app:
```
./todo
```

🛠️ Usage
```bash
todo add "Buy groceries"
todo list
todo done 2
todo delete 3
```

Example
```bash
$ todo add "Finish homework"
Task added: "Finish homework"

$ todo list
1. [ ] Buy groceries
2. [ ] Finish homework

$ todo done 1
Task marked as completed: "Buy groceries"

$ todo list
1. [x] Buy groceries
2. [ ] Finish homework
🧾 File Structure
main.go – Entry point for the application

task.go – Task model and data handling

storage.go – File read/write logic
```

💾 Data Persistence
Tasks are stored in a local JSON file (e.g., tasks.json) in the same directory as the executable. This ensures your task list is saved even after restarting the program.

✅ Requirements
Go 1.18+

Compatible with Linux, macOS, and Windows

📄 License
This project is licensed under the MIT License. See the LICENSE file for details.

Happy coding! 🚀


