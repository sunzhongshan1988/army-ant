package scalar

import (
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"io"
)

func MarshalBrokerScalar(b *model.Broker) graphql.Marshaler {
	jsonByte, err := json.Marshal(b)
	return graphql.WriterFunc(func(w io.Writer) {
		if err == nil {
			_, _ = w.Write(jsonByte)
		} else {
			_, _ = w.Write(nil)
		}
	})
}

func UnmarshalBrokerScalar(v interface{}) (*model.Broker, error) {
	b := &model.Broker{}

	str, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("must be a string")
	}
	err := json.Unmarshal([]byte(str), &b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
