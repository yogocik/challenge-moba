package creature

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
