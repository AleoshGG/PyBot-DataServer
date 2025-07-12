-- FUNCION PARA OBTENER EL LA ULTIMA HORA REGISTRADA Y TOMARLO COMO EL FINAL DEL PERIODO PASADO
CREATE OR REPLACE FUNCTION getLastHourPeriod()
RETURNS TABLE (
    period_id INTEGER,
    hour_period TIMESTAMPTZ
    ) 
    LANGUAGE SQL
    STABLE
AS $$
    SELECT period_id, hour_period 
    FROM weight_data
    ORDER BY hour_period DESC
    LIMIT 1;
$$; 

-- FUNCION PARA DEVOLVER EL ULTIMO PESO REGISTRADO DE LA TABLA
CREATE OR REPLACE FUNCTION getLastWeight()
RETURNS DECIMAL(10,4)
LANGUAGE SQL
STABLE
AS $$
    SELECT weight
    FROM weight_data
    ORDER BY hour_period DESC
    LIMIT 1;
$$;

-- FUNCION PARA CALCULAR LA DISTANCIA TOTAL RECORRIDA Y RETORNARLA
CREATE OR REPLACE FUNCTION calcular_distancia_total(p_period INTEGER)
RETURNS NUMERIC(12,4)
LANGUAGE plpgsql
STABLE
AS $$
DECLARE
  total_dist NUMERIC := 0;
BEGIN
  WITH pares AS (
    SELECT
      latitude  AS lat2,
      longitude AS lon2,
      LAG(latitude ) OVER w AS lat1,
      LAG(longitude) OVER w AS lon1
    FROM gps_data
    WHERE period_id = p_period
    WINDOW w AS (ORDER BY date_gps, hour_UTC)
  )
  SELECT
    SUM(
      -- Fórmula de Haversine
      2 * 6371
      * ASIN(
          SQRT(
            POWER(SIN(RADIANS((lat2 - lat1) / 2)), 2)
            + COS(RADIANS(lat1))
              * COS(RADIANS(lat2))
              * POWER(SIN(RADIANS((lon2 - lon1) / 2)), 2)
          )
        )
    )
  INTO total_dist
  FROM pares
  WHERE lat1 IS NOT NULL AND lon1 IS NOT NULL;

  RETURN total_dist;
END;
$$;


