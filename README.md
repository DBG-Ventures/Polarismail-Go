# Polarismail-Go

> This is a work in progress, do not use it in production

PolarisMail-Go is a community driven sdk for the Polarismail REST API written in Go. The SDK strives to normalize the data return from the API to make it more human readable and provide an intuitive interface structure for interacting with your email systems.

## Getting Started

To add the SDK to your project, simply create a `Credentials` object and pass that to the `NewClient` method.

```go
import (
 "github.com/dbgventures/polarismail-go/polarismail"
 "github.com/dbgventures/polarismail-go/types"
)

func main() {
    creds := types.Credentials{
        Username: "your_username",
        Password: "your_password",
    }

    client, err := polarismail.NewClient(&creds)
    if err != nil {
        log.Fatal(err)
    }
}
```

> Better examples will be provided once the SDK is completed

## `Client`

```go
type Client interface {
    Admin() Admin
    Users() Users
    Domains() Domains
    Lists() Lists
    Aliases() Aliases
}
```

### `Admin`

```go
type Admin interface {
    GetStats() (AdminStats, error)
    GetBrandInfo() (AdminBrandInfo, error)
    GetActionHistory() ([]AdminActionHistory, error)
    UpdatePassword(newPassword string) error
    UpdateEmail(newEmail string) error
}
```

### `Domains`

```go
type Domains interface {
    List() ([]DomainsList, error)
    CheckAvailable(newDomain string) bool
    Add(newDomain string) (Domain, error)
    Get(domain string) Domain
}
```

#### `Domain`

```go
type Domain interface {
    Info() (DomainInfo, error)
    Edit() DomainEdit
    DKIM() DomainDKIM
    Footer() DomainFooter
    DomainAliases() DomainAliases
    Delete() error
}
```

##### `DomainEdit`

```go
type DomainEdit interface {
    DomainEnabled(enabled bool) error
    CatchAll(status string) error
    Timezone(timezone string) error
    ExchangeEnabled(enabled bool) error
}
```

##### `DomainDKIM`

```go
type DomainDKIM interface {
    GetStatus() (DomainDKIMStatus, error)
    Enable() (DomainDKIMStatus, error)
    Disable() error
}
```

##### `DomainFooter`

```go
type DomainFooter interface {
    Get() (string, error)
    Set(footer string) error
}
```

##### `DomainAliases`

```go
type DomainAliases interface {
    List() ([]DomainAlias, error)
    Add(alias string) error
    Delete(alias string) error
}
```
