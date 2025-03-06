DROP DATABASE gobankdb;
CREATE DATABASE gobankdb;

CREATE TYPE account_status_enum AS ENUM ('active', 'inactive');
CREATE TYPE transaction_type_enum AS ENUM ('deposit', 'withdraw', 'none');

CREATE TABLE customers (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    nik VARCHAR(20) UNIQUE NOT NULL,
    phone_number VARCHAR(15) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE accounts (
    id VARCHAR(50) PRIMARY KEY,
    customer_id VARCHAR(50) NOT NULL UNIQUE,
    account_number VARCHAR(20) UNIQUE NOT NULL,
    balance DECIMAL(15,2) DEFAULT 0.00 NOT NULL,
    status account_status_enum NOT NULL DEFAULT 'inactive', 
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE
);

CREATE TABLE transactions (
    id VARCHAR(50) PRIMARY KEY,
    account_id VARCHAR(50) NOT NULL,
    account_number VARCHAR(20) NOT NULL,
    amount DECIMAL(15,2) NOT NULL,
    transaction_type transaction_type_enum NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE
);


CREATE INDEX idx_customers_nik ON customers(nik);
CREATE INDEX idx_customers_phone ON customers(phone_number);
CREATE INDEX idx_accounts_number ON accounts(account_number);
CREATE INDEX idx_accounts_customer_id ON accounts(customer_id);
CREATE INDEX idx_transactions_account_id ON transactions(account_id);

