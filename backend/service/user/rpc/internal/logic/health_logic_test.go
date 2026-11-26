package logic

import (
	"context"
	"testing"

	"backend/service/user/rpc/internal/config"
	"backend/service/user/rpc/internal/svc"
	"backend/service/user/rpc/user"
)

func TestHealth(t *testing.T) {
	c := config.Config{}
	c.Name = "user.rpc"

	logic := NewHealthLogic(context.Background(), svc.NewServiceContext(c))
	resp, err := logic.Health(&user.HealthRequest{})
	if err != nil {
		t.Fatalf("Health() error = %v", err)
	}
	if resp.Service != "user.rpc" {
		t.Fatalf("Health() service = %q, want %q", resp.Service, "user.rpc")
	}
	if resp.Status != "ok" {
		t.Fatalf("Health() status = %q, want %q", resp.Status, "ok")
	}
}
