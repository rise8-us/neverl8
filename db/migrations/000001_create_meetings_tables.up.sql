CREATE TABLE IF NOT EXISTS meetings (
    id SERIAL PRIMARY KEY,
    candidate_id INTEGER NOT NULL,
    calendar TEXT NOT NULL,
    duration INTEGER NOT NULL,
    title TEXT NOT NULL,
    description TEXT DEFAULT 'no description provided',
    has_bot_guest BOOLEAN DEFAULT FALSE,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS hosts (
    id SERIAL PRIMARY KEY,
    host_name TEXT NOT NULL,
    last_meeting_time TIMESTAMP DEFAULT '1970-01-01 00:00:00'
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
  start_window TEXT DEFAULT '00:00',
  end_window TEXT DEFAULT '00:00' -- TODO: consider using a duration instead of an end time
);

CREATE TABLE IF NOT EXISTS calendars (
  id SERIAL PRIMARY KEY,
  google_calendar_id TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS host_calendars (
  calendar_id INTEGER NOT NULL,
  host_id INTEGER NOT NULL,
  PRIMARY KEY (calendar_id, host_id),
  CONSTRAINT fk_calendar
    FOREIGN KEY (calendar_id)
    REFERENCES calendars(id)
    ON DELETE CASCADE,
  CONSTRAINT fk_host
    FOREIGN KEY (host_id)
    REFERENCES hosts(id)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS sample_meetings (
  id SERIAL PRIMARY KEY, 
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  duration INTEGER NOT NULL
);