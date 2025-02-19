package events

import (
	"sync"
	"time"
)

// EventInterface define o contrato para um evento.
// Um evento deve poder fornecer seu nome, a data/hora em que ocorreu e seu payload.
type EventInterface interface {
	GetName() string         // Retorna o nome do evento.
	GetDateTime() time.Time  // Retorna a data e hora em que o evento ocorreu.
	GetPayload() interface{} // Retorna os dados associados ao evento.
}

// EventHandlerInterface define o contrato para um manipulador de eventos.
// Um handler deve ser capaz de lidar com um evento recebido.
type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup) // Processa o evento.
}

// EventDispatcherInterface define o contrato para um despachante de eventos.
// Ele deve permitir registrar, remover, verificar, despachar e limpar handlers para eventos.
type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Dispatch(event EventInterface) error
	Clear() error
}
