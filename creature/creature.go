package creature

import "fmt"

//////////////// Model Struct
type Hero struct {
	name        string
	Hp          int
	baseDamage  int
	finalDamage int
	status      string
	weapon      Weapon
	item        Item
	effect      Effect
}

type Monster struct {
	name        string
	Hp          int
	baseDamage  int
	finalDamage int
	status      string
	weapon      Weapon
	item        Item
	effect      Effect
}

type Soldier struct {
	name        string
	Hp          int
	baseDamage  int
	finalDamage int
	status      string
	weapon      Weapon
	item        Item
	effect      Effect
}

type Tower struct {
	name        string
	Hp          int
	baseDamage  int
	finalDamage int
	status      string
	weapon      Weapon
	item        Item
	effect      Effect
}

type Weapon struct {
	name   string
	damage int
}

type Effect struct {
	passiveDamage   int
	attackReduction int
}

type Item struct {
	name       string
	attackUp   int
	attackNull int
}

type BasicInfo struct {
	name        string
	Hp          int
	status      string
	weaponName  string
	itemName    string
	finalDamage int
}

////////////// Interface

type HeroTarget interface {
	ReceiveHeroDamage(target *Hero)
}

////////////// Functions
func CreateHero(name string, hp int, damage int) *Hero {
	return &Hero{name: name, Hp: hp, baseDamage: damage, finalDamage: damage, status: "alive"}
}
func CreateMonster(name string, hp int, damage int) *Monster {
	return &Monster{name: name, Hp: hp, baseDamage: damage, finalDamage: damage, status: "alive"}
}
func CreateSoldier(name string, hp int, damage int) *Soldier {
	return &Soldier{name: name, Hp: hp, baseDamage: damage, finalDamage: damage, status: "alive"}
}
func CreateTower(name string, hp int, damage int) *Tower {
	return &Tower{name: name, Hp: hp, baseDamage: damage, status: "alive"}
}
func CreateItem(name string, attackIncrease int, attackNulify int) *Item {
	return &Item{name: name, attackUp: attackIncrease, attackNull: attackNulify}
}
func CreateWeapon(name string, damage int) *Weapon {
	return &Weapon{name: name, damage: damage}
}

func changeStatus(hp *int, status *string) {
	if *hp <= 0 {
		*hp = 0
		*status = "defeated"
	}
}

////////////// Hero Methods

func (hero *Hero) sumAttack() int {
	hero.finalDamage += hero.effect.passiveDamage
	return hero.finalDamage
}
func (hero *Hero) ReceiveHeroDamage(opponent *Hero) {
	hero.Hp -= (opponent.sumAttack() - hero.effect.attackReduction)
	changeStatus(&hero.Hp, &hero.status)
}

func (monster *Monster) ReceiveHeroDamage(opponent *Hero) {
	monster.Hp -= (opponent.sumAttack() - monster.effect.attackReduction)
	changeStatus(&monster.Hp, &monster.status)
}

func (soldier *Soldier) ReceiveHeroDamage(opponent *Hero) {
	soldier.Hp -= (opponent.sumAttack() - soldier.effect.attackReduction)
	changeStatus(&soldier.Hp, &soldier.status)
}

func (tower *Tower) ReceiveHeroDamage(opponent *Hero) {
	tower.Hp -= opponent.finalDamage
	opponent.Hp -= tower.baseDamage
}

func (hero *Hero) UseItem(items *Item) {
	hero.finalDamage += items.attackUp - hero.item.attackUp
	hero.effect.attackReduction = items.attackNull - hero.item.attackNull
	hero.item = *items
}

func (hero *Hero) AttachWeapon(weapons *Weapon) {
	hero.finalDamage += weapons.damage - hero.weapon.damage
	hero.weapon = *weapons
}

func (hero *Hero) Punch(target HeroTarget) {
	target.ReceiveHeroDamage(hero)
}

func (hero *Hero) GetBasicInfo() BasicInfo {
	return BasicInfo{name: hero.name, Hp: hero.Hp, status: hero.status,
		weaponName: hero.weapon.name, itemName: hero.item.name, finalDamage: hero.finalDamage}
}

func (hero *Hero) ShowBasicInfo() {
	fmt.Printf("name: %v, Hp: %v, status: %v, weaponName: %v, itemName: %v, finalDamage: %v\n",
		hero.name, hero.Hp, hero.status, hero.weapon.name, hero.item.name, hero.finalDamage)
}
