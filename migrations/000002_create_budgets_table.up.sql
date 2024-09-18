CREATE TABLE budgets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),  
    user_id UUID NOT NULL,
    category_id UUID REFERENCES categories(id),                      
    amount BIGINT NOT NULL,                       
    spent BIGINT DEFAULT 0,                        
    currency VARCHAR(3) NOT NULL                    
);
