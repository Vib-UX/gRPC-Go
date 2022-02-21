protoc -I greet\ --go-grpc_out=greet\ greet\greetpb\greet.proto
protoc -I greet\ --go_out=greet\ greet\greetpb\greet.proto

protoc -I calculator\ --go-grpc_out=calculator\ calculator\calcpb\calc.proto
protoc -I calculator\ --go_out=calculator\ calculator\calcpb\calc.proto