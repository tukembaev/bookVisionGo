-- Enums
CREATE TYPE user_role AS ENUM ('user', 'moderator', 'admin');
CREATE TYPE age_rating AS ENUM ('6+', '12+', '16+', '18+');
CREATE TYPE verification_type AS ENUM ('AI', 'Community');
 
-- Users таблица
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    avatar_url TEXT,
    role user_role NOT NULL DEFAULT 'user',
    created_at TIMESTAMP DEFAULT NOW(),
    books_read INTEGER DEFAULT 0,
    reviews_count INTEGER DEFAULT 0,
    likes_received INTEGER DEFAULT 0,
    profile_visibility VARCHAR(10) DEFAULT 'public',
    activity_visibility VARCHAR(10) DEFAULT 'public'
);
 
-- Books таблица
CREATE TABLE books (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    original_title VARCHAR(255),
    author VARCHAR(255) NOT NULL,
    year INTEGER,
    genres TEXT[],
    age_rating age_rating,
    author_country VARCHAR(100),
    description TEXT NOT NULL,
    cover_url TEXT,
    pages_count INTEGER NOT NULL,
    verified BOOLEAN DEFAULT false,
    verification_type verification_type,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW(),
    tags TEXT[],
    average_rating DECIMAL(3,1) DEFAULT 0,
    rating_count INTEGER DEFAULT 0
);
 
-- Book Parts таблица
CREATE TABLE book_parts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    book_id UUID REFERENCES books(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    order_num INTEGER NOT NULL,
    page_start INTEGER,
    page_end INTEGER,
    mood_tags TEXT[],
    average_rating DECIMAL(3,1),
    UNIQUE(book_id, order_num)
);
 
-- Characters таблица
CREATE TABLE characters (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    book_id UUID REFERENCES books(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    source VARCHAR(20) NOT NULL DEFAULT 'community',
    verified BOOLEAN DEFAULT false,
    popularity_score INTEGER DEFAULT 0
);
 
-- Reviews таблица
CREATE TABLE reviews (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    book_id UUID REFERENCES books(id) ON DELETE CASCADE,
    rating INTEGER CHECK (rating >= 1 AND rating <= 10),
    text TEXT NOT NULL,
    liked_characters TEXT[],
    disliked_characters TEXT[],
    best_parts TEXT[],
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(user_id, book_id)
);
 
-- Comments таблица
CREATE TABLE comments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    book_id UUID REFERENCES books(id) ON DELETE CASCADE,
    part_id UUID REFERENCES book_parts(id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    likes INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    parent_comment_id UUID REFERENCES comments(id),
    reply_to_user_id UUID REFERENCES users(id)
);
 
-- User Book Progress таблица
CREATE TABLE user_book_progress (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    book_id UUID REFERENCES books(id) ON DELETE CASCADE,
    completed_part_ids TEXT[] DEFAULT '{}',
    current_part_id UUID REFERENCES book_parts(id),
    is_completed BOOLEAN DEFAULT false,
    completed_at TIMESTAMP,
    UNIQUE(user_id, book_id)
);
 
-- Playlists таблица (для музыкальных плейлистов)
CREATE TABLE playlists (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    mood_tag VARCHAR(50) NOT NULL,
    tracks TEXT[],
    created_by VARCHAR(10) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP DEFAULT NOW()
);
 
-- Challenges таблица
CREATE TABLE challenges (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    type VARCHAR(10) NOT NULL CHECK (type IN ('books', 'reviews')),
    target_count INTEGER NOT NULL,
    reward_points INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
 
-- Quotes таблица
CREATE TABLE quotes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    book_id UUID REFERENCES books(id) ON DELETE CASCADE,
    part_id UUID REFERENCES book_parts(id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
 
-- Character Illustrations таблица
CREATE TABLE character_illustrations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    image_url TEXT,
    author_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
 
-- Character Profiles таблица (расширенная информация о персонажах)
CREATE TABLE character_profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    book_id UUID REFERENCES books(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    aliases TEXT[],
    image_url TEXT,
    age TEXT,
    height TEXT,
    weight TEXT,
    social_status TEXT,
    description_no_spoilers TEXT NOT NULL,
    description_spoilers TEXT NOT NULL,
    quotes_no_spoilers TEXT[],
    quotes_spoilers TEXT[],
    favorited_by_user_ids TEXT[],
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
 
-- Character Profile Illustrations связь
CREATE TABLE character_profile_illustrations (
    character_profile_id UUID REFERENCES character_profiles(id) ON DELETE CASCADE,
    illustration_id UUID REFERENCES character_illustrations(id) ON DELETE CASCADE,
    PRIMARY KEY (character_profile_id, illustration_id)
);
 
-- Articles таблица
CREATE TABLE articles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    type VARCHAR(20) NOT NULL CHECK (type IN ('shouldRead', 'analysis', 'review', 'collection', 'guide', 'comparison', 'discussion')),
    author_id UUID REFERENCES users(id),
    book_id UUID REFERENCES books(id) ON DELETE CASCADE,
    excerpt TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    likes INTEGER DEFAULT 0,
    views INTEGER DEFAULT 0,
    reading_minutes INTEGER,
    cover_url TEXT,
    verified BOOLEAN DEFAULT false,
    verification_type verification_type,
    no_spoilers BOOLEAN NOT NULL DEFAULT true,
    readiness VARCHAR(10) CHECK (readiness IN ('must', 'maybe', 'no')),
    content JSONB
);
 
-- User Reading Sessions таблица (для отслеживания сессий чтения)
CREATE TABLE user_reading_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    book_id UUID REFERENCES books(id) ON DELETE CASCADE,
    part_id UUID REFERENCES book_parts(id) ON DELETE CASCADE,
    started_at TIMESTAMP DEFAULT NOW(),
    ended_at TIMESTAMP,
    pages_read INTEGER DEFAULT 0,
    duration_minutes INTEGER
);
 
-- User Favorites таблица (избранные книги)
CREATE TABLE user_favorites (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    book_id UUID REFERENCES books(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(user_id, book_id)
);
 
-- User Character Favorites таблица
CREATE TABLE user_character_favorites (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    character_id UUID REFERENCES characters(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(user_id, character_id)
);
 
-- Индексы для производительности
CREATE INDEX idx_books_author ON books(author);
CREATE INDEX idx_books_genres ON books USING GIN(genres);
CREATE INDEX idx_books_rating ON books(average_rating DESC);
CREATE INDEX idx_books_created_by ON books(created_by);
CREATE INDEX idx_books_tags ON books USING GIN(tags);
 
CREATE INDEX idx_reviews_book_id ON reviews(book_id);
CREATE INDEX idx_reviews_user_id ON reviews(user_id);
CREATE INDEX idx_reviews_rating ON reviews(rating);
 
CREATE INDEX idx_comments_book_id ON comments(book_id);
CREATE INDEX idx_comments_part_id ON comments(part_id);
CREATE INDEX idx_comments_user_id ON comments(user_id);
CREATE INDEX idx_comments_parent ON comments(parent_comment_id);
CREATE INDEX idx_comments_created_at ON comments(created_at DESC);
 
CREATE INDEX idx_user_progress_user_id ON user_book_progress(user_id);
CREATE INDEX idx_user_progress_book_id ON user_book_progress(book_id);
 
CREATE INDEX idx_characters_book_id ON characters(book_id);
CREATE INDEX idx_characters_popularity ON characters(popularity_score DESC);
 
CREATE INDEX idx_book_parts_book_id ON book_parts(book_id);
CREATE INDEX idx_book_parts_order ON book_parts(book_id, order_num);
 
CREATE INDEX idx_quotes_book_id ON quotes(book_id);
CREATE INDEX idx_quotes_user_id ON quotes(user_id);
 
CREATE INDEX idx_articles_author_id ON articles(author_id);
CREATE INDEX idx_articles_book_id ON articles(book_id);
CREATE INDEX idx_articles_type ON articles(type);
CREATE INDEX idx_articles_created_at ON articles(created_at DESC);
CREATE INDEX idx_articles_likes ON articles(likes DESC);
 
CREATE INDEX idx_character_profiles_book_id ON character_profiles(book_id);
CREATE INDEX idx_character_profiles_name ON character_profiles(name);
 
CREATE INDEX idx_user_favorites_user_id ON user_favorites(user_id);
CREATE INDEX idx_user_favorites_book_id ON user_favorites(book_id);
 
CREATE INDEX idx_user_reading_sessions_user_id ON user_reading_sessions(user_id);
CREATE INDEX idx_user_reading_sessions_book_id ON user_reading_sessions(book_id);
CREATE INDEX idx_user_reading_sessions_started_at ON user_reading_sessions(started_at DESC);
 
-- Индексы для массивов
CREATE INDEX idx_users_books_read ON users(books_read);
CREATE INDEX idx_users_reviews_count ON users(reviews_count);
CREATE INDEX idx_users_likes_received ON users(likes_received);