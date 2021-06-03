# Dev
## grpc
`protoc.exe --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  helloworld/helloworld.proto`


# graphql
1.update ./broker/graph/schema.graphqls, in this path:

` go generate`

#Build
## Change Target OS
`go env`

`$env:GOARCH = "amd64"`

`$env:GOOS = "linux"`

# Deploy
## Deploy to Linux
Create a service file in /etc/systemd/system. Let's call it go.service. Let the contents be:
```
[Unit]
Description=ArmyAnt Service

[Service]
User=root
# The configuration file application.properties should be here:
WorkingDirectory=/home/myuser/my-apps 
ExecStart=/opt/armyant-broker
TimeoutStopSec=10
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
```
