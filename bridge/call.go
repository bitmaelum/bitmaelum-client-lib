package bitmaelumClientBridge

import (
	"encoding/json"
	"fmt"

	"github.com/bitmaelum/bitmaelum-suite/pkg/vault"
)

// Call ...
func Call(name string, payload []byte) ([]byte, error) {
	instance := NewInstance()

	var output map[string]interface{}
	switch name {
	case "openVault":
		output = instance.openVault(payload)
	default:
		return nil, fmt.Errorf("not implemented: %s", name)
	}

	return json.Marshal(output)
}

type instance struct {
}

func NewInstance() *instance {
	return &instance{}
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

	v, err := vault.Open(arguments["path"], arguments["password"])
	if err != nil {
		return map[string]interface{}{
			"error":    "error opening vault, check your password",
			"response": nil,
		}
	}

	result := make([]map[string]interface{}, len(v.Store.Accounts))

	for i, acc := range v.Store.Accounts {
		pk := acc.GetActiveKey().PrivKey
		privkey := pk.String()

		result[i] = map[string]interface{}{
			"address":     acc.Address.String(),
			"hash":        acc.Address.Hash().String(),
			"name":        acc.Name,
			"private_key": privkey,
		}
	}

	return map[string]interface{}{
		"error":    nil,
		"response": result,
	}
}
