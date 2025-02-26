package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

// Função init é executada antes da função main
func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"AKIAIOSFODNN7EXAMPLE",
				"wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
				"",
			),
		},
	)

	if err != nil {
		panic(err)
	}

	s3Client = s3.New(sess)
	s3Bucket = "goexpert-bucket-exemplo"
}

func main() {
	dir, err := os.Open("./temp")

	if err != nil {
		panic(err)
	}
	defer dir.Close()

	// Controla uploads ativos no limite maximo definido
	uploadControl := make(chan struct{}, 15)
	errorFileUpload := make(chan string, 5)

	go func() {
		for {
			select {
			case fileName := <-errorFileUpload:
				uploadControl <- struct{}{}
				wg.Add(1)
				go uploadFile(fileName, uploadControl, errorFileUpload)
			}
		}
	}()

	for {
		files, err := dir.ReadDir(1)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %v\n", err)
			continue
		}
		wg.Add(1)
		// Vai preenchendo o channel e aguarda esvaziar para continuar adicionando
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Wait()

}

func uploadFile(filename string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done()
	completeFileName := fmt.Sprintf(".temp/%s", filename)
	fmt.Printf("Uploading file %s\n", completeFileName)
	f, err := os.Open(completeFileName)

	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", completeFileName, err)
		<-uploadControl // Esvazia o canal
		errorFileUpload <- completeFileName
		return
	}
	defer f.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})

	if err != nil {
		fmt.Printf("Error uploading file %s: %v\n", completeFileName, err)
		<-uploadControl // Esvazia o canal
		errorFileUpload <- completeFileName
		return
	}
	fmt.Printf("File %s uploaded successfully\n", completeFileName)
	<-uploadControl // Esvazia o canal
}
