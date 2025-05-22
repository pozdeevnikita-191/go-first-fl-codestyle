/ Вы познакомились с основными правилами документирования исходного кода.
// Ниже будут отображаться заголовок и ссылки.
//
// Смотрите также
//
// https://pkg.go.dev/
//
// https://elliotchance.medium.com/godoc-tips-tricks-cda6571549b  
//
package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// ___________RANDINT_____________________________________________________________________________________________
func randint(min, max int) int {
	return rand.Intn(max-min) + min
}

// ___________ATTACK_____________________________________________________________________________________________
func attack(name, class string) string {
	switch class {
	case "warrior":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", name, 5+randint(3, 5))
	case "mage":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", name, 5+randint(5, 10))
	case "healer":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", name, 5+randint(-3, -1))
	}
	return "неизвестный класс персонажа"
}

// ___________DEFENCE_____________________________________________________________________________________________
// обратите внимание на "if else" и на "else"
func defence(name, class string) string {
	if class == "warrior" {
		return fmt.Sprintf("%s блокировал %d урона.", name, 10+randint(5, 10))
	} else if class == "mage" {
		return fmt.Sprintf("%s блокировал %d урона.", name, 10+randint(-2, 2))
	} else if class == "healer" {
		return fmt.Sprintf("%s блокировал %d урона.", name, 10+randint(2, 5))
	}
	return "неизвестный класс персонажа"

}

//___________SPECIAL_____________________________________________________________________________________________

// обратите внимание на "if else" и на "else"
func special(name, class string) string {
	if class == "warrior" {
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", name, 80+25)
	}

	if class == "mage" {
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", name, 5+40)
	}

	if class == "healer" {
		return fmt.Sprintf("%s применил специальное умение `Защита %d`", name, 10+30)
	}
	return "неизвестный класс персонажа"
}

//___________START TRAINING_____________________________________________________________________________________________

// здесь обратите внимание на имена параметров
func training(name, class string) string {
	switch class {
	case "warrior":
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", name)
	case "mage":
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", name)
	case "healer":
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", name)
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)
		switch cmd {
		case "attack":
			fmt.Println(attack(name, class))
		case "defence":
			fmt.Println(defence(name, class))
		case "special":
			fmt.Println(special(name, class))
		default:
			fmt.Println("неизвестная команда")
		}
	}
	return "тренировка окончена"
}

//___________CHOISE_CHAR_CLASS__________________________________________________________________________________

// обратите внимание на имя функции и имена переменных
func choisClass() string {
	var approveChoice string
	var charClass string

	for approveChoice != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &charClass)
		if charClass == "warrior" {
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		} else if charClass == "mage" {
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		} else if charClass == "healer" {
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		}
		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &approveChoice)
		approveChoice = strings.ToLower(approveChoice)
	}
	return charClass
}

//___________MAIN_____________________________________________________________________________________________

// обратите внимание на имена переменных
func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var name string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &name)

	fmt.Printf("Здравствуй, %s\n", name)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	class := choisClass()

	fmt.Println(training(name, class))
}
