package main

import "cli-task-manager/cmd"

func main() {
    cli := cmd.NewCLI()
    cli.Run()
}