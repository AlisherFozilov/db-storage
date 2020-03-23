CREATE TABLE IF NOT EXISTS messages
(
    id          TEXT NOT NULL PRIMARY KEY,
    type        INT  NOT NULL,
    sender_id   INT  NOT NULL,
    receiver_id INT  NOT NULL,
    data        TEXT NOT NULL,
    timestamp   INT  NOT NULL,
    removed     BOOLEAN DEFAULT FALSE
);

INSERT INTO messages (id, type, sender_id, receiver_id, data, timestamp)
VALUES (?,?,?,?,?,?);

DELETE FROM messages;