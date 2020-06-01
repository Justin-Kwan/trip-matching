package internal

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"order-matching/internal/config"
  "order-matching/internal/order"
	"order-matching/internal/storage/redis"
)


var (
	_orderManager *OrderManager

  testOrders = []order.Order{
    order.Order{
      Location: order.OrderLocation{
  			Lon: 43.45123431,
  			Lat: 75.13124123,
  		},
  		Id:            "test_order_id1",
   	  Desc:          "test_order_description1",
  		TimeRequested: "test_order_time_requested1",
  		Duration:      "test_order_duration1",
  		ConsumerId:    "test_order_consumer_id1",
  		BidPrice:      1.234,
  	},
    order.Order{
  		Location: order.OrderLocation{
  			Lon: 2.23,
  			Lat: 2.43,
  		},
  		Id:            "test_order_id2",
  		Desc:          "test_order_description2",
  		TimeRequested: "test_order_time_requested2",
  		Duration:      "test_order_duration2",
  		ConsumerId:    "test_order_consumer_id2",
  		BidPrice:      1.236,
  	},
    order.Order{
  		Location: order.OrderLocation{
  			Lon: 0,
  			Lat: 0,
  		},
  		Id:            "",
  		Desc:          "",
  		TimeRequested: "",
  		Duration:      "",
  		ConsumerId:    "",
  		BidPrice:      0,
  	},
  }
)

type orderManagerTestConstants struct {
	configFilePath string
	env            string
	dbNum          int
	setIndex       string
}

func newOrderManagerTestConstants() *orderManagerTestConstants {
	return &orderManagerTestConstants{
		configFilePath: "../",
		env:            "test",
	}
}

func setupTests() func() {
	tc := newOrderManagerTestConstants()

	testCfg, _ := config.NewConfig(tc.configFilePath, tc.env)
	testRedisCfg := &(*testCfg).Redis
	redisPool, _ := redis.NewPool(testRedisCfg)

	geoDB := redis.NewGeoDB(redisPool, tc.dbNum, tc.setIndex)
  keyDB := redis.NewKeyDB(redisPool, tc.dbNum)

  geoDB.Clear()
  keyDB.Clear()

  _orderManager = NewOrderManager(keyDB, geoDB)

	return func() {
		_orderManager.Clear()
	}
}

func TestAddOrder(t *testing.T) {
  teardownTests := setupTests()
  defer teardownTests()

  for _, testOrder := range testOrders {
    // function under test
    if err := _orderManager.AddNewOrder(&testOrder); err != nil {
      log.Fatalf(err.Error())
    }

    // assert correct order was added
    order, err := _orderManager.GetOrder(testOrder.Id)
    if err != nil {
      log.Fatalf(err.Error())
    }

    assert.Equal(t, testOrder, *order, "should add and then get order")
  }
}

func TestGetOrder(t *testing.T) {
  teardownTests := setupTests()
  defer teardownTests()

  // function under test
  _, err := _orderManager.GetOrder("non_existent_id")
  assert.EqualError(t, err, "Error getting value using key 'non_existent_id': redigo: nil returned")

  // function under test
  _, err = _orderManager.GetOrder("test_order_id1 ")
  assert.EqualError(t, err, "Error getting value using key 'test_order_id1 ': redigo: nil returned")

  // function under test
  _, err = _orderManager.GetOrder(" ")
  assert.EqualError(t, err, "Error getting value using key ' ': redigo: nil returned")
}

func TestOrderExists(t *testing.T) {
  teardownTests := setupTests()
  defer teardownTests()

  for _, testOrder := range testOrders {
    // setup
    if err := _orderManager.AddNewOrder(&testOrder); err != nil {
      log.Fatalf(err.Error())
    }

    // function under test
    orderExists, err := _orderManager.OrderExists(testOrder.Id)
    if err != nil {
      log.Fatalf(err.Error())
    }

    assert.True(t, orderExists, "order should exist")
  }

  // function under test
  orderExists, err := _orderManager.OrderExists("non_existent_id")
  if err != nil {
    log.Fatalf(err.Error())
  }

  assert.False(t, orderExists, "order should not exist")

  // function under test
  orderExists, err = _orderManager.OrderExists("test_order_id1 ")
  if err != nil {
    log.Fatalf(err.Error())
  }

  assert.False(t, orderExists, "order should not exist")

  // function under test
  orderExists, err = _orderManager.OrderExists("test_ord er_id2")
  if err != nil {
    log.Fatalf(err.Error())
  }

  assert.False(t, orderExists, "order should not exist")

  // function under test
  orderExists, err = _orderManager.OrderExists(" ")
  if err != nil {
    log.Fatalf(err.Error())
  }

  assert.False(t, orderExists, "order should not exist")
}

func TestCountOrders(t *testing.T) {
  teardownTests := setupTests()
  defer teardownTests()

  // function under test
  orderCount, err := _orderManager.CountOrders()
  if err != nil {
    log.Fatalf(err.Error())
  }

  assert.Equal(t, 0, orderCount)

  for i, testOrder := range testOrders {
    // setup
    ordersAdded := i + 1
    if err := _orderManager.AddNewOrder(&testOrder); err != nil {
      log.Fatalf(err.Error())
    }

    // function under test
    orderCount, err := _orderManager.CountOrders()
    if err != nil {
      log.Fatalf(err.Error())
    }

    assert.Equal(t, ordersAdded, orderCount)
  }
}
