package events

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	Types "github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
)

func CreateCustomMetric(cw *cloudwatch.Client) (*cloudwatch.PutMetricDataOutput, error) {

	dimensions := []types.Dimension{
		{
			Name:  aws.String("InstanceId"),
			Value: aws.String("i-1234567890abcdef0"), // Replace with your resource ID.
		},
	}

	input := &cloudwatch.PutMetricDataInput{
		Namespace: aws.String("MyCustomNameSpace"),
		MetricData: []types.MetricDatum{
			{
				MetricName: aws.String("MyCustomMetric"),
				Dimensions: dimensions,
				Value:      aws.Float64(42.0),
				Unit:       types.StandardUnit("Count"),
				Timestamp:  aws.Time(time.Now()),
			},
		},
	}

	output, err := cw.PutMetricData(context.TODO(), input)

	if err != nil {
		return nil, err
	}

	return output, nil
}

func CreateEnableMetricAlarm(cw *cloudwatch.Client) (*cloudwatch.PutMetricAlarmOutput, error) {

	input := &cloudwatch.PutMetricAlarmInput{
		AlarmName:          aws.String("MyAlarmName"),
		MetricName:         aws.String("MyMetricName"),
		Namespace:          aws.String("MyNameSpace"),
		Threshold:          aws.Float64(50.0),
		ComparisonOperator: types.ComparisonOperatorGreaterThanOrEqualToThreshold,
		AlarmDescription:   aws.String("My custom metric alarm description"),
		ActionsEnabled:     aws.Bool(true), // Enable actions.

	}

	output, err := cw.PutMetricAlarm(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil

}

func DisableMetricAlarm(cw *cloudwatch.Client) (*cloudwatch.DisableAlarmActionsOutput, error) {

	input := &cloudwatch.DisableAlarmActionsInput{
		AlarmNames: []string{"CustomAlarm"},
	}

	output, err := cw.DisableAlarmActions(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func ListMetrics(cw *cloudwatch.Client) (*cloudwatch.ListMetricsOutput, error) {

	input := &cloudwatch.ListMetricsInput{}
	output, err := cw.ListMetrics(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil

}

func PutEvents(cwe *cloudwatchevents.Client) (*cloudwatchevents.PutEventsOutput, error) {

	input := &cloudwatchevents.PutEventsInput{
		Entries: []Types.PutEventsRequestEntry{
			{
				EventBusName: aws.String(""),
				DetailType:   aws.String(""),
				Detail:       aws.String(""),
				Source:       aws.String("my-custom-application"),
				Time:         aws.Time(time.Now()), // Use the current timestamp or specify your own.
				Resources:    []string{"resource-arn"},
			},
		},
	}

	output, err := cwe.PutEvents(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}
