package redis

// type RedisGeoDb struct {
//   num int
//   index string
//
// }
//
// func (gdb *GeoDB) InsertPOI(keyId string, lon float64, lat float64) error {
// 	conn := gdb.pool.Get()
//   conn.Do("SELECT", gdb.num)
// 	defer conn.Close()
//
// 	_, err := conn.Do("GEOADD", gdb.index, lat, lon, keyId)
// 	if err != nil {
// 		return errors.Errorf("Error adding POI with key '%s': %v", keyId, err)
//   }
// 	return nil
// }
//
// func (gdb *GeoDB) SelectAllInRadius()  {
//   conn := gdb.pool.Get()
//   conn.Do("SELECT", gdb.num)
// 	defer conn.Close()
//
// 	_, err := conn.Do("GEODIST", gdb.index, lat, lon, keyId)
// 	if err != nil {
// 		return "", errors.Errorf("Error adding POI with key '%s': %v", keyId, err)
// 	}
// 	return val, nil
// }
