// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// [START storage_quickstart]
package main

import (
  "fmt"
  "golang.org/x/net/context"
  "log"

  // Imports the Google Cloud client library
  "cloud.google.com/go/storage"
)

func main() {
  ctx := context.Background()

  // Your Google Cloud Platform project ID
  projectId := "YOUR_PROJECT_ID"

  // Instantiates a client
  client, err := storage.NewClient(ctx)
  if err != nil {
    log.Fatalf("Failed to create client: %v", err)
  }

  // The name for the new bucket
  bucketName := "my-new-bucket"

  // Prepares a new bucket
  bucket := client.Bucket(bucketName)

  // Creates the new bucket
  if err := bucket.Create(ctx, projectId, nil); err != nil {
    log.Fatalf("Failed to create bucket: %v", err)
  }

  fmt.Printf("Bucket %v created.", bucketName)
}
// [END storage_quickstart]
