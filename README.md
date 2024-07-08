# Super Roller
Just a really simple dice rolling command line app.  Saves some time when rolling out characters in various older tabletop RPGs such as Rolemaster, where you need to roll 10 * d10 to get your character stats.

## Usage
```
Usage of ./roller:
  -b	Whether you want a breakdown of the dice rolls
  -d int
    	Dice selection, out of: D4, D6, D8, D10, D12, D20 and D100 (default 6)
  -m int
    	Any modifier you wish to add
  -n int
    	The number of dice you want to roll (default 1)
```

## Default roll with no settings
```
##################################################################
			SUPER ROLLER v1.0
------------------------------------------------------------------
	Die: 1 * d6      | Modifier: 0   | Breakdown: false
##################################################################

Total: 3

```

## Roll with various settings and breakdown 
useful when you dont actually need the total, but the separate rolls.

```
##################################################################
			SUPER ROLLER v1.0
------------------------------------------------------------------
	Die: 10 * d10     | Modifier: 0   | Breakdown: true
##################################################################

Rolled: 5  		 Running total: 5
Rolled: 8  		 Running total: 13
Rolled: 8  		 Running total: 21
Rolled: 5  		 Running total: 26
Rolled: 7  		 Running total: 33
Rolled: 1  		 Running total: 34
Rolled: 7  		 Running total: 41
Rolled: 2  		 Running total: 43
Rolled: 2  		 Running total: 45
Rolled: 7  		 Running total: 52
Total: 52

```
