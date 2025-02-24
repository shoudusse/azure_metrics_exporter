package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

// PrintPrettyJSON - Prints structs nicely for debugging.
func PrintPrettyJSON(input map[string]interface{}) {
	out, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		log.Fatalf("Error indenting JSON: %v", err)
	}
	fmt.Println(string(out))
}

// GetTimes - Returns the endTime and startTime used for querying Azure Metrics API
func GetTimes() (string, string) {
	// Make sure we are using UTC
	now := time.Now().UTC()

	// Use query delay of 3 minutes when querying for latest metric data
	endTime := now.Add(time.Minute * time.Duration(-3)).Format(time.RFC3339)
	startTime := now.Add(time.Minute * time.Duration(-4)).Format(time.RFC3339)
	return endTime, startTime
}

// CreateResourceLabels - Returns resource labels for a give resource ID.
func CreateResourceLabels(resourceID string) map[string]string {
	labels := make(map[string]string)
	resource := strings.Split(resourceID, "/")
	labels["resource_group"] = resource[4]
	labels["resource_name"] = resource[8]
	if len(resource) > 13 {
		labels["sub_resource_name"] = resource[10]
	}

	return labels
}

func hasAggregation(aggregations []string, aggregation string) bool {
	if len(aggregations) == 0 {
		return true
	}

	for _, aggr := range aggregations {
		if aggr == aggregation {
			return true
		}
	}
	return false
}

func filterAggregations(aggregations []string) []string {
	base := []string{"Total", "Average", "Minimum", "Maximum"}
	if len(aggregations) > 0 {
		return aggregations
	}
	return base
}
