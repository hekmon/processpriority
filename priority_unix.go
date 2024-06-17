//go:build !windows

package processpriority

import (
	"fmt"
	"syscall"
)

// Opiniated values
const (
	UnixPriorityIdle        = 19
	UnixPriorityBelowNormal = 5
	UnixPriorityNormal      = 0
	UnixPriorityAboveNormal = -5
	UnixPriorityHigh        = -10
	UnixPriorityRealTime    = -20
)

func getOS(pid int) (priority ProcessPriority, rawPriority int, err error) {
	if rawPriority, err = GetRAW(pid); err != nil {
		return
	}
	switch rawPriority {
	case UnixPriorityIdle:
		priority = PriorityIdle
	case UnixPriorityBelowNormal:
		priority = PriorityBelowNormal
	case UnixPriorityNormal:
		priority = PriorityNormal
	case UnixPriorityAboveNormal:
		priority = PriorityAboveNormal
	case UnixPriorityHigh:
		priority = PriorityHigh
	case UnixPriorityRealTime:
		priority = PriorityRealTime
	default:
		priority = PriorityCustom
	}
	return
}

func setOS(pid int, priority ProcessPriority) error {
	var unixPriority int
	switch priority {
	case PriorityIdle:
		unixPriority = 19
	case PriorityBelowNormal:
		unixPriority = 5
	case PriorityNormal:
		unixPriority = 0
	case PriorityAboveNormal:
		unixPriority = -5
	case PriorityHigh:
		unixPriority = -10
	case PriorityRealTime:
		unixPriority = -20
	default:
		return fmt.Errorf("unknown universal priority: %d", priority)
	}
	return SetRAW(pid, unixPriority)
}

// GetRAW is an OS specific function to get the priority of a process.
// As priority values are not the same on all OSes, you should use the universal function Get() instead to be platform agnostic.
func GetRAW(pid int) (priority int, err error) {
	return syscall.Getpriority(syscall.PRIO_PROCESS, pid)
}

// SetRAW is an OS specific function to set the priority of a process.
// As priority values are not the same on all OSes, you should use the universal function Set() instead to be platform agnostic.
func SetRAW(pid, priority int) (err error) {
	return syscall.Setpriority(syscall.PRIO_PROCESS, pid, priority)
}
