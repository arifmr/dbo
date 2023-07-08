CREATE TABLE customers (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  address TEXT,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  customer_id BIGINT NOT NULL,
  product_name VARCHAR(255) NOT NULL,
  quantity BIGINT NOT NULL,
  total_price DOUBLE PRECISION NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  FOREIGN KEY (customer_id) REFERENCES customers (id)
);

CREATE TABLE login_data (
  id SERIAL PRIMARY KEY,
  customer_id BIGINT NOT NULL,
  login_time TIMESTAMPTZ DEFAULT NOW(),
  ip_address VARCHAR(45),
  token VARCHAR(255) NOT NULL,
  FOREIGN KEY (customer_id) REFERENCES customers (id)
);
