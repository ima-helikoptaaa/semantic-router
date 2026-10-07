package kvtransfer

import (
	"github.com/vllm-project/semantic-router/src/semantic-router/pkg/config"
	"testing"
	"time"
)

func TestAddressRegistryConnectionTimeoutUsesSeconds(t *testing.T) {
	redis := &config.RedisConfig{}
	redis.Connection.Timeout = 3
	valkey := &config.ValkeyConfig{}
	valkey.Connection.Timeout = 4
	if got := redisOptionsFromConfig(redis).Timeout; got != 3*time.Second {
		t.Fatalf("redis timeout = %s", got)
	}
	if got := valkeyOptionsFromConfig(valkey).Timeout; got != 4*time.Second {
		t.Fatalf("valkey timeout = %s", got)
	}
}
