package main

import (
	"fmt"
	"slices"
	"strings"
)

type Location struct {
	items       []map[string][]string
	motions     []string
	name        string
	description string
}

type Player struct {
	items    []string
	missions []string
}

type World struct {
	Player          Player
	CurrentLocation Location
}

var w = World{}
var Locations = []Location{}

func initGame() {
	kitchen := Location{
		name:        "кухня",
		description: "ты находишься на кухне",
		items: []map[string][]string{
			{
				"на столе": {"чай"},
			},
		},
		motions: []string{"коридор"},
	}

	hall := Location{
		name:        "коридор",
		description: "ничего интересного",
		items:       nil,
		motions:     []string{"кухня", "комната", "улица"},
	}

	room := Location{
		name:        "комната",
		description: "ты в своей комнате",
		items: []map[string][]string{
			{
				"на столе": {"ключи", "конспекты"},
			},
			{
				"на стуле": {"рюкзак"},
			},
		},
		motions: []string{"коридор"},
	}

	outside := Location{
		name:        "улица",
		description: "на улице весна",
		motions:     []string{"домой"},
	}

	Locations = []Location{kitchen, hall, room, outside}

	w = World{
		Player: Player{
			missions: []string{"собрать рюкзак", "идти в универ"},
		},
		CurrentLocation: kitchen,
	}

}

func isMapEmpty(items []map[string][]string) bool {
	for _, item := range items {
		if len(item) > 0 {
			return false
		}
	}
	return true
}

func deleteItem(items *[]map[string][]string, name string) {

	for _, itemPlaces := range *items {
		for place, items := range itemPlaces {
			var newValue []string
			found := false

			for _, item := range items {
				if item != name {
					newValue = append(newValue, item)
				} else {
					found = true
				}
			}

			if found {
				if len(newValue) == 0 {
					delete(itemPlaces, place)
				} else {
					itemPlaces[place] = newValue
				}
			}
		}
	}

}

// Вспомогательная функция для соединения массива строк с использованием союза
func joinWithConjunction(items []string, conjunction string) string {
	if len(items) == 0 {
		return ""
	}
	if len(items) == 1 {
		return items[0]
	}

	lastIndex := len(items) - 1
	joined := strings.Join(items[:lastIndex], ", ")

	return fmt.Sprintf("%s %s %s", joined, conjunction, items[lastIndex])
}

func lookCommand() string {
	answer := w.CurrentLocation.description

	if !isMapEmpty(w.CurrentLocation.items) {
		for _, itemMap := range w.CurrentLocation.items {
			for place, items := range itemMap {
				answer += fmt.Sprintf(", %s: %s", place, strings.Join(items, ", "))
			}
		}
	} else if w.CurrentLocation.name != "коридор" {
		answer += ", пустая комната"
	}

	if w.CurrentLocation.name == "кухня" && len(w.Player.missions) > 0 {
		answer += fmt.Sprintf(", надо %s", joinWithConjunction(w.Player.missions, "и"))
	}

	if len(w.CurrentLocation.motions) > 0 {
		answer += fmt.Sprintf(", можно пройти - %s", strings.Join(w.CurrentLocation.motions, ", "))
	}

	return answer
}

func goCommand(place string) string {
	pathExist := false
	for _, motion := range w.CurrentLocation.motions {
		if motion == place {
			pathExist = true
			break
		}
	}

	if pathExist == true {
		for _, location := range Locations {
			if location.name == place {
				w.CurrentLocation = location
			}
		}
	} else {
		return "нет пути в " + fmt.Sprintf("%s", place)
	}

	answer := w.CurrentLocation.description

	if place == "кухня" {
		answer += fmt.Sprintf(", ничего интересного")
	}

	answer += fmt.Sprintf(", можно пройти - %s", strings.Join(w.CurrentLocation.motions, ", "))
	return answer
}

func donCommand(obj string) string {
	itemExist := false
	for _, maps := range w.CurrentLocation.items {
		for _, items := range maps {
			if slices.Contains(items, obj) {
				itemExist = true
				deleteItem(&w.CurrentLocation.items, obj)
			}
		}
	}
	if itemExist == true {
		w.Player.items = append(w.Player.items, obj)
		answer := fmt.Sprintf("вы надели: %s", obj)
		return answer
	}
	return "нет такого"
}

func getCommand(obj string) string {

	if slices.Contains(w.Player.items, "рюкзак") == false {
		return "некуда класть"
	}

	itemExist := false
	for _, maps := range w.CurrentLocation.items {
		for _, items := range maps {
			if slices.Contains(items, obj) {
				itemExist = true
				deleteItem(&w.CurrentLocation.items, obj)
				break
			}
		}
	}
	if itemExist == true {
		w.Player.items = append(w.Player.items, obj)
		answer := fmt.Sprintf("предмет добавлен в инвентарь: %s", obj)
		return answer
	}
	return "нет такого"
}

func applyCommand(obj string, sub string) string {
	answer := ""
	if slices.Contains(w.Player.items, obj) == false {
		answer = fmt.Sprintf("нет предмета в инвентаре - %s", obj)
		return answer
	} else {
		if obj == "ключи" && sub == "дверь" {
			answer = fmt.Sprintf("%s открыта", sub)
		} else {
			answer = "не к чему применить"
		}
	}
	return answer
}

func handleCommand(command string) string {
	parts := strings.Split(command, " ")
	if parts[0] == "осмотреться" {
		return lookCommand()
	} else if parts[0] == "идти" {
		return goCommand(parts[1])
	} else if parts[0] == "надеть" {
		return donCommand(parts[1])
	} else if parts[0] == "взять" {
		return getCommand(parts[1])
	} else if parts[0] == "применить" {
		return applyCommand(parts[1], parts[2])
	}
	return "неизвестная команда"
}

func main() {

}
