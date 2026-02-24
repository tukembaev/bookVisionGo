package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tukembaev/bookVisionGo/internal/models"
)

// SeedBooks - заполнение базы данных начальными книгами
func SeedBooks(ctx context.Context, pool *pgxpool.Pool) error {
	books := []models.Book{
		{
			ID:               "b0000000-0000-0000-0000-000000000001",
			Title:            "Отцы и дети",
			OriginalTitle:    stringPtr("Fathers and Sons"),
			Author:           "Иван Тургенев",
			Year:             intPtr(1862),
			Genres:           []string{"Роман", "Драма", "Философия"},
			AgeRating:        ageRatingPtr(models.AgeRating12),
			AuthorCountry:    stringPtr("Россия"),
			Description:      "Роман о конфликте поколений и столкновении идей в России середины XIX века.",
			CoverURL:         stringPtr("https://cdn.azbooka.ru/cv/w1100/61a7ee1f-7c15-412b-b3b0-37aec56fabc2.jpg"),
			PagesCount:       320,
			Verified:         true,
			VerificationType: verificationTypePtr(models.VerificationTypeAI),
			CreatedBy:        nil,
			CreatedAt:        time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			Tags:             []string{"философия", "общество", "конфликт"},
			AverageRating:    8.6,
			RatingCount:      1240,
		},
		{
			ID:               "b0000000-0000-0000-0000-000000000002",
			Title:            "Преступление и наказание",
			Author:           "Фёдор Достоевский",
			Year:             intPtr(1866),
			Genres:           []string{"Роман", "Психология", "Философия"},
			AgeRating:        ageRatingPtr(models.AgeRating16),
			AuthorCountry:    stringPtr("Россия"),
			Description:      "История внутреннего кризиса и морального выбора, разворачивающаяся вокруг преступления.",
			CoverURL:         stringPtr("https://flibusta.su/b/img/big/208394.jpg"),
			PagesCount:       560,
			Verified:         true,
			VerificationType: verificationTypePtr(models.VerificationTypeAI),
			CreatedBy:        nil,
			CreatedAt:        time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC),
			Tags:             []string{"психология", "вина", "искупление"},
			AverageRating:    9.1,
			RatingCount:      2305,
		},
		{
			ID:               "b0000000-0000-0000-0000-000000000003",
			Title:            "Мастер и Маргарита",
			Author:           "Михаил Булгаков",
			Year:             intPtr(1967),
			Genres:           []string{"Роман", "Фантастика", "Философия"},
			AgeRating:        ageRatingPtr(models.AgeRating16),
			AuthorCountry:    stringPtr("Россия"),
			Description:      "Сатира и мистический roman, переплетающий несколько линий и смысловых пластов.",
			CoverURL:         stringPtr("https://cdn.azbooka.ru/cv/w1100/98fa6b42-e86d-4f17-9376-25e98cc784e5.jpg"),
			PagesCount:       410,
			Verified:         true,
			VerificationType: verificationTypePtr(models.VerificationTypeCommunity),
			CreatedBy:        nil,
			CreatedAt:        time.Date(2026, 1, 3, 0, 0, 0, 0, time.UTC),
			Tags:             []string{"мистика", "сатира", "любовь"},
			AverageRating:    9.0,
			RatingCount:      3102,
		},
		{
			ID:               "b0000000-0000-0000-0000-000000000004",
			Title:            "Отверженные",
			OriginalTitle:    stringPtr("Les Misérables"),
			Author:           "Виктор Гюго",
			Year:             intPtr(1862),
			Genres:           []string{"Роман", "Драма", "Исторический"},
			AgeRating:        ageRatingPtr(models.AgeRating12),
			AuthorCountry:    stringPtr("Франция"),
			Description:      "Эпический роман о милосердии, справедливости и судьбах людей на фоне эпохи.",
			CoverURL:         stringPtr("https://cdn.azbooka.ru/cv/w1100/8e4f70bd-f412-4f3c-a9cf-b1755601fb97.jpg"),
			PagesCount:       1240,
			Verified:         true,
			VerificationType: verificationTypePtr(models.VerificationTypeAI),
			CreatedBy:        nil,
			CreatedAt:        time.Date(2026, 1, 4, 0, 0, 0, 0, time.UTC),
			Tags:             []string{"эпос", "общество", "история"},
			AverageRating:    8.9,
			RatingCount:      980,
		},
		{
			ID:               "b0000000-0000-0000-0000-000000000005",
			Title:            "Норвежский лес",
			OriginalTitle:    stringPtr("Norwegian Wood"),
			Author:           "Харуки Мураками",
			Year:             intPtr(1987),
			Genres:           []string{"Роман", "Драма"},
			AgeRating:        ageRatingPtr(models.AgeRating16),
			AuthorCountry:    stringPtr("Япония"),
			Description:      "Тихий роман о взрослении, памяти и потерях.",
			CoverURL:         stringPtr("https://flibusta.su/b/img/big/589.jpg"),
			PagesCount:       384,
			Verified:         true,
			VerificationType: verificationTypePtr(models.VerificationTypeCommunity),
			CreatedBy:        nil,
			CreatedAt:        time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC),
			Tags:             []string{"взросление", "память"},
			AverageRating:    8.1,
			RatingCount:      1405,
		},
	}

	// Добавляем части книг
	bookParts := []models.BookPart{
		// Отцы и дети
		{ID: "bp0000000-0000-0000-0000-000000000001", BookID: "b0000000-0000-0000-0000-000000000001", Title: "Глава 1. Приезд в усадьбу", Content: "Май 1859 года. Помещик Николай Петрович Кирсанов с нетерпением ждет приезда сына Аркадия, который окончил университет в Петербурге. Он встречает его на почтовой станции, гордится успехами и знакомит с управляющим Василием Ивановичем. По дороге в усадьбу Марьино Николай Петрович рассказывает сыну о своей жизни, о любви к крестьянке Фенечке и о рождении второго сына. Аркадий представляет отца нового друга - Евгения Базарова, нигилиста и студента-медика, которого просит разрешить пожить в Марьино.", OrderNum: 1, PageStart: intPtr(1), PageEnd: intPtr(45), MoodTags: []string{"знакомство", "семья", "ожидание"}, AverageRating: float64Ptr(8.2)},
		{ID: "bp0000000-0000-0000-0000-000000000002", BookID: "b0000000-0000-0000-0000-000000000001", Title: "Глава 2. Спор о нигилизме", Content: "Базаров знакомится с бытом Марьино и вступает в философские споры с Павлом Петровичем Кирсановым, дядей Аркадия. Базаров отстаивает свои нигилистические взгляды - отрицание авторитетов, искусств, романтических чувств. Павел Петрович, аристократ до мозга костей, не может принять идеи Базарова. Споры становятся все более острыми, раскрывая фундаментальные различия в мировоззрении двух поколений. Аркадий пытается найти компромисс, но все больше склоняется на сторону Базарова.", OrderNum: 2, PageStart: intPtr(46), PageEnd: intPtr(89), MoodTags: []string{"философия", "конфликт", "идеи"}, AverageRating: float64Ptr(8.8)},
		{ID: "bp0000000-0000-0000-0000-000000000003", BookID: "b0000000-0000-0000-0000-000000000001", Title: "Глава 3. В городе", Content: "Базаров и Аркадий отправляются в губернский город, где знакомятся с местным обществом. Они посещают бал у губернатора, где Базаров ведет себя вызывающе и иронично. Там же они встречают Анну Сергеевну Одинцову, красивую и умную вдову. Базаров, несмотря на свои теоретические убеждения, испытывает к ней интерес. Начинается сложная психологическая игра между ними - Одинцова intriguered by his intelligence and confidence, while Базаров struggles with his unexpected feelings.", OrderNum: 3, PageStart: intPtr(90), PageEnd: intPtr(134), MoodTags: []string{"светская жизнь", "любовь", "ревность"}, AverageRating: float64Ptr(8.5)},
		{ID: "bp0000000-0000-0000-0000-000000000004", BookID: "b0000000-0000-0000-0000-000000000001", Title: "Глава 4. Дуэль", Content: "После возвращения в Марьино напряжение между Базаровым и Павлом Петровичем достигает предела. Павел Петрович, оскорбленный насмешками Базарова и его отношением к Фенечке, вызывает его на дуэль. Дуэль проходит в роще - Павел Петрович легко ранен в ногу. Этот incident становится поворотным моментом. Базаров проявляет неожиданное благородство, оказывая первую помощь противнику. После дуэли он понимает, что должен покинуть Марьино.", OrderNum: 4, PageStart: intPtr(135), PageEnd: intPtr(178), MoodTags: []string{"трагедия", "честь", "потеря"}, AverageRating: float64Ptr(9.1)},
		{ID: "bp0000000-0000-0000-0000-000000000005", BookID: "b0000000-0000-0000-0000-000000000001", Title: "Глава 5. Прощание", Content: "Базаров уезжает в родительский дом. Перед отъездом он прощается со всеми, но особенно трогательным становится его прощание с Аркадием. Их дружба проходит проверку - Аркадий все больше отдаляется от нигилизма и возвращается к традиционным ценностям. Базаров чувствует свое одиночество и понимает, что его теория не дает ему счастья. Он возвращается к родителям, пытаясь найти утешение в работе и семейном тепле.", OrderNum: 5, PageStart: intPtr(179), PageEnd: intPtr(220), MoodTags: []string{"расставание", "одиночество", "размышления"}, AverageRating: float64Ptr(8.7)},
		{ID: "bp0000000-0000-0000-0000-000000000006", BookID: "b0000000-0000-0000-0000-000000000001", Title: "Эпилог", Content: "Проходит несколько месяцев. Базаров, работая врачом, случайно заражается тифом и умирает. Перед смертью он просит отца послать за Анной Одинцовой, но она приезжает уже после его смерти. На могиле Базарова стоят только его старые родители. Аркадий женится на Кате, сестре Анны, и счастливо живет в Марьино. Николай Петрович женится на Фенечке. Павел Петрович уезжает за границу. Жизнь продолжается, но память о Базарове остается как символ трагической судьбы человека, опередившего свое время.", OrderNum: 6, PageStart: intPtr(221), PageEnd: intPtr(320), MoodTags: []string{"итог", "время", "память"}, AverageRating: float64Ptr(8.9)},

		// Преступление и наказание
		{ID: "bp0000000-0000-0000-0000-000000000007", BookID: "b0000000-0000-0000-0000-000000000002", Title: "Часть первая. Замысел", Content: "Санкт-Петербург, июль 1865 года. Бывший студент Родион Раскольников живет в крошечной каморке в бедном районе. Он размышляет о своем плане убить старуху-процентщицу Алёну Ивановну, чтобы забрать ее деньги и спасти мать и сестру от нищеты. Раскольников разрабатывает теорию о делении людей на 'обыкновенных' и 'необыкновенных', которым позволено переступать через закон. Он знакомится с Семеном Мармеладовым, пьющим чиновником, и узнает о трагической судьбе его семьи, особенно его дочери Сони.", OrderNum: 1, PageStart: intPtr(1), PageEnd: intPtr(95), MoodTags: []string{"бедность", "план", "тревога"}, AverageRating: float64Ptr(9.0)},
		{ID: "bp0000000-0000-0000-0000-000000000008", BookID: "b0000000-0000-0000-0000-000000000002", Title: "Часть вторая. Преступление", Content: "Раскольников осуществляет свой план. Он приходит к старухе-процентщице под предлогом закладки вещей и убивает ее топором. Возвращается ее сестра Лизавета, и Раскольников в панике убивает и ее. Он успевает взять немного ценностей и скрыться. После убийства его мучают лихорадка и кошмары. Он почти не помнит деталей преступления, но чувствует непреодолимое отвращение к себе. Порфирий Петрович, следователь, начинает расследование, и Раскольников понимает, что его могут вычислить.", OrderNum: 2, PageStart: intPtr(96), PageEnd: intPtr(190), MoodTags: []string{"насилие", "ужас", "раскаяние"}, AverageRating: float64Ptr(9.3)},
		{ID: "bp0000000-0000-0000-0000-000000000009", BookID: "b0000000-0000-0000-0000-000000000002", Title: "Часть третья. Расследование", OrderNum: 3, PageStart: intPtr(191), PageEnd: intPtr(285), MoodTags: []string{"подозрение", "психология", "напряжение"}, AverageRating: float64Ptr(9.1)},
		{ID: "bp0000000-0000-0000-0000-000000000010", BookID: "b0000000-0000-0000-0000-000000000002", Title: "Часть четвертая. Страдания", OrderNum: 4, PageStart: intPtr(286), PageEnd: intPtr(380), MoodTags: []string{"муки", "совесть", "боль"}, AverageRating: float64Ptr(9.2)},
		{ID: "bp0000000-0000-0000-0000-000000000011", BookID: "b0000000-0000-0000-0000-000000000002", Title: "Часть пятая. Искупление", OrderNum: 5, PageStart: intPtr(381), PageEnd: intPtr(475), MoodTags: []string{"надежда", "любовь", "покаяние"}, AverageRating: float64Ptr(9.4)},
		{ID: "bp0000000-0000-0000-0000-000000000012", BookID: "b0000000-0000-0000-0000-000000000002", Title: "Эпилог. Возрождение", OrderNum: 6, PageStart: intPtr(476), PageEnd: intPtr(560), MoodTags: []string{"искупление", "новая жизнь", "свет"}, AverageRating: float64Ptr(9.0)},

		// Мастер и Маргарита
		{ID: "bp0000000-0000-0000-0000-000000000013", BookID: "b0000000-0000-0000-0000-000000000003", Title: "Часть первая. Понтий Пилат", Content: "Москва, 1930-е годы. На Патриарших прудах встречаются редактор журнала Михаил Берлиоз и поэт Иван Бездомный. Они знакомятся с загадочным иностранцем, который оказывается Воландом - самим Дьяволом. Воланд предсказывает Берлиозу gruesome death, которая тут же сбывается. Параллельно разворачивается история в древнем Ершалаиме, где прокуратор Иудеи Понтий Пилат судит Иешуа Га-Ноцри. Пилат понимает невиновность Иешуа, но, боясь потерять власть, отправляет его на казнь.", OrderNum: 1, PageStart: intPtr(1), PageEnd: intPtr(68), MoodTags: []string{"мистика", "история", "власть"}, AverageRating: float64Ptr(9.2)},
		{ID: "bp0000000-0000-0000-0000-000000000014", BookID: "b0000000-0000-0000-0000-000000000003", Title: "Часть вторая. Воланд и свита", Content: "Воланд и его свита - кот Бегемот, Коровьев-Фагот и Азазелло - поселяются в квартире профессора психиатрии Стравинского. Они начинают устраивать в Москве хаос: устраивают сеанс черной магии в варьете, превращают управляющего дома Никанора Ивановича в контрабандиста, преследуют председателя жилтоварищества Берлиоза. Иван Бездомный, пытаясь разоблачить Воланда, попадает в психиатрическую лечебницу, где встречает Мастера - автора романа о Понтии Пилате.", OrderNum: 2, PageStart: intPtr(69), PageEnd: intPtr(136), MoodTags: []string{"сатира", "магия", "хаос"}, AverageRating: float64Ptr(9.5)},
		{ID: "bp0000000-0000-0000-0000-000000000015", BookID: "b0000000-0000-0000-0000-000000000003", Title: "Часть третья. Мастер", OrderNum: 3, PageStart: intPtr(137), PageEnd: intPtr(204), MoodTags: []string{"творчество", "любовь", "безумие"}, AverageRating: float64Ptr(9.3)},
		{ID: "bp0000000-0000-0000-0000-000000000016", BookID: "b0000000-0000-0000-0000-000000000003", Title: "Часть четвертая. Маргарита", OrderNum: 4, PageStart: intPtr(205), PageEnd: intPtr(272), MoodTags: []string{"преданность", "сила", "магия"}, AverageRating: float64Ptr(9.4)},
		{ID: "bp0000000-0000-0000-0000-000000000017", BookID: "b0000000-0000-0000-0000-000000000003", Title: "Часть пятая. Бал у Сатаны", OrderNum: 5, PageStart: intPtr(273), PageEnd: intPtr(340), MoodTags: []string{"пир", "тайна", "воскрешение"}, AverageRating: float64Ptr(9.6)},
		{ID: "bp0000000-0000-0000-0000-000000000018", BookID: "b0000000-0000-0000-0000-000000000003", Title: "Эпилог. Покой", OrderNum: 6, PageStart: intPtr(341), PageEnd: intPtr(410), MoodTags: []string{"гармония", "вечность", "мир"}, AverageRating: float64Ptr(9.1)},

		// Отверженные
		{ID: "bp0000000-0000-0000-0000-000000000019", BookID: "b0000000-0000-0000-0000-000000000004", Title: "Том первый. Каторжник", OrderNum: 1, PageStart: intPtr(1), PageEnd: intPtr(207), MoodTags: []string{"несправедливость", "милосердие", "страдания"}, AverageRating: float64Ptr(8.8)},
		{ID: "bp0000000-0000-0000-0000-000000000020", BookID: "b0000000-0000-0000-0000-000000000004", Title: "Том второй. Козетта", OrderNum: 2, PageStart: intPtr(208), PageEnd: intPtr(414), MoodTags: []string{"невинность", "забота", "надежда"}, AverageRating: float64Ptr(8.9)},
		{ID: "bp0000000-0000-0000-0000-000000000021", BookID: "b0000000-0000-0000-0000-000000000004", Title: "Том третий. Мариус", OrderNum: 3, PageStart: intPtr(415), PageEnd: intPtr(621), MoodTags: []string{"юность", "любовь", "идеалы"}, AverageRating: float64Ptr(8.7)},
		{ID: "bp0000000-0000-0000-0000-000000000022", BookID: "b0000000-0000-0000-0000-000000000004", Title: "Том четвертый. Идилла улицы Плюмер", OrderNum: 4, PageStart: intPtr(622), PageEnd: intPtr(828), MoodTags: []string{"счастье", "семья", "спокойствие"}, AverageRating: float64Ptr(9.0)},
		{ID: "bp0000000-0000-0000-0000-000000000023", BookID: "b0000000-0000-0000-0000-000000000004", Title: "Том пятый. Жан Вальжан", OrderNum: 5, PageStart: intPtr(829), PageEnd: intPtr(1035), MoodTags: []string{"жертва", "долг", "искупление"}, AverageRating: float64Ptr(9.2)},
		{ID: "bp0000000-0000-0000-0000-000000000024", BookID: "b0000000-0000-0000-0000-000000000004", Title: "Том шестой. Барьер", OrderNum: 6, PageStart: intPtr(1036), PageEnd: intPtr(1240), MoodTags: []string{"революция", "героизм", "память"}, AverageRating: float64Ptr(9.1)},

		// Норвежский лес
		{ID: "bp0000000-0000-0000-0000-000000000025", BookID: "b0000000-0000-0000-0000-000000000005", Title: "Глава 1. Воспоминания", OrderNum: 1, PageStart: intPtr(1), PageEnd: intPtr(64), MoodTags: []string{"ностальгия", "прошлое", "тишина"}, AverageRating: float64Ptr(8.3)},
		{ID: "bp0000000-0000-0000-0000-000000000026", BookID: "b0000000-0000-0000-0000-000000000005", Title: "Глава 2. Наоко", OrderNum: 2, PageStart: intPtr(65), PageEnd: intPtr(128), MoodTags: []string{"любовь", "хрупкость", "глубина"}, AverageRating: float64Ptr(8.5)},
		{ID: "bp0000000-0000-0000-0000-000000000027", BookID: "b0000000-0000-0000-0000-000000000005", Title: "Глава 3. Мидори", OrderNum: 3, PageStart: intPtr(129), PageEnd: intPtr(192), MoodTags: []string{"жизнь", "энергия", "противоречия"}, AverageRating: float64Ptr(8.2)},
		{ID: "bp0000000-0000-0000-0000-000000000028", BookID: "b0000000-0000-0000-0000-000000000005", Title: "Глава 4. Санаторий", OrderNum: 4, PageStart: intPtr(193), PageEnd: intPtr(256), MoodTags: []string{"болезнь", "уединение", "размышления"}, AverageRating: float64Ptr(8.6)},
		{ID: "bp0000000-0000-0000-0000-000000000029", BookID: "b0000000-0000-0000-0000-000000000005", Title: "Глава 5. Выбор", OrderNum: 5, PageStart: intPtr(257), PageEnd: intPtr(320), MoodTags: []string{"решение", "взросление", "потеря"}, AverageRating: float64Ptr(8.8)},
		{ID: "bp0000000-0000-0000-0000-000000000030", BookID: "b0000000-0000-0000-0000-000000000005", Title: "Эпилог. Прощание", OrderNum: 6, PageStart: intPtr(321), PageEnd: intPtr(384), MoodTags: []string{"принятие", "память", "жизнь"}, AverageRating: float64Ptr(8.4)},
	}

	for _, book := range books {
		err := insertBook(ctx, pool, &book)
		if err != nil {
			return fmt.Errorf("ошибка вставки книги %s: %w", book.Title, err)
		}
		fmt.Printf("Книга '%s' успешно добавлена\n", book.Title)
	}

	for _, part := range bookParts {
		err := insertBookPart(ctx, pool, &part)
		if err != nil {
			return fmt.Errorf("ошибка вставки части книги %s: %w", part.Title, err)
		}
		fmt.Printf("Часть книги '%s' успешно добавлена\n", part.Title)
	}

	return nil
}

func insertBook(ctx context.Context, pool *pgxpool.Pool, book *models.Book) error {
	query := `
		INSERT INTO books (
			id, title, original_title, author, year, genres, age_rating,
			author_country, description, cover_url, pages_count, tags,
			verified, verification_type, created_by, created_at,
			average_rating, rating_count
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18
		) ON CONFLICT (id) DO NOTHING`

	_, err := pool.Exec(ctx, query,
		book.ID,
		book.Title,
		book.OriginalTitle,
		book.Author,
		book.Year,
		book.Genres,
		book.AgeRating,
		book.AuthorCountry,
		book.Description,
		book.CoverURL,
		book.PagesCount,
		book.Tags,
		book.Verified,
		book.VerificationType,
		book.CreatedBy,
		book.CreatedAt,
		book.AverageRating,
		book.RatingCount,
	)

	return err
}

func insertBookPart(ctx context.Context, pool *pgxpool.Pool, part *models.BookPart) error {
	query := `
		INSERT INTO book_parts (
			id, book_id, title, content, order_num, page_start, page_end, mood_tags, average_rating
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		) ON CONFLICT (id) DO NOTHING`

	_, err := pool.Exec(ctx, query,
		part.ID,
		part.BookID,
		part.Title,
		part.Content,
		part.OrderNum,
		part.PageStart,
		part.PageEnd,
		part.MoodTags,
		part.AverageRating,
	)

	return err
}

// Вспомогательные функции для создания указателей
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func float64Ptr(f float64) *float64 {
	return &f
}

func ageRatingPtr(ar models.AgeRating) *models.AgeRating {
	return &ar
}

func verificationTypePtr(vt models.VerificationType) *models.VerificationType {
	return &vt
}
