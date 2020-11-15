# uuid

A wrapper of [github.com/google/uuid](https://github.com/google/uuid) and provides deployable serverless handler on [@vercel](https://github.com/vercel).

## Using this as a Go library
Get it using `go get`:
```sh
go get -u github.com/husniadil/uuid
```

Import it from your code:
```go
import "github.com/husniadil/uuid/uuid"

...
...
size := 1 // how many UUIDs to generate?
hypen := true // should UUID be formatted with hypen?
uppercase := false // should UUID be formatted in uppercase?

// UUID generation request
req := uuid.Request{
  Version:   version,   // UUID version to generate
  Domain:    domain,    // param for version 2 UUID
  ID:        id,        // param for version 2 UUID
  Namespace: namespace, // param for version 3 and 5 UUID
  Data:      data,      // param for version 3 and 5 UUID
}

// generates a list of UUIDs
uuids, err := uuid.Generate(size, hypen, uppercase, req)
...
...

// get UUID's metadata
metadata, err := uuid.Parse("a51a2ef7-f80d-4152-bdbd-abeb6579ee3d")
...
...
```

## Deployment on vercel
Import this repo on vercel, no fancy settings are required. For more detail, refer to the [Vercel Go Runtime](https://vercel.com/docs/runtimes#official-runtimes/go) docs.

## For reading
https://en.wikipedia.org/wiki/Universally_unique_identifier#Versions
