# URL Shortener (Go + Redis)

A high-performance URL shortening service built with **Go (Fiber)**, **Redis (Upstash)**, and **React (Vite)**.

**Live Demo:** https://url-shortener-go-deployed.vercel.app/

---

## Build Instructions & Source Code

For local development setup, Docker instructions, and complete source code, please visit the main repository:

**https://github.com/SKjustSK/url-shortener-go**

---

## Tech Stack

* **Frontend:** React + Vite + Tailwind CSS (Hosted on Vercel)
* **Backend:** Go (Golang) + Fiber (Hosted on Render)
* **Database:** Redis (Upstash Serverless)
* **Infrastructure:** Docker & Docker Compose

---

## API Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `POST` | `/api/shorten` | JSON body `{ "url": "https://google.com", "short": "goog", "expiry": 24 }` |
| `GET` | `/:url` | Redirects to the original long URL. |

---

## Deployment Configuration

### Backend (Render)
The Go backend is deployed as a Web Service on **Render**.

**Environment Variables:**
* `API_QUOTA`: `10` (Rate limit per IP)
* `APP_PORT`: `3000` (Internal port)
* `DOMAIN`: `url-shortener-go-deployed.onrender.com` (Used for generating short links)
* `FRONTEND_DOMAIN`: `https://url-shortener-go-deployed.vercel.app` (Allowed CORS origin)
* `REDIS_URL`: `rediss://default:password@...` (Upstash connection string)

### Frontend (Vercel)
The React frontend is deployed on **Vercel**.

**Environment Variables:**
* `VITE_GO_SHORTEN_URL`: `https://url-shortener-go-deployed.onrender.com/api/shorten`
* `VITE_GO_DOMAIN`: `https://url-shortener-go-deployed.onrender.com`

---

## API Limits

To prevent abuse, the API implements rate limiting using Redis.
* **Default Quota:** 10 requests per user per IP.
* **Reset Period:** 30 minutes.