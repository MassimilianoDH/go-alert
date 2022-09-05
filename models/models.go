package models

// Google Cloud Monitoring Alert
// https://cloud.google.com/monitoring/support/notification-options#webhooks

type Alert struct {
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

type Message struct {
	ProjectID    string
	ResourceType string
	PolicyName   string
	ThreatLevel  string
	Summary      string
	URL          string
}
