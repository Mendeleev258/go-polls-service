DROP INDEX IF EXISTS idx_votes_poll_user;
DROP INDEX IF EXISTS idx_polls_active_created;
DROP INDEX IF EXISTS idx_polls_author;

DROP TABLE IF EXISTS votes;
DROP TABLE IF EXISTS options;
DROP TABLE IF EXISTS polls;
DROP TABLE IF EXISTS users;