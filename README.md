<!-- togo-brand -->
<p align="center">
  <img src=".github/assets/togo-mark.svg" width="96" alt="togo" />
</p>
<h1 align="center">auth-dev</h1>
<p align="center"><sub>part of the <a href="https://github.com/togo-framework">togo-framework</a> — the full-stack Go + React framework</sub></p>

Developer login for togo [auth](https://github.com/togo-framework/auth) — a
one-click "Login as developer" that issues an admin session. **Disabled in
production** (APP_ENV=production). For local development only.

```bash
togo install togo-framework/auth-dev
```

Env: `DEV_LOGIN_EMAIL` (default dev@togo.local). Adds a "dev" login method and
`POST /api/auth/dev/login`.


---

## 💎 Premium sponsors

togo is proudly sponsored by **ID8 Media** and **One Studio**.

<p align="center">
  <a href="https://id8media.com"><img src=".github/assets/id8media.svg" height="44" alt="ID8 Media" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://one-studio.co"><img src=".github/assets/one-studio.jpeg" height="44" alt="One Studio" /></a>
</p>
