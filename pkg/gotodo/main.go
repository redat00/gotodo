package gotodo

import (
	"fmt"
	"log"
	"os"

	"github.com/redat00/gotodo/internal/models"
	"gopkg.in/yaml.v3"
)

var baseDir string = fmt.Sprintf("%s/.gotodo", getUserHomeDir())

//func OutputTable() {
//	t := tabby.New()
//	t.AddHeader("Name", "Username")
//	t.AddLine("Renaud", "redat00")
//	t.Print()
//}

// Check if a file is empty
func fileEmpty(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Panicln(err)
	}
	if fileInfo.Size() < 1 {
		return true
	}
	return false
}

// Check if a file exist
func fileExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// Used to create a file for a given path
func createFile(path string) {
	fmt.Printf("%s does not exist. Creating it.\n", path)
	if _, err := os.Create(path); err != nil {
		log.Panicln(err)
	}
}

func getUserHomeDir() string {
	// Obtain user home directory
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Panicln(err)
	}
	return homedir
}

func createBaseFileStructure() {
	baseConfFile := fmt.Sprintf("%s/gotodo.yml", baseDir)
	baseDatabase := fmt.Sprintf("%s/gotodo.db", baseDir)

	// Create base directory
	if !fileExist(baseDir) {
		fmt.Printf("%s does not exist. Creating it.\n", baseDir)
		if err := os.Mkdir(baseDir, os.ModePerm); err != nil {
			log.Panicln(err)
		}
	}
	// Create base conf file
	if !fileExist(baseConfFile) {
		createFile(baseConfFile)
	}
	// Create database file
	if !fileExist(baseDatabase) {
		createFile(baseDatabase)
	}

}

// Load configuration from configuration file
func loadConfigurationFromFile(configurationPath string) *models.GotodoConfiguration {
	data, err := os.ReadFile(configurationPath)
	if err != nil {
		log.Panicln(err)
	}

	config := models.GotodoConfiguration{}
	yaml.Unmarshal(data, &config)

	return &config
}

func createBaseConfiguration() {

}

// Used to init application
func InitApp() {
	createBaseFileStructure()
	config := loadConfigurationFromFile("/home/redat/.gotodo/gotodo.yml")
	fmt.Println(config)
}
