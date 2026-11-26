package health

import (
	"context"
	"testing"

	"backend/service/user/api/internal/config"
	"backend/service/user/api/internal/svc"
)

func TestHealth(t *testing.T) {
	c := config.Config{}
	c.Name = "user-api"

	logic := NewHealthLogic(context.Background(), svc.NewServiceContext(c))
	resp, err := logic.Health()
	if err != nil {
		t.Fatalf("Health() error = %v", err)
	}
	if resp.Service != "user-api" {
		t.Fatalf("Health() service = %q, want %q", resp.Service, "user-api")
	}
	if resp.Status != "ok" {
		t.Fatalf("Health() status = %q, want %q", resp.Status, "ok")
	}
}
