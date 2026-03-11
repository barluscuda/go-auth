# GoAuth System by Barluscuda

## Description

An user management system built with Go, featuring authentication, authorization, and account management powered by GORM, Cache, and JWT

## Environment Variables

| Variable | Category | Description | Default |
|----------|----------|-------------|---------|
| `SERVER_MODE` | Server | Application mode — `dev` or `release` | `dev` |
| `SERVER_PORT` | Server | HTTP server listening port | `3200` |
| `DB_URL` | Database | Full database connection string e.g. `postgres://user:pass@localhost:5432/db` | **required** |
| `JWT_ACCESS_SECRET` | JWT | Secret key for signing access tokens | **required** |
| `JWT_REFRESH_SECRET` | JWT | Secret key for signing refresh tokens | **required** |
| `JWT_ACCESS_EXPIRE` | JWT | Access token expiration duration | `24h` |
| `JWT_REFRESH_EXPIRE` | JWT | Refresh token expiration duration | `720h` (30 days) |
