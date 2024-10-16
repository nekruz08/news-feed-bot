# News Feed Bot
Этот бот для Telegram, написанный на Go, позволяет пользователям получать обновления новостей из различных источников. Бот использует PostgreSQL для хранения данных и поддерживает управление источниками новостей через команды для администраторов.

#  Функциональность
Добавление, удаление и управление источниками новостей
Периодическое получение новостей и уведомление пользователей
Проверка работоспособности через HTTP
Использование
Настройте токен бота и параметры базы данных в конфигурационном файле.
Запустите бота и наслаждайтесь обновлениями новостей!


# DOCKER
docker-compose -f docker-compose.yml down -v    	
docker-compose -f docker-compose.yml up        	 	
docker-compose -f docker-compose.yml ps         	


# GOOSE
goose postgres "host=localhost port=5433 user=postgres dbname=news_feed_bot password=postgres sslmode=disable" status
goose postgres "host=localhost port=5433 user=postgres dbname=news_feed_bot password=postgres sslmode=disable" up
goose postgres "host=localhost port=5433 user=postgres dbname=news_feed_bot password=postgres sslmode=disable" down

# BOT COMMANDS																	
/addsource {"name":"#golang on hashnode.dev", "url": "https://hashnode.com/n/golang/rss"}

/addsource {"name":"#golang on dev.to", "url": "https://dev.to/feed/tag/golang"}

/listsources

/getsource 1    существующий

/getsource 1    несуществующий

/getsource abc  неверный формат ввода

/setpriority {"source_id": 1, "priority": 5}

/deletesource 123										

