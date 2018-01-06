package sirena

import "encoding/xml"

// VehiclesResponse is a response to all vehicles request
type VehiclesResponse struct {
	Answer  VehiclesAnswer `xml:"answer"`
	XMLName xml.Name       `xml:"sirena" json:"-"`
}

// VehiclesAnswer is an <answer> section in all vehicles response
type VehiclesAnswer struct {
	Vehicles VehiclesAnswerDetails `xml:"describe"`
}

// VehiclesAnswerDetails is a <describe> section in all vehicles response
type VehiclesAnswerDetails struct {
	Data []VehiclesAnswerData `xml:"data"`
}

// VehiclesAnswerData is a <data> section in all vehicles response
type VehiclesAnswerData struct {
	Code []VehiclesAnswerDataCode `xml:"code"`
	Name []VehiclesAnswerDataName `xml:"name"`
}

// VehiclesAnswerDataCode represents <code> entry in <data> section
type VehiclesAnswerDataCode struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

// VehiclesAnswerDataName represents <name> entry in <data> section
type VehiclesAnswerDataName struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}