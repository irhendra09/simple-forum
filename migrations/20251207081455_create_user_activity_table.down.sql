ALTER TABLE user_activities DROP CONSTRAINT IF EXISTS fk_post_id_user_activities;
ALTER TABLE user_activities DROP CONSTRAINT IF EXISTS fk_user_id_user_activities;
DROP TABLE IF EXISTS user_activities;