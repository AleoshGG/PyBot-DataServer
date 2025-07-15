ALTER TABLE waste_collection
 ALTER COLUMN amount TYPE INTEGER DEFAULT 0;

INSERT INTO waste_types (waste_id, waste_type) VALUES (DEFAULT, 'PET');
INSERT INTO waste_types (waste_id, waste_type) VALUES (DEFAULT, 'CAN');