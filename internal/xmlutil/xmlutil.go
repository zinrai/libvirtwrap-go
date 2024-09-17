package xmlutil

import (
	"encoding/xml"
)

// ParseXML parses XML data into the provided interface
func ParseXML(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

// GenerateXML generates XML from the provided interface
func GenerateXML(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}
