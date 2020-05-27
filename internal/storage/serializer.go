package storage

type Serializer interface {
  Serialize() ()
  Derialize() (*DBKey, error)
}

type DBKey struct {
  keyId string
  value string
}

// func Serlialize(keyId string, value []byte) *DBKey {
//   return &DBKey{
//     keyId: keyId,
//     value: value
//   }
// }
