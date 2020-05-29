package order

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// test order json payloads to create new order structs
	_testOrderParamBytes1 = []byte(`{"location":{"lon":43.45123431,"lat":75.13124123},"description":"test_order_description1","consumerId":"test_order_consumer_id1","bidPrice":100.23}`)
	_testOrderParamBytes2 = []byte(`{"location":{"lon":43.45123432,"lat":75.13124122},"description":"test_order_description2","consumerId":"test_order_consumer_id2","bidPrice":200.23}`)
	_testOrderParamBytes3 = []byte(`{"location":{"lon":0,"lat":0},"description":"","consumerId":"","bidPrice":0}`)

	// test order json payloads to marshal to order structs
	_testOrderBytes1 = []byte(`{"location":{"lon":1.23,"lat":1.63},"id":"test_order_id1","description":"test_order_description1","timeRequested":"test_order_time_requested1","duration":"test_order_duration1","consumerId":"test_order_consumer_id1","bidPrice":1.234}`)
	_testOrderBytes2 = []byte(`{"location":{"lon":2.23,"lat":2.43},"id":"test_order_id2","description":"test_order_description2","timeRequested":"test_order_time_requested2","duration":"test_order_duration2","consumerId":"test_order_consumer_id2","bidPrice":1.236}`)
	_testOrderBytes3 = []byte(`{"location":{"lon":0,"lat":0},"id":"","description":"","timeRequested":"","duration":"","consumerId":"","bidPrice":0}`)
	_testBadOrderBytes4 = []byte(`{"location":{"lon": "should not be a string","lat":0},"id":"test_order_id2","description":"test_order_description2","timeRequested":"test_order_time_requested2","duration":"test_order_duration2","consumerId":"test_order_consumer_id2","bidPrice":654.2312451232}`)
	_testBadOrderBytes5 = []byte(`{"location":{"lon":123,"lat":234},"id":"test_order_id2","description":"test_order_description2","timeRequested":"test_order_time_requested2","duration":"test_order_duration2","consumerId":"test_order_consumer_id2","bidPrice":"should not be a string"}`)
)

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
	order1 := NewOrder(_testOrderParamBytes1)
	assertOrderValid(t, &Order{
		Location: OrderLocation{
			Lon: 43.45123431,
			Lat: 75.13124123,
		},
		Desc:       "test_order_description1",
		ConsumerId: "test_order_consumer_id1",
		BidPrice:   100.23,
	}, order1)

	// new test
	// function under test
	order2 := NewOrder(_testOrderParamBytes2)
	assertOrderValid(t, &Order{
		Location: OrderLocation{
			Lon: 43.45123432,
			Lat: 75.13124122,
		},
		Desc:       "test_order_description2",
		ConsumerId: "test_order_consumer_id2",
		BidPrice:   200.23,
	}, order2)

	// new test
	// function under test
	order3 := NewOrder(_testOrderParamBytes3)
	assertOrderValid(t, &Order{
		Location: OrderLocation{
			Lon: 0,
			Lat: 0,
		},
		Desc:       "",
		ConsumerId: "",
		BidPrice:   0,
	}, order3)
}

func TestMarshalJSON(t *testing.T) {
	// new test
	// setup
	order1 := &Order{
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
	// function under test
	orderBytes1, err := order1.MarshalJSON()
	if err != nil {
		log.Printf(err.Error())
	}
	assert.Equal(t, string(_testOrderBytes1), string(orderBytes1), "should marshal order struct to json")

	// new test
	// setup
	order2 := &Order{
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
	// function under test
	orderBytes2, err := order2.MarshalJSON()
	if err != nil {
		log.Printf(err.Error())
	}
	assert.Equal(t, _testOrderBytes2, orderBytes2, "should marshal order struct to json")

	// new test
	// setup
	order3 := &Order{
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
	// function under test
	orderBytes3, err := order3.MarshalJSON()
	if err != nil {
		log.Printf(err.Error())
	}
	assert.Equal(t, _testOrderBytes3, orderBytes3, "should marshal order struct with empty string fields and zeroes to json")
}

func TestUnmarshalJSON(t *testing.T) {
	// new test
	order := &Order{}
	// function under test
	if err := order.UnmarshalJSON(_testOrderBytes1); err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, &Order{
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
	}, order, "should unmarshal json to order struct")

	// new test
	order = &Order{}
	// function under test
	if err := order.UnmarshalJSON(_testOrderBytes2); err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, &Order{
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
	}, order, "should unmarshal json to order struct")

	// new test
	order = &Order{}
	// function under test
	if err := order.UnmarshalJSON(_testOrderBytes3); err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, &Order{
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
	}, order, "should unmarshal json to order struct with empty string fields and zeroes")

	// new test
	order = &Order{}
	// function under test
	err := order.UnmarshalJSON(_testBadOrderBytes4)
	assert.EqualError(t, err, "Error unmarshalling to order struct json: cannot unmarshal string into Go struct field OrderLocation.lon of type float64")

	// new test
	order = &Order{}
	// function under test
	err = order.UnmarshalJSON(_testBadOrderBytes5)
	assert.EqualError(t, err, "Error unmarshalling to order struct json: cannot unmarshal string into Go struct field OrderCopy.bidPrice of type float64")
}
