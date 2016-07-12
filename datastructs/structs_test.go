package datastructs

import "testing"

func TestDataStructs_NewDataServiceMember(t *testing.T) {
	t.Parallel()
	dataServiceMemberStr := `
	{
		"role": "master",
		"state": "running",
		"conn_url": "postgres://appuser:rminKTk9kOLWWlvh@10.244.21.7:32768/postgres",
		"api_url": "http://10.244.21.7:32769/patroni",
		"xlog_location": 6593446208
	}
  `
	dataServiceMember, err := NewDataServiceMember(dataServiceMemberStr)
	if err != nil {
		t.Fatalf("NewDataServiceMember error: %v", err)
	}

	expectedRole := "master"
	if dataServiceMember.Role != expectedRole {
		t.Fatalf("Service member role should be %s", expectedRole)
	}

	expectedAPI := "http://10.244.21.7:32769/patroni"
	if dataServiceMember.APIURL != expectedAPI {
		t.Fatalf("Service member API should be %s", expectedAPI)
	}

	expectedRootAPI := "http://10.244.21.7:32769/"
	if dataServiceMember.RootAPIURL != expectedRootAPI {
		t.Fatalf("Service member root API should be %s", expectedRootAPI)
	}
}
