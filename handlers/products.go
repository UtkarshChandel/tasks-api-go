package handlers

import (
	"log"
	"net/http"

	"github.com/UtkarshChandel/tasks-api-go/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal the JSON", http.StatusInternalServerError)
	}

}


func (p *Products) getProducts(rw http.ResponseWriter,h* http.Request )