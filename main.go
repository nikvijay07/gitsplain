package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("hello from gitsplain")

	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments. usage: gitsplain <command>")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "status":
		fmt.Println("Argument: " + cmd)
		output, err := git("status")
		if err != nil {
			fmt.Println("Error executing git status: " + err.Error())
			os.Exit(1)
		}
		fmt.Println("Git status:\n" + output)
	case "history":
		fmt.Println("Argument: " + cmd)
		output, err := getGitCommitHistory()
		if err != nil {
			fmt.Println("Error executing git status: " + err.Error())
			os.Exit(1)
		}
		fmt.Println("Git history:\n" + output)
	default:
		fmt.Println("Invalid argument")
	}
}

func git(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	fmt.Println("COMMAND: ", cmd)
	err := cmd.Run()
	return strings.TrimSpace(out.String()), err
}

// git rev-list --parents --max-count=15 HEAD

func getGitCommitHistory() (string, error) {
	commit_dag, err_1 := git("rev-list", "--parents", "--max-count=15", "HEAD")
	messages, err_2 := git("log", "-n", "15", `--pretty=format:%h %ad %s`, "--date=short")

	commit_message_map := make(map[string]string)

	if err_1 != nil && err_2 != nil {
		fmt.Println("COMMIT DAG: ", commit_dag)
		fmt.Println("COMMIT messages: ", messages)
		fmt.Println("COMMIT MAP: ", commit_message_map)
	}

	return commit_dag, err_1

	// if err_1 == nil and err_2 == nil {
	// 	for i := 0; i < len(commit_dag); i++ {
	// 		commit_message_map[]
	// 	}
	// }

}
