package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"io/ioutil"
	"os"
)

func main(pemfile string, encryptfile string, kms_arn string) {
	// Define the KMS key ID
	keyID := "arn:aws:kms:your-region:your-account-id:key/your-key-id"

	// Read the file content
	filePath := pemfile
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Load the default configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("your-region"))
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	// Create a new KMS client
	svc := kms.NewFromConfig(cfg)

	// Encrypt the file content using AWS KMS
	input := &kms.EncryptInput{
		KeyId:     aws.String(keyID),
		Plaintext: fileContent,
	}
	result, err := svc.Encrypt(context.TODO(), input)
	if err != nil {
		fmt.Println("Error encrypting file:", err)
		return
	}

	// Write the encrypted content to a new file
	encryptedFilePath := encryptfile
	err = ioutil.WriteFile(encryptedFilePath, result.CiphertextBlob, 0644)
	if err != nil {
		fmt.Println("Error writing encrypted file:", err)
		return
	}

	fmt.Println("File encrypted successfully and saved to:", encryptedFilePath)
}
