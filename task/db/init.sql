-- Создание таблицы пользователей
CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     username TEXT UNIQUE NOT NULL,
                                     balance INT DEFAULT 0 CHECK (balance >= 0)
);

-- Создание таблицы инвентаря
CREATE TABLE IF NOT EXISTS inventory (
                                         id SERIAL PRIMARY KEY,
                                         user_id INT NOT NULL,
                                         item_type TEXT NOT NULL,
                                         quantity INT DEFAULT 1 CHECK (quantity >= 0),
                                         FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Создание таблицы истории переводов монет
CREATE TABLE IF NOT EXISTS coin_history (
                                            id SERIAL PRIMARY KEY,
                                            from_user_id INT NOT NULL,
                                            to_user_id INT NOT NULL,
                                            amount INT CHECK (amount > 0),
                                            timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                            FOREIGN KEY (from_user_id) REFERENCES users(id) ON DELETE CASCADE,
                                            FOREIGN KEY (to_user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Индексы для ускорения поиска по истории транзакций
CREATE INDEX idx_coin_history_from ON coin_history (from_user_id);
CREATE INDEX idx_coin_history_to ON coin_history (to_user_id);