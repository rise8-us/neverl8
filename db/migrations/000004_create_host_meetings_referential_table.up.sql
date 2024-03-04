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