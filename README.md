# Internal Audit Evaluation System API

The **Internal Audit Evaluation System API** is a backend service built with [Go (Golang)](https://golang.org/) and the [Gin Web Framework](https://github.com/gin-gonic/gin), using [GORM](https://gorm.io/) as the ORM and PostgreSQL for data storage.

It provides a RESTful API for managing internal audit performance evaluations for roles such as Lead Auditors, Team Leaders, and Internal Auditors. The system supports structured evaluations, standardized criteria, and historical tracking.

---

## 📌 Features

- RESTful API for internal audit evaluations
- Role-based assessment system with evaluator/evaluatee logic
- Standard 4-point rating scale (1 = Strongly Disagree to 4 = Strongly Agree)
- Evaluation criteria and scoring per period
- PostgreSQL support via GORM
- Modular, production-ready Go codebase

---

## 🧱 Tech Stack

- **Language**: Go
- **Framework**: [Gin](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: PostgreSQL
- **Env Config**: `.env` using [`godotenv`](https://github.com/joho/godotenv)

---

## 📂 API Endpoints (Examples)

| Method | Endpoint               | Description                      |
|--------|------------------------|----------------------------------|
| GET    | `/api/users`           | List all users                   |
| POST   | `/api/evaluations`     | Create a new evaluation          |
| GET    | `/api/evaluations/:id` | Get evaluation details           |
| GET    | `/api/criteria`        | List evaluation criteria         |
| POST   | `/api/responses`       | Submit evaluation responses      |

---

## 🗃️ Database Schema Overview

Core Tables:

- `users`: Stores all system users (auditors, evaluators)
- `evaluations`: Links evaluator, evaluatee, and metadata
- `criteria`: Predefined questions per role
- `evaluation_responses`: Scores linked to evaluation sessions

---

## 🚀 Getting Started

### 1. Clone the Repository

```bash
    git clone git@github.com:john-henry-alfeche/internal-audit-evaluation-system-api.git
    cd internal-audit-evaluation-system-api
