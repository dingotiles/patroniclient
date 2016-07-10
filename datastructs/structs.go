package datastructs

import (
	"encoding/json"
	"io"
	"strings"
	"time"
)

type DataServiceMember struct {
	Key           string    `json:"key"`
	JSONValue     string    `json:"value"`
	Expiration    time.Time `json:"expiration"`
	TTL           int       `json:"ttl"`
	ModifiedIndex int       `json:"modifiedIndex"`
	CreatedIndex  int       `json:"createdIndex"`
	Value         DataServiceMemberValue
}

type DataServiceMemberValue struct {
	Role         string `json:"role"`
	State        string `json:"state"`
	ConnURL      string `json:"conn_url"`
	APIURL       string `json:"api_url"`
	XlogLocation int64  `json:"xlog_location"`
}

func NewDataServiceMember(jsonStream string) (dataServiceMember *DataServiceMember, err error) {
	dataServiceMember = &DataServiceMember{}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	if err = dec.Decode(&dataServiceMember); err == io.EOF {
		return
	} else if err != nil {
		return
	}

	dec = json.NewDecoder(strings.NewReader(dataServiceMember.JSONValue))
	err = dec.Decode(&dataServiceMember.Value)

	return
}
