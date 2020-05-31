package internal

// import (
// 	"os"
// 	"log"
// 	"testing"
//
// 	"github.com/stretchr/testify/assert"
//
// 	"order-matching/internal/config"
//   "order-matching/internal/order"
// 	"order-matching/internal/storage/redis"
// )
//
// const (
// 	// test config dependencies
// 	_env            = "test"
// 	_configFilePath = "../"
// )
//
// var (
// 	_orderManager *OrderManager
// 	_rdb *redis.RedisDb
//
//   _testOrder1 = &order.Order{
//     Location: order.OrderLocation{
// 			Lon: 43.45123431,
// 			Lat: 75.13124123,
// 		},
// 		Id:            "test_order_id1",
//  	Desc:          "test_order_description1",
// 		TimeRequested: "test_order_time_requested1",
// 		Duration:      "test_order_duration1",
// 		ConsumerId:    "test_order_consumer_id1",
// 		BidPrice:      1.234,
// 	}
//   _testOrder2 = &order.Order{
// 		Location: order.OrderLocation{
// 			Lon: 2.23,
// 			Lat: 2.43,
// 		},
// 		Id:            "test_order_id2",
// 		Desc:          "test_order_description2",
// 		TimeRequested: "test_order_time_requested2",
// 		Duration:      "test_order_duration2",
// 		ConsumerId:    "test_order_consumer_id2",
// 		BidPrice:      1.236,
// 	}
// 	_testOrder3 = &order.Order{
// 		Location: order.OrderLocation{
// 			Lon: 0,
// 			Lat: 0,
// 		},
// 		Id:            "",
// 		Desc:          "",
// 		TimeRequested: "",
// 		Duration:      "",
// 		ConsumerId:    "",
// 		BidPrice:      0,
// 	}
// )
//
// func TestMain(m *testing.M) {
// 	beforeAll()
// 	code := m.Run()
// 	afterAll()
// 	os.Exit(code)
// }
//
// func beforeAll() {
// 	testCfg, _ := config.NewConfig(_configFilePath, _env)
// 	testRedisCfg := &(*testCfg).Redis
// 	_rdb, _ = redis.NewRedisDb(testRedisCfg)
// 	_rdb.Clear()
// 	_orderManager = NewOrderManager(_rdb)
// }
//
// func afterAll() {
// 	_rdb.Clear()
// }
//
// func TestAddOrder(t *testing.T) {
//   // new test
//   // function under test
//   if err := _orderManager.AddNewOrder(_testOrder1); err != nil {
//     log.Fatalf(err.Error())
//   }
//   // function under test
//   order1, err := _orderManager.GetOrder(_testOrder1.Id)
//   if err != nil {
//     log.Fatalf(err.Error())
//   }
//   assert.Equal(t, _testOrder1, order1, "should add and then get order")
//
//   // new test
//   // function under test
//   if err := _orderManager.AddNewOrder(_testOrder2); err != nil {
//     log.Fatalf(err.Error())
//   }
//   // function under test
//   order2, err := _orderManager.GetOrder(_testOrder2.Id)
//   if err != nil {
//     log.Fatalf(err.Error())
//   }
//   assert.Equal(t, _testOrder2, order2, "should add and then get order")
//
//   // new test
//   // function under test
//   if err := _orderManager.AddNewOrder(_testOrder3); err != nil {
//     log.Fatalf(err.Error())
//   }
//   // function under test
//   order3, err := _orderManager.GetOrder(_testOrder3.Id)
//   if err != nil {
//     log.Fatalf(err.Error())
//   }
//   assert.Equal(t, _testOrder3, order3, "should add and then get order")
// }
//
// func TestGetOrder(t *testing.T) {
//   // new test
//   // function under test
//   _, err := _orderManager.GetOrder("non_existent_keyid")
//   assert.EqualError(t, err, "Error getting value using key 'non_existent_keyid': redigo: nil returned")
//
//   // new test
//   // function under test
//   _, err = _orderManager.GetOrder("test_order_id1 ")
//   assert.EqualError(t, err, "Error getting value using key 'test_order_id1 ': redigo: nil returned")
//
//   // new test
//   // function under test
//   _, err = _orderManager.GetOrder(" ")
//   assert.EqualError(t, err, "Error getting value using key ' ': redigo: nil returned")
// }
//
// func TestOrderExists(t *testing.T) {
//   // new test
//   // setup
//   if err := _orderManager.AddNewOrder(_testOrder1); err != nil {
//     log.Fatalf(err.Error())
//   }
//   // function under test
//   orderExists1 := _orderManager.OrderExists(_testOrder1.Id)
//   assert.True(t, orderExists1, "order should exist")
//
//   // new test
//   // setup
//   if err := _orderManager.AddNewOrder(_testOrder2); err != nil {
//     log.Fatalf(err.Error())
//   }
//   // function under test
//   orderExists2 := _orderManager.OrderExists(_testOrder2.Id)
//   assert.True(t, orderExists2, "order should exist")
//
//   // new test
//   // setup
//   if err := _orderManager.AddNewOrder(_testOrder3); err != nil {
//     log.Fatalf(err.Error())
//   }
//   // function under test
//   orderExists3 := _orderManager.OrderExists(_testOrder3.Id)
//   assert.True(t, orderExists3, "order should exist")
//
//   // new test
//   // function under test
//   orderExists4 := _orderManager.OrderExists("non_existent_keyid")
//   assert.False(t, orderExists4, "order should not exist")
//
//   // new test
//   // function under test
//   orderExists5 := _orderManager.OrderExists("test_order_id1 ")
//   assert.False(t, orderExists5, "order should not exist")
//
//   // new test
//   // function under test
//   orderExists6 := _orderManager.OrderExists("test_ord er_id2")
//   assert.False(t, orderExists6, "order should not exist")
//
//   // new test
//   // function under test
//   orderExists7 := _orderManager.OrderExists(" ")
//   assert.False(t, orderExists7, "order should not exist")
// }
//
// func TestCountOrders(t *testing.T) {
// 	_rdb.Clear()
//
// 	// new test
// 	// function under test
// 	orderCount0, _ := _orderManager.CountOrders()
// 	assert.Equal(t, 0, orderCount0, "order count should be 0")
//
// 	// new test
//   // setup
//   if err := _orderManager.AddNewOrder(_testOrder1); err != nil {
//     log.Fatalf(err.Error())
//   }
// 	// function under test
// 	orderCount1, _ := _orderManager.CountOrders()
// 	assert.Equal(t, 1, orderCount1, "order count should be 1")
//
// 	_rdb.Clear()
//
// 	// new test
//   // setup
//   if err := _orderManager.AddNewOrder(_testOrder1); err != nil {
//     log.Fatalf(err.Error())
//   }
// 	if err := _orderManager.AddNewOrder(_testOrder2); err != nil {
//     log.Fatalf(err.Error())
//   }
// 	// function under test
// 	orderCount2, _ := _orderManager.CountOrders()
// 	assert.Equal(t, 2, orderCount2, "order count should be 2")
//
// 	_rdb.Clear()
//
// 	// new test
// 	// setup
// 	if err := _orderManager.AddNewOrder(_testOrder1); err != nil {
// 		log.Fatalf(err.Error())
// 	}
// 	if err := _orderManager.AddNewOrder(_testOrder2); err != nil {
// 		log.Fatalf(err.Error())
// 	}
// 	if err := _orderManager.AddNewOrder(_testOrder3); err != nil {
// 		log.Fatalf(err.Error())
// 	}
// 	// function under test
// 	orderCount3, _ := _orderManager.CountOrders()
// 	assert.Equal(t, 3, orderCount3, "order count should be 3")
//
// 	_rdb.Clear()
// }
