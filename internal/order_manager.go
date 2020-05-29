package internal

import (
	"github.com/pkg/errors"
  "order-matching/internal/order"
)

type DB interface {
	Get(keyId string) (string, error)
	Add(keyId string, value string) error
	Delete(keyId string) error
  Exists(keyId string) (bool, error)
  Clear() error
}

// type OrderManager interface {
// 	AddOrder(consumerId string, params string) (string, error)
// 	// IsOrderAccepted(string) (bool, error)
// }

type OrderManager struct {
	db DB
}

func OrderManager(db DB) OrderManager {
	return &orderManager{
		db: db
	}
}

func (m *orderManager) AddOrder(params []byte) (string, error) {
  // order validation
  // - validate lon, lat
  // - validate price given > 0 and 2 decimal places (in schema definition? with minimum price)

  // create order struct
  o := order.NewOrder(params)

	orderExists, err := m.db.Exists(o.Id)
	if err != nil {
		return "", err
	}
	if orderExists {
		return "", errors.Errorf("Error, order with id '%s' already exists", o.Id)
	}

  m.db.Add(o.Id, o.marshalJson())
  return o.Id
}
