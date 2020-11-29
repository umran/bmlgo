# BMLGo
Go bindings for the BML internet banking API

## Getting Started
Install the package using the following command
```
go get -u github.com/umran/bmlgo
```

## Usage
### Instantiating a Client
Before interacting with the API a Client must be instantiated
```go
import "github.com/umran/bmlgo"

func main() {
    client, err := bmlgo.NewClient("username", "password")
    if err != nil {
        panic("failed to create client", err)
    }
}
```

### Getting New Statement Items
New items posted to an account's statement can be retrieved by calling the client's `GetNewStatementItems` method.
This method returns either all items posted to the account within the past year (up to the current date) or all items posted after a specified item. This method also returns a cursor, which is the ID of the most recently posted item.

To get all items posted after a specified item, we must provide the item's ID as the cursor
```go
items, nextCursor, err := client.GetNewStatementItems("accountID", "cursor")
if err != nil {
    panic("failed to get new statement items", err)
}

for _, item := range items {
    fmt.Println(item.Amount)
}
```

To get all items posted within the past year (up to the current date), we must leave the cursor blank
```go
client.GetNewStatementItems("accountID", "")
```

### Making Money Transfers
Money transfers can be instantiated by calling the client's `InitiateTransfer` method.