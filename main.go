package main

import (
	"fmt"
	. "moba/creature"
)

func main() {
	var hero *Hero = CreateHero("HERO", 1000, 50)
	var hero2 *Hero = CreateHero("HERO2", 1000, 50)
	var monster *Monster = CreateMonster("Monster", 100, 20)
	var monster2 *Monster = CreateMonster("Monster2", 100, 20)
	var soldier *Soldier = CreateSoldier("Soldier", 50, 10)
	var soldier2 *Soldier = CreateSoldier("Soldier2", 50, 10)
	sword := CreateWeapon("Excalibur", 200)
	elixir := CreateItem("Elixir", 50, 20)

	hero.Punch(monster)
	fmt.Println("Hero attack monster =>", monster.GetBasicInfo())
	hero.Punch(soldier)
	fmt.Println("Hero attack soldier =>", soldier.GetBasicInfo())
	hero.Punch(hero2)
	fmt.Println("Hero attack hero2 =>", hero2.GetBasicInfo())

	soldier.Punch(hero)
	fmt.Println("Soldier attack hero =>", hero.GetBasicInfo())
	soldier.Punch(soldier2)
	fmt.Println("Soldier attack soldier2 =>", soldier2.GetBasicInfo())
	soldier.Punch(monster)
	fmt.Println("Soldier attack monster =>", monster.GetBasicInfo())

	monster.Punch(hero)
	fmt.Println("Monster attack hero =>", hero.GetBasicInfo())
	monster.Punch(soldier)
	fmt.Println("Monster attack soldier =>", soldier.GetBasicInfo())
	monster.Punch(monster2)
	fmt.Println("Monster attack monster2 =>", monster2.GetBasicInfo())

	hero.AttachWeapon(sword)
	hero.UseItem(elixir)

	hero.ShowBasicInfo()
	hero.Punch(monster)
	fmt.Println("Monster after attacked by armed hero")
	monster.ShowBasicInfo()
}
