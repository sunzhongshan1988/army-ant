# Dev
## grpc
`protoc.exe --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  helloworld/helloworld.proto`


# graphql
1.update ./broker/graph/schema.graphqls, in this path:

` go generate`