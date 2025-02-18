package events

import (
	"errors"
	"sync"
)

// Erro retornado se um handler já estiver registrado para um evento.
var ErrHandlerAlreadyRegistered = errors.New("this handler is already registered")

// EventDispatcher gerencia a inscrição e o despacho de eventos.
// Ele mantém um mapa de "eventName" para uma lista de EventHandlerInterface.
type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

// NewEventDispatcher cria e retorna um novo EventDispatcher.
func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

// Dispatch
func (eventDispatcher *EventDispatcher) Dispatch(event EventInterface) error {
	// Verifica se houver handlers registrados para este evento.
	if handlers, ok := eventDispatcher.handlers[event.GetName()]; ok {
		// Despacha o evento para todos os handlers registrados.
		waitgroup := &sync.WaitGroup{}
		for _, handler := range handlers {
			waitgroup.Add(1)
			go handler.Handle(event, waitgroup)
		}
		waitgroup.Wait()
	}
	return nil
}

// Register adiciona um novo handler para um evento específico.
// Se o handler já estiver registrado para o mesmo evento, retorna um erro.
func (eventDispatcher *EventDispatcher) Register(eventName string, newHandler EventHandlerInterface) error {
	// Se já houver handlers registrados para este evento, verifica se o novo já existe.
	if _, ok := eventDispatcher.handlers[eventName]; ok {
		for _, handler := range eventDispatcher.handlers[eventName] {
			if handler == newHandler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	// Adiciona o novo handler à lista de handlers para este evento.
	eventDispatcher.handlers[eventName] = append(eventDispatcher.handlers[eventName], newHandler)
	return nil
}

// Remove um handler de um evento.
func (eventDispatcher *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	// Verifica se houver handlers registrados para este evento.
	if _, ok := eventDispatcher.handlers[eventName]; ok {
		// Remove o handler da lista de handlers para este evento.
		for i, registeredHandler := range eventDispatcher.handlers[eventName] {
			// Utiliza o Index do loop para remover a posição correta.
			if registeredHandler == handler {
				eventDispatcher.handlers[eventName] = append(eventDispatcher.handlers[eventName][:i], eventDispatcher.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}

// Limpa todos eventos
func (eventDispatcher *EventDispatcher) Clear() error {
	eventDispatcher.handlers = make(map[string][]EventHandlerInterface)
	return nil
}

// Verifica se um handler está registrado para um evento.
func (eventDispatcher *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	// Verifica se houver handlers registrados para este evento.
	if _, ok := eventDispatcher.handlers[eventName]; ok {
		// Verifica se o handler está registrado para este evento.
		for _, registeredHandler := range eventDispatcher.handlers[eventName] {
			if registeredHandler == handler {
				return true
			}
		}
	}
	return false
}
