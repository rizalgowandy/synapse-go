package api

import (
	"crypto/sha256"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/segmentio/ksuid"
)

// Reference: https://docs.synapsefi.com/intro-to-apis

func NewClient(cfg *Config) *Client {
	return &Client{
		cfg: cfg,
		client: resty.New().
			SetHostURL(cfg.HostURL).
			SetHeaders(DefaultHeaders(cfg.ID, cfg.Key)).
			SetTimeout(cfg.Timeout).
			SetRetryCount(cfg.RetryCount).
			SetRetryMaxWaitTime(cfg.RetryMaxWaitTime).
			SetDebug(cfg.Debug),
	}
}

type Client struct {
	cfg    *Config
	client *resty.Client
}

func (c *Client) UserHeader(userOAuthKey, userID string) map[string]string {
	return map[string]string{
		"X-SP-USER": fmt.Sprintf("%s|%s", userOAuthKey, c.GenerateUserDeviceFingerprint(c.cfg.ID, c.cfg.Key, userID)),
	}
}

func (c *Client) UserIPHeader(ip string) map[string]string {
	return map[string]string{
		"X-SP-USER-IP": ip,
	}
}

func (c *Client) IdempotencyHeader() map[string]string {
	return map[string]string{
		"X-SP-IDEMPOTENCY-KEY": ksuid.New().String(),
	}
}

func (c *Client) GenerateUserDeviceFingerprint(id, key, userID string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%s_%s_%s", userID, id, key))))
}

func DefaultHeaders(id, key string) map[string]string {
	return map[string]string{
		"X-SP-GATEWAY": fmt.Sprintf("%s|%s", id, key),
	}
}
