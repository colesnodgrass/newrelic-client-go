package cloud

import (
	"encoding/json"
	"fmt"

	"github.com/newrelic/newrelic-client-go/internal/http"
	"github.com/newrelic/newrelic-client-go/internal/logging"
	"github.com/newrelic/newrelic-client-go/pkg/config"
)

type Cloud struct {
	client http.Client
	config config.Config
	logger logging.Logger
	pager  http.Pager
}

func New(config config.Config) Cloud {

	client := http.NewClient(config)
	client.SetAuthStrategy(&http.PersonalAPIKeyCapableV2Authorizer{})

	pkg := Cloud{
		client: client,
		config: config,
		logger: config.GetLogger(),
		pager:  &http.LinkHeaderPager{},
	}

	return pkg
}
func (r *CloudLinkedAccount) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	fmt.Printf("cloudLinkedAccount: objMap: %+v\n", objMap)

	// var cloudLinkedAccount CloudLinkedAccount
	// err = json.Unmarshal(b, &cloudLinkedAccount)
	// if err != nil {
	// 	fmt.Printf("err, cloudLinkedAccount: %+v\n", cloudLinkedAccount)
	//
	// 	return err
	// }

	return nil
}

func (r *linkedAccountsResponse) UnmarshalJSONx(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	var actor Actor
	err = json.Unmarshal(*objMap["actor"], &actor)
	if err != nil {

		fmt.Printf("err, actor: %+v\n", actor)

		return err
	}
	fmt.Printf("actor: %+v\n", actor)

	fmt.Printf("objMap: %+v\n", objMap)

	var rawMessagesForActor map[string]*json.RawMessage
	err = json.Unmarshal(*objMap["actor"], &rawMessagesForActor)
	if err != nil {
		return err
	}

	fmt.Printf("rawMessagesForActor: %+v\n", rawMessagesForActor)
	// fmt.Printf("rawMessagesForActor: %+v", string(*rawMessagesForActor))

	var rawMessagesForCloud map[string]*json.RawMessage
	err = json.Unmarshal(*rawMessagesForActor["cloud"], &rawMessagesForCloud)
	if err != nil {
		return err
	}

	fmt.Printf("rawMessagesForCloud: %+v\n", rawMessagesForCloud)

	var rawMessagesForLinkedAccounts []*json.RawMessage
	err = json.Unmarshal(*rawMessagesForCloud["linkedAccounts"], &rawMessagesForLinkedAccounts)
	if err != nil {
		return err
	}

	fmt.Printf("rawMessagesForLinkedAccounts: %+v\n", rawMessagesForLinkedAccounts)

	// // Let's add a place to store our de-serialized Plant and Animal structs
	// r.Actor = make([]ColoredThing, len(rawMessagesForColoredThings))
	//
	var m map[string]*json.RawMessage
	for i, rawMessageForLinkedAccount := range rawMessagesForLinkedAccounts {
		// fmt.Printf("rawMessageForLinkedAccount: %d %+v\n", i, string(*rawMessageForLinkedAccount))

		err = json.Unmarshal(*rawMessageForLinkedAccount, &m)
		if err != nil {
			return err
		}

		fmt.Printf("m: %d %+v\n", i, m)

		// 	// Depending on the type, we can run json.Unmarshal again on the same byte slice
		// 	// But this time, we'll pass in the appropriate struct instead of a map
		// 	if m["type"] == "plant" {
		// 		var p Plant
		// 		err := json.Unmarshal(*rawMessage, &p)
		// 		if err != nil {
		// 			return err
		// 		}
		// 		// After creating our struct, we should save it
		// 		ce.Things[index] = &p
		// 	} else if m["type"] == "animal" {
		// 		var a Animal
		// 		err := json.Unmarshal(*rawMessage, &a)
		// 		if err != nil {
		// 			return err
		// 		}
		// 		// After creating our struct, we should save it
		// 		ce.Things[index] = &a
		// 	} else {
		// 		return errors.New("Unsupported type found!")
		// 	}
	}

	return nil
}
