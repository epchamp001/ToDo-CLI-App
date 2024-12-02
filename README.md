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
![image](https://github.com/user-attachments/assets/2563707d-ae4c-47c2-bcd1-2da384d8f483)

2. We look at the information on the team of interest:
```bash
./todo create --help
```
![image](https://github.com/user-attachments/assets/2e8f2e02-44c0-4cea-8dfd-1b66c447cc07)

3. Create a task:
```bash
./todo create -d "Write a README" -l "02-12-2024"
```
![image_2024-12-02_17-18-59](https://github.com/user-attachments/assets/a702cdb4-93d5-4834-9b85-a811c530cee9)

4. Let's create some more tasks:
```bash
./todo create -d "Make a training plan" -l "05-12-2024"
```
```bash
./todo create -d "Cook dinner" -l "02-12-2024"
```
![image](https://github.com/user-attachments/assets/431f625f-c8c5-4eec-9a9f-df12bb91128f)

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
![image](https://github.com/user-attachments/assets/b41767da-e77e-4262-a8d1-08428587e3b2)

6. Set the status of the subtask completed
```bash
./todo status -i 5 -s done
```
![image](https://github.com/user-attachments/assets/6e6b72e1-ab42-46c9-8b1c-91f69d958008)

7. Set the status of the subtask to unnecessary
```bash
./todo status -i 4 -s unnecessary
```
![image](https://github.com/user-attachments/assets/f63a3c3e-51f9-4897-862b-215bf72d50ec)

8. Set the status of the task completed
```bash
./todo status -i 3 -s done
```
![image](https://github.com/user-attachments/assets/45d5efbf-e9ec-45b8-9894-efee2bf57fd8)

9. Clear the list of current tasks from completed ones
```bash
./todo clear
```
![image](https://github.com/user-attachments/assets/f41e8fa0-6ef4-4a74-ba68-39b828ac6eb1)

10. Let's look at the history
```bash
./todo history
```
![image](https://github.com/user-attachments/assets/34df1d2f-5380-4b15-94c8-12ae5e13c05e)

11. Clear all data:
```bash
./todo clear all
```
All data files have been successfully cleared.


