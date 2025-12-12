package actioninfo

import "log"

//Создайте интерфейс DataParser, в котором объявите сигнатуры методов Parse() и ActionInfo().
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

//Функция принимает слайс строк с данными о тренировках или прогулках и экземпляр одной из ваших структур Training или DaySteps
//Это возможно, потому что они обе реализуют интерфейс DataParser,
//то есть для каждой из этих структур вы реализовали методы, которые описаны в интерфейсе.
func Info(dataset []string, dp DataParser) {
	//Перебрать все значения слайса dataset в цикле.
	//Распарсить каждое значение с помощью метода Parse().
	//Обработать ошибку парсинга. Если она возникает, нужно залогировать ошибку и перейти к следующей итерации цикла.
	//Сформировать и вывести строку с информацией об активности с помощью метода ActionInfo(). При возникновении ошибки ее нужно залогировать.
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Println("Ошибка парсинга:", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("Ошибка формирования:", err)
			continue
		}

		log.Println(info)
	}
}
