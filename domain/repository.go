package domain

type OrderRepository interface {
	NextIdentity() string
	FormID(order_id string) (*Order, error)
	Save(entity Order) error
}
