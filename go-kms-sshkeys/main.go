package main

import (
	"encoding/json"
	"fmt"
	"go-kms-sshkeys/src/types"
	"log"
	"os"
)

var (
	jsonFile types.JsonFile
	fileByte []byte
	err      error
)

func init() {

	fileByte, err = os.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(fileByte, &jsonFile)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	log.Println("\033[1;32m[+]\033[0m jsonFile: ", jsonFile.Project.Name)
	log.Println("\033[1;32m[+]\033[0m Author: ", jsonFile.Project.Author)
	log.Println("\033[1;32m[+]\033[0m Description: ", jsonFile.Project.Description)
	log.Println("\033[1;32m[+]\033[0m Version: ", jsonFile.Project.Version)
	fmt.Println()
	fmt.Println()
	log.Println("\033[1;32m[+]\033[0m Starting...")
	fmt.Println(jsonFile.Aws.Kmsid)
	fmt.Println(jsonFile.Aws.Region)

}
