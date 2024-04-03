package go_session_auto_config

import (
	"github.com/spf13/cast"
	fastconfig "github.com/zj-kenzhou/fast-config"
	"github.com/zj-kenzhou/go-session/config"
)

func init() {
	hostAny := fastconfig.GetValue("session.host", []any{})
	hostAnyList, ok := hostAny.([]any)
	if !ok {
		panic("redis host is not valid")
	}
	if len(hostAnyList) == 0 {
		panic("redis host is empty")
	}
	var hostSlice []string
	for _, anyItem := range hostAnyList {
		hostSlice = append(hostSlice, cast.ToString(anyItem))
	}

	config.SetConfig(config.SessionConfig{
		Timeout:          cast.ToInt64(fastconfig.GetValue("session.timeout", 604800)),
		ActivityTimeout:  cast.ToInt64(fastconfig.GetValue("session.activity-timeout", 7200)),
		KeyPrefix:        fastconfig.GetString("session.key-prefix", ""),
		Host:             hostSlice,
		Username:         fastconfig.GetString("session.username", ""),
		Password:         fastconfig.GetString("session.password", ""),
		SentinelUsername: fastconfig.GetString("session.sentinel-username", ""),
		SentinelPassword: fastconfig.GetString("session.sentinel-password", ""),
		Db:               fastconfig.GetInt("session.db", 0),
		MasterName:       fastconfig.GetString("session.master-name", ""),
		ClientName:       fastconfig.GetString("session.client-name", ""),
		PoolSize:         fastconfig.GetInt("session.pool-size", 150),
		PoolTimeout:      fastconfig.GetInt("session.pool-timeout", 10),
		MinIdleConns:     fastconfig.GetInt("session.min-idle-conns", 10),
		MaxIdleConns:     fastconfig.GetInt("session.max-idle-conns", 20),
		ConnMaxIdleTime:  fastconfig.GetInt("session.conn-max-idle-time", 30),
		ConnMaxLifetime:  fastconfig.GetInt("session.conn-max-lifetime", 240),
	})
}
