package db

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tukembaev/bookVisionGo/internal/models"
)

// SeedArticles - заполнение базы данных начальными статьями
func SeedArticles(ctx context.Context, pool *pgxpool.Pool) error {
	articles := []models.Article{
		// --- Связанные с книгами (60% = 12 статей) ---

		// Книга: Отцы и дети (b0000000-0000-0000-0000-000000000001)
		{
			ID:                  uuid.New().String(),
			Title:               "Нигилизм в романе 'Отцы и дети'",
			Type:                models.ArticleTypeAnalysis,
			BookID:              stringPtr("b0000000-0000-0000-0000-000000000001"),
			Excerpt:             "Глубокий разбор философии Базарова и его влияния на русскую литературу.",
			CreatedAt:           time.Now(),
			Likes:               150,
			Views:               1200,
			ReadingMinutes:      intPtr(10),
			Verified:            true,
			VerificationType:    verificationTypePtr(models.VerificationTypeAI),
			NoSpoilers:          false,
			ShouldReadReadiness: models.ArticleReadinessMust,
			Content: []map[string]interface{}{
				{"block_type": "h2", "text": "Базаров как зеркало эпохи"},
				{"block_type": "p", "text": "Евгений Базаров — персонаж, который перевернул представление о герое своего времени..."},
				{"block_type": "quote", "text": "Природа не храм, а мастерская, и человек в ней работник."},
			},
		},
		{
			ID:                  uuid.New().String(),
			Title:               "Конфликт поколений: тогда и сейчас",
			Type:                models.ArticleTypeDiscussion,
			BookID:              stringPtr("b0000000-0000-0000-0000-000000000001"),
			Excerpt:             "Актуален ли конфликт Кирсановых и Базарова в XXI веке?",
			CreatedAt:           time.Now(),
			Likes:               85,
			Views:               900,
			ReadingMinutes:      intPtr(7),
			Verified:            true,
			VerificationType:    verificationTypePtr(models.VerificationTypeCommunity),
			NoSpoilers:          true,
			ShouldReadReadiness: models.ArticleReadinessMaybe,
			Content: []map[string]interface{}{
				{"block_type": "p", "text": "Многие читатели задаются вопросом, насколько изменились отношения родителей и детей за последние 150 лет..."},
			},
		},
		{
			ID:                  uuid.New().String(),
			Title:               "Гид по творчеству Тургенева",
			Type:                models.ArticleTypeGuide,
			BookID:              stringPtr("b0000000-0000-0000-0000-000000000001"),
			Excerpt:             "С чего начать знакомство с автором и какое место занимает 'Отцы и дети' в его библиографии.",
			CreatedAt:           time.Now(),
			ReadingMinutes:      intPtr(12),
			ShouldReadReadiness: models.ArticleReadinessMust,
			Content: []map[string]interface{}{
				{"block_type": "h3", "text": "Ранние повести"},
				{"block_type": "p", "text": "Прежде всего стоит обратить внимание на 'Записки охотника'..."},
			},
		},

		// Книга: Преступление и наказание (b0000000-0000-0000-0000-000000000002)
		{
			ID:                  uuid.New().String(),
			Title:               "Психология Раскольникова",
			Type:                models.ArticleTypeAnalysis,
			BookID:              stringPtr("b0000000-0000-0000-0000-000000000002"),
			Excerpt:             "Почему теория о 'право имеющих' привела к катастрофе.",
			CreatedAt:           time.Now(),
			Likes:               300,
			Views:               5000,
			ReadingMinutes:      intPtr(15),
			Verified:            true,
			VerificationType:    verificationTypePtr(models.VerificationTypeAI),
			NoSpoilers:          false,
			ShouldReadReadiness: models.ArticleReadinessMust,
			Content: []map[string]interface{}{
				{"block_type": "h2", "text": "Тварь ли я дрожащая или право имею?"},
				{"block_type": "p", "text": "Достоевский виртуозно описывает процесс разложения человеческой души под гнетом ложной идеи..."},
			},
		},
		{
			ID:                  uuid.New().String(),
			Title:               "Петербург Достоевского: город как персонаж",
			Type:                models.ArticleTypeCollection,
			BookID:              stringPtr("b0000000-0000-0000-0000-000000000002"),
			Excerpt:             "Маршрут по местам действия романа в современном Санкт-Петербурге.",
			CreatedAt:           time.Now(),
			Likes:               210,
			Views:               2800,
			ReadingMinutes:      intPtr(10),
			NoSpoilers:          true,
			ShouldReadReadiness: models.ArticleReadinessMaybe,
			Content: []map[string]interface{}{
				{"block_type": "p", "text": "Сенная площадь, Столярный переулок, дом Раскольникова — эти места до сих пор хранят атмосферу романа..."},
			},
		},
		{
			ID:                  uuid.New().String(),
			Title:               "Сравнение экранизаций 'Преступления и наказания'",
			Type:                models.ArticleTypeComparison,
			BookID:              stringPtr("b0000000-0000-0000-0000-000000000002"),
			Excerpt:             "От классики Кулиджанова до современных интерпретаций.",
			CreatedAt:           time.Now(),
			ReadingMinutes:      intPtr(8),
			ShouldReadReadiness: models.ArticleReadinessNo,
			Content: []map[string]interface{}{
				{"block_type": "p", "text": "Каждая эпоха видит Раскольникова по-своему. Давайте сравним самые значимые работы кинорежиссеров..."},
			},
		},

		// Книга: Мастер и Маргарита (b0000000-0000-0000-0000-000000000003)
		{
			ID:                  uuid.New().String(),
			Title:               "Мистика и реальность в романе Булгакова",
			Type:                models.ArticleTypeAnalysis,
			BookID:              stringPtr("b0000000-0000-0000-0000-000000000003"),
			Excerpt:             "Разбор символики Воланда и его свиты в контексте советской Москвы.",
			CreatedAt:           time.Now(),
			Likes:               450,
			Views:               6000,
			ReadingMinutes:      intPtr(20),
			Verified:            true,
			VerificationType:    verificationTypePtr(models.VerificationTypeCommunity),
			NoSpoilers:          false,
			ShouldReadReadiness: models.ArticleReadinessMust,
			Content: []map[string]interface{}{
				{"block_type": "h2", "text": "Явление Воланда"},
				{"block_type": "p", "text": "Булгаков использует сатанинскую свиту для обнажения пороков общества..."},
			},
		},
		{
			ID:                  uuid.New().String(),
			Title:               "Почему 'Мастер и Маргарита' — роман в романе?",
			Type:                models.ArticleTypeAnalysis,
			BookID:              stringPtr("b0000000-0000-0000-0000-000000000003"),
			Excerpt:             "Структура произведения и связь ершалаимских глав с московскими.",
			CreatedAt:           time.Now(),
			Likes:               120,
			Views:               1500,
			ReadingMinutes:      intPtr(11),
			NoSpoilers:          true,
			ShouldReadReadiness: models.ArticleReadinessMust,
			Content: []map[string]interface{}{
				{"block_type": "p", "text": "Параллелизм двух миров — древнего Иерусалима и Москвы 30-х годов — создает уникальное полотно..."},
			},
		},

		// Книга: Отверженные (b0000000-0000-0000-0000-000000000004)
		{
			ID:                  uuid.New().String(),
			Title:               "Жан Вальжан: путь к искуплению",
			Type:                models.ArticleTypeReview,
			BookID:              stringPtr("b0000000-0000-0000-0000-000000000004"),
			Excerpt:             "Как одна встреча с епископом может изменить жизнь преступника навсегда.",
			CreatedAt:           time.Now(),
			Likes:               180,
			Views:               2000,
			ReadingMinutes:      intPtr(12),
			Verified:            true,
			VerificationType:    verificationTypePtr(models.VerificationTypeAI),
			NoSpoilers:          false,
			ShouldReadReadiness: models.ArticleReadinessMust,
			Content: []map[string]interface{}{
				{"block_type": "p", "text": "Гюго создал один из самых мощных образов трансформации личности в мировой литературе..."},
			},
		},
		{
			ID:                  uuid.New().String(),
			Title:               "Исторический фон 'Отверженных'",
			Type:                models.ArticleTypeGuide,
			BookID:              stringPtr("b0000000-0000-0000-0000-000000000004"),
			Excerpt:             "Июньское восстание 1832 года и реалии Франции XIX века.",
			CreatedAt:           time.Now(),
			ReadingMinutes:      intPtr(14),
			ShouldReadReadiness: models.ArticleReadinessMaybe,
			Content: []map[string]interface{}{
				{"block_type": "h3", "text": "Баррикады Парижа"},
				{"block_type": "p", "text": "Чтобы понять действия героев, нужно знать контекст политической нестабильности Франции того времени..."},
			},
		},

		// Книга: Норвежский лес (b0000000-0000-0000-0000-000000000005)
		{
			ID:                  uuid.New().String(),
			Title:               "Меланхолия и джаз: атмосфера Мураками",
			Type:                models.ArticleTypeAnalysis,
			BookID:              stringPtr("b0000000-0000-0000-0000-000000000005"),
			Excerpt:             "Как музыка влияет на восприятие романа 'Норвежский лес'.",
			CreatedAt:           time.Now(),
			Likes:               340,
			Views:               4200,
			ReadingMinutes:      intPtr(9),
			Verified:            true,
			NoSpoilers:          true,
			ShouldReadReadiness: models.ArticleReadinessMust,
			Content: []map[string]interface{}{
				{"block_type": "p", "text": "Связь названия с песней The Beatles и постоянное присутствие музыки создает неповторимый ритм текста..."},
			},
		},
		{
			ID:                  uuid.New().String(),
			Title:               "Ватанабэ между двух огней",
			Type:                models.ArticleTypeDiscussion,
			BookID:              stringPtr("b0000000-0000-0000-0000-000000000005"),
			Excerpt:             "Выбор между прошлым (Наоко) и будущим (Мидори).",
			CreatedAt:           time.Now(),
			ReadingMinutes:      intPtr(10),
			ShouldReadReadiness: models.ArticleReadinessMaybe,
			Content: []map[string]interface{}{
				{"block_type": "p", "text": "Главный герой оказывается в ситуации сложного экзистенциального выбора..."},
			},
		},

		// --- Самостоятельные (40% = 8 статей) ---

		{
			ID:                  uuid.New().String(),
			Title:               "Как читать больше книг в год?",
			Type:                models.ArticleTypeGuide,
			Excerpt:             "Практические советы по скорочтению и планированию времени.",
			CreatedAt:           time.Now(),
			Likes:               500,
			Views:               15000,
			ReadingMinutes:      intPtr(5),
			Verified:            true,
			VerificationType:    verificationTypePtr(models.VerificationTypeAI),
			NoSpoilers:          true,
			ShouldReadReadiness: models.ArticleReadinessMust,
			Content: []map[string]interface{}{
				{"block_type": "p", "text": "Чтение — это навык. И как любой навык, его можно тренировать..."},
			},
		},
		{
			ID:                  uuid.New().String(),
			Title:               "Топ-10 книг для отдыха",
			Type:                models.ArticleTypeCollection,
			Excerpt:             "Легкие произведения, которые помогут расслабиться после рабочего дня.",
			CreatedAt:           time.Now(),
			Likes:               120,
			Views:               3000,
			ReadingMinutes:      intPtr(6),
			NoSpoilers:          true,
			ShouldReadReadiness: models.ArticleReadinessMaybe,
			Content: []map[string]interface{}{
				{"block_type": "p", "text": "В этом списке мы собрали романы, которые читаются на одном дыхании..."},
			},
		},
		{
			ID:                  uuid.New().String(),
			Title:               "Почему бумажные книги все еще популярны?",
			Type:                models.ArticleTypeDiscussion,
			Excerpt:             "Битва форматов: бумага, электронные книги и аудиокниги.",
			CreatedAt:           time.Now(),
			ReadingMinutes:      intPtr(8),
			ShouldReadReadiness: models.ArticleReadinessMaybe,
			Content: []map[string]interface{}{
				{"block_type": "p", "text": "Несмотря на цифровизацию, запах бумаги и тактильные ощущения остаются важными для читателей..."},
			},
		},
		{
			ID:                  uuid.New().String(),
			Title:               "Забытые классики: кого стоит перечитать?",
			Type:                models.ArticleTypeCollection,
			Excerpt:             "Авторы, которые были популярны раньше, но сейчас оказались в тени.",
			CreatedAt:           time.Now(),
			ReadingMinutes:      intPtr(12),
			ShouldReadReadiness: models.ArticleReadinessMust,
			Content: []map[string]interface{}{
				{"block_type": "p", "text": "Некоторые писатели незаслуженно забыты. Мы решили вспомнить их имена..."},
			},
		},
		{
			ID:                  uuid.New().String(),
			Title:               "Влияние литературы на кино",
			Type:                models.ArticleTypeAnalysis,
			Excerpt:             "Как великие романы формируют современный кинематограф.",
			CreatedAt:           time.Now(),
			Likes:               90,
			Views:               1100,
			ReadingMinutes:      intPtr(15),
			ShouldReadReadiness: models.ArticleReadinessNo,
			Content: []map[string]interface{}{
				{"block_type": "p", "text": "Голливуд уже давно черпает вдохновение в классической литературе. Но всегда ли это удачно?"},
			},
		},
		{
			ID:                  uuid.New().String(),
			Title:               "Искусственный интеллект и писательство",
			Type:                models.ArticleTypeDiscussion,
			Excerpt:             "Заменит ли нейросеть автора бестселлеров?",
			CreatedAt:           time.Now(),
			Likes:               600,
			Views:               8000,
			ReadingMinutes:      intPtr(10),
			ShouldReadReadiness: models.ArticleReadinessMust,
			Content: []map[string]interface{}{
				{"block_type": "p", "text": "С появлением ChatGPT мир литературы столкнулся с новым вызовом..."},
			},
		},
		{
			ID:                  uuid.New().String(),
			Title:               "Библиотеки будущего",
			Type:                models.ArticleTypeGuide,
			Excerpt:             "Как меняются современные пространства для чтения.",
			CreatedAt:           time.Now(),
			ReadingMinutes:      intPtr(6),
			ShouldReadReadiness: models.ArticleReadinessMaybe,
			Content: []map[string]interface{}{
				{"block_type": "p", "text": "Современная библиотека — это уже не просто хранилище книг, а коворкинг и культурный центр..."},
			},
		},
		{
			ID:                  uuid.New().String(),
			Title:               "Почему важно читать классику?",
			Type:                models.ArticleTypeAnalysis,
			Excerpt:             "Аргументы в пользу школьной программы и не только.",
			CreatedAt:           time.Now(),
			Likes:               45,
			Views:               600,
			ReadingMinutes:      intPtr(12),
			ShouldReadReadiness: models.ArticleReadinessNo,
			Content: []map[string]interface{}{
				{"block_type": "p", "text": "Классическая литература дает нам базу для понимания культуры и человеческой природы..."},
			},
		},
	}

	for _, article := range articles {
		err := insertArticle(ctx, pool, &article)
		if err != nil {
			return fmt.Errorf("ошибка вставки статьи %s: %w", article.Title, err)
		}
		fmt.Printf("Статья '%s' успешно добавлена\n", article.Title)
	}

	return nil
}

func insertArticle(ctx context.Context, pool *pgxpool.Pool, article *models.Article) error {
	query := `
		INSERT INTO articles (
			id, title, type, author_id, book_id, excerpt,
			created_at, likes, views, reading_minutes, cover_url,
			verified, verification_type, no_spoilers, readiness, content
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
		) ON CONFLICT (id) DO NOTHING`

	_, err := pool.Exec(ctx, query,
		article.ID,
		article.Title,
		article.Type,
		article.AuthorID,
		article.BookID,
		article.Excerpt,
		article.CreatedAt,
		article.Likes,
		article.Views,
		article.ReadingMinutes,
		article.CoverURL,
		article.Verified,
		article.VerificationType,
		article.NoSpoilers,
		article.ShouldReadReadiness,
		article.Content,
	)

	return err
}
