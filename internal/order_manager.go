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
func (om *OrderManager) AddNewOrder(order *order.Order) error {
	orderStr, err := order.MarshalJSON()
	if err != nil {
		return err
	}

	if err := om.keyDB.Insert(order.Id, orderStr); err != nil {
		return err
	}

	coords := map[string]float64{
		"lon": order.Location.Lon,
		"lat": order.Location.Lat,
	}

	return om.geoDB.Insert(order.Id, coords)
}

// tested, TODO: make name better
// is this function needed?
func (om *OrderManager) GetOrder(orderId string) (*order.Order, error) {
	orderStr, err := om.keyDB.Select(orderId)
	if err != nil {
		return nil, err
	}

	order := &order.Order{}
	order.UnmarshalJSON(orderStr)
	return order, nil
}

// func (om *OrderManager) GetOrdersInRadius(lon float64, lat float64, radius float64) {
// }

// tested
func (om *OrderManager) DeleteOrder(orderId string) error {
	if err := om.keyDB.Delete(orderId); err != nil {
		return err
	}

  return om.geoDB.Delete(orderId)
}

// tested
func (om *OrderManager) CountOrders() (int, error) {
	return om.keyDB.CountKeys()
}

// tested
func (om *OrderManager) OrderExists(orderId string) (bool, error) {
	return om.keyDB.Exists(orderId)
}

// TODO: test
func (om *OrderManager) Clear() error {
	if err := om.keyDB.Clear(); err != nil {
		return err
	}
	return om.geoDB.Clear()
}
