package tax

import "testing"

/*
[ Testar ]
> go test . (Diretorio atual)
> go test (Pacote atual)
> go test -coverprofile=coverage.out (Gera arquivo coverage e testa se engloba todas os cases)
> go tool cover -html=coverage.out (Lê arquivo coverage gerado e abre página web mostrando funções que faltam)
> go test -bench=. (Diretorio atual, Mostra o tempo de cada teste e configurações do sistema, como cpu, goos, goarch, pacote, nucleos, etc.)
> go test -bench=. -run=^# (Roda todos os testes que tenha o nome do regex)
> go test -bench=. -run=^# -count=10 (Roda todos os testes que tenha o nome do regex e executa 10 vezes)
> go test -bench=. -run=^# -count=10 -benchtime=3s (Continua rodando por 3 segundos cada função)
> go test -bench=. -run=^# -benchmem (Mostra o quanto de memória foi usado)
Flags:
-v (Executa os testes em modo verboso, mostrando detalhes de cada execução)
-coverprofile=coverage.out (Gera um relatório de cobertura que pode ser visualizado em um navegador para identificar quais partes do código foram testadas.)
-bench=. (Executa todos os benchmarks no pacote atual.)
-run=^# (Utilizado para rodar testes específicos com base em regex.)
-count=10 (Define o número de vezes que os testes devem ser executados.)
-benchtime=3s (Define o tempo de execução de cada benchmark.)
-benchmem (Mostra a quantidade de memória alocada durante a execução do benchmark.)
*/

// Toda função teste tem que começar com "Test" antes do nome da função
// A função `TestCalculateTax` é um teste unitário básico para a função `CalculateTax`.
// O valor `500.0` é passado como entrada, e espera-se que o imposto seja `5.0`.
// Se o resultado não for igual ao esperado, o teste falha e uma mensagem de erro é exibida.
func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0 // Teste será aprovado, mas pode-se mudar para 6.0 para testar falha

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

// `TestCalculateTaxBatch` realiza testes em batch usando uma tabela de testes.
// Uma tabela de testes é uma maneira eficiente de executar múltiplos casos de teste com diferentes entradas e saídas esperadas.
// Aqui, a tabela define vários cenários de teste com valores e resultados esperados diferentes.
func TestCalculateTaxBatch(t *testing.T) {

	// Define a estrutura para a tabela de testes, com duas propriedades: amount e expect.
	type calcTax struct {
		amount, expect float64
	}

	// Tabela de testes com vários casos de teste.
	table := []calcTax{
		{500.0, 5.0},   // Valor médio, imposto esperado é 5.0
		{1000.0, 10.0}, // Valor na borda superior, imposto esperado é 10.0
		{1500.0, 10.0}, // Valor acima de 1000, imposto esperado é 10.0
		{0.0, 0.0},     // Valor zero, imposto esperado é 0.0
	}

	// Loop sobre a tabela de testes para executar cada caso.
	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expect {
			t.Errorf("Expected %f, but got %f", item.expect, result)
		}
	}
}

// `BenchmarkCalculateTax` é um teste de benchmark que mede o desempenho da função `CalculateTax`.
// O teste roda a função várias vezes (baseado em `b.N`) e mede o tempo gasto.
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0) // Roda a função com o valor fixo de 500.0
	}
}

// `BenchmarkCalculateTax2` é um teste de benchmark para a função `CalculateTax2`, que inclui um delay artificial.
// Isso é útil para comparar o desempenho de funções semelhantes com e sem delays.
func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}
