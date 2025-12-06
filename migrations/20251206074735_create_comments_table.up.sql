CREATE TABLE IF NOT EXISTS comments (
                                        id BIGSERIAL PRIMARY KEY,
                                        post_id INT NOT NULL,
                                        user_id BIGINT NOT NULL,
                                        comment_content TEXT NOT NULL,
                                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                        created_by TEXT NOT NULL,
                                        updated_by TEXT NOT NULL,
                                        CONSTRAINT fk_post_id_comments FOREIGN KEY (post_id) REFERENCES posts(id),
                                        CONSTRAINT fk_user_id_comments FOREIGN KEY (user_id) REFERENCES users(id)
);
