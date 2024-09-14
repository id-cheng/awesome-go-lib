## Project: "MyNode"

### Generated with
 - Types for the network messaging: false
 - Enabled Observer (http://localhost:9911): true

### Supervision Tree

Applications
 - `MyApp{}` mynode/apps/myapp/myapp.go
   - `MySup{}` mynode/apps/myapp/mysup.go
     - `MyActor{}` mynode/apps/myapp/myactor.go

Process list that is starting by node directly
 - `MyActor2{}` mynode/cmd/myactor2.go
 - `MyWeb{}` mynode/cmd/myweb.go


#### Used command

This project has been generated with the `ergo` tool. To install this tool, use the following command:

`$ go install ergo.services/tools/ergo@latest`

Below the command that was used to generate this project:

```$ ergo -init MyNode -with-app MyApp -with-sup MyApp:MySup -with-actor MySup:MyActor -with-web MyWeb -with-actor MyActor2 -with-observer ```
