package bitmaelumClient

import (
	"encoding/json"

	"github.com/bitmaelum/bitmaelum-suite/pkg/vault"
)

// OpenVault ...
func (b *BitMaelumClient) OpenVault(payload []byte) map[string]interface{} {
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
