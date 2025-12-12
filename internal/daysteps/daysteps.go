package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Метод парсит строку с данными формата "678,0h50m" и записывает данные в соответствующие поля структуры DaySteps.
func (ds *DaySteps) Parse(datastring string) (err error) {
	//Метод принимает строку формата "678,0h50m". Мы используем указатель на структуру (ds *DaySteps), так как будем записывать в неё данные.
	//Метод возвращает ошибку.

	//Подсказка: Используйте полученные знания и навыки при реализации подобного метода в пакете trainings и при выполнении задания 4 спринта.

	//1.	Разделить строку datastring на слайс строк.
	//2.	Проверить, чтобы длина слайса была равна 2, так как в строке данных у нас количество шагов и продолжительность.
	//3.	Преобразовать первый элемент слайса (количество шагов) в тип int. Обработать возможные ошибки.
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("неверный формат строки")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("неверное количество шагов")
	}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}

	if duration <= 0 {
		return errors.New("неверная продолжительность")
	}
	ds.Steps = steps
	ds.Duration = duration

	return nil
}

//Метод ничего не принимает, так как все данные содержатся в структуре DaySteps. Метод должен возвращать строку вида:

// Количество шагов: 792.
// Дистанция составила 0.51 км.
// Вы сожгли 221.33 ккал.
func (ds DaySteps) ActionInfo() (string, error) {
	//1.	Вычислите дистанцию.
	//2.	Вычислите количество сожжённых калорий. При возникновении ошибки верните пустую строку и ошибку.
	//3.	Сформируйте и верните строку с информацией.
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	), nil
}
