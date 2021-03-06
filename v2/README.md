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
This method returns either all items posted to the account within the past year (up to the current date) or all items posted after a specified item. This method also returns a cursor, which is the ID of the most recently posted item. The final argument is an optional filter over statement items and ensures that only items that pass the filter are returned.

To get all items posted after a specified item, we must provide the item's ID as the cursor
```go
items, nextCursor, err := client.GetNewStatementItems("accountID", "cursor", nil)
if err != nil {
    panic("failed to get new statement items", err)
}

for _, item := range items {
    fmt.Println(item.Amount)
}
```

To get all items posted within the past year (up to the current date), we must leave the cursor blank
```go
allItems, nextCursor := client.GetNewStatementItems("accountID", "", nil)
```

To get all non-negative entries (inflows) that were posted after a specified item
```go
nonNegativeEntries, nextCursor := client.GetNewStatementItems("accountID", "cursor", func(item *bmlgo.HistoryItem) bool {
    return !item.Minus
})
```

### Making Money Transfers
Money transfers can be initiated by calling the client's `InitiateTransfer` method. The amount of the transfer must be provided in Laari as an integer. The debit account and credit accounts must also be provided.

To complete a transfer we must call the client's `CompleteTransfer` method with the transfer form returned by `InitiateTransfer` and the OTP received via email or SMS.

Here's an example where we transfer 1 Rufiyaa (100 Laari) from "AccountA" to "AccountB"
```go
transferForm, err := client.InitiateTransfer(100, "AccountA", "AccountB")
if err != nil {
    panic("couldn't initiate transfer", err)
}

// then when the OTP is available
transferRecord, err := client.CompleteTransfer(transferForm, "OTP")
if err != nil {
    panic("couldn't complete transfer", err)
}

fmt.Println(transferRecord.Reference)
```