package mocks

//go:generate go run github.com/golang/mock/mockgen@latest -source ../transfer.go -package mocks -destination transfer_mock.go
//go:generate go run github.com/golang/mock/mockgen@latest -source ../adaptor.go -package mocks -destination adaptor_mock.go
