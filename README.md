# Resume Cli

![made-with-go](https://img.shields.io/badge/Made%20with-Go-00ADD8.svg)
![vscode](https://img.shields.io/badge/GoLand-202020?logo=GoLand&logoColor=white)

This is a **Go-based** Command-Line Interface (CLI) portfolio. It offers a straightforward and interactive method for presenting your internships, projects, and skills within a terminal environment.

---

## Installation

**Prerequisite:** Ensure you have **Go (version 1.25 or higher)** installed on your system.

1.  Clone the repository using `git`:
    ```shell
    git clone https://github.com/Prakashdeveloper03/Resume-Cli.git
    ```

2.  Navigate into the cloned directory:
    ```shell
    cd Resume-Cli
    ```

3.  Tidy up the dependencies:
    ```shell
    go mod tidy
    ```

---

## Usage

You can run the application in two ways.

**Option 1: Run directly**

Execute the following command to run the program without compiling an executable:
```shell
go run main.go
````

**Option 2: Build and run the executable**

First, build the application:

```shell
go build
```

Then, run the compiled binary:

```shell
# On Linux/macOS
./Resume-Cli

# On Windows
Resume-Cli.exe
```