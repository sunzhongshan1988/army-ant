# Dev
## grpc
`protoc.exe --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  helloworld/helloworld.proto`


# graphql
1.update ./broker/graph/schema.graphqls, in this path:

` go generate`

#Build
## Change Target OS
```
go env
$env:GOARCH = "amd64"
$env:GOOS = "linux"

go build
```
# Deploy
## Environment Variables
For Broker and Worker, 'label' and 'address:port' is the only basis for distinguishing the same instance,
 and the change will be regarded as a new instance.
### Broker
    AAB_LABEL           Broekr label
	AAB_ADDRESS         Broker address
	AAB_GRPC_PORT       Broker Grpc listen port
	AAB_GRAPHQL_PORT    Broekr GraphQL listen port
### Worker
    AAW_BROKER_LINK     Worker link broker
    AAW_ADDRESS         Worker address
    AAW_PORT            Worker listen port
    AAW_LABEL           Worker label
## Deploy to Linux
Create a service file in /etc/systemd/system. Let's call it go.service. Let the contents be:
### Broker
```
[Unit]
Description=ArmyAnt Broker Service

[Service]
User=root
Environment=AAB_LABEL=system
Environment=AAB_ADDRESS=127.0.0.1
Environment=AAB_GRPC_PORT=50051
Environment=AAB_GRAPHQL_PORT=8080
WorkingDirectory=/root 
ExecStart=/opt/armyant-broker
TimeoutStopSec=10
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
```

### Worker
```
[Unit]
Description=ArmyAnt Worker Service

[Service]
User=root
Environment=AAW_BROKER_LINK=127.0.0.1:50051
Environment=AAW_ADDRESS=127.0.0.1
Environment=AAW_PORT=50052
Environment=AAW_LABEL=worker01   
WorkingDirectory=/root 
ExecStart=/opt/armyant-worker
TimeoutStopSec=10
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
```