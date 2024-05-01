CREATE TABLE jobs(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    salary FLOAT,
    level VARCHAR(10) NOT NULL, -- Junior, Middle, Senior, Staff, ...
    location_type VARCHAR(7) NOT NULL, -- On-site, Remote, Hybrid, ...
    employment_type VARCHAR(10) NOT NULL, -- Full-Time, Part-Time, Internship, Temporary, Contract, ...
    address TEXT NOT NULL,
    company TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE client_jobs(
    client_id UUID NOT NULL,
    job_id UUID NOT NULL,
    start_date TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    end_date TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    FOREIGN KEY (job_id) REFERENCES jobs(id)
);
