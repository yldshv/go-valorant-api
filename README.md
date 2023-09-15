

<img align="right" src="https://i.ibb.co/7px2VF7/valogoapi.png" height="200" width="200">




# GO-VALORANT-API (WIP!)

GO-VALORANT-API is a go-based wrapper for the following Valorant Rest API:

https://github.com/Henrik-3/unofficial-valorant-api

This API is free and freely accessible for everyone. An API key is optional but not mandatory.

This is the first version. There could be some bugs, unexpected exceptions or similar.

### API key

You can request an API key on [Henrik's discord server](https://discord.com/invite/X3GaVkX2YN) <br> It is NOT required to use an API key though!

## Summary

1. [Installation](#installation)
2. [Example](#example)
3. [Documentation](#documentation)

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
  vapi := govapi.New() //You can pass in your API-Key as String (if you dont have a key just leave it empty)

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
  log.Fatal(err)
}

fmt.Printf("%+v", mmrHistory)
```

## Documentation

The detailed documentations are still in progress.
