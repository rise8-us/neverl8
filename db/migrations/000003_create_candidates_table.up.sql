CREATE TABLE IF NOT EXISTS candidates (
    id SERIAL PRIMARY KEY,
    candidate_name TEXT NOT NULL,
    role TEXT DEFAULT 'unknown role',
    email TEXT DEFAULT 'unknown email',
    phone_number TEXT DEFAULT 'unknown phone number',
    interview_status TEXT DEFAULT 'unknown interview status' -- TODO: interview status should be an ENUM
);