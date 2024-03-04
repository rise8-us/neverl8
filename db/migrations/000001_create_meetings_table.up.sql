CREATE TABLE IF NOT EXISTS meetings (
    id SERIAL PRIMARY KEY,
    candidate_id INTEGER NOT NULL,
    calendar TEXT NOT NULL,
    duration INTEGER NOT NULL,
    title TEXT NOT NULL,
    description TEXT DEFAULT 'no description provided',
    has_bot_guest BOOLEAN DEFAULT FALSE,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
