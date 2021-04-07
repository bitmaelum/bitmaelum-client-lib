package bitmaelumClientBridge

import (
	"encoding/json"
	"sync"

	bitmaelumClient "github.com/bitmaelum/bitmaelum-suite/library"
)

type instance struct {
	client *bitmaelumClient.BitMaelumClient
}

var (
	clientInstance *bitmaelumClient.BitMaelumClient
	once           sync.Once
)

func NewInstance() *instance {
	return &instance{client: bitmaelumClient.NewBitMaelumClient()}
}

func GetInstance() *instance {
	once.Do(func() {

		clientInstance = bitmaelumClient.NewBitMaelumClient()

	})

	return &instance{client: clientInstance}
}

// Call ...
func Call(name string, payload []byte) ([]byte, error) {
	instance := GetInstance()

	var output map[string]interface{}
	switch name {
	case "openVault":
		output = instance.openVault(payload)
	case "sendSimpleMessage":
		output = instance.sendSimpleMessage(payload)
	case "setClientFromVault":
		output = instance.setClientFromVault(payload)
	case "setClientFromMnemonic":
		output = instance.setClientFromMnemonic(payload)
	case "setClientFromPrivateKey":
		output = instance.setClientFromPrivateKey(payload)
	default:
		return json.Marshal(map[string]interface{}{
			"error":    "not implemented",
			"response": nil,
		})
	}

	return json.Marshal(output)
}

func (m instance) openVault(payload []byte) map[string]interface{} {
	var arguments map[string]string

	err := json.Unmarshal(payload, &arguments)
	if err != nil {
		return map[string]interface{}{
			"error":    "failed to decode arguments",
			"response": nil,
		}
	}

	v, err := m.client.OpenVault(arguments["path"], arguments["password"])
	if err != nil {
		return map[string]interface{}{
			"error":    err.Error(),
			"response": nil,
		}
	}

	return map[string]interface{}{
		"error":    nil,
		"response": v,
	}
}

func (m instance) sendSimpleMessage(payload []byte) map[string]interface{} {
	var arguments map[string]string

	err := json.Unmarshal(payload, &arguments)
	if err != nil {
		return map[string]interface{}{
			"error": "failed to decode arguments",
		}
	}

	err = m.client.SendSimpleMessage(arguments["recipient"], arguments["subject"], arguments["body"])
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	return map[string]interface{}{
		"error": nil,
	}
}

func (m instance) setClientFromVault(payload []byte) map[string]interface{} {
	var arguments map[string]string

	err := json.Unmarshal(payload, &arguments)
	if err != nil {
		return map[string]interface{}{
			"error": "failed to decode arguments",
		}
	}

	err = m.client.SetClientFromVault(arguments["account"])
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	return map[string]interface{}{
		"error": nil,
	}

}

func (m instance) setClientFromMnemonic(payload []byte) map[string]interface{} {
	var arguments map[string]string

	err := json.Unmarshal(payload, &arguments)
	if err != nil {
		return map[string]interface{}{
			"error": "failed to decode arguments",
		}
	}

	err = m.client.SetClientFromMnemonic(arguments["account"], arguments["name"], arguments["mnemonic"])
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	return map[string]interface{}{
		"error": nil,
	}

}

func (m instance) setClientFromPrivateKey(payload []byte) map[string]interface{} {
	var arguments map[string]string

	err := json.Unmarshal(payload, &arguments)
	if err != nil {
		return map[string]interface{}{
			"error": "failed to decode arguments",
		}
	}

	err = m.client.SetClientFromPrivateKey(arguments["account"], arguments["name"], arguments["private_key"])
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	return map[string]interface{}{
		"error": nil,
	}

}
