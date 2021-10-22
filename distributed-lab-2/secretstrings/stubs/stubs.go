package stubs

var ReverseHandler = "SecretStringOperations.Reverse"
var PremiumReverseHandler = "SecretStringOperations.FastReverse"

//response struct and the request struct
//nboth have a single string variable which is caled the message

type Response struct {
	Message string
}

type Request struct {
	Message string
}
