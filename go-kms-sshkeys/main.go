package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-kms-sshkeys/src/kms"
	"go-kms-sshkeys/src/types"
	"log"
	"os"

	arg "github.com/alexflint/go-arg"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var (
	jsonFile types.JsonFile
	fileByte []byte
	err      error
	args     types.Args
	cfg      aws.Config
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

	// Load the default configuration
	cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithRegion(jsonFile.Aws.Region))
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	arg.MustParse(&args)
}

func main() {

	log.Println("\033[1;32m[+]\033[0m Starting...")
	if args.Import == "true" {
		log.Println("\033[1;32m[+]\033[0m Importing keys...")
		kms.Encrypt(cfg, jsonFile.Project.Pem_file, jsonFile.Project.Encrypt_file, jsonFile.Aws.Kms_arn, jsonFile.Aws.Region)
	} else if args.Export == "true" {
		log.Println("\033[1;32m[+]\033[0m Exporting keys...")
		kms.Decrypt(cfg, jsonFile.Project.Encrypt_file, jsonFile.Project.Decrypt_file, jsonFile.Aws.Region)
	} else {
		log.Fatal("None arguments were given, use -h for options\nExample: ./main.go -h")
	}
}
