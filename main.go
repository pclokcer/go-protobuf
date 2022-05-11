package main

import (
	"fmt"

	alert_service_proto "go_protobuf/protos"

	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	// Encoding
	var alertData alert_service_proto.AlertStreamData
	alertData.Operator = "test_Operator"
	alertData.Value = "5555"
	encoded, _ := protojson.Marshal(&alertData)
	fmt.Println(encoded)

	// Decoding
	var newData alert_service_proto.AlertStreamData
	protojson.Unmarshal(encoded, &newData)

	fmt.Println(newData.GetOperator(), newData.GetValue())
}
