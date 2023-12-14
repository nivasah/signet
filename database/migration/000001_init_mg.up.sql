CREATE TABLE users
(
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  email VARCHAR(50) NOT NULL,
  password VARCHAR(50) NOT NULL
);

CREATE TABLE directories
(
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  description VARCHAR(500),
  user_id BIGINT NOT NULL
);

ALTER TABLE directories ADD FOREIGN KEY(user_id) REFERENCES users(id);

CREATE TABLE resources
(
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  location VARCHAR(200) NOT NULL,
  description VARCHAR(500),
  directory_id BIGINT NOT NULL
);

ALTER TABLE resources ADD FOREIGN KEY(directory_id) REFERENCES directories(id);
