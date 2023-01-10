package config

// DEFAULT CONFIG
var DefaultConfig = map[string]interface{}{
	"name":        "INVITATION-GATEWAY",
	"port":        ":4000",
	"rpc_product": "localhost:3001",
	"rpc_order":   ":3002",
	"rpc_auth":    ":3003",
}
