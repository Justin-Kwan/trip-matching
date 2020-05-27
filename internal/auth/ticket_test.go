package auth

// import (
//   "log"
//   "testing"
//
//   "github.com/stretchr/testify/assert"
// )
//
// var (
//   _testTicket1 = Ticket{
//     Id:       "test_id1",
//     IpAddr:   "test_ip_address1",
//     UserId:   "test_user_id1",
//     IssuedAt: "test_time_of_issue1",
//     Exp:      "test_time_of_expiry1",
//   }
//   _testTicket2 = Ticket{
//     Id:       "test_id2",
//     IpAddr:   "test_ip_address2",
//     UserId:   "test_user_id2",
//     IssuedAt: "test_time_of_issue2",
//     Exp:      "test_time_of_expiry2",
//   }
//
//   _testTicket3 = Ticket{
//     Id:       "",
//     IpAddr:   "",
//     UserId:   "",
//     IssuedAt: "",
//     Exp:      "",
//   }
// )
//
// const (
//   _testTicketStr1 = `{"ticketId":"test_id1","ipAddress":"test_ip_address1","userId":"test_user_id1","issuedAt":"test_time_of_issue1","expiry":"test_time_of_expiry1"}`
//   _testTicketStr2 = `{"ticketId":"test_id2","ipAddress":"test_ip_address2","userId":"test_user_id2","issuedAt":"test_time_of_issue2","expiry":"test_time_of_expiry2"}`
//   _testTicketStr3 = `{"ticketId":"","ipAddress":"","userId":"","issuedAt":"","expiry":""}`
//   _testBadTicketStr4 = `{"ticketId":"","ipAddress":false,"userId":"","issuedAt":"","expiry":""}`
//   _testBadTicketStr5 = `{"ticketId":"","ipAddress":"","userId":"","issuedAt":1000,"expiry":""}`
// )
//
// func TestMarshalJSON(t *testing.T) {
//   // new test
//   // function under test
//   ticketBytes, err := _testTicket1.marshalJSON()
//   if err != nil {
//     log.Fatalf(err.Error())
//   }
//   assert.Equal(t, _testTicketStr1, string(ticketBytes), "should marshal struct to json")
//
//   // new test
//   // function under test
//   ticketBytes, err = _testTicket2.marshalJSON()
//   if err != nil {
//     log.Fatalf(err.Error())
//   }
//   assert.Equal(t, _testTicketStr2, string(ticketBytes), "should marshal struct to json")
//
//   // new test
//   // function under test
//   ticketBytes, err = _testTicket3.marshalJSON()
//   if err != nil {
//     log.Fatalf(err.Error())
//   }
//   assert.Equal(t, _testTicketStr3, string(ticketBytes), "should marshal struct to json")
// }
//
// func TestUnmarshalJSON(t *testing.T) {
//   // new test
//   tic := Ticket{}
//   // function under test
//   if err := tic.unmarshalJSON([]byte(_testTicketStr1)); err != nil {
//     log.Fatal(err.Error())
//   }
//   assert.Equal(t, _testTicket1, tic, "should unmarshal json to ticket struct")
//
//   // new test
//   tic = Ticket{}
//   // function under test
//   if err := tic.unmarshalJSON([]byte(_testTicketStr2)); err != nil {
//     log.Fatal(err.Error())
//   }
//   assert.Equal(t, _testTicket2, tic, "should unmarshal json to ticket struct")
//
//   // new test
//   tic = Ticket{}
//   // function under test
//   if err := tic.unmarshalJSON([]byte(_testTicketStr3)); err != nil {
//     log.Fatal(err.Error())
//   }
//   assert.Equal(t, _testTicket3, tic, "should unmarshal json to ticket struct")
//
//   // new test
//   tic = Ticket{}
//   // function under test
//   err := tic.unmarshalJSON([]byte(_testBadTicketStr4))
//   assert.EqualError(t, err, "Error unmarshalling json: cannot unmarshal bool into Go struct field Ticket.ipAddress of type string")
//
//   // new test
//   tic = Ticket{}
//   // function under test
//   err = tic.unmarshalJSON([]byte(_testBadTicketStr5))
//   assert.EqualError(t, err, "Error unmarshalling json: cannot unmarshal number into Go struct field Ticket.issuedAt of type string")
// }
//
// func TestNewTicket(t *testing.T) {
//
// }
