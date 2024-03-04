CREATE TABLE IF NOT EXISTS time_preferences (
  id SERIAL PRIMARY KEY,
  host_id INTEGER REFERENCES hosts(id),
  start_window TIMESTAMP,
  end_window TIMESTAMP -- TODO: consider using a duration instead of an end time
);
