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

CREATE TABLE IF NOT EXISTS hosts (
    id SERIAL PRIMARY KEY,
    host_name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS candidates (
    id SERIAL PRIMARY KEY,
    candidate_name TEXT NOT NULL,
    role TEXT DEFAULT 'unknown role',
    email TEXT DEFAULT 'unknown email',
    phone_number TEXT DEFAULT 'unknown phone number',
    interview_status TEXT DEFAULT 'unknown interview status' -- TODO: interview status should be an ENUM
);

CREATE TABLE IF NOT EXISTS host_meetings (
    host_id INTEGER NOT NULL,
    meeting_id INTEGER NOT NULL,
    PRIMARY KEY (host_id, meeting_id),
    CONSTRAINT fk_host
        FOREIGN KEY (host_id)
        REFERENCES hosts(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_meeting
        FOREIGN KEY (meeting_id)
        REFERENCES meetings(id)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS time_preferences (
  id SERIAL PRIMARY KEY,
  host_id INTEGER REFERENCES hosts(id),
  start_window TIMESTAMP,
  end_window TIMESTAMP -- TODO: consider using a duration instead of an end time
);
