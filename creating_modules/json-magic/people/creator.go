package people

import "encoding/json"

func CreateHusband(h Husband) []byte {
	jsonHusband, _ := json.Marshal(h)
	return jsonHusband
} 