package main

import (
	"github.com/jessevdk/go-flags"
	"os/user"
	"path"
	"os"
	"bufio"
	"regexp"
	"encoding/json"
	"strings"
	"fmt"
	"io/ioutil"
)

type CliOptions struct {
	Verbose         bool   `short:"v" long:"verbose" description:"Verbose logging to stderr"`
	ConfigFile      string `long:"config-file" description:"The location of the config file"`
	CredentialsFile string `long:"credentials-file" description:"The location of the credentials file"`
	OutputFile      string `long:"output-file" description:"If omitted, the JSON body is printed to stdout"`
}

type ParsedAWS struct {
	Config      map[string]map[string]interface{} `json:"config"`
	Credentials map[string]map[string]interface{} `json:"credentials"`
}

func main() {
	options := CliOptions{
		Verbose:    false,
		OutputFile: "",
	}
	if _, err := flags.Parse(&options); err != nil {
		// No need to print anything, go-flags will do so for us.
		return
	} else {
		// Set the defaults if they're empty
		setCliDefaults(&options)
	}

	aws := ParsedAWS{
		Config:      make(map[string]map[string]interface{}),
		Credentials: make(map[string]map[string]interface{}),
	}

	// Let's read all the files!
	readValues(options.ConfigFile, &aws.Config, options.Verbose)
	readValues(options.CredentialsFile, &aws.Credentials, options.Verbose)

	// Write it out to stdout
	bytes, err := json.Marshal(aws)
	if err != nil {
		panic(err)
	}
	if options.OutputFile == "" {
		jsonBody := string(bytes)
		fmt.Println(jsonBody)
	} else {
		ioutil.WriteFile(options.OutputFile, bytes, 0644)
	}
}

func readValues(filename string, destination *map[string]map[string]interface{}, verbose bool) {
	if verbose {
		stderr(fmt.Sprintf("Reading from: %s", filename))
	}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	profileRegex := regexp.MustCompile("\\[(?:profile\\s*)?(.*)\\]")
	currentProfile := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := strings.TrimSpace(scanner.Text())
		if currentLine == "" {
			continue
		}
		if verbose {
			stderr(currentLine)
		}
		submatch := profileRegex.FindStringSubmatch(currentLine)
		isProfile := len(submatch) > 1
		if isProfile {
			currentProfile = submatch[1]
			(*destination)[currentProfile] = make(map[string]interface{})
		} else {
			tuple := strings.SplitN(currentLine, "=", 2)
			key := strings.TrimSpace(tuple[0])
			value := strings.TrimSpace(tuple[1])
			(*destination)[currentProfile][key] = value
		}
	}
	if verbose {
		stderr("")
	}
}

func stderr(out string) (int, error) {
	return fmt.Fprintf(os.Stderr, "%s\n", out)
}

func setCliDefaults(options *CliOptions) {
	if options.ConfigFile == "" || options.CredentialsFile == "" {
		usr, err := user.Current()
		if err != nil {
			panic(err)
		}
		home := usr.HomeDir

		if options.ConfigFile == "" {
			options.ConfigFile = path.Join(home, ".aws", "config")
		}
		if options.CredentialsFile == "" {
			options.CredentialsFile = path.Join(home, ".aws", "credentials")
		}
	}
}
