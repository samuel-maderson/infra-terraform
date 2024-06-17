package kms

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

func Encrypt(pemfile string, encryptfile string, kms_arn string, region string) {
	// Define the KMS key ID
	keyID := kms_arn

	// Read the file content
	filePath := pemfile
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Load the default configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
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

	b64data := base64.StdEncoding.EncodeToString(result.CiphertextBlob)

	// Write the encrypted content to a new file
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	encryptedFilePath := fmt.Sprintf("%s/data/%s", dir, encryptfile)
	err = ioutil.WriteFile(encryptedFilePath, []byte(b64data), 0644)
	if err != nil {
		fmt.Println("Error writing encrypted file:", err)
		return
	}

	fmt.Println("File encrypted successfully and saved to:", encryptedFilePath)
}
