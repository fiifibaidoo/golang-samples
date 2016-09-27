// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// [START datastore_quickstart]
package main

import (
  "fmt"
  "golang.org/x/net/context"
  "log"

  // Imports the Google Cloud client library
  "cloud.google.com/go/datastore"
)

type Entity struct {
  Value string
}

func main() {
  ctx := context.Background()

  // Your Google Cloud Platform project ID
  projectId := "YOUR_PROJECT_ID"

  // Instantiates a client
  client, err := datastore.NewClient(ctx, projectId)
  if err != nil {
    log.Fatalf("Failed to create client: %v", err)
  }

  // The kind of the entity to retrieve
  kind := "Task"
  // The id of the entity to retrieve
  var id int64 = 1234567890
  // The Datastore key for the entity
  key := datastore.NewKey(ctx, kind, "", id, nil)

  entity := new(Entity)

  // Retrieves the entity
  if err := client.Get(ctx, key, entity); err != nil {
    log.Fatalf("Failed to get entity: %v", err)
  }

  fmt.Printf("Got entity: %v", key.String())
}
// [END datastore_quickstart]
