package db

// import (
// 	"database/sql/driver"
// 	"encoding/json"
// )

// type NullInt struct {
// 	Int int
// }

// var nullJson = []byte("null")

// func (v *NullInt) Scan(value any) error {
// 	if value == nil {
// 		return nil
// 	}
// 	v.Int = int(value.(int64))
// 	return nil
// }

// func (v NullInt) Value() (driver.Value, error) {
// 	if v.Int == 0 {
// 		return nil, nil
// 	}
// 	return v.Int, nil
// }

// func (v NullInt) MarshalJSON() ([]byte, error) {
// 	if v.Int == 0 {
// 		return nullJson, nil
// 	} else {
// 		return json.Marshal(v.Int)
// 	}
// }

// func (v *NullInt) UnmarshalJSON(data []byte) error {
// 	if err := json.Unmarshal(data, &v.Int); err != nil {
// 		return err
// 	}
// 	return nil
// }

// type NullString struct {
// 	Str string
// }

// func (v *NullString) Scan(value any) error {
// 	if value == nil {
// 		return nil
// 	}
// 	switch typ := value.(type) {
// 	case string:
// 		v.Str = typ
// 	case []uint8:
// 		v.Str = string(typ)
// 	default:
// 		panic("unknown type received in Scan NullString")
// 	}
// 	return nil
// }

// func (v NullString) Value() (driver.Value, error) {
// 	if v.Str == "" {
// 		return nil, nil
// 	}
// 	return v.Str, nil
// }

// func (v NullString) MarshalJSON() ([]byte, error) {
// 	if v.Str == "" {
// 		return nullJson, nil
// 	} else {
// 		return []byte("\"" + v.Str + "\""), nil
// 	}
// }

// func (v *NullString) UnmarshalJSON(data []byte) error {
// 	if err := json.Unmarshal(data, &v.Str); err != nil {
// 		return err
// 	}
// 	return nil
// }
