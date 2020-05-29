package order

import (
  "time"
	"encoding/json"

	"github.com/buger/jsonparser"
	"github.com/pkg/errors"
	"github.com/rs/xid"
)

type Order struct {
	location      OrderLocation `json:"location"`
	id            string         `json:"id"`
	desc          string         `json:"description"`
	timeRequested string         `json:"timeRequested"`
	duration      string         `json:"duration"`
	consumerId    string         `json:"consumerId"`
	bidPrice      float64        `json:"bidPrice"`
}

type OrderLocation struct {
	lon float64 `json:"longitude"`
	lat float64 `json:"latitude"`
}

// Returns a new order struct with generated id.
func NewOrder(params []byte) *Order {
	o := &Order{}
	o.id = xid.New().String()
	o.timeRequested = time.Now().String()
	o.desc, _ = jsonparser.GetString(params, "orderInfo", "description")
	o.bidPrice, _ = jsonparser.GetFloat(params, "orderInfo", "bidPrice")
	o.consumerId, _ = jsonparser.GetString(params, "orderInfo", "consumerId")
	o.location.lon, _ = jsonparser.GetFloat(params, "orderInfo", "location", "lon")
	o.location.lat, _ = jsonparser.GetFloat(params, "orderInfo", "location", "lat")
	return o
}

// time unit??
func (o *Order) SetEstDuration(dur string) {
	o.duration = dur
}

func (o *Order) MarshalJSON() ([]byte, error) {
	return json.Marshal(*o)
}

func (o *Order) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &o); err != nil {
		return errors.Errorf("Error unmarshalling to order struct %v", err)
	}
	return nil
}
