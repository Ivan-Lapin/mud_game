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
	mision      string
}

type Player struct {
	itmes []string
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
		description: "ты находишься на кухне, ",
		items: []map[string][]string{
			{
				"на столе": {"чай"},
			},
		},
		motions: []string{"коридор"},
		mision:  "надо собрать рюкзак и идти в универ",
	}

	hall := Location{
		name:        "коридор",
		description: "ничего интересного, ",
		items:       nil,
		motions:     []string{"кухня", "комната", "улица"},
	}

	room := Location{
		name:        "комната",
		description: "ты в своей комнате, ",
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
		description: "на улице весна, ",
		motions:     []string{"домой"},
	}

	Locations = []Location{kitchen, hall, room, outside}

	w = World{
		Player:          Player{},
		CurrentLocation: kitchen,
	}

}

func emptyMap(items []map[string][]string) bool {
	for _, item := range items {
		if len(item) > 0 {
			return true
		}
	}
	return false
}

func deleteItem(items *[]map[string][]string, name string) {

	for _, itemMap := range *items {
		for key, value := range itemMap {
			var newValue []string
			found := false

			for _, item := range value {
				if item != name {
					newValue = append(newValue, item)
				} else {
					found = true
				}
			}

			if found {
				if len(newValue) == 0 {
					delete(itemMap, key)
				} else {
					itemMap[key] = newValue
				}
			}
		}
	}

}

func lookCommand() string {
	answer := w.CurrentLocation.description
	fmt.Printf("Items: %v\n", w.CurrentLocation.items)
	if emptyMap(w.CurrentLocation.items) == false {
		if w.CurrentLocation.name == "коридор" {
			answer += fmt.Sprintf(", ")
		} else {
			answer += fmt.Sprintf("пустая комната, ")
		}
	} else {
		for _, maps := range w.CurrentLocation.items {
			for place, items := range maps {
				answer += fmt.Sprintf("%s: %s, ", place, strings.Join(items, ", "))
			}
		}
	}
	if len(w.CurrentLocation.mision) != 0 {
		answer += fmt.Sprintf("%s, ", w.CurrentLocation.mision)
	}
	answer += fmt.Sprintf("можно пройти - %s", strings.Join(w.CurrentLocation.motions, ", "))
	return answer

}

func goCommand(place string) string {
	flag := false
	for _, motion := range w.CurrentLocation.motions {
		if motion == place {
			flag = true
			break
		}
	}
	if flag == true {
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
		answer += "ничего интересного, "
	}
	answer += fmt.Sprintf("можно пройти - %s", strings.Join(w.CurrentLocation.motions, ", "))
	return answer
}

func donCommand(obj string) string {
	flag := false
	for _, maps := range w.CurrentLocation.items {
		for _, items := range maps {
			if slices.Contains(items, obj) {
				flag = true
				deleteItem(&w.CurrentLocation.items, obj)
			}
		}
	}
	if flag == true {
		w.Player.itmes = append(w.Player.itmes, obj)
		answer := fmt.Sprintf("вы надели: %s", obj)
		return answer
	}
	return "нет такого"
}

func getCommand(obj string) string {

	if slices.Contains(w.Player.itmes, "рюкзак") == false {
		return "некуда класть"
	}

	flag := false
	for _, maps := range w.CurrentLocation.items {
		for _, items := range maps {
			if slices.Contains(items, obj) {
				flag = true
				deleteItem(&w.CurrentLocation.items, obj)
			}
		}
	}
	if flag == true {
		answer := fmt.Sprintf("предмет добавлен в инвентарь: %s", obj)
		return answer
	}
	return "нет такого"
}

func applyCommand(obj string, sub string) string {
	answer := ""
	if slices.Contains(w.Player.itmes, obj) == false {
		answer = fmt.Sprintf("нет предмета в инвентаре - %s", obj)
		return answer
	} else {
		if obj == "ключи" && sub == "дверь" {
			answer = fmt.Sprintf("%s открыта", sub)
		} else {
			answer = "не к чему применить"
		}
	}
	return ""
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
