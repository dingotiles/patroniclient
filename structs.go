package patroniclient

import (
	"encoding/json"
	"io"
	"strings"
	"time"
)

type ServiceMember struct {
	Key           string    `json:"key"`
	JSONValue     string    `json:"value"`
	Expiration    time.Time `json:"expiration"`
	TTL           int       `json:"ttl"`
	ModifiedIndex int       `json:"modifiedIndex"`
	CreatedIndex  int       `json:"createdIndex"`
	Value         ServiceMemberValue
}

type ServiceMemberValue struct {
	Role         string `json:"role"`
	State        string `json:"state"`
	ConnURL      string `json:"conn_url"`
	APIURL       string `json:"api_url"`
	XlogLocation int64  `json:"xlog_location"`
}

func NewServiceMember(jsonStream string) (serviceMember *ServiceMember, err error) {
	serviceMember = &ServiceMember{}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	if err = dec.Decode(&serviceMember); err == io.EOF {
		return
	} else if err != nil {
		return
	}

	dec = json.NewDecoder(strings.NewReader(serviceMember.JSONValue))
	err = dec.Decode(&serviceMember.Value)

	return
}
