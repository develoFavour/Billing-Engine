-- Create Customers Table
CREATE TABLE IF NOT EXISTS customers (
    id UUID PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create Pricing Tiers Table
CREATE TABLE IF NOT EXISTS pricing_tiers (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    resource_type VARCHAR(50) NOT NULL,
    model VARCHAR(20) NOT NULL, -- 'flat', 'tiered'
    unit_price DECIMAL(12, 4),
    tiers JSONB, -- For tiered pricing logic
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Associate Customers with Pricing Tiers
ALTER TABLE customers ADD COLUMN pricing_tier_id UUID REFERENCES pricing_tiers(id);

-- Create Usage Records Table (The High-Volume Table)
CREATE TABLE IF NOT EXISTS usage_records (
    id UUID PRIMARY KEY,
    customer_id UUID REFERENCES customers(id) ON DELETE CASCADE,
    resource_type VARCHAR(50) NOT NULL,
    quantity DECIMAL(15, 4) NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for performance
CREATE INDEX idx_usage_customer_time ON usage_records(customer_id, timestamp);
CREATE INDEX idx_usage_resource_type ON usage_records(resource_type);
