# Cloud Cost Tracker API

A lightweight and extensible API to **track, analyze, and manage cloud costs** across multiple cloud providers.

---

## 🚀 Features

- **Project Management**

  - Register and manage tracked GCP projects
  - Link projects with their GCP Project IDs

- **Cost Tracking**

  - Query daily costs per project
  - Filter by service, region, and date range
  - View aggregated cost summaries (group by project, service, or region)
  - Analyze cost trends over time (for charts and forecasting)

- **Budgets (Optional Extension)**

  - Set budgets per project
  - Track budget usage vs. actual spending
  - Trigger alerts when spending exceeds the budget

- **Reports**

  - Generate monthly and weekly reports
  - Export reports in JSON or CSV format
  - Provide summary breakdowns for finance teams

- **Authentication & Security**

  - JWT-based authentication for API access
  - Role-based access control (future roadmap)

- **Scalability (Future)**
  - Redis caching for frequent queries
  - Multi-cloud support (AWS, Azure)

---

## 🛠️ Tech Stack

- **Language:** Golang
- **Framework:** Gin (HTTP server)
- **ORM:** GORM
- **Database:** PostgreSQL
- **Cache (optional):** Redis
- **Authentication:** JWT
- **Cloud Integration:** GCP Billing Export (BigQuery)

---

## 📌 Roadmap

- [ ] **MVP (PoC Stage)**

  - [ ] Setup project structure with Gin + Clean Architecture
  - [ ] Implement PostgreSQL schema (projects, costs, budgets)
  - [ ] Add basic API endpoints (projects, costs, budgets)
  - [ ] Provide JWT authentication for secured routes

- [ ] **Phase 2: GCP Integration**

  - [ ] Integrate with GCP Billing Export (BigQuery)
  - [ ] Build background job to sync cost data daily
  - [ ] Add filtering support (service, region, project, date range)
  - [ ] Implement Redis caching for frequent queries

- [ ] **Phase 3: Reporting & Visualization**

  - [ ] Implement cost summaries and trend analysis
  - [ ] Export reports in CSV/JSON format
  - [ ] Integrate with Grafana dashboards for visualization
  - [ ] Add API for monthly/weekly reports

- [ ] **Phase 4: Budgets & Alerts**

  - [ ] Add budget tracking per project
  - [ ] Implement alerting when usage exceeds budget
  - [ ] Integrate with Slack/Email notifications

- [ ] **Future Enhancements**
  - [ ] Support multi-cloud (AWS, Azure)
  - [ ] Add RBAC for user roles (Admin, Viewer, Finance)
  - [ ] Build web UI frontend (React/Next.js)
  - [ ] Add cost forecasting using ML models
