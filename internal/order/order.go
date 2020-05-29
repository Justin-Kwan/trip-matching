package order

import (
	"encoding/json"
	"time"

	"github.com/buger/jsonparser"
	"github.com/pkg/errors"
	"github.com/rs/xid"
)

type Order struct {
	Location      OrderLocation `json:"location"`
	Id            string        `json:"id"`
	Desc          string        `json:"description"`
	TimeRequested string        `json:"timeRequested"`
	Duration      string        `json:"duration"`
	ConsumerId    string        `json:"consumerId"`
	BidPrice      float64       `json:"bidPrice"`
}

type OrderLocation struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

// Returns a new order struct with generated id.
func NewOrder(params []byte) *Order {
	o := &Order{}
	o.Id = xid.New().String()
	o.TimeRequested = time.Now().String()
	o.Desc, _ = jsonparser.GetString(params, "description")
	o.BidPrice, _ = jsonparser.GetFloat(params, "bidPrice")
	o.ConsumerId, _ = jsonparser.GetString(params, "consumerId")
	o.Location.Lon, _ = jsonparser.GetFloat(params, "location", "lon")
	o.Location.Lat, _ = jsonparser.GetFloat(params, "location", "lat")
	return o
}

// time unit??
func (o *Order) SetEstDuration(duration string) {
	o.Duration = duration
}

func (o *Order) MarshalJSON() ([]byte, error) {
	return json.Marshal(*o)
}

func (o *Order) UnmarshalJSON(orderParams []byte) error {
  // define new type locally to avoid recursion
  type OrderCopy Order
  var oCopy OrderCopy

	if err := json.Unmarshal(orderParams, &oCopy); err != nil {
		return errors.Errorf("Error unmarshalling to order struct %v", err)
	}
  // cast back to Order struct type and assign to current Order struct
  *o = Order(oCopy)
	return nil
}
