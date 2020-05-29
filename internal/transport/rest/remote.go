package rest

// import (
// 	"net/http"
//
//   "github.com/pkg/errors"
// )
//
// type RemoteApi struct {
//   apiKey string
//   url string
// }
//
// func NewRemoteApi(apiKey string) {
//
// }
//
// func (r *RemoteApi) MakeGetReq() {
// 	res, err := http.Get(url)
// 	if err != nil {
// 		return errors.Errorf("Error making GET request to '%s': %v", r.url, err)
// 	}
// 	defer res.Body.Close()
// 	return res.Body
// }
