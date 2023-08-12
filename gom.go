package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gom <command>")
		return
	}

	if err := checkInit(); err != nil {
		fmt.Println("Please run git init first")
		return
	}

	if os.Args[1] == "init" {
		if err := gmInit(); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func checkInit() error {
	_, err := exec.Command("git", "status").Output()
	if err != nil {
		return err
	}
	return nil
}

func gmInit() error {
	out, err := exec.Command("git", "remote", "-v").Output()
	if err != nil {
		return errors.New("git remote not found")
	}

	if IsModule() {
		return errors.New("go.mod already exists")
	}

	origin := ""
	for _, line := range strings.Split(string(out), "\n") {
		fields := strings.Fields(line)
		if len(fields) >= 2 && fields[0] == "origin" {
			origin = fields[1]
			break
		}
	}
	if origin == "" {
		return errors.New("origin not found")
	}

	origin = strings.Replace(origin, "git@", "", 1)
	origin = strings.Replace(origin, "https://", "", 1)
	origin = strings.Replace(origin, ".git", "", 1)

	_, err = exec.Command("go", "mod", "init", origin).Output()
	if err != nil {
		return err
	}

	return nil
}

func IsModule() bool {
	_, err := os.Stat("go.mod")
	return err == nil
}
