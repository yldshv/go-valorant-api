

<img align="right" src="https://i.ibb.co/7px2VF7/valogoapi.png" height="200" width="200">




# GO-VALORANT-API (WIP!) [![Go Reference](https://pkg.go.dev/badge/github.com/yldshv/go-valorant-api.svg)](https://pkg.go.dev/github.com/yldshv/go-valorant-api)

GO-VALORANT-API is a go-based wrapper for the following Valorant Rest API:

https://github.com/Henrik-3/unofficial-valorant-api v3.0.0

This API is free and freely accessible for everyone. An API key is optional but not mandatory.

This is the first version. There could be some bugs, unexpected exceptions or similar.

### API key

You can request an API key on [Henrik's discord server](https://discord.com/invite/X3GaVkX2YN) <br> It is NOT required to use an API key though!

## Installation

```bash
go get -u github.com/yldshv/go-valorant-api
```

### Example

Get Account Information
```go
package main

import (
  govapi "github.com/yldshv/go-valorant-api"
)

func main() {
  vapi := govapi.New() 
  // vapi := govapi.New(govapi.WithKey("xxx")) <- If you have a key.

  acc, err := vapi.GetAccountByName(govapi.GetAccountByNameParams{
    Name: "xxx",
    Tag: "xxx",
  })
  if err != nil {
    // handle the error
  }

  // Do whatever you want with acc
  fmt.Printf("%+v", acc.Data.Puuid)
}
```

```go
mmrHistory, err := vapi.GetLifetimeMMRHistoryByPUUID(govapi.GetLifetimeMMRHistoryByPUUIDParams{
  Affinity: "eu",
  Puuid: acc.Data.Puuid, //here you can see we used the value from above
  Page: "2",
  Size: "10",
})
if err != nil {
  // handle the error
}

fmt.Printf("%+v", mmrHistory)
```

## Documentation

https://pkg.go.dev/github.com/yldshv/go-valorant-api#VAPI

Methodnaming is heavily inspired by the endpoint naming:<br>
[Swagger](https://app.swaggerhub.com/apis-docs/Henrik-3/HenrikDev-API/3.0.0)
