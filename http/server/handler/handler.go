package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/RuGoffer/gofr/internal/service"
)

// создадим новый тип для добавления middleware к обработчикам
type Decorator func(http.Handler) http.Handler

// объект для хранения состояния игры
type LifeStates struct {
	service.LifeService
}

func New(_ context.Context,
	lifeService service.LifeService,
) (http.Handler, error) {
	serveMux := http.NewServeMux()

	lifeState := LifeStates{
		LifeService: lifeService,
	}

	serveMux.HandleFunc("/nextstate", lifeState.nextState)

	return serveMux, nil
}

// функция добавления middleware
func Decorate(next http.Handler, ds ...Decorator) http.Handler {
	decorated := next
	for d := len(ds) - 1; d >= 0; d-- { // why reversed???
		decorated = ds[d](decorated)
	}

	return decorated
}

// получение очередного состояния игры
func (ls *LifeStates) nextState(w http.ResponseWriter, _ *http.Request) {
	worldState := ls.LifeService.NewState()

	err := json.NewEncoder(w).Encode(worldState.Cells)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
