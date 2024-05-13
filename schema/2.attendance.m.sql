-- Copyright (©) 2024 - Shivam Kumar Jha - All Rights Reserved, Proprietary and confidential
-- Unauthorised copying of this file, via any medium is strictly prohibited

CREATE TABLE attendance (
  studentid INTEGER REFERENCES student (id),
  date      INTEGER,
  sectionid SMALLINT NOT NULL REFERENCES section (id),
  isabsent  BOOLEAN,
  ishalfday BOOLEAN,
  islate    BOOLEAN,
  PRIMARY KEY (studentid, date)
);

CREATE TABLE offday (
  date        INTEGER,
  classid     SMALLINT REFERENCES class (id),
  sectionid   SMALLINT REFERENCES section (id),
  isholiday   BOOLEAN,
  isweekend   BOOLEAN,
  ishalfday   BOOLEAN,
  description TEXT CHECK (LENGTH(description) <= 100)
);
