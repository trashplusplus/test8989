package main

type ABStruct struct {
	A int `json:"a"`
	B int `json:"b"`
}

type ResponseError struct {
	Error string `json:"error"`
}