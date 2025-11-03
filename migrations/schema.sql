CREATE TABLE IF NOT EXISTS categories (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  created_by VARCHAR(255),
  modified_at TIMESTAMP,
  modified_by VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  created_by VARCHAR(255),
  modified_at TIMESTAMP,
  modified_by VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS books (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  image_url TEXT,
  release_year INT,
  price INT,
  total_page INT,
  thickness VARCHAR(50),
  category_id INT REFERENCES categories(id) ON DELETE SET NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  created_by VARCHAR(255),
  modified_at TIMESTAMP,
  modified_by VARCHAR(255)
);
