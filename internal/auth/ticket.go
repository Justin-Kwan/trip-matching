package auth

// import (
//   // "time"
//   "encoding/json"
//
//   "github.com/pkg/errors"
//   // "github.com/google/uuid"
// )

// type Ticket struct {
//   Id       string  `json:"ticketId"`
// 	IpAddr   string  `json:"ipAddress"`
// 	UserId   string  `json:"userId"`
// 	IssuedAt string  `json:"issuedAt"`
// 	Exp      string  `json:"expiry"`
// }

// type Handler interface {
// }

// need to check if ticket is expired.
// check if issue time is before current time and.
// func newTicket(user *user) *Ticket {
//   return &Ticket{
//     id:       uuid.NewUUID(),
//     ipAddr:   user.ipAddr,
//     userId:   user.id,
//     issuedAt: time.Now().String(),
//     exp:      time.Now().Add(time.Minute * 15).String(),
//   }
// }

// func (t *Ticket) marshalJSON() ([]byte, error) {
//   return json.Marshal(*t)
// }
//
// func (t *Ticket) unmarshalJSON(data []byte) error {
//   if err := json.Unmarshal(data, &t); err != nil {
//     return errors.Errorf("Error unmarshalling %v", err)
//   }
//   return nil
// }
//
// func (t *Ticket) isValid() bool {
//   doTicketsMatch := t.Id == newT.Id && t.IpAddress == newT.IpAddress
//   && t.userId == newT.userId &&
//
//   isTicketTimeValid :=
//
//   return doTicketsMatch
// }
