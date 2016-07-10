package apistructs

import "testing"

func TestAPIStructs_NewAPIPatroniState(t *testing.T) {
	t.Parallel()
	apiStatusStr := `
  {
    "database_system_identifier": "6304830698674094206",
    "postmaster_start_time": "2016-07-08 06:03:24.859 UTC",
    "xlog": {
      "received_location": 6643777856,
      "replayed_timestamp": null,
      "paused": false,
      "replayed_location": 6643777856
    },
    "patroni": {
      "scope": "025ea0b0-710e-4da2-890d-f245a4d35259",
      "version": "0.90"
    },
    "state": "running",
    "role": "replica",
    "server_version": 90503
  }
  `
	apiStatus, err := NewAPIPatroniState(apiStatusStr)
	if err != nil {
		t.Fatalf("NewAPIPatroniState error: %v", err)
	}

	expectedRole := "replica"
	if apiStatus.Role != expectedRole {
		t.Fatalf("Member role should be %s", expectedRole)
	}
}
