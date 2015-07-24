# Cloth

[![Build Status](https://img.shields.io/circleci/project/abema/cloth/master.svg?style=flat)](https://circleci.com/gh/abema/cloth)
[![Coverage Status](https://img.shields.io/codecov/c/github/abema/cloth/master.svg?style=flat)](https://codecov.io/github/abema/cloth)

**Under development**

ORM over Cloud Bigtable by Golang.

## Installation

```
$ go get -u github.com/abema/cloth
```

## Quick Start

```go
package main

import (
    "fmt"
    "time"

    "github.com/abema/cloth"
    "golang.org/x/net/context"
    "google.golang.org/cloud/bigtable"
)

type User struct {
    ID        string `bigtable:",rowkey"`
    Name      string `bigtable:"name"`
    Content   string `bigtable:"content,omitempty"`
    Status    int    `bigtable:"-"`
    CreatedAt int64  `bigtable:"createdAt"`
}

var client *bigtable.Client

func init() {
    // Initialize client.
}

func main() {

    u := entity.User{
        ID:        "john_doe",
		Name:      "John Doe",
        CreatedAt: time.Now().Unix(),
	}

    // Genetate Mutation
    m, err := cloth.GenerateMutation("FAMILY_COLUMN_NAME", bigtable.Now(), &u)
    if err != nil {
	    panic(err)
    }

    // Apply for Cloud Bigtable
    err = client.Open("TABLE_NAME").Apply(context.Background(), u.ID, m)
    if err != nil {
        panic(err)
    }

    // Read rows from Cloud Bigtable
    err = client.Open("TABLE_NAME").ReadRows(context.Background(), bigtable.PrefixRange(""), func(r bigtable.Row) bool {
        u := User{}
        cloth.ReadItems("FAMILY_COLUMN_NAME", &u)
        fmt.Println("Read row: ", u)
        return true
	})
    if err != nil {
        panic(err)
    }
}
```

## License

Released under the [MIT License](https://github.com/osamingo/cloth/blob/master/LICENSE).
