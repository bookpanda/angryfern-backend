-- Create "scores" table
CREATE TABLE "scores" (
    "code" VARCHAR(10) NOT NULL,
    -- "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    -- "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "country_name" VARCHAR(100) NOT NULL,
    "click_count" INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY ("code")
)