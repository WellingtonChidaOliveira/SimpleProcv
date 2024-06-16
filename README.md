# SimpleProcv

SimpleProcv is a Go-based application designed to manipulate Excel files. It allows you to match columns, get values in rows, and set values in rows.

## Prerequisites

- Go (version 1.16 or later)
- Excelize Go package

## Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/wellingtonchida/simpleprocv.git
```

Usage
The application uses Excel files located in the templates directory. You can add your own Excel files to this directory.
When referencing these files in the application, use the relative path format ../../templates/filename.xlsx.  To run the
application, navigate to the directory containing the main.go file:

```bash
cd cmd/run
```

Then, run the application:

```bash
go run main.go
```

You will be prompted to enter the following information:  
- Index columns
- Main file path
- Second file path
- Sheet name
- Column to get value
- Column to set value

Enter the requested information, and the application will perform the specified operations on the Excel files.