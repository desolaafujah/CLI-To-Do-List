package main
import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	ID	int
	Title	string
	Done	bool
}

var tasks []Task 
var nextID = 1

func main() {
	fmt.Println("Welcome to Your CLI To-Do List!")
	menu()
}


func menu() {
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
			case 1:
				addTask()
			case 2:
				listTasks()
			case 3:
				// markTask()
			case 4:
				// deleteTask()
			case 5:
				// saveTasksToFile()
				fmt.Println("Bye!")
				return
			// default:
			// 	fmt.Println("Invalid choice. Pick something else ")
		}
	}
}

func addTask() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Enter task name: ")
	title, _ := reader.ReadString('\n')
	title = title[:len(title)-1]

	task:= Task{
		ID: nextID,
		Title: title,
		Done: false,
	}
	tasks = append(tasks, task)
	nextID++

	fmt.Println("Task added successfully")
}

func listTasks() {
	var size = len(tasks)
	for i := 0; i < size; i++ {
		fmt.Println(tasks[i].Title)
	}
}