package synapse

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/kokizzu/gotro/L"
	"github.com/rizalgowandy/synapse-go/pkg/api"
	"github.com/rizalgowandy/synapse-go/pkg/entity"
	"github.com/stretchr/testify/assert"
)

// How to run all integration test:
// $ ID=REAL_API_ID KEY=REAL_API_KEY go test -v . -run . -integration

var (
	integration bool
	client      *Client
)

func TestMain(m *testing.M) {
	flag.BoolVar(&integration, "integration", false, "enable integration test")
	flag.Parse()

	if !integration {
		os.Exit(m.Run())
	}

	var err error
	client, err = NewClient(&api.Config{
		ID:      os.Getenv("ID"),
		Key:     os.Getenv("KEY"),
		Debug:   true,
		HostURL: "https://uat-api.synapsefi.com",
	})
	if L.IsError(err, "client: create failure") {
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestClient_GenerateRefreshToken_Integration(t *testing.T) {
	if !integration {
		return
	}

	type args struct {
		ctx context.Context
		req *entity.GenerateRefreshTokenReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				req: &entity.GenerateRefreshTokenReq{
					UserID:        "5fd7acfd8677040053ad486d",
					UserIPAddress: "255.127.79.76",
					Email:         "test@synapsepay.com",
					Password:      "test1234",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client
			got, err := c.GenerateRefreshToken(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			L.Describe(got, err)
			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestClient_AllowedEntityScopes_Integration(t *testing.T) {
	if !integration {
		return
	}

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client
			got, err := c.AllowedEntityScopes(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllowedEntityScopes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			L.Describe(got, err)
			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestClient_AllowedEntityTypes_Integration(t *testing.T) {
	if !integration {
		return
	}

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client
			got, err := c.AllowedEntityTypes(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllowedEntityTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			L.Describe(got, err)
			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestClient_AllowedDocumentTypes_Integration(t *testing.T) {
	if !integration {
		return
	}

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client
			got, err := c.AllowedDocumentTypes(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllowedDocumentTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			L.Describe(got, err)
			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}
