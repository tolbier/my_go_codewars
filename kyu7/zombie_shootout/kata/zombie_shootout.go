package kata

import "fmt"

func Zombie_shootout(initialZombies, initialRange, initialAmmo int) string {
	killedZombies := 0

	ammo := initialAmmo
	if killedZombies == initialZombies {
		return fmt.Sprintf("You shot all %d zombies.", killedZombies)
	}
	for _range := float64(initialRange); _range > 0.0; _range -= 0.5 {
		if ammo == 0 {
			return fmt.Sprintf("You shot %d zombies before being eaten: ran out of ammo.", killedZombies)
		}
		ammo--
		killedZombies++
		if killedZombies == initialZombies {
			return fmt.Sprintf("You shot all %d zombies.", killedZombies)
		}
	}
	return fmt.Sprintf("You shot %d zombies before being eaten: overwhelmed.", killedZombies)
}

/*
func Zombie_shootout(z, r, a int) string {
	x := r * 2

	if x >= z && z <= a {
		return fmt.Sprintf("You shot all %d zombies.", z)
	}

	if x <= a {
		return fmt.Sprintf("You shot %d zombies before being eaten: overwhelmed.", x)
	}

	return fmt.Sprintf("You shot %d zombies before being eaten: ran out of ammo.", a)

}*/
