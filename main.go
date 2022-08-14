package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
)

var (
	// глобальная переменная в которой храним токен
	telegramBotToken = "///"
)

func main() {
	// используя токен создаем новый инстанс бота
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// u - структура с конфигом для получения апдейтов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates, err := bot.GetUpdatesChan(u)

	// в канал updates прилетают структуры типа Update
	// вычитываем их и обрабатываем
	for update := range updates {
		// универсальный ответ на любое сообщение
		reply := "Используйте команду /help"
		if update.Message == nil {
			continue
		}

		// логируем от кого какое сообщение пришло
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// свитч на обработку комманд
		// комманда - сообщение, начинающееся с "/"
		switch update.Message.Command() {
		case "help":
			reply = "При регистрации и создании аккаунта сотрудником безопасности ФИО следует вводить русскими буквами. Поле 'Отчество' может быть пустым.\n" +
				"Поле 'Серия паспорта' должно содержать строго 4 цифры, поле 'Номера паспорта' должно содержать 6 цифр.\n" +
				"Пароль должен содержать только цифры.\n" +
				"Логин должен содержать только английские буквы."
		case "details":
			reply = "Кошелек создается автоматически при регистрации.\n" +
				"После регистрации Вам следует войти в ваш аккаунт, где содержится информация о балансе и реквизитах кошелька."
		case "start":
			reply = "Команды бота: '/help' - информация по использованию\n '/details' - детли программы\n '/start' - команды бота"
		}

		// создаем ответное сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		// отправляем
		bot.Send(msg)
	}
}
