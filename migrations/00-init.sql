-- Create "scoreboard" table
CREATE TABLE "scoreboard" (
    "code" VARCHAR(10) NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "name" VARCHAR(100) NOT NULL,
    "count" INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY ("code")
)