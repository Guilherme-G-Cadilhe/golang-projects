package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// uma implementação de EventInterface para testes.
type TestEvent struct {
	Name    string      // Nome do evento
	Payload interface{} // Dados do evento
}

// retorna o nome do evento.
func (e *TestEvent) GetName() string {
	return e.Name
}

// retorna a data/hora atual (para fins de teste).
func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

// retorna os dados associados ao evento.
func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

// implementa a interface de manipulador de eventos para testes.
type TestEventHandler struct {
	ID int // Identificador único para diferenciar os handlers
}

func (e *TestEventHandler) Handle(event EventInterface) {
	// do something
}

// Agrupa os testes para o EventDispatcher.
type EventDispatcherTestSuite struct {
	suite.Suite
	event      TestEvent
	event2     TestEvent
	handler    TestEventHandler
	handler2   TestEventHandler
	handler3   TestEventHandler
	dispatcher *EventDispatcher
}

// SetupTest é o nome do metodo executado antes de cada teste na suíte.
// Equivalente a um "beforeEach" em frameworks de teste do JavaScript.
func (suite *EventDispatcherTestSuite) SetupTest() {
	// Cria um varias variaveis de teste dentro do contexto da Suite, que utilizam as interfaces de evento
	suite.event = TestEvent{Name: "test", Payload: "payload"}
	suite.event2 = TestEvent{Name: "test2", Payload: "payload2"}
	suite.handler = TestEventHandler{
		ID: 1,
	}
	suite.handler2 = TestEventHandler{
		ID: 2,
	}
	suite.handler3 = TestEventHandler{
		ID: 3,
	}
	suite.dispatcher = NewEventDispatcher()
}

func TestSuite(t *testing.T) {
	// Inicializa e executa a suíte de testes.
	suite.Run(t, new(EventDispatcherTestSuite))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	// Adiciona um handler para um evento chamado "test" criado no setup
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	// Verifica se a lista de handlers para o evento "test" contém 1 elemento.
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	// Registra outro handler para o mesmo evento "test".
	err = suite.dispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	// Verifica se agora a lista contém 2 handlers.
	suite.Equal(2, len(suite.dispatcher.handlers[suite.event.GetName()]))

	// Verifica se os handlers foram adicionados na ordem correta.
	assert.Equal(suite.T(), &suite.handler, suite.dispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.dispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_DuplicateHandler() {
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	// Registra o mesmo handler para o mesmo evento.
	err = suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	// Verifica se o erro retornado é o esperado.
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	//  Verifica se a lista de handlers ainda contém apenas 1 handler.
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event2.GetName()]))

	suite.dispatcher.Clear()
	suite.Equal(0, len(suite.dispatcher.handlers))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_has() {
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.dispatcher.handlers[suite.event.GetName()]))

	assert.True(suite.T(), suite.dispatcher.Has(suite.event.GetName(), &suite.handler))
	assert.True(suite.T(), suite.dispatcher.Has(suite.event.GetName(), &suite.handler2))
	assert.False(suite.T(), suite.dispatcher.Has(suite.event.GetName(), &suite.handler3))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface) {
	m.Called(event)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eventH := &MockHandler{}
	eventH.On("Handle", &suite.event)
	suite.dispatcher.Register(suite.event.GetName(), eventH)

	suite.dispatcher.Dispatch(&suite.event)
	eventH.AssertExpectations(suite.T())
	eventH.AssertNumberOfCalls(suite.T(), "Handle", 1)

}
