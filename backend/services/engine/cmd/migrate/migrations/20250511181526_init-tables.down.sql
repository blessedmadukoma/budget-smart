ALTER TABLE "ml_models"
DROP CONSTRAINT IF EXISTS "ml_models_user_id_fkey";

ALTER TABLE "transactions"
DROP CONSTRAINT IF EXISTS "transactions_category_id_fkey";

ALTER TABLE "transactions"
DROP CONSTRAINT IF EXISTS "transactions_budget_id_fkey";

ALTER TABLE "categories"
DROP CONSTRAINT IF EXISTS "categories_user_id_fkey";

ALTER TABLE "budgets"
DROP CONSTRAINT IF EXISTS "budgets_user_id_fkey";

-- Drop tables
DROP TABLE IF EXISTS "ml_models";

DROP TABLE IF EXISTS "transactions";

DROP TABLE IF EXISTS "categories";

DROP TABLE IF EXISTS "budgets";

DROP TABLE IF EXISTS "users";