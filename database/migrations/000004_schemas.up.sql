ALTER TABLE waste_collection
 ALTER COLUMN amount TYPE INTEGER DEFAULT 1;

INSERT INTO waste_types (waste_id, waste_types) VALUES (DEFAULT, "PET");
INSERT INTO waste_types (waste_id, waste_types) VALUES (DEFAULT, "CAN");