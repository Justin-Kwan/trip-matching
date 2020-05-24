package auth

// import (
// 	"crypto/sha256"
// 	"encoding/hex"
// )

// type HashConfig struct {
// 	secretKey: string
// }
//
// type TicketInfo struct {
// 	IpAddress: string,
// 	UserId: string,
// 	IssuedAt: time.Time
// 	ExpiresAt: time.Time
// }
//
// type Handler interface {
// 	HashAsString() string
//
// }
//
// type TickerInformer interface {
//
// }
//
// func NewTicket() Ticket {
//
// }
//
// func IsExpired(ticketInfo *TicketInfo) bool {
//
// }
//
// func HashAsString(ticketInfo *TicketInfo) string {
// 	hasher := sha256.New()
//
// 	// marshal to json string first
//
// 	ticketInfoJson := ....()
// 	ticketHashBytes := sha256.Sum256([]byte(ticketInfoJson))
// 	ticketHashString := hex.EncodeToString(ticketHashBytes[:])
// 	return ticketHashString
// }
//
// func Clear(ticketInfo *TicketInfo) {
//
// }
