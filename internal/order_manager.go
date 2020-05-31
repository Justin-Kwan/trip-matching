package internal

import (
	// "github.com/pkg/errors"
  "order-matching/internal/order"
)

type KeyStore interface {
	Select(keyId string) (string, error)
	Insert(keyId string, val string) error
	Delete(keyId string) error
  Exists(keyId string) bool
	CountKeys() (int, error)
  Clear() error
}

// type OrderManager interface {
// 	AddOrder(consumerId string, params string) (string, error)
// 	// IsOrderAccepted(string) (bool, error)
// }

type OrderManager struct {
	db KeyStore
}

func NewOrderManager(ks KeyStore) *OrderManager {
	return &OrderManager{
		db: ks,
	}
}

// tested
func (om *OrderManager) GetOrder(orderId string) (*order.Order, error) {
	orderStr, err := om.db.Select(orderId)
	if err != nil {
		return nil, err
	}

	order := &order.Order{}
	order.UnmarshalJSON(orderStr)
	return order, nil
}

func (om *OrderManager) CountOrders() (int, error) {
	return om.db.CountKeys()
}

// accepts order struct (deals with orders)
func (om *OrderManager) AddNewOrder(order *order.Order) error {
  // order validation
  // - validate lon, lat
  // - validate price given > 0 and 2 decimal places (in schema definition? with minimum price)

	orderStr, err := order.MarshalJSON()
	if err != nil {
		return err
	}



	return om.db.Insert(order.Id, orderStr)
}

// func (om *OrderManager) OrderAccepted(orderId string) error {
//
// }

// tested
func (om *OrderManager) OrderExists(orderId string) bool {
	return om.db.Exists(orderId)
}
