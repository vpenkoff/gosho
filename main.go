package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type FlagBool struct {
	ShortName bool
	LongName  bool
}

func (f FlagBool) Passed() bool {
	return f.ShortName || f.LongName
}

var eFlag, cFlag, hFlag FlagBool

func printDefaults() {
	fmt.Fprintf(os.Stderr, `
gosho
-c  --connect  connect to host
-e  --edit     edit ssh config
-h  --help     print this message

`)
}

func init() {
	flag.BoolVar(&eFlag.ShortName, "e", false, "edit ssh config")
	flag.BoolVar(&eFlag.LongName, "edit", false, "edit ssh config")
	flag.BoolVar(&cFlag.ShortName, "c", false, "connect to host")
	flag.BoolVar(&cFlag.LongName, "connect", false, "connect to host")
	flag.BoolVar(&hFlag.ShortName, "h", false, "print this message")
	flag.BoolVar(&hFlag.LongName, "help", false, "print this message")
	flag.Parse()
}

func main() {
	config_path, err := getConfigPath()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(config_path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	hosts, err := readConfig(file)
	if err != nil {
		log.Fatal(err)
	}

	if cFlag.Passed() {
		prompt := promptui.Select{
			Label: "Select Host",
			Items: hosts,
		}

		_, selected_host, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", selected_host)

		cmd := exec.Command("ssh", selected_host)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}

		return
	}

	if eFlag.Passed() {
		cmd := exec.Command("vi", config_path)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}

		return
	}

	printDefaults()
}

func getConfigPath() (string, error) {
	config_path, err := filepath.Abs(os.Getenv("HOME") + "/.ssh/config")
	if err != nil {
		return "", err
	}

	return config_path, nil
}

func readConfig(file *os.File) ([]string, error) {
	var hosts []string
	target_line_regex := regexp.MustCompile(`^\bHost\b.+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if target_line_regex.MatchString(line) {
			host := strings.Split(line, " ")[1]
			hosts = append(hosts, host)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	sort.Strings(hosts)

	return hosts, nil
}
