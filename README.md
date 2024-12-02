# Todo CLI
Todo CLI is a user–friendly console application for managing tasks, storing history, and clearing data. The utility supports the creation of tasks, subtasks, status management and the history of completed actions.

## Features
* Create tasks with a description and a deadline.
* Create and manage subtasks.
* Deleting tasks and subtasks
* Changing task statuses.
* View a list of tasks with deadlines.
* Current tasks can be swapped in the list
* Task history management (viewing, filtering, cleaning).
* Clearing tasks, history and all data.
* All tasks are sorted by deadline

## Установка
1. Make sure that you have Go installed (version 1.23.3+).
2. Clone the repository
3. Compile the utility:
```bash
go build -o todo main.go
```
4. Launch the application:
```bash
./todo
```

## Usage
### Commands for working with tasks
* Show help:
```bash
./todo --help
```
`--help` can also be applied to other commands to see detailed information about them
```bash
./todo create --help
```

* View all current tasks:
```bash
./todo ./todo tasks
```
* Show the list of tasks for today:
```bash
./todo today
```
* Show a detailed description of the task by id:
```bash
./todo task --id=2
```
* Creating a task:
```bash
./todo create --description="Task name" --deadline="DD-MM-YYYY"
```
* Creating a subtask:
```bash
./todo create subtask --parentid=1 --description="Subtask name"
```
* Deleting a task:
```bash
./todo delete --parentid=1
```
* Deleting a subtask:
```bash
./todo delete subtask --parentid=1 --subtaskid=2
```
* Changing the status of the task:
```bash
./todo status --id=1 --status=done
```
* Change the order of two tasks in the list in places (change their deadline):
```bash
./todo swap --id1=1 --id2=2
```
* Cleaning up all completed and unnecessary tasks:
```bash
./todo clear
```

### Commands for working with history
* View the entire history:
```bash
./todo history
```
* View the history for the last N days:
```bash
./todo history days --days=7
```
* View history for a specific date:
```bash
./todo history date --date="DD-MM-YYYY"
```
* Detailed information about the task from the EntryID history:
```bash
./todo history task --id=1
```
* Clearing the entire history:
```bash
./todo history clear
```
* Clearing old records from history (older than 7 days):
```bash
./todo history clear old
```
### Command to clear all data
* Complete data cleanup (id.json, tasks.json, history.json):
```bash
./todo clear all
```

## Usage example
1. Look at the help:
```bash
./todo --help
```

2. We look at the information on the team of interest:
```bash
./todo create --help
```

3. Create a task:
```bash
./todo create -d "Write a README" -l "02-12-2024"
```
4. Let's create some more tasks:
```bash
./todo create -d "Make a training plan" -l "05-12-2024"
```
```bash
./todo create -d "Cook dinner" -l "02-12-2024"
```

5. Let's add some subtasks:
```bash
./todo create subtask -p 2 -d "Include leg exercises in your workout"
```
```bash
./todo create subtask -p 3 -d "Buy products"
```
```bash
./todo create subtask -p 3 -d "Make a salad" 
```
```bash
./todo create subtask -p 3 -d "Cook the baked chicken" 
```

6. Set the status of the subtask completed
```bash
./todo status -i 5 -s done
```

7. Set the status of the subtask to unnecessary
```bash
./todo status -i 4 -s unnecessary
```

8. Set the status of the task completed
```bash
./todo status -i 3 -s done
```

9. Clear the list of current tasks from completed ones
```bash
./todo clear
```

10. Let's look at the history
```bash
./todo history
```

11. Clear all data:
```bash
./todo clear all
```
All data files have been successfully cleared.


