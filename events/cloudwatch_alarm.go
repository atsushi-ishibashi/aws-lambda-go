package events

import (
	"encoding/json"
	"errors"
	"time"
)

type CloudWatchAlarmStateChangeTime time.Time

const cloudWatchAlarmStateChangeTimeReference = "\"2006-01-02T15:04:05.000-0700\""

func (t CloudWatchAlarmStateChangeTime) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(t).Format(cloudWatchAlarmStateChangeTimeReference)), nil
}

func (t *CloudWatchAlarmStateChangeTime) UnmarshalJSON(data []byte) error {
	if t == nil {
		return errors.New("CloudWatchAlarmStateChangeTime cannot be nil")
	}

	pt, err := time.Parse(cloudWatchAlarmStateChangeTimeReference, string(data))
	if err == nil {
		*t = CloudWatchAlarmStateChangeTime(pt)
	}
	return err
}

// CloudWatchAlarm is the outer structure of an event sent via CloudWatch Alarm.
type CloudWatchAlarm struct {
	AlarmName        string                         `json:"AlarmName"`
	AlarmDescription string                         `json:"AlarmDescription"`
	AWSAccountID     string                         `json:"AWSAccountId"`
	NewStateValue    string                         `json:"NewStateValue"`
	NewStateReason   string                         `json:"NewStateReason"`
	StateChangeTime  CloudWatchAlarmStateChangeTime `json:"StateChangeTime"`
	Region           string                         `json:"Region"`
	OldStateValue    string                         `json:"OldStateValue"`
	Trigger          CloudWatchAlarmTrigger         `json:"Trigger"`
}

// CloudWatchAlarmTrigger is Trigger of CloudWatchAlarm
type CloudWatchAlarmTrigger struct {
	MetricName                       string          `json:"MetricName"`
	Namespace                        string          `json:"Namespace"`
	StatisticType                    string          `json:"StatisticType"`
	Statistic                        string          `json:"Statistic"`
	Unit                             *string         `json:"Unit"`
	Dimensions                       json.RawMessage `json:"Dimensions"`
	Period                           int64           `json:"Period"`
	EvaluationPeriods                int64           `json:"EvaluationPeriods"`
	ComparisonOperator               string          `json:"ComparisonOperator"`
	Threshold                        float64         `json:"Threshold"`
	TreatMissingData                 string          `json:"TreatMissingData"`
	EvaluateLowSampleCountPercentile string          `json:"EvaluateLowSampleCountPercentile"`
}
