CREATE TABLE IF NOT EXISTS work_periods (
    period_id    SERIAL PRIMARY KEY,
    start_hour   TIMESTAMPTZ NOT NULL,
    end_hour     TIMESTAMPTZ NOT NULL,
    day_work     VARCHAR(15) NOT NULL,
    prototype_id VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS readings (
    period_id         SERIAL PRIMARY KEY,
    distance_traveled DECIMAL(10,4) NOT NULL,
    weight_waste      DECIMAL(10,4) NOT NULL,
    FOREIGN KEY (period_id) REFERENCES work_periods(period_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS waste_types (
    waste_id   SERIAL PRIMARY KEY,
    waste_type VARCHAR(50) NOT NULL
); 

CREATE TABLE IF NOT EXISTS waste_collection (
    waste_collection_id     SERIAL PRIMARY KEY,
    period_id   INTEGER NOT NULL,
    amount      INTEGER NOT NULL,
    waste_id    INTEGER NOT NULL,
    FOREIGN KEY (waste_id) REFERENCES waste_types(waste_id) ON DELETE CASCADE,
    FOREIGN KEY (period_id) REFERENCES readings(period_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS weight_data (
    weight_data_id     SERIAL PRIMARY KEY,
    period_id   INTEGER NOT NULL,
    hour_period TIMESTAMPTZ,
    weight      DECIMAL(10,4),
    FOREIGN KEY (period_id) REFERENCES readings(period_id) ON DELETE CASCADE 
);

CREATE TABLE IF NOT EXISTS gps_data (
    gps_data_id     SERIAL PRIMARY KEY,
    period_id   INTEGER NOT NULL,
    latitude    DECIMAL(10,10),
    longitude   DECIMAL(10,10),
    altitude    DECIMAL(10,10),
    speed       DECIMAL(10,10),
    date_gps    DATE,
    hour_UTC    TIMESTAMPTZ,
    FOREIGN KEY (period_id) REFERENCES readings(period_id) ON DELETE CASCADE
);