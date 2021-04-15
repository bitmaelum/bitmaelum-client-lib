package bitmaelumClientBridge

import (
	"encoding/json"
	"sync"
	"time"

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
	case "sendMessage":
		output = instance.sendMessage(payload)
	case "sendSimpleMessage":
		output = instance.sendSimpleMessage(payload)
	case "setClientFromVault":
		output = instance.setClientFromVault(payload)
	case "setClientFromMnemonic":
		output = instance.setClientFromMnemonic(payload)
	case "setClientFromPrivateKey":
		output = instance.setClientFromPrivateKey(payload)
	case "listMessages":
		output = instance.listMessages(payload)
	case "readBlock":
		output = instance.readBlock(payload)
	case "saveAttachment":
		output = instance.saveAttachment(payload)
	default:
		return json.Marshal(map[string]interface{}{
			"error":    "the function " + name + " is not implemented on GO side",
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

func (m instance) sendMessage(payload []byte) map[string]interface{} {
	type messageArguments struct {
		Recipient   string            `json:"recipient"`
		Subject     string            `json:"subject"`
		Blocks      map[string]string `json:"blocks"`
		Attachments []string          `json:"attachments"`
	}

	var arguments messageArguments

	err := json.Unmarshal(payload, &arguments)
	if err != nil {
		return map[string]interface{}{
			"error": "failed to decode arguments",
		}
	}

	err = m.client.SendMessage(arguments.Recipient, arguments.Subject, arguments.Blocks, arguments.Attachments)
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

func (m instance) listMessages(payload []byte) map[string]interface{} {
	type listArguments struct {
		Since string `json:"since"`
		Box   int    `json:"box"`
	}

	var arguments listArguments

	err := json.Unmarshal(payload, &arguments)
	if err != nil {
		return map[string]interface{}{
			"error": "failed to decode arguments",
		}
	}

	since, err := time.Parse(time.RFC3339Nano, arguments.Since)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	messages, err := m.client.ListMessages(since, arguments.Box)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	return map[string]interface{}{
		"response": messages,
		"error":    nil,
	}
}

func (m instance) readBlock(payload []byte) map[string]interface{} {
	type listArguments struct {
		MsgID   string `json:"msgid"`
		BlockID string `json:"blockid"`
	}

	var arguments listArguments

	err := json.Unmarshal(payload, &arguments)
	if err != nil {
		return map[string]interface{}{
			"error": "failed to decode arguments",
		}
	}

	block, err := m.client.ReadBlock(arguments.MsgID, arguments.BlockID)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	return map[string]interface{}{
		"response": block,
		"error":    nil,
	}
}

func (m instance) saveAttachment(payload []byte) map[string]interface{} {
	type listArguments struct {
		MsgID        string `json:"msgid"`
		AttachmentID string `json:"attachmentid"`
		Path         string `json:"path"`
		Overwrite    bool   `json:"overwrite"`
	}

	var arguments listArguments

	err := json.Unmarshal(payload, &arguments)
	if err != nil {
		return map[string]interface{}{
			"error": "failed to decode arguments",
		}
	}

	response, err := m.client.SaveAttachment(arguments.MsgID, arguments.AttachmentID, arguments.Path, arguments.Overwrite)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	return map[string]interface{}{
		"response": response,
		"error":    nil,
	}
}
