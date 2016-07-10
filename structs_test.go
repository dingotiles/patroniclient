package patroniclient

import "testing"

func TestStructs_NewServiceMember(t *testing.T) {
	t.Parallel()
	serviceMemberStr := `
    {
    "key": "/service/025ea0b0-710e-4da2-890d-f245a4d35259/members/49a91b69-264b-4df9-a836-523072b2f778",
    "value": "{\"role\":\"master\",\"state\":\"running\",\"conn_url\":\"postgres://appuser:rminKTk9kOLWWlvh@10.244.21.7:32768/postgres\",\"api_url\":\"http://10.244.21.7:32769/patroni\",\"xlog_location\":6593446208}",
    "expiration": "2016-07-10T22:48:49.164821357Z",
    "ttl": 27,
    "modifiedIndex": 100148,
    "createdIndex": 100148
  }
  `
	serviceMember, err := NewServiceMember(serviceMemberStr)
	if err != nil {
		t.Fatalf("NewServiceMember error: %v", err)
	}

	expectedKey := "/service/025ea0b0-710e-4da2-890d-f245a4d35259/members/49a91b69-264b-4df9-a836-523072b2f778"
	if serviceMember.Key != expectedKey {
		t.Fatalf("Service member key should be %s", expectedKey)
	}

	expectedRole := "master"
	if serviceMember.Value.Role != expectedRole {
		t.Fatalf("Service member role should be %s", expectedRole)
	}
}
