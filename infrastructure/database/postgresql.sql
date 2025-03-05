CREATE DATABASE banking_services;

USE banking_services;

CREATE TABLE customers (
    id VARCHAR(50) PRIMARY KEY DEFAULT,
    name VARCHAR(100) NOT NULL,
    nik VARCHAR(20) UNIQUE NOT NULL,
    phone_number VARCHAR(15) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE accounts (
    id VARCHAR(50) PRIMARY KEY DEFAULT,
    customer_id VARCHAR(50) NOT NULL,
    account_number VARCHAR(20) UNIQUE NOT NULL,
    balance DECIMAL(15,2) DEFAULT 0.00 NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE
);


CREATE TABLE transactions (
    id VARCHAR(50) PRIMARY KEY DEFAULT,
    account_id VARCHAR(50) NOT NULL,
    amount DECIMAL(15,2) NOT NULL,
    transaction_type VARCHAR(10) CHECK (transaction_type IN ('deposit', 'withdraw')),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE
);


CREATE INDEX idx_customers_nik ON customers(nik);
CREATE INDEX idx_customers_phone ON customers(phone_number);
CREATE INDEX idx_accounts_number ON accounts(account_number);
CREATE INDEX idx_accounts_customer_id ON customers(customer_id);
CREATE INDEX idx_transactions_account_id ON transactions(account_id);

