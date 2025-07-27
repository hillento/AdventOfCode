package fifteen

import (
	"fmt"
	"strconv"
	"strings"
)

type Spell struct {
	Name   string
	Cost   int
	Damage int
	Heal   int
	Armor  int
	Turns  int
	Mana   int
}

type BattleStatus struct {
	PlayerHp       int
	PlayerMana     int
	PlayerArmor    int
	bossHp         int
	bossDamage     int
	manaSpent      int
	activeStatuses []Spell
}

// This is the smae for everyone, it isn't passed as input
var Spells = [5]Spell{
	{Name: "Magic Missile", Cost: 53, Damage: 4, Heal: 0, Armor: 0, Turns: 0, Mana: 0},
	{Name: "Drain", Cost: 73, Damage: 2, Heal: 2, Armor: 0, Turns: 0, Mana: 0},
	{Name: "Shield", Cost: 113, Damage: 0, Heal: 0, Armor: 7, Turns: 6, Mana: 0},
	{Name: "Poison", Cost: 173, Damage: 3, Heal: 0, Armor: 0, Turns: 6, Mana: 0},
	{Name: "Recharge", Cost: 229, Damage: 0, Heal: 0, Armor: 0, Turns: 5, Mana: 101},
}

func Day22(input string) {
	var status BattleStatus
	status.PlayerHp = 50
	status.PlayerMana = 500
	status.PlayerArmor = 0
	for line := range strings.Lines(input) {
		stat := strings.Split(strings.TrimSpace(line), ": ")
		if stat[0] == "Damage" {
			status.bossDamage, _ = strconv.Atoi(stat[1])
		} else {
			status.bossHp, _ = strconv.Atoi(stat[1])
		}
	}
	fmt.Println(status)
}
