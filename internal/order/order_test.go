package order

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	// test order json payloads to create new order structs
	// (with id and timeRequested fields omitted)
	_testOrderParamStr1 = `{"location":{"lon":43.45123431,"lat":75.13124123},"description":"test_order_description1","consumerId":"test_order_consumer_id1","bidPrice":100.23}`
	_testOrderParamStr2 = `{"location":{"lon":43.45123432,"lat":75.13124122},"description":"test_order_description2","consumerId":"test_order_consumer_id2","bidPrice":200.23}`
	_testOrderParamStr3 = `{"location":{"lon":0,"lat":0},"description":"","consumerId":"","bidPrice":0}`

	// test order json payloads to marshal to order structs
	_testOrderStr1    = `{"location":{"lon":1.23,"lat":1.63},"id":"test_order_id1","description":"test_order_description1","timeRequested":"test_order_time_requested1","duration":"test_order_duration1","consumerId":"test_order_consumer_id1","bidPrice":1.234}`
	_testOrderStr2    = `{"location":{"lon":2.23,"lat":2.43},"id":"test_order_id2","description":"test_order_description2","timeRequested":"test_order_time_requested2","duration":"test_order_duration2","consumerId":"test_order_consumer_id2","bidPrice":1.236}`
	_testOrderStr3    = `{"location":{"lon":0,"lat":0},"id":"","description":"","timeRequested":"","duration":"","consumerId":"","bidPrice":0}`
	_testBadOrderStr4 = `{"location":{"lon": "should not be a string","lat":0},"id":"test_order_id2","description":"test_order_description2","timeRequested":"test_order_time_requested2","duration":"test_order_duration2","consumerId":"test_order_consumer_id2","bidPrice":654.2312451232}`
	_testBadOrderStr5 = `{"location":{"lon":123,"lat":234},"id":"test_order_id2","description":"test_order_description2","timeRequested":"test_order_time_requested2","duration":"test_order_duration2","consumerId":"test_order_consumer_id2","bidPrice":"should not be a string"}`
)

var (
	// test order structs for NewOrder function
	// (with id and timeRequested fields omitted)
	_testOrderParam1 = &Order{
		Location: OrderLocation{
			Lon: 43.45123431,
			Lat: 75.13124123,
		},
		Desc:       "test_order_description1",
		ConsumerId: "test_order_consumer_id1",
		BidPrice:   100.23,
	}
	_testOrderParam2 = &Order{
		Location: OrderLocation{
			Lon: 43.45123432,
			Lat: 75.13124122,
		},
		Desc:       "test_order_description2",
		ConsumerId: "test_order_consumer_id2",
		BidPrice:   200.23,
	}
	_testOrderParam3 = &Order{
		Location: OrderLocation{
			Lon: 0,
			Lat: 0,
		},
		Desc:       "",
		ConsumerId: "",
		BidPrice:   0,
	}

	// test order structs for json marshal/unmarshal functions
	_testOrder1 = &Order{
		Location: OrderLocation{
			Lon: 1.23,
			Lat: 1.63,
		},
		Id:            "test_order_id1",
		Desc:          "test_order_description1",
		TimeRequested: "test_order_time_requested1",
		Duration:      "test_order_duration1",
		ConsumerId:    "test_order_consumer_id1",
		BidPrice:      1.234,
	}
	_testOrder2 = &Order{
		Location: OrderLocation{
			Lon: 2.23,
			Lat: 2.43,
		},
		Id:            "test_order_id2",
		Desc:          "test_order_description2",
		TimeRequested: "test_order_time_requested2",
		Duration:      "test_order_duration2",
		ConsumerId:    "test_order_consumer_id2",
		BidPrice:      1.236,
	}
	_testOrder3 = &Order{
		Location: OrderLocation{
			Lon: 0,
			Lat: 0,
		},
		Id:            "",
		Desc:          "",
		TimeRequested: "",
		Duration:      "",
		ConsumerId:    "",
		BidPrice:      0,
	}
)

// asserts actual order struct's fields matches the expected order
// struct's fields, except for id and timeRequested (which are dynamic).
func assertOrderValid(t *testing.T, actual *Order, expected *Order) {
	assert.Equal(t, actual.Location.Lon, expected.Location.Lon)
	assert.Equal(t, actual.Location.Lat, expected.Location.Lat)
	assert.NotEmpty(t, expected.Id)
	assert.Equal(t, actual.Desc, expected.Desc)
	assert.NotEmpty(t, expected.TimeRequested)
	assert.Equal(t, actual.ConsumerId, expected.ConsumerId)
	assert.Equal(t, actual.BidPrice, expected.BidPrice)
}

func TestNewOrder(t *testing.T) {
	// new test
	// function under test
	order1 := NewOrder(_testOrderParamStr1)
	assertOrderValid(t, _testOrderParam1, order1)

	// new test
	// function under test
	order2 := NewOrder(_testOrderParamStr2)
	assertOrderValid(t, _testOrderParam2, order2)

	// new test
	// function under test
	order3 := NewOrder(_testOrderParamStr3)
	assertOrderValid(t, _testOrderParam3, order3)
}

func TestMarshalJSON(t *testing.T) {
	// new test
	// function under test
	orderStr1, err := _testOrder1.MarshalJSON()
	if err != nil {
		log.Printf(err.Error())
	}
	assert.Equal(t, _testOrderStr1, orderStr1, "should marshal order struct to json")

	// new test
	// function under test
	orderStr2, err := _testOrder2.MarshalJSON()
	if err != nil {
		log.Printf(err.Error())
	}
	assert.Equal(t, _testOrderStr2, orderStr2, "should marshal order struct to json")

	// new test
	// function under test
	orderStr3, err := _testOrder3.MarshalJSON()
	if err != nil {
		log.Printf(err.Error())
	}
	assert.Equal(t, _testOrderStr3, orderStr3, "should marshal order struct with empty string fields and zeroes to json")
}

func TestUnmarshalJSON(t *testing.T) {
	// new test
	// setup
	order1 := &Order{}
	// function under test
	if err := order1.UnmarshalJSON(_testOrderStr1); err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, _testOrder1, order1, "should unmarshal json to order struct")

	// new test
	// setup
	order2 := &Order{}
	// function under test
	if err := order2.UnmarshalJSON(_testOrderStr2); err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, _testOrder2, order2, "should unmarshal json to order struct")

	// new test
	// setup
	order3 := &Order{}
	// function under test
	if err := order3.UnmarshalJSON(_testOrderStr3); err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, _testOrder3, order3, "should unmarshal json to order struct with empty string fields and zeroes")

	// new test
	// setup
	order4 := &Order{}
	// function under test
	err := order4.UnmarshalJSON(_testBadOrderStr4)
	assert.EqualError(t, err, "Error unmarshalling to order struct: json: cannot unmarshal string into Go struct field OrderLocation.lon of type float64")

	// new test
	// setup
	order5 := &Order{}
	// function under test
	err = order5.UnmarshalJSON(_testBadOrderStr5)
	assert.EqualError(t, err, "Error unmarshalling to order struct: json: cannot unmarshal string into Go struct field OrderCopy.bidPrice of type float64")
}
