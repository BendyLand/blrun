package main

import (
	"fmt"
	"bufio"
	"github.com/BurntSushi/toml"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Hi blrun!")
	file, err := getConfigFile()
	if err != nil {
		fmt.Println("Error getting config file:", err)
		return
	}
	fmt.Println(file)
	// config, err := constructConfig(file)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// execBuild(config)
}

/*
Steps:
Checks for config file (probably toml again).
	If no config file found:
		Prompt for details.
		Write config file in root dir.
	Otherwise:
		Parse config file.
Construct build command.
Construct run command.
Execute build command.
Execute run command.
*/

func execBuild(config Config) {
	command := constructBuildCmd(config)
	fmt.Println("Running:", command)
	cmd := exec.Command("sh", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing build command:", err)
		return
	}
	fmt.Println("Project built successfully!")
}

type Config struct {
	Compiler string
	Path     string
	Files    []string
	Extras   string
	Run      string
}

func constructConfig(config string) (Config, error) {
	var result Config
	_, err := toml.Decode(config, &result)
	if err != nil {
		e := fmt.Errorf("Error parsing TOML file: %s\n", err)
		return Config{}, e
	}
	return result, nil
}

func getMissingConfigFields() []string {
	stdin := bufio.NewReader(os.Stdin)
	fmt.Println("What compiler would you like to use?")
	compiler, _ := stdin.ReadString('\n')
	compiler = strings.Trim(compiler, "\n")
	
	fmt.Println("What is the path from the root directory to where the files are located?")
	fmt.Println("Please keep it blank if they are in the root directory.")
	path, _ := stdin.ReadString('\n')
	path = strings.Trim(path, "\n")

	fmt.Println("Please enter all of your files separated by spaces, and then a newline.")
	fileStr, _ := stdin.ReadString('\n')
	files := strings.Split(fileStr, " ")
	for i := range len(files) {
		files[i] = strings.Trim(files[i], "\n")
		files[i] = "\"" + files[i] + "\""
	}
	temp := strings.Trim(strings.Join(files, ", "), "\n")
	fileStr = "[" + temp + "]"

	fmt.Println("Please enter any extra flags or options to include in your build command.")
	extras, _ := stdin.ReadString('\n')
	extras = strings.Trim(extras, "\n")

	fmt.Println("Please enter the command that you want to use to run your program.")
	run, _ := stdin.ReadString('\n')
	run = strings.Trim(run, "\n")

	lines := []string{compiler, path, fileStr, extras, run}
	return lines
}

func createMissingConfigFile() string {
	var result string
	lines := getMissingConfigFields()
	for i, line := range lines {
		switch i {
		case 0:
			result += "compiler = \"" + line + "\"\n"
		case 1:
			result += "path = \"" + line + "\"\n"
		case 2:
			result += "files = " + line + "\n"
		case 3:
			result += "extras = \"" + line + "\"\n"
		case 4:
			result += "run = \"" + line + "\"\n"
		}
	}
	file, err := os.Create("blrun.toml")
	if err != nil {
		fmt.Println("Error automatically creating build file.\nPlease make your own to avoid re-entering details each time.")
	} else {
		file.WriteString(result)
		file.Close()
	}
	return result
}

func getConfigFile() (string, error) {
	path, err := os.Getwd()
	configPath := filepath.Join(path, "blrun.toml")
	_, err = os.Stat(configPath)
	if err == nil {
		result, err := os.ReadFile(configPath)
		if err != nil {
			e := fmt.Errorf("Error reading file:%s\n", err)
			return "", e
		}
		return string(result), nil
	}
	fmt.Println("No config file detected. Let's create one:")
	config := createMissingConfigFile()
	return config, nil
}

func constructBuildCmd(config Config) string {
	result := config.Compiler + " "
	for _, file := range config.Files {
		temp := filepath.Join(config.Path, file)
		result += temp + " "
	}
	result += config.Extras
	return result
}
