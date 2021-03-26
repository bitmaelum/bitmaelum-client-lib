package bitmaelumClientBridge

import (
	"encoding/json"

	bitmaelumClient "github.com/bitmaelum/bitmaelum-client-lib/bitmaelum"
)

// Call ...
func Call(name string, payload []byte) ([]byte, error) {
	instance := NewInstance()

	var output map[string]interface{}
	switch name {
	case "openVault":
		output = instance.openVault(payload)
	default:
		return json.Marshal(map[string]interface{}{
			"error":    "not implemented",
			"response": nil,
		})
	}

	return json.Marshal(output)
}

type instance struct {
	instance *bitmaelumClient.BitMaelumClient
}

func NewInstance() *instance {
	return &instance{instance: bitmaelumClient.NewBitMaelumClient()}
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

	v, err := m.instance.OpenVault(arguments["path"], arguments["password"])
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
