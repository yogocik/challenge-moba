package main

import (
	. "moba/creature"
)

func main() {
	var hero *Hero = CreateHero("HERO", 1000, 50)
	var hero2 *Hero = CreateHero("HERO2", 1000, 50)
	var monster *Monster = CreateMonster("Monster", 100, 20)
	var monster2 *Monster = CreateMonster("Monster2", 100, 20)
}
