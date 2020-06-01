package internal

import (
	// "github.com/pkg/errors"
  "order-matching/internal/order"
)

type KeyDB interface {
  Insert(keyId string, val string) error
	Select(keyId string) (string, error)
	Delete(keyId string) error
  Exists(keyId string) (bool, error)
	CountKeys() (int, error)
  Clear() error
}

type GeoDB interface {
  Insert(keyId string, coord map[string]float64) error
  Select(keyId string) (map[string]float64, error)
	Delete(keyId string) error
  Clear() error
}

type OrderManager struct {
	keyDB KeyDB
  geoDB GeoDB
}

func NewOrderManager(keyDB KeyDB, geoDB GeoDB) *OrderManager {
	return &OrderManager{
		keyDB: keyDB,
    geoDB: geoDB,
	}
}

// tested
func (om *OrderManager) GetOrder(orderId string) (*order.Order, error) {
	orderStr, err := om.keyDB.Select(orderId)
	if err != nil {
		return nil, err
	}

	order := &order.Order{}
	order.UnmarshalJSON(orderStr)
	return order, nil
}

func (om *OrderManager) CountOrders() (int, error) {
	return om.keyDB.CountKeys()
}

func (om *OrderManager) AddNewOrder(order *order.Order) error {
  // order validation
  // - validate lon, lat
  // - validate price given > 0 and 2 decimal places (in schema definition? with minimum price)

	orderStr, err := order.MarshalJSON()
	if err != nil {
		return err
	}

	return om.keyDB.Insert(order.Id, orderStr)
}

// func (om *OrderManager) OrderAccepted(orderId string) error {
//
// }

// tested
func (om *OrderManager) OrderExists(orderId string) (bool, error) {
	return om.keyDB.Exists(orderId)
}

func (om *OrderManager) Clear() error {
  if err := om.keyDB.Clear(); err != nil {
    return err
  }
  return om.geoDB.Clear()
}
