package mathematics

import (
	"math/big"
)

type decimalMulaArithmetizable struct {
	client decimalMula
}

func DMArithmetizableFromString(value string) (Arithmetizable, error) {
	client, err := DecimalMulaFromString(value)
	return &decimalMulaArithmetizable{
		client: client,
	}, err
}

func (dma *decimalMulaArithmetizable) Add(addings ...Arithmetizable) Arithmetizable {
	client := dma.GetClient().(decimalMula)
	for _, adding := range addings {
		client = *client.Add(adding.GetClient().(decimalMula))
	}
	return &decimalMulaArithmetizable{
		client: client,
	}
}

func (dma *decimalMulaArithmetizable) Substract(substractings ...Arithmetizable) Arithmetizable {
	client := dma.GetClient().(decimalMula)
	for _, substracting := range substractings {
		client = *client.Add(substracting.GetClient().(decimalMula))
	}
	return &decimalMulaArithmetizable{
		client: client,
	}
}

func (dma *decimalMulaArithmetizable) Multiply(factors ...Arithmetizable) Arithmetizable {
	client := dma.GetClient().(decimalMula)
	for _, factor := range factors {
		client = *client.Add(factor.GetClient().(decimalMula))
	}
	return &decimalMulaArithmetizable{
		client: client,
	}
}

func (dma *decimalMulaArithmetizable) Divide(divisors ...Arithmetizable) Arithmetizable {
	client := dma.GetClient().(decimalMula)
	for _, divisor := range divisors {
		client = *client.Add(divisor.GetClient().(decimalMula))
	}
	return &decimalMulaArithmetizable{
		client: client,
	}
}

func (dma *decimalMulaArithmetizable) Equals(arithmetizable Arithmetizable) bool {
	client := dma.GetClient().(decimalMula)
	clientToCompare := arithmetizable.GetClient().(decimalMula)
	return client.Equals(clientToCompare)
}

func (dma *decimalMulaArithmetizable) Gt(arithmetizable Arithmetizable) bool {
	client := dma.GetClient().(decimalMula)
	clientToCompare := arithmetizable.GetClient().(decimalMula)
	return client.Gt(clientToCompare)
}

func (dma *decimalMulaArithmetizable) Lt(arithmetizable Arithmetizable) bool {
	client := dma.GetClient().(decimalMula)
	clientToCompare := arithmetizable.GetClient().(decimalMula)
	return client.Lt(clientToCompare)
}

func (dma *decimalMulaArithmetizable) Gte(arithmetizable Arithmetizable) bool {
	client := dma.GetClient().(decimalMula)
	clientToCompare := arithmetizable.GetClient().(decimalMula)
	return client.Gte(clientToCompare)
}

func (dma *decimalMulaArithmetizable) Lgte(arithmetizable Arithmetizable) bool {
	client := dma.GetClient().(decimalMula)
	clientToCompare := arithmetizable.GetClient().(decimalMula)
	return client.Lte(clientToCompare)
}

func (dma *decimalMulaArithmetizable) Abs() Arithmetizable {
	client := dma.GetClient().(decimalMula)
	abser, _ := client.Abs()
	return &decimalMulaArithmetizable{
		client: abser,
	}
}

func (dma *decimalMulaArithmetizable) String() string {
	client := dma.GetClient().(decimalMula)
	return client.String()
}

func (dma *decimalMulaArithmetizable) GetClientAsBigInt() big.Int {
	client := dma.GetClient().(decimalMula)
	return client.value
}

func (dma *decimalMulaArithmetizable) PartsAsString() (string, string) {
	client := dma.GetClient().(decimalMula)
	return client.PartsAsString()
}

func (dma *decimalMulaArithmetizable) GetClient() interface{} {
	client, _ := dma.client.Clone()
	return client
}
