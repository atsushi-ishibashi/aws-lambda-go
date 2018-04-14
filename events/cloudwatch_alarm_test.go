package events

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events/test"
)

func TestCloudWatchAlarmMarshaling(t *testing.T) {
	inputJSON := readJsonFromFile(t, "./testdata/cloudwatch-alarm.json")

	var inputEvent CloudWatchAlarm
	if err := json.Unmarshal(inputJSON, &inputEvent); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}

	outputJSON, err := json.Marshal(inputEvent)
	if err != nil {
		t.Errorf("could not marshal event. details: %v", err)
	}

	test.AssertJsonsEqual(t, inputJSON, outputJSON)
}
