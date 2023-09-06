package main

import (
	"context"
	"fmt"

	events "github.com/Asad2730/KMS_Example/Events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err.Error())
	}

	cwClient := cloudwatch.NewFromConfig(cfg)
	cweClient := cloudwatchevents.NewFromConfig(cfg)

	res, err := events.CreateCustomMetric(cwClient)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(res.ResultMetadata)
	rs, err := events.PutEvents(cweClient)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(rs.Entries)
}
