INSERT INTO clients (id, first_name, last_name, age, gender, phone_number, address, email, password, status, refresh, created_at, updated_at, deleted_at)
VALUES
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'John', 'Doe', 30, 'Male', '1234567890', '123 Main St, City, Country', 'john.doe@example.com', 'password123', TRUE, 'refresh_token_1', '2024-05-01 08:00:00', '2024-05-01 08:00:00', NULL),
    ('c20ad4d7-90b9-44f9-8c7e-b02c40bcdbb4', 'Jane', 'Smith', 25, 'Female', '0987654321', '456 Elm St, City, Country', 'jane.smith@example.com', 'pass1234', TRUE, 'refresh_token_2', '2024-05-01 09:00:00', '2024-05-01 09:00:00', NULL),
    ('c51ce410c124a10e0db5e4b97fc2af39', 'Alice', 'Johnson', 35, 'Female', '1112223333', '789 Oak St, City, Country', 'alice.johnson@example.com', 'abcde12345', TRUE, 'refresh_token_3', '2024-05-01 10:00:00', '2024-05-01 10:00:00', NULL),
    ('c74d97b01eae257e44aa9d5bade97baf', 'Michael', 'Williams', 28, 'Male', '5556667777', '101 Pine St, City, Country', 'michael.williams@example.com', 'qwerty9876', TRUE, 'refresh_token_4', '2024-05-01 11:00:00', '2024-05-01 11:00:00', NULL),
    ('70efdf2ec9b086079795c442636b55fb', 'Emily', 'Brown', 32, 'Female', '9998887777', '202 Maple St, City, Country', 'emily.brown@example.com', 'ilovecats123', TRUE, 'refresh_token_5', '2024-05-01 12:00:00', '2024-05-01 12:00:00', NULL),
    ('8e296a067a37563370ded05f5a3bf3ec', 'Daniel', 'Martinez', 40, 'Male', '3332221111', '303 Cedar St, City, Country', 'daniel.martinez@example.com', 'securepass789', TRUE, 'refresh_token_6', '2024-05-01 13:00:00', '2024-05-01 13:00:00', NULL),
    ('e4da3b7fbbce2345d7772b0674a318d5', 'Sarah', 'Jones', 27, 'Female', '7778889999', '505 Walnut St, City, Country', 'sarah.jones@example.com', 'password4321', TRUE, 'refresh_token_7', '2024-05-01 14:00:00', '2024-05-01 14:00:00', NULL),
    ('1679091c5a880faf6fb5e6087eb1b2dc', 'David', 'Taylor', 45, 'Male', '6665554444', '707 Oak St, City, Country', 'david.taylor@example.com', 'passpass123', TRUE, 'refresh_token_8', '2024-05-01 15:00:00', '2024-05-01 15:00:00', NULL),
    ('8f14e45fceea167a5a36dedd4bea2543', 'Jessica', 'Lee', 22, 'Female', '2223334444', '909 Maple St, City, Country', 'jessica.lee@example.com', 'passpass456', TRUE, 'refresh_token_9', '2024-05-01 16:00:00', '2024-05-01 16:00:00', NULL),
    ('c9f0f895fb98ab9159f51fd0297e236d', 'Matthew', 'Davis', 38, 'Male', '4445556666', '404 Pine St, City, Country', 'matthew.davis@example.com', 'mysecurepass', TRUE, 'refresh_token_10', '2024-05-01 17:00:00', '2024-05-01 17:00:00', NULL);

INSERT INTO jobs (id, name, salary, level, location_type, employment_type, address, company)
VALUES
    ('f86cf9e4-6e28-4b3c-b1b5-5426570d9d2a', 'Software Developer', 85000.00, 'Senior', 'Hybrid', 'Full-Time', '123 Main St, Anytown, USA', 'Tech Innovations Inc.'),
    ('73efc46b-1b11-4b23-b0a7-63b02a226b7e', 'Marketing Manager', 75000.00, 'Senior', 'On-site', 'Full-Time', '456 Elm St, Otherville, USA', 'Brand Boosters LLC'),
    ('e8897b9b-0728-4e3c-83cc-af032b059c39', 'Data Scientist', 90000.00, 'Senior', 'Remote', 'Full-Time', '789 Oak St, Somewhere, USA', 'Data Insights Co.'),
    ('21b7c24d-8a5f-4b6a-94e3-4b9e5c52f362', 'Web Developer', 80000.00, 'Middle', 'Hybrid', 'Full-Time', '101 Pine St, Nowhere, USA', 'WebSolutions Inc.'),
    ('a7d17e0d-0d5e-4a89-a8d3-9fe510b3fd1b', 'Graphic Designer', 60000.00, 'Junior', 'On-site', 'Full-Time', '246 Maple St, Here, USA', 'DesignWorks LLC'),
    ('b476745e-2460-44d1-9a2a-f8bbdc1f72c8', 'Financial Analyst', 85000.00, 'Senior', 'Hybrid', 'Full-Time', '369 Cherry St, There, USA', 'FinanceGenius Inc.'),
    ('c536b2f0-91fd-48e1-bf46-d036fb60a7d7', 'Project Manager', 90000.00, 'Senior', 'On-site', 'Full-Time', '789 Pine St, Everywhere, USA', 'Project Management Co.'),
    ('d279cd20-1c14-4d10-9a26-9c3e8f4b51ac', 'Customer Support Specialist', 55000.00, 'Staff', 'Remote', 'Full-Time', '963 Elm St, Anywhere, USA', 'Support Solutions LLC'),
    ('e5db5f17-5bbd-4e08-846a-f537a1067b9d', 'Sales Representative', 70000.00, 'Middle', 'Hybrid', 'Full-Time', '852 Oak St, Everywhere, USA', 'Sales Solutions Inc.'),
    ('f8e90531-3c8b-4565-b35c-7271c8752d78', 'Human Resources Coordinator', 65000.00, 'Middle', 'On-site', 'Full-Time', '741 Maple St, Nowhere, USA', 'HR Management Co.');

INSERT INTO client_jobs (client_id, job_id, start_date)
VALUES
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'f86cf9e4-6e28-4b3c-b1b5-5426570d9d2a', CURRENT_TIMESTAMP),
    ('c20ad4d7-90b9-44f9-8c7e-b02c40bcdbb4', '73efc46b-1b11-4b23-b0a7-63b02a226b7e', CURRENT_TIMESTAMP),
    ('c51ce410c124a10e0db5e4b97fc2af39', 'e8897b9b-0728-4e3c-83cc-af032b059c39', CURRENT_TIMESTAMP),
    ('c74d97b01eae257e44aa9d5bade97baf', '21b7c24d-8a5f-4b6a-94e3-4b9e5c52f362', CURRENT_TIMESTAMP),
    ('70efdf2ec9b086079795c442636b55fb', 'a7d17e0d-0d5e-4a89-a8d3-9fe510b3fd1b', CURRENT_TIMESTAMP),
    ('8e296a067a37563370ded05f5a3bf3ec', 'b476745e-2460-44d1-9a2a-f8bbdc1f72c8', CURRENT_TIMESTAMP),
    ('e4da3b7fbbce2345d7772b0674a318d5', 'c536b2f0-91fd-48e1-bf46-d036fb60a7d7', CURRENT_TIMESTAMP),
    ('1679091c5a880faf6fb5e6087eb1b2dc', 'd279cd20-1c14-4d10-9a26-9c3e8f4b51ac', CURRENT_TIMESTAMP),
    ('8f14e45fceea167a5a36dedd4bea2543', 'e5db5f17-5bbd-4e08-846a-f537a1067b9d', CURRENT_TIMESTAMP),
    ('c9f0f895fb98ab9159f51fd0297e236d', 'f8e90531-3c8b-4565-b35c-7271c8752d78', CURRENT_TIMESTAMP);