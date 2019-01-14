package lib

//Output is main output struct
type Output struct {
	ID         string
	Identifier string `json:",omitempty"`
	Status     Status
	Fields     map[string]interface{}
	Bid        int32  `json:",omitempty"`
	Error      string `json:",omitempty"`
}
