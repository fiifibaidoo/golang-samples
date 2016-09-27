// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// [START pubsub_quickstart]
package main

import (
  "fmt"
  "golang.org/x/net/context"
  "log"

  // Imports the Google Cloud client library
  "cloud.google.com/go/pubsub"
)

func main() {
  ctx := context.Background()

  // Your Google Cloud Platform project ID
  projectId := "YOUR_PROJECT_ID"

  // Instantiates a client
  client, err := pubsub.NewClient(ctx, projectId)
  if err != nil {
    log.Fatalf("Failed to create client: %v", err)
  }

  // The name for the new topic
  topicName := "my-new-topic"

  // Creates the new topic
  topic, err := client.CreateTopic(ctx, topicName)
  if err != nil {
    log.Fatalf("Failed to create topic: %v", err)
  }

  fmt.Printf("Topic %v created.", topic)
}
// [END pubsub_quickstart]
