ALTER TABLE waste_collection
    ALTER COLUMN amount TYPE INTEGER,
    ALTER COLUMN amount SET DEFAULT 0;

INSERT INTO waste_types (waste_id, waste_type) VALUES (1, 'PET');
INSERT INTO waste_types (waste_id, waste_type) VALUES (2, 'CANS');