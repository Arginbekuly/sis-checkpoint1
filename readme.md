# SIS Project - Checkpoint 1

## What does the program (Domain) do

Program works as a text analyzer:\
Firstly sent line with a text\
The grogram checks the text,splits, it and counts the statistics
  - the number of words;
  - the actual number of characters (Unicode-runes);
  - the number of spaces.

Task LifeStyle

Each task goes through a path of three states:

[ PENDING ]  --->  ( Process() )  --->  [ PROCESSING ]  --->  ( Успех )  --->  [ COMPLETED ]
(Ожидает)                              (В процессе)                           (Завершена)

1. **PENDING:** The task has been created, the input text has been verified, and the task is in memory and awaiting processing.
2. **PROCESSING (In process):** The text processing function is called, statistics are being calculated.
3. **COMPLETED:** Statistics are calculated, the result is saved in the task.


## Rules and validation

- **Checking the input:** The text cannot be empty or consist only of spaces ('ErrEmptyInput').
- **Protection against re-processing:** If the task has already been completed (`COMPLETED`), it cannot be started again — the program returns the error `ErrInvalidTransition`
- **Memory search:** If you request a task with a non-existent ID, the repository returns the error `ErrTaskNotFound`


## How to check the project

Commands for the terminal:

```bash
# Formatting the code
gofmt -l .

# Checking for Go errors
go vet ./...

# Running automated tests
go test -v ./...