package bitmaelumClient

import (
	"github.com/bitmaelum/bitmaelum-suite/pkg/vault"
)

// OpenVault ...
func (b *BitMaelumClient) OpenVault(path, password string) (interface{}, error) {
	v, err := vault.Open(path, password)
	if err != nil {
		return nil, err
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

	return result, nil
}
