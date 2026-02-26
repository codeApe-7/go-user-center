# Go User Center (Learning Project)

A small **user center** web app for learning Go.

- Backend: Go + Gin + GORM + MySQL + JWT
- Frontend: Vue 3 + Vite + TypeScript + Element Plus

## Features

- Register / Login
- View & update profile (nickname / avatar / bio)
- Change password

## Quick Start (Docker)

```bash
cd go-user-center
docker compose up -d --build
```

- Frontend: http://127.0.0.1:3301
- Backend: http://127.0.0.1:18601

## Local Dev

### Backend

```bash
cd backend
cp .env.example .env
# edit DB_DSN if needed
# make sure you have Go in PATH

go mod tidy
DB_DSN="uc_user:uc_pass@tcp(127.0.0.1:13306)/go_user_center?charset=utf8mb4&parseTime=True&loc=Local" \
JWT_SECRET="dev-secret-change-me" \
SERVER_ADDR=":18601" \
go run ./cmd/server
```

### Frontend

```bash
cd frontend
npm install
# Vite dev server
VITE_API_BASE="http://127.0.0.1:18601/api" npm run dev -- --port 3301
```

## API

- `POST /api/auth/register` { email, password, nickname }
- `POST /api/auth/login` { email, password } -> { accessToken, user }
- `GET /api/me` (Bearer token)
- `PUT /api/me` (Bearer token)
- `PUT /api/me/password` (Bearer token)

## Notes

This repo is intentionally simple but structured like a real project.
