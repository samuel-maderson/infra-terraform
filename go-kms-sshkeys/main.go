package main

import (
	"encoding/json"
	"fmt"
	"go-kms-sshkeys/src/kms"
	"go-kms-sshkeys/src/types"
	"log"
	"os"

	arg "github.com/alexflint/go-arg"
)

var (
	jsonFile types.JsonFile
	fileByte []byte
	err      error
	args     types.Args
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

	arg.MustParse(&args)

	if args.Import == "" {
		log.Fatal("None arguments were given, use -h for options\nExample: ./main.go -h")
	}
}

func main() {

	log.Println("\033[1;32m[+]\033[0m jsonFile: ", jsonFile.Project.Name)
	log.Println("\033[1;32m[+]\033[0m Author: ", jsonFile.Project.Author)
	log.Println("\033[1;32m[+]\033[0m Description: ", jsonFile.Project.Description)
	log.Println("\033[1;32m[+]\033[0m Version: ", jsonFile.Project.Version)
	fmt.Println()
	log.Println("\033[1;32m[+]\033[0m Starting...")
	if args.Import == "true" {
		log.Println("\033[1;32m[+]\033[0m Importing keys...")
		kms.Encrypt(jsonFile.Project.Pem_file, jsonFile.Project.Encrypt_file, jsonFile.Aws.Kms_arn, jsonFile.Aws.Region)
	}

	// else if args.Export == "true" {
	// 	log.Println("\033[1;32m[+]\033[0m Exporting keys...")
	// 	kms.Decrypt(jsonFile.Project.Encrypt_file, jsonFile.Project.Decrypt_file, jsonFile.Aws.Kms_arn, jsonFile.Aws.Region)
	// } else {
	// 	log.Fatal("None arguments were given, use -h for options\nExample: ./main.go -h")
	// }
}
