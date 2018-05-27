package common

/*
	Your object to be marshaled
 	as JSON from raw bytes.

 	Add fields as you need this
 	object is shared between the
 	client and server.

	This struct is parsed and modified in common.go
*/
type MyObject struct {
	Magic uint8
	Msg   string
	Response string
}
