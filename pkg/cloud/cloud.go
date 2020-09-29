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
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	r.CreatedAt = EpochSeconds(int(raw["createdAt"].(float64)))
	r.UpdatedAt = EpochSeconds(int(raw["updatedAt"].(float64)))
	r.AuthLabel = raw["authLabel"].(string)
	r.ExternalId = raw["externalId"].(string)
	r.ID = int(raw["id"].(float64))
	r.Disabled = raw["disabled"].(bool)
	r.Name = raw["name"].(string)
	r.NrAccountId = int(raw["nrAccountId"].(float64))
	// TODO: Provider


	rawIntegrations := raw["integrations"].([]interface{})
	for _, ri := range rawIntegrations {
		rawIntegration := ri.(map[string]interface{})
		integrationType := rawIntegration["__typename"].(string)
		switch integrationType {
		case "CloudTrustedadvisorIntegration":
			integration := &CloudTrustedadvisorIntegration{}

			integration.CreatedAt = EpochSeconds(int(rawIntegration["createdAt"].(float64)))
			integration.UpdatedAt = EpochSeconds(int(rawIntegration["updatedAt"].(float64)))
			integration.Name = rawIntegration["name"].(string)
			integration.NrAccountId = int(rawIntegration["nrAccountId"].(float64))

			if rawIntegration["inventoryPollingInterval"] != nil {
				integration.InventoryPollingInterval = int(rawIntegration["inventoryPollingInterval"].(float64))
			}

			if rawIntegration["metricsPollingInterval"] != nil {
				integration.MetricsPollingInterval = int(rawIntegration["metricsPollingInterrval"].(float64))
			}

			// TODO: LinkedAccount

			r.Integrations = append(r.Integrations, integration)
		default:
			//fmt.Printf("integration type %s not recognized", integrationType)
		}
	}

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
