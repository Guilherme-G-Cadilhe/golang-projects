// File: tax_extended_test.go

package tax

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestParallelCalculateTax demonstra como executar testes em paralelo.
// Isso é útil para testes que não dependem um do outro e podem ser executados simultaneamente.
func TestParallelCalculateTax(t *testing.T) {
	// Marca este teste e seus subtestes para execução paralela
	t.Parallel()

	// Define uma série de casos de teste
	testCases := []struct {
		name        string  // Nome descritivo do caso de teste
		input       float64 // Valor de entrada para CalculateTax
		expectedTax float64 // Valor esperado de imposto
	}{
		{"Case 1", 500, 5},
		{"Case 2", 1000, 10},
		{"Case 3", 20000, 20},
	}

	// Itera sobre cada caso de teste
	for _, tc := range testCases {
		// Importante: cria uma nova variável tc para cada iteração
		// Isso evita problemas de concorrência em testes paralelos
		tc := tc
		// Cria um subteste para cada caso
		t.Run(tc.name, func(t *testing.T) {
			// Marca este subteste para execução paralela
			t.Parallel()
			// Chama a função que estamos testando
			tax, err := CalculateTax(tc.input)
			// Verifica se não houve erro
			assert.NoError(t, err)
			// Verifica se o resultado é o esperado
			assert.Equal(t, tc.expectedTax, tax)
		})
	}
}

// TestCalculateTaxWithTimeout demonstra como usar um contexto com timeout em testes.
// Isso é útil para garantir que testes longos não executem indefinidamente.
func TestCalculateTaxWithTimeout(t *testing.T) {
	// Cria um contexto com timeout de 100 milissegundos
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	// Garante que a função cancel será chamada ao final da função
	defer cancel()

	// Cria um canal para sinalizar a conclusão do teste
	done := make(chan struct{})

	// Inicia uma goroutine para executar o teste
	go func() {
		// Fecha o canal done quando a goroutine terminar
		defer close(done)
		// Executa o teste
		tax, err := CalculateTax(1000)
		assert.NoError(t, err)
		assert.Equal(t, 10.0, tax)
	}()

	// Espera pelo primeiro evento: conclusão do teste ou timeout
	select {
	case <-ctx.Done():
		// Se o contexto expirou, o teste falhou
		t.Fatal("Test timed out")
	case <-done:
		// Se o canal done foi fechado, o teste concluiu a tempo
	}
}

// TestCalculateTaxWithSetupTeardown demonstra como usar setup e teardown em testes.
// Setup prepara o ambiente para o teste, e teardown limpa após o teste.
func TestCalculateTaxWithSetupTeardown(t *testing.T) {
	// Setup: prepara o ambiente de teste
	repo := &TaxRepositoryMock{}
	repo.On("SaveTax", 10.0).Return(nil)

	// Teardown: limpa o ambiente após o teste
	// defer garante que esta função será chamada ao final do teste
	defer func() {
		repo.AssertExpectations(t)
	}()

	// Teste principal
	err := CalculateTaxAndSave(1000, repo)
	assert.NoError(t, err)
}

// TestCalculateTaxSubtests demonstra como usar subtestes.
// Subtestes permitem agrupar testes relacionados e executá-los individualmente.
func TestCalculateTaxSubtests(t *testing.T) {
	// Grupo de testes para valores válidos
	t.Run("Valid amounts", func(t *testing.T) {
		// Subteste para valores baixos
		t.Run("Low range", func(t *testing.T) {
			tax, err := CalculateTax(500)
			require.NoError(t, err)
			assert.Equal(t, 5.0, tax)
		})
		// Subteste para valores médios
		t.Run("Mid range", func(t *testing.T) {
			tax, err := CalculateTax(10000)
			require.NoError(t, err)
			assert.Equal(t, 10.0, tax)
		})
		// Subteste para valores altos
		t.Run("High range", func(t *testing.T) {
			tax, err := CalculateTax(30000)
			require.NoError(t, err)
			assert.Equal(t, 20.0, tax)
		})
	})

	// Grupo de testes para valores inválidos
	t.Run("Invalid amounts", func(t *testing.T) {
		// Subteste para valor zero
		t.Run("Zero", func(t *testing.T) {
			_, err := CalculateTax(0)
			assert.Error(t, err)
		})
		// Subteste para valor negativo
		t.Run("Negative", func(t *testing.T) {
			_, err := CalculateTax(-100)
			assert.Error(t, err)
		})
	})
}

// TestExpectedFailure demonstra como testar casos que devem falhar.
// É importante testar não apenas os casos de sucesso, mas também os de erro.
func TestExpectedFailure(t *testing.T) {
	_, err := CalculateTax(-1)
	// Verifica se um erro foi retornado para um valor negativo
	if err == nil {
		t.Error("Expected an error for negative amount, but got nil")
	}
}

// TestCalculateTaxAndSaveOrder demonstra como usar mocks para verificar a ordem das chamadas.
// Isso é útil quando a sequência de operações é importante.
func TestCalculateTaxAndSaveOrder(t *testing.T) {
	// Cria um mock do repositório
	repo := &TaxRepositoryMock{}
	// Configura o mock para esperar duas chamadas em ordem específica
	repo.On("SaveTax", 10.0).Return(nil).Once()
	repo.On("SaveTax", 20.0).Return(nil).Once()

	// Executa a primeira operação
	err := CalculateTaxAndSave(1000, repo)
	assert.NoError(t, err)

	// Executa a segunda operação
	err = CalculateTaxAndSave(20000, repo)
	assert.NoError(t, err)

	// Verifica se todas as expectativas do mock foram atendidas
	repo.AssertExpectations(t)
}

// TestCalculateTaxPanic demonstra como testar funções que podem causar panic.
// Em Go, panic é usado para erros irrecuperáveis, e é importante saber testá-los.
func TestCalculateTaxPanic(t *testing.T) {
	// assert.Panics verifica se a função passada causa um panic
	assert.Panics(t, func() {
		// Simula uma situação que causaria panic
		// Nota: Este é apenas um exemplo, a função real não causa panic
		CalculateTax(1000000000)
	})
}

// Conceitos importantes demonstrados nestes testes:
// 1. Testes paralelos (t.Parallel())
// 2. Testes com timeout (context.WithTimeout)
// 3. Setup e teardown de testes
// 4. Subtestes (t.Run())
// 5. Testes de casos de erro
// 6. Uso de mocks para simular dependências
// 7. Verificação de panics
//
// Bibliotecas utilizadas:
// - "testing": pacote padrão do Go para testes
// - "github.com/stretchr/testify/assert": facilita asserções em testes
// - "github.com/stretchr/testify/require": similar ao assert, mas para por falha
// - "github.com/stretchr/testify/mock": auxilia na criação de mocks
//
// Dica: Para executar estes testes, use o comando `go test` no terminal
// dentro do diretório do projeto. Para ver mais detalhes, use `go test -v`.
