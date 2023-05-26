# Overview

* Go client for the Proxmox-VE REST API (https://pve.proxmox.com/wiki/Proxmox_VE_API)

* This code is based on Go package [go-proxmox](https://github.com/luthermonson/go-proxmox)


# Proxmox API Client Go Package
A Go package containing a client for [Proxmox VE](https://www.proxmox.com/). The client implements [/api2/json](https://pve.proxmox.com/pve-docs/api-viewer/index.html) and aims to provide better sdk solution for especially [cluster-api-provider-proxmox](https://github.com/sp-yduck/cluster-api-provider-proxmox) project.

## Usage
Create a client and use the public methods to access Proxmox resources.

### Basic usage with login credentials
```go
package main

import (
	"fmt"
	"github.com/sp-yduck/proxmox"
)

func main() {
    client := proxmox.NewClient("https://localhost:8006/api2/json")
    if err := client.Login("root@pam", "password"); err != nil {
        panic(err)
    }
    version, err := client.Version()
    if err != nil {
        panic(err)
    }
    fmt.Println(version.Release) // 6.3
}
```

### Usage with Client Options
```go
package main

import (
	"fmt"
	"github.com/sp-yduck/proxmox"
)

func main() {
    insecureHTTPClient := http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: true,
            },
        },
    }
    tokenID := "root@pam!mytoken"
    secret := "somegeneratedapitokenguidefromtheproxmoxui"
    
    client := proxmox.NewClient("https://localhost:8006/api2/json",
        proxmox.WithClient(&insecureHTTPClient),
        proxmox.WithAPIToken(tokenID, secret),
    )
    
    version, err := client.Version()
    if err != nil {
        panic(err)
    }
    fmt.Println(version.Release) // 6.3
}
```

# Developing
This project relies on [Mage](https://magefile.org/) for cross os/arch compatibility, please see their installation guide. 

## Unit Testing
Run `mage test` to run the unit tests in the root directory.

## Integration Testing
To run the integration testing suite against an existing Proxmox API set some env vars in your shell before running `mage testIntegration`. The integration tests will test logging in and using an API token credentials so make sure you set all five env vars before running tests for them to pass.

Please leave no trace when developing integration tests. All tests should create and remove all testing data they generate then they can be repeatably run against the same proxmox environment. Most people working on this package will likely use their personal Proxmox VE home lab and consuming extra resources via tests will lead to frustration.

### Bash
```shell
export PROXMOX_URL="https://192.168.1.6:8006/api2/json"
export PROXMOX_USERNAME="root@pam"
export PROXMOX_PASSWORD="password"
export PROXMOX_TOKENID="root@pam!mytoken"
export PROXMOX_SECRET="somegeneratedapitokenguidefromtheproxmoxui"

mage test:integration
```

### Powershell
```powershell
$Env:PROXMOX_URL = "https://192.168.1.6:8006/api2/json"
$Env:PROXMOX_USERNAME = "root@pam"
$Env:PROXMOX_PASSWORD = "password"
$Env:PROXMOX_TOKENID = "root@pam!mytoken"
$Env:PROXMOX_SECRET = "somegeneratedapitokenguidefromtheproxmoxui"

mage test:integration
```


