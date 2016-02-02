package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

// Checks error for nil and if true, displays it and stops the program with error code
func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Checks error for nil and if true, displays it as fatal error and then stops the program with error code
func CheckFatalError(err error) {
	if err != nil {
		CheckError(fmt.Errorf("Fatal error: %v", err))
	}
}

// Checks whether a file or directory exists
func FileExists(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

// Writes current PID to the provided file
func WritePid(pidfile string) {
	dirname := path.Dir(pidfile)
	if !FileExists(dirname) {
		CheckFatalError(fmt.Errorf("Directory \"%s\" doesn't exist", dirname))
	}

	if FileExists(pidfile) {
		CheckFatalError(fmt.Errorf("File \"%s\" already exist", pidfile))
	}

	pid := os.Getpid()
	err := ioutil.WriteFile(pidfile, []byte(strconv.Itoa(pid)), os.FileMode(0644))
	CheckFatalError(err)
}

// Removes file with current PID and stops the program with provided code
func StopExecution(code int, pidfile string) {
	os.Remove(pidfile)
	os.Exit(code)
}
