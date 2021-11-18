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
type MonsterTarget interface {
	ReceiveMonsterDamage(target *Monster)
}

type SoldierTarget interface {
	ReceiveSoldierDamage(target *Soldier)
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

////////////// Monster Methods
func (monster *Monster) sumAttack() int {
	monster.finalDamage += monster.effect.passiveDamage
	return monster.finalDamage
}

func (hero *Hero) ReceiveMonsterDamage(opponent *Monster) {
	hero.Hp -= (opponent.sumAttack() - hero.effect.attackReduction)
	changeStatus(&hero.Hp, &hero.status)
}

func (monster *Monster) ReceiveMonsterDamage(opponent *Monster) {
	monster.Hp -= (opponent.sumAttack() - monster.effect.attackReduction)
	changeStatus(&monster.Hp, &monster.status)
}

func (soldier *Soldier) ReceiveMonsterDamage(opponent *Monster) {
	soldier.Hp -= (opponent.sumAttack() - soldier.effect.attackReduction)
	changeStatus(&soldier.Hp, &soldier.status)
}

func (monster *Monster) UseItem(items *Item) {
	monster.finalDamage += items.attackUp - monster.item.attackUp
	monster.effect.attackReduction = items.attackNull - monster.item.attackNull
	monster.item = *items
}

func (monster *Monster) Punch(target MonsterTarget) {
	target.ReceiveMonsterDamage(monster)
}

func (monster *Monster) GetBasicInfo() BasicInfo {
	return BasicInfo{name: monster.name, Hp: monster.Hp, status: monster.status,
		weaponName: monster.weapon.name, itemName: monster.item.name, finalDamage: monster.finalDamage}
}

func (monster *Monster) ShowBasicInfo() {
	fmt.Printf("name: %v, Hp: %v, status: %v, weaponName: %v, itemName: %v, finalDamage: %v\n",
		monster.name, monster.Hp, monster.status, monster.weapon.name, monster.item.name, monster.finalDamage)
}

//////////////// Soldier Methods

func (soldier *Soldier) sumAttack() int {
	soldier.finalDamage += soldier.effect.passiveDamage
	return soldier.finalDamage
}

func (hero *Hero) ReceiveSoldierDamage(opponent *Soldier) {
	hero.Hp -= (opponent.sumAttack() - hero.effect.attackReduction)
	changeStatus(&hero.Hp, &hero.status)
}

func (monster *Monster) ReceiveSoldierDamage(opponent *Soldier) {
	monster.Hp -= (opponent.sumAttack() - monster.effect.attackReduction)
	changeStatus(&monster.Hp, &monster.status)
}

func (soldier *Soldier) ReceiveSoldierDamage(opponent *Soldier) {
	soldier.Hp -= (opponent.sumAttack() - soldier.effect.attackReduction)
	changeStatus(&soldier.Hp, &soldier.status)
}

func (soldier *Soldier) UseItem(items *Item) {
	soldier.finalDamage += items.attackUp - soldier.item.attackUp
	soldier.effect.attackReduction = items.attackNull - soldier.item.attackNull
	soldier.item = *items
}

func (soldier *Soldier) AttachSoldier(weapons *Weapon) {
	soldier.finalDamage += weapons.damage - soldier.weapon.damage
	soldier.weapon = *weapons
}

func (soldier *Soldier) Punch(target SoldierTarget) {
	target.ReceiveSoldierDamage(soldier)
}

func (soldier *Soldier) GetBasicInfo() BasicInfo {
	return BasicInfo{name: soldier.name, Hp: soldier.Hp, status: soldier.status,
		weaponName: soldier.weapon.name, itemName: soldier.item.name, finalDamage: soldier.finalDamage}
}

func (soldier *Soldier) ShowBasicInfo() {
	fmt.Printf("name: %v, Hp: %v, status: %v, weaponName: %v, itemName: %v, finalDamage: %v\n",
		soldier.name, soldier.Hp, soldier.status, soldier.weapon.name, soldier.item.name, soldier.finalDamage)
}
