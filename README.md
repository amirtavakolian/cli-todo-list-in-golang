
# To-Do List CLI App in Golang

🔴 To simulate the cookie, the CPU serial number is used, 
which is obtained using `wmic cpu get ProcessorId` <br> 
Therefore, to ensure the software works correctly, please run the code on a <mark><strong>Windows operating system</strong></mark>

<hr>

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


💾 Data Persistence
Tasks are stored in a local JSON file (e.g., tasks.json) in the same directory as the executable. This ensures your task list is saved even after restarting the program.

✅ Requirements
Go 1.18+

Compatible with Linux, macOS, and Windows

📄 License
This project is licensed under the MIT License. See the LICENSE file for details.

Happy coding! 🚀


