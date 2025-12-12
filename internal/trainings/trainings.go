package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Метод парсит строку с данными формата "3456,Ходьба,3h00m" и записывает данные в соответствующие поля структуры Training.
func (t *Training) Parse(datastring string) (err error) {
	//1.	Разделить строку datastring на слайс строк.
	//2.	Проверить, чтобы длина слайса была равна 3, так как в строке данных у нас количество шагов, вид активности и продолжительность.
	//3.	Преобразовать первый элемент слайса (количество шагов) в тип int. Обработать возможные ошибки.
	//		При возникновении ошибки вернуть её из метода. Сохранить полученное значение в соответствующем поле структуры Training.
	//4.	Сохранить значение типа тренировки в соответствующем поле структуры Training.
	//5.	Преобразовать третий элемент слайса в time.Duration. В пакете time есть метод для парсинга строки в time.Duration.
	//		Обработать возможные ошибки. При их возникновении вернуть ошибку,
	//		в противном случае сохранить полученное значение в соответствующем поле структуры Training.
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("неверный формат строки")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("неверное количество шагов")
	}
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return err
	}

	if duration <= 0 {
		return errors.New("неверная продолжительность")
	}
	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = duration

	return nil
}

//  1. Вычислить дистанцию, используя функцию из пакета spentenergy.
//  2. Вычислить среднюю скорость, используя функцию из пакета spentenergy.
//  3. Проверить, какой вид тренировки содержится в структуре Training.
//     Для каждого из видов тренировок рассчитать калории, используя функцию из пакета spentenergy.
//  4. Сформируйте и верните строку, образец которой был выше.
//  5. Если был передан неизвестный тип тренировки, верните ошибку с текстом неизвестный тип тренировки.
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(
			t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(
			t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calories,
	), nil
}
