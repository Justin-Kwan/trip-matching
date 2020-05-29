package order

import (
    // "log"
    "testing"

    "github.com/stretchr/testify/assert"
)

var (
    testOrderBytes1 = [] byte(`{
    "orderInfo": {
      "location":{
        "lon": 43.45123431,
        "lat": 75.13124123
      },
      "description": "test_order_description1",
      "consumerId": "test_order_consumer_id1",
      "bidPrice": 100.23
    }
  }`)
  testOrderBytes2 = [] byte(`{
    "orderInfo": {
      "location":{
        "lon": 43.45123432,
        "lat": 75.13124122
      },
      "description": "test_order_description2",
      "consumerId": "test_order_consumer_id2",
      "bidPrice": 200.23
    }
  }`)
  testOrderBytes3 = [] byte(`{
    "orderInfo": {
      "location":{
        "lon": 0,
        "lat": 0
      },
      "description": "",
      "duration???"
      "consumerId": "",
      "bidPrice": 0
    }
  }`)
)

func assertOrderValid(t * testing.T, actual * Order, expected * Order) {
    assert.Equal(t, actual.location.lon, expected.location.lon)
    assert.Equal(t, actual.location.lat, expected.location.lat)
    assert.NotEmpty(t, expected.id)
    assert.Equal(t, actual.desc, expected.desc)
    assert.NotEmpty(t, expected.timeRequested)
    assert.Equal(t, actual.consumerId, expected.consumerId)
    assert.Equal(t, actual.bidPrice, expected.bidPrice)
}

func TestNewOrder(t * testing.T) {
    // new test
    // function under test
    o1: = NewOrder(testOrderBytes1)
    assertOrderValid(t, & Order {
        location: OrderLocation {
            lon: 43.45123431,
            lat: 75.13124123,
        },
        desc: "test_order_description1",
        consumerId: "test_order_consumer_id1",
        bidPrice: 100.23,
    }, o1)

    // new test
    // function under test
    o2: = NewOrder(testOrderBytes2)
    assertOrderValid(t, & Order {
        location: OrderLocation {
            lon: 43.45123432,
            lat: 75.13124122,
        },
        desc: "test_order_description2",
        consumerId: "test_order_consumer_id2",
        bidPrice: 200.23,
    }, o2)

    // new test
    // function under test
    o3: = NewOrder(testOrderBytes3)
    assertOrderValid(t, & Order {
        location: OrderLocation {
            lon: 0,
            lat: 0,
        },
        desc: "",
        consumerId: "",
        bidPrice: 0,
    }, o3)
}
