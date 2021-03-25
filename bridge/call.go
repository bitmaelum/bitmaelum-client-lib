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
	return m.instance.OpenVault(payload)
}
