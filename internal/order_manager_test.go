package internal

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"order-matching/internal/config"
	"order-matching/internal/order"
	"order-matching/internal/storage/redis"
)

const (
	_floatMaxDelta = 0.01001
)

var (
	_om    *OrderManager
	_geoDB *redis.GeoDB
	_keyDB *redis.KeyDB

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
	setIndex       string
}

func newOrderManagerTestConstants() *orderManagerTestConstants {
	return &orderManagerTestConstants{
		configFilePath: "../",
		env:            "test",
		setIndex:       "test_index",
	}
}

func setupTests() func() {
	tc := newOrderManagerTestConstants()

	cfg, _ := config.NewConfig(tc.configFilePath, tc.env)
	keyDBPool := redis.NewPool(&(*cfg).RedisKeyDB)
	geoDBPool := redis.NewPool(&(*cfg).RedisGeoDB)

	_keyDB = redis.NewKeyDB(keyDBPool)
	_geoDB = redis.NewGeoDB(geoDBPool, tc.setIndex)

	_keyDB.Clear()
	_geoDB.Clear()

	_om = NewOrderManager(_keyDB, _geoDB)

	return func() {
		_om.Clear()
	}
}

func TestAddOrder(t *testing.T) {
	teardownTests := setupTests()
	defer teardownTests()

	for _, testOrder := range testOrders {
		// function under test
		if err := _om.AddNewOrder(&testOrder); err != nil {
			log.Fatalf(err.Error())
		}

		// assert correct order was added in key value db
		orderStr, err := _keyDB.Select(testOrder.Id)
		if err != nil {
			teardownTests()
			log.Fatalf(err.Error())
		}

		order := &order.Order{}
		if err := order.UnmarshalJSON(orderStr); err != nil {
			log.Fatalf(err.Error())
		}

		assert.Equal(t, testOrder, *order, "should add correct order in key value db")

		// assert correct order was added in geo db
		coord, err := _geoDB.Select(testOrder.Id)
		if err != nil {
			teardownTests()
			log.Fatal(err.Error())
		}

		assert.InDelta(t, testOrder.Location.Lon, coord["lon"], _floatMaxDelta)
		assert.InDelta(t, testOrder.Location.Lat, coord["lat"], _floatMaxDelta)
	}
}

func TestGetOrder(t *testing.T) {
	teardownTests := setupTests()
	defer teardownTests()

	// function under test
	_, err := _om.GetOrder("non_existent_id")
	assert.EqualError(t, err, "Error getting value using key 'non_existent_id': redigo: nil returned")

	// function under test
	_, err = _om.GetOrder("test_order_id1 ")
	assert.EqualError(t, err, "Error getting value using key 'test_order_id1 ': redigo: nil returned")

	// function under test
	_, err = _om.GetOrder(" ")
	assert.EqualError(t, err, "Error getting value using key ' ': redigo: nil returned")
}

func TestOrderExists(t *testing.T) {
	teardownTests := setupTests()
	defer teardownTests()

	for _, testOrder := range testOrders {
		// setup
		if err := _om.AddNewOrder(&testOrder); err != nil {
			log.Fatalf(err.Error())
		}

		// function under test
		orderExists, err := _om.OrderExists(testOrder.Id)
		if err != nil {
			teardownTests()
			log.Fatalf(err.Error())
		}

		assert.True(t, orderExists, "order should exist")
	}

	// function under test
	orderExists, err := _om.OrderExists("non_existent_id")
	if err != nil {
		teardownTests()
		log.Fatalf(err.Error())
	}

	assert.False(t, orderExists, "order should not exist")

	// function under test
	orderExists, err = _om.OrderExists("test_order_id1 ")
	if err != nil {
		teardownTests()
		log.Fatalf(err.Error())
	}

	assert.False(t, orderExists, "order should not exist")

	// function under test
	orderExists, err = _om.OrderExists("test_ord er_id2")
	if err != nil {
		teardownTests()
		log.Fatalf(err.Error())
	}

	assert.False(t, orderExists, "order should not exist")

	// function under test
	orderExists, err = _om.OrderExists(" ")
	if err != nil {
		teardownTests()
		log.Fatalf(err.Error())
	}

	assert.False(t, orderExists, "order should not exist")
}

func TestDeleteOrder(t *testing.T) {
	for _, testOrder := range testOrders {
		// setup
		if err := _om.AddNewOrder(&testOrder); err != nil {
			log.Fatalf(err.Error())
		}

		// function under test
		_om.DeleteOrder(testOrder.Id)

		// assert order deleted from key value db
		_, err := _keyDB.Select(testOrder.Id)
		assert.EqualError(t, err, "Error getting value using key '"+testOrder.Id+"': redigo: nil returned")

		// assert order deleted from geo db
		_, err = _geoDB.Select(testOrder.Id)
		assert.EqualError(t, err, "Error selecting POI with key '"+testOrder.Id+"'")
	}
}

func TestCountOrders(t *testing.T) {
	teardownTests := setupTests()
	defer teardownTests()

	// function under test
	orderCount, err := _om.CountOrders()
	if err != nil {
		teardownTests()
		log.Fatalf(err.Error())
	}

	assert.Equal(t, 0, orderCount)

	for i, testOrder := range testOrders {
		// setup
		ordersAdded := i + 1
		if err := _om.AddNewOrder(&testOrder); err != nil {
			log.Fatalf(err.Error())
		}

		// function under test
		orderCount, err := _om.CountOrders()
		if err != nil {
			teardownTests()
			log.Fatalf(err.Error())
		}

		assert.Equal(t, ordersAdded, orderCount, "order count should match number of orders inserted")
	}
}
