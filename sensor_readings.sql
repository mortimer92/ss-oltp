DROP INDEX IF EXISTS sensor_readings_idx;
DROP TABLE IF EXISTS sensor_readings;
CREATE TABLE sensor_readings
(
	sensor_reading_id	serial NOT NULL,
	temp_value			double precision,
	recorded_at 		timestamp,
	CONSTRAINT sensor_readings_pk PRIMARY KEY (sensor_reading_id)
);

CREATE INDEX sensor_readings_idx ON sensor_readings (recorded_at);

--wont be needing the ID column in the insert since Postgres auto increments the serial
-- (also there is no need for a sequence)
/*
INSERT INTO sensor_readings
	(temp_value, recorded_at)
VALUES
	( 23.86, now());
*/
