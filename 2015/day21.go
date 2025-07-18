// Turns: Player = n; Boss = n -1;
// Boss damage - player armor * n-1 < player hp
// player damage - boss armor * n  > = player hp
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const PLAYER_HP = 100

type Stats struct {
	armor  int
	damage int
	hp     int
}

type Item struct {
	category string
	name     string
	cost     int
	damage   int
	armor    int
}

type Loadout struct {
	items   []Item
	cost    int
	damage  int
	armor   int
	winning bool
}

var shop = `Weapons:    Cost  Damage  Armor
Dagger        8     4       0
Shortsword   10     5       0
Warhammer    25     6       0
Longsword    40     7       0
Greataxe     74     8       0

Armor:      Cost  Damage  Armor
Leather      13     0       1
Chainmail    31     0       2
Splintmail   53     0       3
Bandedmail   75     0       4
Platemail   102     0       5

Rings:      Cost  Damage  Armor
Damage +1    25     1       0
Damage +2    50     2       0
Damage +3   100     3       0
Defense +1   20     0       1
Defense +2   40     0       2
Defense +3   80     0       3`

func Day21(input string, part int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var boss Stats
	for _, v := range lines {
		stat := strings.Split(v, ":")
		num, _ := strconv.Atoi(strings.TrimSpace(stat[1]))
		switch stat[0] {
		case "Hit Points":
			boss.hp = num
		case "Damage":
			boss.damage = num
		case "Armor":
			boss.armor = num
		}
	}
	loadouts := getLoadouts(boss)

	if part == 1 {
		cheapestWin := math.MaxInt32
		var cheapLoadout Loadout
		for _, v := range loadouts {
			if v.winning && v.cost < cheapestWin {
				cheapestWin = v.cost
				cheapLoadout = v
			}
		}
		fmt.Printf("Day 21 Part 1 - Cheapest win is: %d with loadout: %v\n", cheapestWin, cheapLoadout)
	} else if part == 2 {
		costlyLoss := 0
		var costlyLoadout Loadout
		for _, v := range loadouts {
			if !v.winning && v.cost > costlyLoss {
				costlyLoss = v.cost
				costlyLoadout = v
			}
		}
		fmt.Printf("Day 21 Part 2 - Most expensive loss is: %d with loadout: %v\n", costlyLoss, costlyLoadout)
	}
}

func getLoadouts(boss Stats) []Loadout {
	weapons, armor, rings := makeShop()
	var loadouts []Loadout
	//Setting negatives so unequiped can exist
	for w := 0; w < len(weapons); w++ {
		for a := -1; a < len(armor); a++ {
			for r1 := -1; r1 < len(rings); r1++ {
				for r2 := -1; r2 < len(rings); r2++ {
					items := []Item{}
					items = append(items, weapons[w])
					if a != -1 {
						items = append(items, armor[a])
					}
					if r1 != -1 {
						items = append(items, rings[r1])
					}
					if r2 != -1 && r2 != r1 {
						items = append(items, rings[r2])
					}

					var loadout Loadout
					loadout.items = items
					for _, v := range items {
						loadout.cost += v.cost
						loadout.armor += v.armor
						loadout.damage += v.damage
						if boss.hp/max((loadout.damage-boss.armor), 1) <= PLAYER_HP/max((boss.damage-loadout.armor), 1) {
							loadout.winning = true
						} else {
							loadout.winning = false
						}
					}
					loadouts = append(loadouts, loadout)
				}
			}
		}
	}
	return loadouts
}

func makeShop() (weapons, armors, rings []Item) {
	items := []Item{}
	sections := strings.Split(strings.TrimSpace(shop), "\n\n")

	for _, section := range sections {
		category := strings.Split(strings.TrimSpace(section), ":")[0]
		for l, line := range strings.Split(section, "\n") {
			if l == 0 {
				continue
			}

			if category == "Rings" {
				line = strings.ReplaceAll(line, "e +", "e+")
			}
			fLine := strings.Fields(line)
			var newItem Item
			newItem.category = category
			newItem.name = string(fLine[0])
			newItem.cost, _ = strconv.Atoi(string(fLine[1]))
			newItem.damage, _ = strconv.Atoi(string(fLine[2]))
			newItem.armor, _ = strconv.Atoi(string(fLine[3]))
			items = append(items, newItem)
		}
	}
	for _, v := range items {
		switch v.category {
		case "Rings":
			rings = append(rings, v)
		case "Weapons":
			weapons = append(weapons, v)
		case "Armor":
			armors = append(armors, v)
		}
	}
	return weapons, armors, rings
}
