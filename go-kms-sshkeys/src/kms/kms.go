package kms

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

func Encrypt(cfg aws.Config, pemfile string, encryptfile string, kms_arn string, region string) {

	// Write the encrypted content to a new file
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	encryptedFilePathFull := fmt.Sprintf("%s/data/%s", dir, encryptfile)

	// Define the KMS key ID
	keyID := kms_arn

	// Read the file content
	filePath := pemfile
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
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

	err = ioutil.WriteFile(encryptedFilePathFull, []byte(b64data), 0644)
	if err != nil {
		fmt.Println("Error writing encrypted file:", err)
		return
	}

	log.Println("\033[1;32m[+]\033[0m File encrypted successfully and saved to:", encryptedFilePathFull)
}

func Decrypt(cfg aws.Config, encryptedFilePath string, decryptedFilePath string, region string) {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	decryptedFilePathFull := fmt.Sprintf("%s/data/%s", dir, decryptedFilePath)
	encryptedFilePathFull := fmt.Sprintf("%s/data/%s", dir, encryptedFilePath)

	// Create a KMS client
	svc := kms.NewFromConfig(cfg)

	// Read the encrypted file
	encryptedData, err := ioutil.ReadFile(encryptedFilePathFull)
	if err != nil {
		log.Fatalf("failed to read encrypted file, %v", err)
	}

	// Assuming the data is base64 encoded
	decodedData, err := base64.StdEncoding.DecodeString(string(encryptedData))
	if err != nil {
		log.Fatalf("failed to decode base64 data, %v", err)
	}

	// Decrypt the data
	decryptInput := &kms.DecryptInput{
		CiphertextBlob: decodedData,
	}
	decryptOutput, err := svc.Decrypt(context.TODO(), decryptInput)
	if err != nil {
		log.Fatalf("failed to decrypt data, %v", err)
	}

	// Write the decrypted data to a new file
	err = ioutil.WriteFile(decryptedFilePathFull, decryptOutput.Plaintext, 0644)
	if err != nil {
		log.Fatalf("failed to write decrypted file, %v", err)
	}

	log.Printf("\033[1;32m[+]\033[0m Decrypted file written to %s\n", decryptedFilePathFull)
}
