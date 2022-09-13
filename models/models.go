package models

import "time"

// Google Cloud Monitoring Alert
// https://cloud.google.com/monitoring/support/notification-options#webhooks

type GoogleAlert struct {
	Incident struct {
		IncidentID              string `json:"incident_id"`
		ScopingProjectID        string `json:"scoping_project_id"`
		ScopingProjectNumber    int    `json:"scoping_project_number"`
		URL                     string `json:"url"`
		StartedAt               int    `json:"started_at"`
		EndedAt                 int    `json:"ended_at"`
		State                   string `json:"state"`
		ResourceID              string `json:"resource_id"`
		ResourceName            string `json:"resource_name"`
		ResourceDisplayName     string `json:"resource_display_name"`
		ResourceTypeDisplayName string `json:"resource_type_display_name"`
		Resource                struct {
			Type   string `json:"type"`
			Labels struct {
				InstanceID string `json:"instance_id"`
				ProjectID  string `json:"project_id"`
				Zone       string `json:"zone"`
			} `json:"labels"`
		} `json:"resource"`
		Metric struct {
			Type        string `json:"type"`
			DisplayName string `json:"displayName"`
			Labels      struct {
				InstanceName string `json:"instance_name"`
			} `json:"labels"`
		} `json:"metric"`
		Metadata struct {
			SystemLabels struct {
				Labelkey string `json:"labelkey"`
			} `json:"system_labels"`
			UserLabels struct {
				Labelkey string `json:"labelkey"`
			} `json:"user_labels"`
		} `json:"metadata"`
		PolicyName       string `json:"policy_name"`
		PolicyUserLabels struct {
			UserLabel1 string `json:"user-label-1"`
			UserLabel2 string `json:"user-label-2"`
		} `json:"policy_user_labels"`
		ConditionName  string `json:"condition_name"`
		ThresholdValue string `json:"threshold_value"`
		ObservedValue  string `json:"observed_value"`
		Condition      struct {
			Name               string `json:"name"`
			DisplayName        string `json:"displayName"`
			ConditionThreshold struct {
				Filter       string `json:"filter"`
				Aggregations []struct {
					AlignmentPeriod  string `json:"alignmentPeriod"`
					PerSeriesAligner string `json:"perSeriesAligner"`
				} `json:"aggregations"`
				Comparison     string  `json:"comparison"`
				ThresholdValue float64 `json:"thresholdValue"`
				Duration       string  `json:"duration"`
				Trigger        struct {
					Count int `json:"count"`
				} `json:"trigger"`
			} `json:"conditionThreshold"`
		} `json:"condition"`
		Summary string `json:"summary"`
	} `json:"incident"`
	Version string `json:"version"`
}

// Azure Cloud Monitoring Alert
// https://docs.microsoft.com/en-us/azure/azure-monitor/alerts/alerts-log-webhook

type AzureAlert struct {
	SchemaID string `json:"schemaId"`
	Data     struct {
		Essentials struct {
			AlertID             string    `json:"alertId"`
			AlertRule           string    `json:"alertRule"`
			Severity            string    `json:"severity"`
			SignalType          string    `json:"signalType"`
			MonitorCondition    string    `json:"monitorCondition"`
			MonitoringService   string    `json:"monitoringService"`
			AlertTargetIDs      []string  `json:"alertTargetIDs"`
			OriginAlertID       string    `json:"originAlertId"`
			FiredDateTime       time.Time `json:"firedDateTime"`
			Description         string    `json:"description"`
			EssentialsVersion   string    `json:"essentialsVersion"`
			AlertContextVersion string    `json:"alertContextVersion"`
		} `json:"essentials"`
		AlertContext struct {
			Properties struct {
				Name1 string `json:"name1"`
				Name2 string `json:"name2"`
			} `json:"properties"`
			ConditionType string `json:"conditionType"`
			Condition     struct {
				WindowSize string `json:"windowSize"`
				AllOf      []struct {
					SearchQuery         string `json:"searchQuery"`
					MetricMeasureColumn string `json:"metricMeasureColumn"`
					TargetResourceTypes string `json:"targetResourceTypes"`
					Operator            string `json:"operator"`
					Threshold           string `json:"threshold"`
					TimeAggregation     string `json:"timeAggregation"`
					Dimensions          []struct {
						Name  string `json:"name"`
						Value string `json:"value"`
					} `json:"dimensions"`
					MetricValue    float64 `json:"metricValue"`
					FailingPeriods struct {
						NumberOfEvaluationPeriods int `json:"numberOfEvaluationPeriods"`
						MinFailingPeriodsToAlert  int `json:"minFailingPeriodsToAlert"`
					} `json:"failingPeriods"`
					LinkToSearchResultsUI          string `json:"linkToSearchResultsUI"`
					LinkToFilteredSearchResultsUI  string `json:"linkToFilteredSearchResultsUI"`
					LinkToSearchResultsAPI         string `json:"linkToSearchResultsAPI"`
					LinkToFilteredSearchResultsAPI string `json:"linkToFilteredSearchResultsAPI"`
				} `json:"allOf"`
				WindowStartTime time.Time `json:"windowStartTime"`
				WindowEndTime   time.Time `json:"windowEndTime"`
			} `json:"condition"`
		} `json:"alertContext"`
	} `json:"data"`
}

// Amazon Web Services SNS Alert
// https://docs.aws.amazon.com/sns/latest/dg/sns-message-and-json-formats.html

type AmazonSubscriptionConfirmation struct {
	Type             string    `json:"Type"`
	MessageID        string    `json:"MessageId"`
	Token            string    `json:"Token"`
	TopicArn         string    `json:"TopicArn"`
	Message          string    `json:"Message"`
	SubscribeURL     string    `json:"SubscribeURL"`
	Timestamp        time.Time `json:"Timestamp"`
	SignatureVersion string    `json:"SignatureVersion"`
	Signature        string    `json:"Signature"`
	SigningCertURL   string    `json:"SigningCertURL"`
}

type AmazonAlert struct {
	Type             string    `json:"Type"`
	MessageID        string    `json:"MessageId"`
	TopicArn         string    `json:"TopicArn"`
	Subject          string    `json:"Subject"`
	Message          string    `json:"Message"`
	Timestamp        time.Time `json:"Timestamp"`
	SignatureVersion string    `json:"SignatureVersion"`
	Signature        string    `json:"Signature"`
	SigningCertURL   string    `json:"SigningCertURL"`
	UnsubscribeURL   string    `json:"UnsubscribeURL"`
}

type AmazonUnsubscribeConfirmation struct {
	Type             string    `json:"Type"`
	MessageID        string    `json:"MessageId"`
	Token            string    `json:"Token"`
	TopicArn         string    `json:"TopicArn"`
	Message          string    `json:"Message"`
	SubscribeURL     string    `json:"SubscribeURL"`
	Timestamp        time.Time `json:"Timestamp"`
	SignatureVersion string    `json:"SignatureVersion"`
	Signature        string    `json:"Signature"`
	SigningCertURL   string    `json:"SigningCertURL"`
}
