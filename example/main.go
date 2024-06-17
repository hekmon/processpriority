package main

import (
	"fmt"
	"os/exec"

	"github.com/hekmon/processpriority"
)

func main() {
	// Run command
	// cmd := exec.Command("sleep", "1")
	cmd := exec.Command("./sleep.exe")
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	// Get its current priority
	prio, rawPrio, err := processpriority.Get(cmd.Process.Pid)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Current process priority is %s (%d)\n", prio, rawPrio)
	// Change its priority
	newPriority := processpriority.BelowNormal
	fmt.Printf("Changing process priority to %s\n", newPriority)
	if err = processpriority.Set(cmd.Process.Pid, newPriority); err != nil {
		panic(err)
	}
	// Verifying
	if prio, rawPrio, err = processpriority.Get(cmd.Process.Pid); err != nil {
		panic(err)
	}
	fmt.Printf("Current process priority is %s (%d)\n", prio, rawPrio)
	// Wait for the cmd to end
	if err = cmd.Wait(); err != nil {
		panic(err)
	}
}
