package backoffice

// go gnerate file to create a pb directory and then use protoc to generate the pb files and
// place them in the pb dicrectory

//go:generate mkdir -p pb
//go:generate protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative ./invoice.proto
