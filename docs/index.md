# auth-dev

Developer login for togo [auth](https://github.com/togo-framework/auth) — a
one-click "Login as developer" that issues an admin session. **Disabled in
production** (APP_ENV=production). For local development only.

```bash
togo install togo-framework/auth-dev
```

Env: `DEV_LOGIN_EMAIL` (default dev@togo.local). Adds a "dev" login method and
`POST /api/auth/dev/login`.
