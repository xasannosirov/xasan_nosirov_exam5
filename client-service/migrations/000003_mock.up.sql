INSERT INTO clients (id, first_name, last_name, age, gender, phone_number, address, email, password, status, refresh)
VALUES
('a5576fd3-e89b-4c01-a357-3f537b72abc2', 'Xasan', 'Nosirov', 18, 'Male', '+998944970514', 'Tashkent/Uzbekistan', 'xasannosirov094@gmail.com', 'Secret12345', true, 'secret.refresh.token');

INSERT INTO jobs (id, name, salary, level, location_type, employment_type, address, company)
VALUES
('a5576fd3-e89b-4c01-a357-3f537b72abc3', 'Back-End Developer', 400, 'Junior', 'On-site', 'Full-Time', 'Tashkent/Uzbekistan', 'Netflix');

INSERT INTO client_jobs (client_id, job_id, start_date)
VALUES
('a5576fd3-e89b-4c01-a357-3f537b72abc2', 'a5576fd3-e89b-4c01-a357-3f537b72abc3', CURRENT_TIMESTAMP);
