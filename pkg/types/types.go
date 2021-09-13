package types

// Amount of money in smallest units (cent, diram etc.)
type Money int64

// Represent the code of currency
type Currency string

// Codes of currency
const (
	TJS Currency = "TJS"
	USD Currency = "USD"
	RUB Currency = "RUB"
)

// Number of card
type PAN string

// Card represents information of card
type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color   string
	Name     string
	Active bool
	MinBalance Money
}

// куда отправили платеж
type Category string

//Payment represents information about payments, but this comment is redundant
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status Status
}

type PaymentSource struct {
	Type string //'card'
	Number string  // номер вида '5058 xxxx xxxx 8888'
	Balance Money
}

type Status string

const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress  = "INPROGRESS"

)