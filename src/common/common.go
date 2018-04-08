package common

// Common return common information
func Common() map[string]string {
	var c = map[string]string{}
	c["ServiceName"] = "manense"
	c["ServiceURL"] = "http://localhost:8000"
	return c
}
