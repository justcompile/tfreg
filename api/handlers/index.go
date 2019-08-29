package handlers

import (
	"encoding/json"
	"net/http"
)

func Index(resp http.ResponseWriter, req *http.Request) {
	data := map[string]interface{}{
		"meta": map[string]interface{}{
			"limit":          2,
			"current_offset": 0,
			"next_offset":    2,
			"next_url":       "/v1/modules?limit=2&offset=2&verified=true",
		},
		"modules": []map[string]interface{}{
			{
				"id":           "GoogleCloudPlatform/lb-http/google/1.0.4",
				"owner":        "",
				"namespace":    "GoogleCloudPlatform",
				"name":         "lb-http",
				"version":      "1.0.4",
				"provider":     "google",
				"description":  "Modular Global HTTP Load Balancer for GCE using forwarding rules.",
				"source":       "https://github.com/GoogleCloudPlatform/terraform-google-lb-http",
				"published_at": "2017-10-17T01:22:17.792066Z",
				"downloads":    213,
				"verified":     true,
			},
			{
				"id":           "terraform-aws-modules/vpc/aws/1.5.1",
				"owner":        "",
				"namespace":    "terraform-aws-modules",
				"name":         "vpc",
				"version":      "1.5.1",
				"provider":     "aws",
				"description":  "Terraform module which creates VPC resources on AWS",
				"source":       "https://github.com/terraform-aws-modules/terraform-aws-vpc",
				"published_at": "2017-11-23T10:48:09.400166Z",
				"downloads":    29714,
				"verified":     true,
			},
		},
	}

	if err := json.NewEncoder(resp).Encode(data); err != nil {
		http.Error(resp, "Internal Server Error", http.StatusInternalServerError)
	}
}
