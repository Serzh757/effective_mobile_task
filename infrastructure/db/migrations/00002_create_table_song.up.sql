-- Создаем таблицу songs
CREATE TABLE IF NOT EXISTS songs
(
    id           SERIAL PRIMARY KEY,    -- Идентификатор песни
    "group"      VARCHAR(255) NOT NULL, -- Название группы
    song         VARCHAR(255) NOT NULL, -- Название песни
    release_date DATE,                  -- Дата релиза
    text         TEXT,                  -- Текст песни
    link         VARCHAR(2083)          -- Ссылка на песню
);

-- Добавляем индекс для быстрого поиска по группе и названию песни
CREATE INDEX idx_group_song ON songs ("group", song);