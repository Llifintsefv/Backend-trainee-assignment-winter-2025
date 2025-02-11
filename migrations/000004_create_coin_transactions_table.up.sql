
CREATE TABLE coin_transactions (
    id UUID PRIMARY KEY,
    sender_user_id UUID,
    receiver_user_id UUID NOT NULL,
    amount INTEGER NOT NULL,
    transaction_type VARCHAR(20) NOT NULL, -- 'send', 'receive', 'purchase'
    transaction_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sender_user_id) REFERENCES users(id),
    FOREIGN KEY (receiver_user_id) REFERENCES users(id)
);