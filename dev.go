// Package authdev adds a developer login to togo auth: one click issues an admin
// session — no credentials. It is DISABLED in production (APP_ENV=production) and
// is intended for local development only. Install: `togo install togo-framework/auth-dev`.
package authdev

import (
	"net/http"
	"os"
	"strings"

	"github.com/togo-framework/auth"
	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("auth-dev", togo.PriorityLate+25, func(k *togo.Kernel) error {
		if isProduction() {
			return nil // never in production
		}
		svc, ok := auth.FromKernel(k)
		if !ok {
			return nil
		}
		k.Router.Post("/api/auth/dev/login", func(w http.ResponseWriter, r *http.Request) {
			id, err := svc.FindOrCreateByEmail(r.Context(), devEmail())
			if err != nil {
				http.Error(w, "dev login failed", http.StatusInternalServerError)
				return
			}
			id.Roles = []string{"admin"}
			id.Permissions = []string{"*"}
			if _, err := svc.IssueSession(w, *id); err != nil {
				http.Error(w, "session failed", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"status":"ok"}`))
		})
		auth.RegisterLoginMethod(auth.LoginMethod{
			Name: "dev", Label: "Login as developer", Type: "dev", URL: "/api/auth/dev/login",
		})
		if k.Log != nil {
			k.Log.Warn("auth-dev: developer login ENABLED (non-production only)")
		}
		return nil
	})
}

func devEmail() string {
	if e := os.Getenv("DEV_LOGIN_EMAIL"); e != "" {
		return e
	}
	return "dev@togo.local"
}

func isProduction() bool {
	for _, k := range []string{"APP_ENV", "ENV", "TOGO_ENV"} {
		switch strings.ToLower(os.Getenv(k)) {
		case "production", "prod":
			return true
		}
	}
	return false
}
