INSERT INTO users (phone, role) 
VALUES ('+79990001122', 'owner')
ON CONFLICT (phone) DO NOTHING;