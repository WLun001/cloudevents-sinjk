package main

import (
	"context"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event/datacodec/json"
	"log"
	"net/http"
)

var eventTypes = []string{
	"dev.tekton.event.pipelinerun.started.v1",
	"dev.tekton.event.pipelinerun.successful.v1",
	"dev.tekton.event.pipelinerun.failed.v1",
}

var pipelineStatus = []string{
	"Started", "Succeeded", "Failed",
}

func main() {
	ctx := context.Background()
	p, err := cloudevents.NewHTTP()
	if err != nil {
		log.Fatalf("failed to create protocol: %s", err.Error())
	}

	c, err := cloudevents.NewClient(p)
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	log.Printf("will listen on :8080\n")
	log.Fatalf("failed to start receiver: %s", c.StartReceiver(ctx, receive))
}

func receive(ctx context.Context, event cloudevents.Event) cloudevents.Result {
	if contains(eventTypes, event.Type()) {

		data := new(PipelineRunModel)
		err := json.Decode(ctx, event.Data(), &data)
		if err != nil {
			return cloudevents.NewHTTPResult(http.StatusInternalServerError, err.Error())

		}
		condition := data.PipelineRun.Status.Conditions[0]
		if contains(pipelineStatus, condition.Type) {
			msg := fmt.Sprintf("%s (%s): %s",
				data.PipelineRun.Spec.PipelineRef.Name, condition.Reason, condition.Message)
			sendMessage(msg)
			return cloudevents.NewHTTPResult(http.StatusOK, "")
		}
	}
	return cloudevents.NewHTTPResult(http.StatusOK, "")
}

func contains[T comparable](arr []T, value T) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
