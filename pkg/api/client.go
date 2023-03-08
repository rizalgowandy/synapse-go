package api

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/segmentio/ksuid"
)

// Reference: https://docs.synapsefi.com/intro-to-apis
func NewRestyClient(cfg Config) *resty.Client {
	return resty.New().
		SetHostURL(cfg.HostURL).
		SetHeaders(DefaultHeaders(cfg.ID, cfg.Key)).
		SetTimeout(cfg.Timeout).
		SetRetryCount(cfg.RetryCount).
		SetRetryMaxWaitTime(cfg.RetryMaxWaitTime).
		SetDebug(cfg.Debug)
}

func DefaultHeaders(id, key string) map[string]string {
	return map[string]string{
		"X-SP-GATEWAY": fmt.Sprintf("%s|%s", id, key),
	}
}

func UserHeader(userOauthKey, userDeviceFingerprint string) map[string]string {
	return map[string]string{
		"X-SP-USER": fmt.Sprintf("%s|%s", userOauthKey, userDeviceFingerprint),
	}
}

func UserIPHeader(ip string) map[string]string {
	return map[string]string{
		"X-SP-USER-IP": ip,
	}
}

func IdempotencyHeader() map[string]string {
	return map[string]string{
		"X-SP-IDEMPOTENCY-KEY": ksuid.New().String(),
	}
}
