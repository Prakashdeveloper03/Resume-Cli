package utils

import (
	"encoding/json"
	"os"
	"os/exec"
	"runtime"
)

// LoadJSON reads a JSON file from the given path and unmarshals it into the provided interface.
// `filePath`: path to the JSON file
// `v`: pointer to the struct or data structure to unmarshal into
func LoadJSON(filePath string, v interface{}) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, v)
}

// --- Terminal Clearing Logic ---

// clear is a map that stores platform-specific clear screen commands.
var clear = map[string]func(){
	"linux": func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	},
	"windows": func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	},
	"darwin": func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	},
}

// CallClear clears the terminal screen based on the user's operating system.
// Panics if the OS is unsupported.
func CallClear() {
	if clearFunc, exists := clear[runtime.GOOS]; exists {
		clearFunc()
	} else {
		panic("Unsupported platform: cannot clear terminal screen")
	}
}

// EnsureDirs checks a list of directory paths and creates any that do not exist.
// Useful for preparing config or data directories at startup.
func EnsureDirs(dirs []string) {
	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			_ = os.MkdirAll(dir, 0755)
		}
	}
}
