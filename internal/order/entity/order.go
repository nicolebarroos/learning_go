package entity

import "errors"

// Interface define um comportamento de um tipo, no caso o comportamento de salvar o objeto
type OrderRepositoryInterface interface {
	Save(order *Order) error
}

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func (o Order) IsValid() error {
	if o.ID == "" {
		return errors.New("invalid id")
	}
	if o.Price == 0 {
		return errors.New("invalid price")
	}
	if o.Tax == 0 {
		return errors.New("invalid tax")
	}
	return nil
}

// (parâmetros de entrada) (retornos)
func NewOrder(id string, price float64, tax float64) (*Order, error) {
	//&: apontando para o endereçamento da variável na memória
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	//Se os is valid retornar um erro, retorno o order em branco e retorno o erro
	//Se os is valid retornar sucesso de fato, retorno minha order e retorno o erro em branco
	err := order.IsValid()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}
