-- Modify "cars" table
ALTER TABLE "public"."cars" ADD COLUMN "brand" character varying NOT NULL, ADD COLUMN "license_plate" character varying NOT NULL;
-- Create index "cars_license_plate_key" to table: "cars"
CREATE UNIQUE INDEX "cars_license_plate_key" ON "public"."cars" ("license_plate");
