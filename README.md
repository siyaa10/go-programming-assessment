# Student Grade Calculator

## Project Overview
A simple command-line application developed using Go. It accepts a student's name and three marks, calculates the total and average, and displays the grade and pass/fail result.

## Features
- Accepts student name and three subject marks.
- Validates marks to ensure they are between 0 and 100.
- Calculates total marks and average.
- Assigns grades based on the average:
  - A: 80 and above
  - B: 60–79
  - C: 40–59
  - F: Below 40
- Displays Pass or Fail status.

## Technologies Used
- Go (Golang)

## How to Run
Open the terminal in the project folder and run:

```bash
go run main.go
```

## Files
- `main.go` – Source code
- `normal-output.png` – Successful program output
- `validation-output.png` – Input validation demonstration