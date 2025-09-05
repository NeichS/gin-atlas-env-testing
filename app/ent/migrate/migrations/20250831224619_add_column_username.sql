-- Modify "users" table
ALTER TABLE "public"."users" ADD COLUMN "username" character varying NOT NULL;
-- Create index "users_username_key" to table: "users"
CREATE UNIQUE INDEX "users_username_key" ON "public"."users" ("username");
