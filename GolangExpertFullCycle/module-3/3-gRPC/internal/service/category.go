package service

import (
	"context"
	"io"

	// Importa a camada de dados e as definições dos serviços (gerados pelo protobuf)
	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-3/3-gRPC/internal/database"
	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-3/3-gRPC/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CategoryService implementa os métodos do serviço gRPC para gerenciar categorias.
// Ele incorpora a implementação padrão (UnimplementedCategoryServiceServer) gerada pelo protobuf,
// o que garante compatibilidade mesmo se novos métodos forem adicionados futuramente.
type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	// Injeta a dependência da camada de dados para operações de persistência.
	CategoryDB database.Category
}

// NewCategoryService cria e retorna uma nova instância de CategoryService.
func NewCategoryService(db database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: db,
	}
}

// CreateCategory implementa o método gRPC para criar uma nova categoria.
// Recebe um contexto (usado para controlar cancelamentos e deadlines, similar ao Request em Express)
// e uma requisição que contém o nome e a descrição da categoria.
func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	// Chama o método Create da camada de dados para inserir a categoria no banco.
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		// Em caso de erro, retorna um erro gRPC com o código Internal.
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Prepara a resposta com os dados da categoria criada.
	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
	// categoryResponse := &pb.CategoryResponse{
	// 	Category: &pb.Category{
	// 		Id:          category.ID,
	// 		Name:        category.Name,
	// 		Description: category.Description,
	// 	},
	// }
	return categoryResponse, nil
}

// ListCategories implementa o método gRPC para listar todas as categorias.
// Recebe um contexto (usado para controlar cancelamentos e deadlines, similar ao Request em Express).
func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	// Chama o método FindAll da camada de dados para obter todas as categorias do banco.
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		// Em caso de erro, retorna um erro gRPC com o código Internal.
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Prepara a resposta com a lista de categorias.
	categoriesResponse := &pb.CategoryList{
		Categories: make([]*pb.Category, len(categories)),
	}
	for i, category := range categories {
		categoriesResponse.Categories[i] = &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}
	}
	return categoriesResponse, nil

	// var categoriesResponse []*pb.Category

	// for _, category := range categories {
	// 	categoryResponse := &pb.Category{
	// 		Id:          category.ID,
	// 		Name:        category.Name,
	// 		Description: category.Description,
	// 	}

	// 	categoriesResponse = append(categoriesResponse, categoryResponse)
	// }

	// return &pb.CategoryList{Categories: categoriesResponse}, nil
}

// GetCategory implementa o método gRPC para obter uma categoria por ID.
// Recebe um contexto (usado para controlar cancelamentos e deadlines, similar ao Request em Express)
// e uma requisição que contém o ID da categoria.
func (c *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryGetRequest) (*pb.Category, error) {
	// Chama o método FindByCourseID da camada de dados para obter a categoria pelo ID.
	category, err := c.CategoryDB.Find(in.Id)
	if err != nil {
		// Em caso de erro, retorna um erro gRPC com o código Internal.
		return nil, status.Error(codes.Internal, err.Error())
	}
	// Prepara a resposta com os dados da categoria encontrada.
	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
	return categoryResponse, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	// Cria uma lista para acumular as categorias criadas.
	categories := &pb.CategoryList{}

	for {
		// Recebe mensagens do cliente. O método stream.Recv() bloqueia até receber uma mensagem.
		in, err := stream.Recv()
		// Se chegar ao final do stream (cliente fechou a conexão), envia a lista acumulada de categorias.
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}

		if err != nil {
			return err
		}

		// Processa cada CreateCategoryRequest, criando uma nova categoria na camada de dados.
		category, err := c.CategoryDB.Create(in.Name, in.Description)
		if err != nil {
			return err
		}

		// Acumula a categoria criada na lista de resposta.
		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
}

func (c *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {
	for {
		// Recebe uma mensagem do cliente.
		in, err := stream.Recv()
		if err == io.EOF {
			// Se o cliente fechar o stream, finaliza o método.
			return nil
		}

		if err != nil {
			return err
		}

		// Cria a categoria a partir da mensagem recebida.
		category, err := c.CategoryDB.Create(in.Name, in.Description)
		if err != nil {
			return err
		}

		// Envia a categoria criada imediatamente para o cliente.
		if err := stream.Send(&pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}); err != nil {
			return err
		}
	}
}
