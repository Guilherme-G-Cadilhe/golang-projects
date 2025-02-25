package main

import (
	"database/sql"
	"fmt"
	"net"

	// Importa a camada de acesso ao banco, a definição dos serviços (gerados via protobuf) e a implementação dos serviços
	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-3/3-gRPC/internal/database"
	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-3/3-gRPC/internal/pb"
	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-3/3-gRPC/internal/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Abre uma conexão com o banco de dados SQLite.
	// No JavaScript, seria semelhante a usar um driver de banco (como sqlite3 para Node.js).
	fmt.Println("Conectando ao banco de dados...")
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// Garante que a conexão será fechada quando main() terminar.
	defer db.Close()

	// Inicializa a camada de dados para a entidade "Category".
	// É equivalente a instanciar um modelo ou repositório em JavaScript.
	categoryDB := database.NewCategory(db)

	// Cria uma instância do serviço de categoria, injetando a dependência da camada de dados.
	// É similar a injetar dependências em frameworks como Express (através de containers ou simples objetos).
	categoryService := service.NewCategoryService(*categoryDB)

	// Cria um novo servidor gRPC.
	grpcServer := grpc.NewServer()

	// Registra o serviço de categoria no servidor gRPC.
	// Isso é como definir um endpoint e associá-lo a um handler em um framework REST (por exemplo, usando express.Router).
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	// Registra o "reflection" no servidor gRPC.
	// O reflection permite que clientes inspecionem os métodos e serviços disponíveis no servidor.
	// Em JavaScript, isso seria similar a ter uma API de documentação dinâmica (como Swagger) que permite a descoberta dos endpoints.
	reflection.Register(grpcServer)

	// Cria um listener TCP na porta 50051 para aceitar conexões.
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Servidor gRPC iniciado na porta 50051")

	// Inicia o servidor gRPC e começa a escutar por requisições.
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Servidor gRPC finalizado")
}
