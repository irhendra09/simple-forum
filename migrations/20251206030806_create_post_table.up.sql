CREATE TABLE IF NOT EXISTS posts (
                                     id SERIAL PRIMARY KEY,
                                     user_id INT NOT NULL,
                                     post_title VARCHAR(250) NOT NULL,
                                     post_content TEXT NOT NULL,
                                     post_hashtags TEXT NOT NULL,
                                     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     created_by TEXT NOT NULL,
                                     updated_by TEXT NOT NULL
);
