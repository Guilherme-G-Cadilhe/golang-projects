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

// A função init é executada antes da main.
// Aqui, ela cria uma sessão AWS e inicializa o cliente S3.
// (Em JavaScript, seria similar a configurar o SDK da AWS antes de usá-lo.)
func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"AKIAIOSFODNN7EXAMPLE",                     // Access Key (exemplo)
				"wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY", // Secret Key (exemplo)
				"",
			),
		},
	)
	if err != nil {
		panic(err)
	}

	// Inicializa o cliente S3 usando a sessão criada.
	s3Client = s3.New(sess)
	// Define o nome do bucket que será usado para o upload.
	s3Bucket = "goexpert-bucket-exemplo"
}

func main() {
	// Abre o diretório onde estão os arquivos a serem enviados.
	dir, err := os.Open("./temp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	// Channel de controle para limitar uploads concorrentes (semelhante a um semaphore).
	// Neste caso, limita a 15 uploads simultâneos.
	uploadControl := make(chan struct{}, 15)
	// Channel para sinalizar arquivos que falharam no upload e precisam ser re-tentados.
	errorFileUpload := make(chan string, 5)

	// Goroutine que monitora erros e reenvia os arquivos que falharam.
	go func() {
		for {
			select {
			// Sempre que um arquivo com erro é enviado para errorFileUpload,
			// ele reenvia esse arquivo para o canal de uploads.
			case fileName := <-errorFileUpload:
				// Garante que não exceda o limite de uploads ativos.
				uploadControl <- struct{}{}
				wg.Add(1)
				go uploadFile(fileName, uploadControl, errorFileUpload)
			}
		}
	}()

	// Loop que lê o diretório e dispara uploads para cada arquivo encontrado.
	for {
		// Lê um arquivo de cada vez.
		files, err := dir.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				break // Finaliza quando não há mais arquivos.
			}
			fmt.Printf("Error reading directory: %v\n", err)
			continue
		}
		wg.Add(1)
		// Preenche o channel para controlar a concorrência e aguarda "vago" para continuar.
		uploadControl <- struct{}{}
		// Lança uma goroutine para fazer o upload do arquivo.
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	// Aguarda que todas as goroutines de upload sejam finalizadas.
	wg.Wait()
}

// Função uploadFile realiza o upload de um arquivo para o S3.
// Ela recebe o nome do arquivo, um canal de controle para gerenciar concorrência e um canal de erro para re-tentativas.
func uploadFile(filename string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done()
	// Monta o caminho completo do arquivo. (Note que a pasta é ".temp", diferente do diretório aberto "temp")
	completeFileName := fmt.Sprintf(".temp/%s", filename)
	fmt.Printf("Uploading file %s\n", completeFileName)

	// Abre o arquivo para leitura.
	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", completeFileName, err)
		<-uploadControl // Libera uma "vaga" no channel de controle.
		// Envia o nome do arquivo para o canal de erro para re-tentativa.
		errorFileUpload <- completeFileName
		return
	}
	defer f.Close()

	// Realiza o upload para o bucket S3.
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("Error uploading file %s: %v\n", completeFileName, err)
		<-uploadControl                     // Libera a vaga no channel de controle.
		errorFileUpload <- completeFileName // Re-tenta o upload.
		return
	}
	fmt.Printf("File %s uploaded successfully\n", completeFileName)
	<-uploadControl // Libera a vaga ao terminar o upload.
}
