-- Seed data for Ottawa bike repair stations
-- Center coordinates: [45.424721, -75.695]

-- Insert repair stations around downtown Ottawa
INSERT INTO repair_station (repair_station_id, address_str, location, created_at) VALUES
    (gen_random_uuid(), 'Parliament Hill, Wellington St, Ottawa, ON', ST_GeogFromText('POINT(-75.6972 45.4236)'), NOW()),
    (gen_random_uuid(), 'Rideau Canal Pathway, Colonel By Dr, Ottawa, ON', ST_GeogFromText('POINT(-75.6950 45.4215)'), NOW()),
    (gen_random_uuid(), 'University of Ottawa Campus, 75 Laurier Ave E, Ottawa, ON', ST_GeogFromText('POINT(-75.6820 45.4231)'), NOW()),
    (gen_random_uuid(), 'ByWard Market, 55 ByWard Market Square, Ottawa, ON', ST_GeogFromText('POINT(-75.6934 45.4288)'), NOW()),
    (gen_random_uuid(), 'Confederation Park, Elgin St, Ottawa, ON', ST_GeogFromText('POINT(-75.6920 45.4198)'), NOW()),
    (gen_random_uuid(), 'National Gallery of Canada, 380 Sussex Dr, Ottawa, ON', ST_GeogFromText('POINT(-75.6991 45.4297)'), NOW()),
    (gen_random_uuid(), 'Lansdowne Park, 1015 Bank St, Ottawa, ON', ST_GeogFromText('POINT(-75.6832 45.3978)'), NOW()),
    (gen_random_uuid(), 'Dow''s Lake Pavilion, 1001 Queen Elizabeth Driveway, Ottawa, ON', ST_GeogFromText('POINT(-75.7050 45.3978)'), NOW()),
    (gen_random_uuid(), 'Carleton University, 1125 Colonel By Dr, Ottawa, ON', ST_GeogFromText('POINT(-75.6972 45.3875)'), NOW()),
    (gen_random_uuid(), 'Little Italy, Preston St & Gladstone Ave, Ottawa, ON', ST_GeogFromText('POINT(-75.7105 45.4065)'), NOW());

-- Get the repair station IDs for reference (using CTEs for clean insertion)
WITH stations AS (
    SELECT repair_station_id FROM repair_station WHERE address_str LIKE '%Parliament Hill%'
),
-- Add comprehensive tools to Parliament Hill station
parliament_tools AS (
    INSERT INTO tool (repair_station_id, tool_type_id)
    SELECT repair_station_id, 'screwdriver_phillips'::tool_type FROM stations
    UNION ALL
    SELECT repair_station_id, 'screwdriver_flat'::tool_type FROM stations
    UNION ALL
    SELECT repair_station_id, 'tire_levers'::tool_type FROM stations
    UNION ALL
    SELECT repair_station_id, 'tire_pump'::tool_type FROM stations
    UNION ALL
    SELECT repair_station_id, 'hex_key_set'::tool_type FROM stations
    RETURNING tool_id
)
SELECT 1; -- Dummy select to complete CTE

-- Add tools to Rideau Canal station
WITH stations AS (
    SELECT repair_station_id FROM repair_station WHERE address_str LIKE '%Rideau Canal%'
)
INSERT INTO tool (repair_station_id, tool_type_id)
SELECT repair_station_id, 'tire_pump'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'tire_levers'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'hex_key_set'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'screwdriver_phillips'::tool_type FROM stations;

-- Add tools to University of Ottawa station
WITH stations AS (
    SELECT repair_station_id FROM repair_station WHERE address_str LIKE '%University of Ottawa%'
)
INSERT INTO tool (repair_station_id, tool_type_id)
SELECT repair_station_id, 'tire_pump'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'tire_levers'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'hex_key_set'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'cone_wrench_8_10mm'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'headset_pedal_wrench'::tool_type FROM stations;

-- Add tools to ByWard Market station
WITH stations AS (
    SELECT repair_station_id FROM repair_station WHERE address_str LIKE '%ByWard Market%'
)
INSERT INTO tool (repair_station_id, tool_type_id)
SELECT repair_station_id, 'tire_pump'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'tire_levers'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'screwdriver_phillips'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'screwdriver_flat'::tool_type FROM stations;

-- Add tools to Confederation Park station
WITH stations AS (
    SELECT repair_station_id FROM repair_station WHERE address_str LIKE '%Confederation Park%'
)
INSERT INTO tool (repair_station_id, tool_type_id)
SELECT repair_station_id, 'tire_pump'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'hex_key_set'::tool_type FROM stations;

-- Add tools to National Gallery station
WITH stations AS (
    SELECT repair_station_id FROM repair_station WHERE address_str LIKE '%National Gallery%'
)
INSERT INTO tool (repair_station_id, tool_type_id)
SELECT repair_station_id, 'tire_pump'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'tire_levers'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'hex_key_set'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'torx_t25'::tool_type FROM stations;

-- Add tools to Lansdowne Park station
WITH stations AS (
    SELECT repair_station_id FROM repair_station WHERE address_str LIKE '%Lansdowne Park%'
)
INSERT INTO tool (repair_station_id, tool_type_id)
SELECT repair_station_id, 'tire_pump'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'tire_levers'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'screwdriver_phillips'::tool_type FROM stations;

-- Add tools to Dow's Lake station
WITH stations AS (
    SELECT repair_station_id FROM repair_station WHERE address_str LIKE '%Dow''s Lake%'
)
INSERT INTO tool (repair_station_id, tool_type_id)
SELECT repair_station_id, 'tire_pump'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'tire_levers'::tool_type FROM stations;

-- Add comprehensive tools to Carleton University station
WITH stations AS (
    SELECT repair_station_id FROM repair_station WHERE address_str LIKE '%Carleton University%'
)
INSERT INTO tool (repair_station_id, tool_type_id)
SELECT repair_station_id, 'screwdriver_phillips'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'screwdriver_flat'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'tire_levers'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'tire_pump'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'hex_key_set'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'cone_wrench_9_11mm'::tool_type FROM stations;

-- Add tools to Little Italy station
WITH stations AS (
    SELECT repair_station_id FROM repair_station WHERE address_str LIKE '%Little Italy%'
)
INSERT INTO tool (repair_station_id, tool_type_id)
SELECT repair_station_id, 'tire_pump'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'hex_key_set'::tool_type FROM stations
UNION ALL
SELECT repair_station_id, 'screwdriver_phillips'::tool_type FROM stations;

-- Add some sample reports (some tools broken/missing, one station missing)
WITH byward_tools AS (
    SELECT t.tool_id
    FROM tool t
    JOIN repair_station rs ON t.repair_station_id = rs.repair_station_id
    WHERE rs.address_str LIKE '%ByWard Market%' AND t.tool_type_id = 'tire_pump'
    LIMIT 1
)
INSERT INTO tool_report (tool_id, created_at, report_type)
SELECT tool_id, NOW() - INTERVAL '2 days', 'tool_broken'::tool_report_type FROM byward_tools;

WITH confed_tools AS (
    SELECT t.tool_id
    FROM tool t
    JOIN repair_station rs ON t.repair_station_id = rs.repair_station_id
    WHERE rs.address_str LIKE '%Confederation Park%' AND t.tool_type_id = 'hex_key_set'
    LIMIT 1
)
INSERT INTO tool_report (tool_id, created_at, report_type)
SELECT tool_id, NOW() - INTERVAL '5 days', 'tool_missing'::tool_report_type FROM confed_tools;

-- Report a working tire pump at Rideau Canal (positive report)
WITH rideau_tools AS (
    SELECT t.tool_id
    FROM tool t
    JOIN repair_station rs ON t.repair_station_id = rs.repair_station_id
    WHERE rs.address_str LIKE '%Rideau Canal%' AND t.tool_type_id = 'tire_pump'
    LIMIT 1
)
INSERT INTO tool_report (tool_id, created_at, report_type)
SELECT tool_id, NOW() - INTERVAL '1 day', 'tool_present'::tool_report_type FROM rideau_tools;

-- Add some images to a few of the stations --
WITH parliament_station AS (
    SELECT repair_station_id FROM repair_station WHERE address_str LIKE '%Parliament Hill%'
)
INSERT INTO repair_station_photo (repair_station_photo_id, repair_station_id, photo_key)
SELECT gen_random_uuid(), repair_station_id, 'photo.png' FROM parliament_station;

WITH byward_station AS (
    SELECT repair_station_id FROM repair_station WHERE address_str LIKE '%ByWard Market%'
)
INSERT INTO repair_station_photo (repair_station_photo_id, repair_station_id, photo_key)
SELECT gen_random_uuid(), repair_station_id, 'photo.png' FROM byward_station;

WITH uottawa_station AS (
    SELECT repair_station_id FROM repair_station WHERE address_str LIKE '%University of Ottawa%'
)
INSERT INTO repair_station_photo (repair_station_photo_id, repair_station_id, photo_key)
SELECT gen_random_uuid(), repair_station_id, 'photo.png' FROM uottawa_station;
