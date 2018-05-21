<p align="center">
    <img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="MIT License"/>
    <a href="https://godoc.org/github.com/mikeyscode/commander"><img src="https://godoc.org/github.com/mikeyscode/commander?status.svg" alt="GoDoc"></a>
</p>


# About Commander
Commander is a basic exec wrapper to provide a cleaner interface for executing multiple commands against the same shell instance.

# Usage
```go
shell := commander.Console{}

shell.Add("cd /home/user/Projects/my-awesome-project")
shell.Add("git status")

result, err := shell.Run()
if err != nil {
    panic(err)
}

fmt.Printf("Output: %s", result)
```

# Contributing
If you wish to contribute to this project, you are welcome to raise a Pull Request into Develop from a feature branch to be reviewed. Outside that you are welcome to create issues for any problems that you come across.