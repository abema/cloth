# Cloth [![Build Status](https://img.shields.io/circleci/project/abema/cloth/master.svg?style=flat)](https://circleci.com/gh/abema/cloth) [![Coverage Status](https://img.shields.io/codecov/c/github/abema/cloth/master.svg?style=flat)](https://codecov.io/github/abema/cloth) [![GoDoc](https://godoc.org/github.com/abema/cloth?status.svg)](https://godoc.org/github.com/abema/cloth) [![GitHub license](https://img.shields.io/github/license/abema/cloth.svg)](https://github.com/abema/cloth/blob/master/LICENSE)

**Under development**

ORM over [Cloud Bigtable](https://cloud.google.com/bigtable/) by Golang

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
	"google.golang.org/cloud"
	"google.golang.org/cloud/bigtable"
	"google.golang.org/cloud/bigtable/bttest"
	"google.golang.org/grpc"
)

const (
	project = "cloth"
	zone    = "local"
	cluster = "dummy"

	tbl = "myTable"
	cf  = "myFamily"
)

// User data model.
type User struct {
	ID         string `bigtable:",rowkey"`
	Name       string `bigtable:"name"`
	Content    string `bigtable:"content, omitempty"`
	Purchased  bool   `bigtable:"purchased"`
	LoggedInAt int64  `bigtable:"-"`
	CreatedAt  int64  `bigtable:"createdAt"`
}

var srv *bttest.Server

func init() {

	var err error

	srv, err = bttest.NewServer()
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(srv.Addr)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	admin, err := bigtable.NewAdminClient(ctx, project, zone, cluster, cloud.WithBaseGRPC(conn))
	if err != nil {
		panic(err)
	}
	defer admin.Close()

	err = admin.CreateTable(ctx, tbl)
	if err != nil {
		panic(err)
	}

	err = admin.CreateColumnFamily(ctx, tbl, cf)
	if err != nil {
		panic(err)
	}
}

func main() {

	// test struct
	user := User{
		ID:         "vytxeTZskVKR7C7WgdSP3d",
		Name:       "osamingo",
		Content:    "WRYYY!",
		Purchased:  true,
		LoggedInAt: time.Now().Unix(),
		CreatedAt:  time.Now().Add(-24 * time.Hour).Unix(),
	}

	// dialing to test server
	conn, err := grpc.Dial(srv.Addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// generate a client
	client, err := bigtable.NewClient(context.Background(), project, zone, cluster, cloud.WithBaseGRPC(conn))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	// generate Mutation
	mutation, err := cloth.GenerateMutation(cf, bigtable.Now(), &user)
	if err != nil {
		fmt.Println(err)
		return
	}

	// apply
	err = client.Open(tbl).Apply(context.Background(), user.ID, mutation)
	if err != nil {
		fmt.Println(err)
		return
	}

	// read a row
	row, err := client.Open(tbl).ReadRow(context.Background(), user.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert
	target := User{}
	err = cloth.ReadItems(row[cf], &target)
	if err != nil {
		fmt.Println(err)
		return
	}

	// check
	if target.ID == user.ID {
		fmt.Println("ID:\t\tOK")
	}
	if target.Name == user.Name {
		fmt.Println("Name:\t\tOK")
	}
	if target.Content == user.Content {
		fmt.Println("Content:\tOK")
	}
	if target.Purchased == user.Purchased {
		fmt.Println("Purchased:\tOK")
	}
	if target.LoggedInAt == 0 {
		fmt.Println("LoggedInAt:\tOK (shoud be zero value)")
	}
	if target.CreatedAt == user.CreatedAt {
		fmt.Println("CreatedAt:\tOK")
	}
}
```

## License

Released under the [MIT License](https://github.com/abema/cloth/blob/master/LICENSE)
