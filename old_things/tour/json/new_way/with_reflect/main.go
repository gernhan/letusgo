package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"time"
)

// map message type with original struct
var TypeDefinitionMap = map[string]reflect.Type{
	"Message": reflect.TypeOf(Message{}),
}

type DataWrapper struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type DataWrapperAlias DataWrapper

type JsonDataWrapper struct {
	DataWrapperAlias
}

func (jsonDataWrapper JsonDataWrapper) DataWrapper() DataWrapper {
	dataWrapper := DataWrapper(jsonDataWrapper.DataWrapperAlias)
	return dataWrapper
}

type Message struct {
	Description string `json:"description"`
}

var _ json.Unmarshaler = (*DataWrapper)(nil)
func (dataWrapper *DataWrapper) UnmarshalJSON(data []byte) error {
	var jsonMessage json.RawMessage
	jsonDataWrapper := JsonDataWrapper{
		DataWrapperAlias: DataWrapperAlias{
			Data: &jsonMessage,
		},
	}
	if err := json.Unmarshal(data, &jsonDataWrapper); err != nil {
		return err
	}

	typeName := jsonDataWrapper.Type
	var dataEntity interface{}
	if dataType, found := TypeDefinitionMap[typeName]; found {
		dataEntity = reflect.New(dataType).Interface()
	}

	// marshal again
	err := json.Unmarshal(jsonMessage, &dataEntity)
	if err != nil {
		return err
	}

	jsonDataWrapper.Data = dataEntity
	*dataWrapper = jsonDataWrapper.DataWrapper()
	return nil
}

func decode(input map[string]interface{}, result interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			toTimeHookFunc()),
		Result: result,
	})
	if err != nil {
		return err
	}

	if err := decoder.Decode(input); err != nil {
		return err
	}
	return err
}

func toTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		switch f.Kind() {
		case reflect.String:
			return time.Parse(time.RFC3339, data.(string))
		case reflect.Float64:
			return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
		case reflect.Int64:
			return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
		default:
			return data, nil
		}
		// Convert it by parsing
	}
}

func main() {

	jsonStr := []byte(`{"type":"Message","data":{"description": "trial message"}}`)
	var dataWrapper DataWrapper
	_ = json.Unmarshal(jsonStr, &dataWrapper)
	//mapOfInterface := dataWrapper.Data.(map[string]interface{})
	//var message Message
	//_ = decode(mapOfInterface, &message)
	message := dataWrapper.Data.(*Message)
	fmt.Println(message.Description)
}
