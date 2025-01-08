DROP TABLE IF EXISTS statuses;
DROP TABLE IF EXISTS applications;

CREATE TABLE IF NOT EXISTS statuses (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  status TEXT NOT NULL
);

INSERT INTO statuses(status) VALUES
  ("Did not apply"),
  ("Applied"),
  ("Rejected"),
  ("Interviewed"),
  ("Offered"),
  ("Offer accepted");

CREATE TABLE IF NOT EXISTS applications (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  company TEXT NOT NULL,
  position TEXT NOT NULL,
  location TEXT,
  date_posted TEXT,
  date_applied TEXT,
  url TEXT,
  notes TEXT,
  status INTEGER REFERENCES statuses(id) NOT NULL
);
