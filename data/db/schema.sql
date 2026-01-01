CREATE EXTENSION postgis;

CREATE TYPE "tool_type" AS ENUM (
	'screwdriver_phillips',
	'screwdriver_flat',
	'tire_levers',
	'tire_pump',
	'headset_pedal_wrench',
	'cone_wrench_8_10mm',
	'cone_wrench_9_11mm',
	'torx_t25',
	'hex_key_set'
);

CREATE TYPE "tool_report_type" AS ENUM (
	'tool_missing',
	'tool_broken',
	'tool_present'
);

CREATE TYPE "repair_station_report_type" AS ENUM (
	'station_missing'
);

CREATE TABLE IF NOT EXISTS "repair_station" (
	"repair_station_id" UUID NOT NULL UNIQUE DEFAULT gen_random_uuid(),
	"address_str" TEXT NOT NULL,
	"location" GEOGRAPHY(POINT, 4326) NOT NULL UNIQUE,
	"created_at" TIMESTAMP NOT NULL,
	PRIMARY KEY("repair_station_id")
);

CREATE TABLE IF NOT EXISTS "tool" (
	"tool_id" UUID NOT NULL DEFAULT gen_random_uuid(),
	"repair_station_id" UUID NOT NULL,
	"tool_type_id" TOOL_TYPE NOT NULL,
	PRIMARY KEY("tool_id")
);

CREATE TABLE IF NOT EXISTS "tool_report" (
	"report_id" UUID NOT NULL DEFAULT gen_random_uuid(),
	"tool_id" UUID NOT NULL,
	"created_at" TIMESTAMP NOT NULL,
	"report_type" TOOL_REPORT_TYPE NOT NULL,
	PRIMARY KEY("report_id")
);

CREATE TABLE IF NOT EXISTS "repair_station_report" (
	"report_id" UUID NOT NULL UNIQUE DEFAULT gen_random_uuid(),
	"repair_station_id" UUID NOT NULL,
	"created_at" TIMESTAMP NOT NULL,
	"report_type" REPAIR_STATION_REPORT_TYPE,
	PRIMARY KEY("report_id")
);

CREATE TABLE IF NOT EXISTS "repair_station_photo" (
	"repair_station_photo_id" UUID NOT NULL UNIQUE,
	"repair_station_id" UUID NOT NULL,
	"photo_key" VARCHAR(512) NOT NULL,
	PRIMARY KEY("repair_station_photo_id")
);

ALTER TABLE "tool"
ADD FOREIGN KEY("repair_station_id") REFERENCES "repair_station"("repair_station_id")
ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "tool_report"
ADD FOREIGN KEY("tool_id") REFERENCES "tool"("tool_id")
ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "repair_station_report"
ADD FOREIGN KEY("repair_station_id") REFERENCES "repair_station"("repair_station_id")
ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "repair_station_photo"
ADD FOREIGN KEY("repair_station_id") REFERENCES "repair_station"("repair_station_id")
ON UPDATE NO ACTION ON DELETE NO ACTION;
