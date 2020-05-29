package internal

import (
	"log"

	"github.com/pkg/errors"
  "order-matching/internal/order"
)

type Db interface {
	Select(keyId string) (string, error)
	Insert(keyId string, value string) error
	Delete(keyId string) error
  Exists(keyId string) (bool, error)
  Clear() error
}

// type OrderManager interface {
// 	AddOrder(consumerId string, params string) (string, error)
// 	// IsOrderAccepted(string) (bool, error)
// }

type OrderManager struct {
	db Db
}

func NewOrderManager(db Db) *OrderManager {
	return &OrderManager{
		db: db,
	}
}

func (om *OrderManager) AddOrder(params []byte) (string, error) {
  // order validation
  // - validate lon, lat
  // - validate price given > 0 and 2 decimal places (in schema definition? with minimum price)

  // create order struct
  order := order.NewOrder(params)

	orderExists, err := om.db.Exists(order.Id)
	if err != nil {
		return "", err
	}
	if orderExists {
		return "", errors.Errorf("Error, order with id '%s' already exists", order.Id)
	}

	orderBytes, err := order.MarshalJSON()
	if err != nil {
		return "", err
	}

	log.Printf("order str: %s", string(orderBytes))

  om.db.Insert(order.Id, string(orderBytes))
  return order.Id, nil
}
