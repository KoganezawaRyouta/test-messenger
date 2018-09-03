package client

import "encoding/json"

type Issues struct {
	Title       string
	Number      int64
}

func GetIssues() []Issues {
	url := "https://gist.githubusercontent.com/punytan/e46fe04ff81c8da636658192c3d5b58d/raw/7e8b2f5543d5aac89a92f678544f693d33c6b465/issues.json"
	byteArray := get(url)
	var t []Issues
	json.Unmarshal([]byte(string(byteArray)), &t)
	return t
}