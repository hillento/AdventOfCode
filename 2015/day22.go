package fifteen

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//Part1
//734 is too low
//1216 is too high

//Part2
//1242 is too high

type Spell struct {
	Name          string
	Cost          int
	InstantDamage int
	Damage        int
	Heal          int
	Armor         int
	Turns         int
	Mana          int
	Index         int
}

type BattleStatus struct {
	PlayerHp     int
	PlayerMana   int
	BossHp       int
	BossDamage   int
	ActiveSpells [5]int
	PlayerTurn   bool
	Depth        int
}

// This is the same for everyone, it isn't passed as input
var Spells = [5]Spell{
	{Name: "Magic Missile", Cost: 53, InstantDamage: 4, Heal: 0, Armor: 0, Turns: 0, Mana: 0, Index: 0},
	{Name: "Drain", Cost: 73, InstantDamage: 2, Damage: 0, Heal: 2, Armor: 0, Turns: 0, Mana: 0, Index: 1},
	{Name: "Shield", Cost: 113, InstantDamage: 0, Damage: 0, Heal: 0, Armor: 7, Turns: 6, Mana: 0, Index: 2},
	{Name: "Poison", Cost: 173, InstantDamage: 0, Damage: 3, Heal: 0, Armor: 0, Turns: 6, Mana: 0, Index: 3},
	{Name: "Recharge", Cost: 229, InstantDamage: 0, Damage: 0, Heal: 0, Armor: 0, Turns: 5, Mana: 101, Index: 4},
}

func Day22(input string, part int) {
	var status BattleStatus
	status.PlayerHp = 50
	status.PlayerMana = 500
	status.ActiveSpells = [5]int{0, 0, 0, 0, 0}
	status.Depth = 0
	status.PlayerTurn = true
	for line := range strings.Lines(input) {
		stat := strings.Split(strings.TrimSpace(line), ": ")
		if stat[0] == "Damage" {
			status.BossDamage, _ = strconv.Atoi(stat[1])
		} else {
			status.BossHp, _ = strconv.Atoi(stat[1])
		}
	}
	minMana := battle(status, map[string]int{}, part)
	fmt.Printf("Day 22 Part %d - The least amount of mana to win the fight is: %d\n", part, minMana)
}

func (status BattleStatus) key() string {
	return fmt.Sprintf("%d_%d_%d_%v_%v", status.PlayerHp, status.PlayerMana, status.BossHp, status.ActiveSpells, status.PlayerTurn)
}

func battle(status BattleStatus, memo map[string]int, part int) (minMana int) {
	key := status.key()
	if v, ok := memo[key]; ok {
		return v
	}

	if part == 2 && status.PlayerTurn {
		status.PlayerHp--
	}

	if status.PlayerHp <= 0 {
		return math.MaxInt32
	}

	//Set the player armor to zero. If the spell is active it will be set to the correct number
	armor := 0
	for _, v := range Spells {
		if status.ActiveSpells[v.Index] > 0 {
			status.ActiveSpells[v.Index]--
			status.BossHp -= v.Damage
			status.PlayerHp += v.Heal
			armor += v.Armor
			status.PlayerMana += v.Mana
		}
	}

	if status.BossHp <= 0 {
		return 0
	}

	minMana = math.MaxInt32
	if status.PlayerTurn {
		var spellCast bool
		for _, v := range Spells {
			if status.ActiveSpells[v.Index] == 0 && status.PlayerMana >= v.Cost {
				spellCast = true
				newActive := status.ActiveSpells
				newActive[v.Index] += v.Turns
				newStatus := BattleStatus{
					PlayerHp:     status.PlayerHp + v.Heal,
					PlayerMana:   status.PlayerMana - v.Cost,
					BossHp:       status.BossHp - v.InstantDamage,
					BossDamage:   status.BossDamage,
					ActiveSpells: newActive,
					PlayerTurn:   false,
					Depth:        status.Depth + 1,
				}

				castResult := v.Cost + battle(newStatus, memo, part)

				if castResult < minMana {
					minMana = castResult
				}
			}
		}
		if !spellCast {
			return math.MaxInt32
		}
	} else {
		bossAttackDamage := int(math.Max(float64(1), float64(status.BossDamage-armor)))

		newStatus := BattleStatus{
			PlayerHp:     status.PlayerHp - bossAttackDamage,
			PlayerMana:   status.PlayerMana,
			BossHp:       status.BossHp,
			BossDamage:   status.BossDamage,
			ActiveSpells: status.ActiveSpells,
			PlayerTurn:   true,
			Depth:        status.Depth + 1,
		}

		bossAttack := battle(newStatus, memo, part)

		if bossAttack < minMana {
			minMana = bossAttack
		}
	}
	memo[key] = minMana
	return minMana
}
