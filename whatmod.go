package whatmod

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

func Get() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	modPath := path.Join(cwd, "go.mod")
	modName, err := GetAt(modPath)
	if err != nil {
		return "", err
	}
	return modName, nil
}

func GetAt(_path string) (string, error) {
	file, err := os.Open(_path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		modPrefix := "module"
		line := scanner.Text()
		if len(line) < len(modPrefix) {
			continue
		}
		if line[:len(modPrefix)] == modPrefix {
			return line[len(modPrefix)+1:], nil
		}
	}
	return "", fmt.Errorf("go.mod file doesn't contains module name")
}
