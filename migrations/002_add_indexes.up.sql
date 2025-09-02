-- Unique provider account
CREATE UNIQUE INDEX ux_provider_external ON provider_accounts(provider, external_id);

-- Projects
CREATE INDEX ix_projects_provider ON projects(provider);
CREATE UNIQUE INDEX ux_project_external ON projects(provider_account_id, external_id);

-- Costs
CREATE UNIQUE INDEX ux_costs_unique_day ON costs(project_id, service, region, date);
CREATE INDEX ix_costs_date ON costs(date);
CREATE INDEX ix_costs_project ON costs(project_id);
