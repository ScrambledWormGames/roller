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
	Die: 2 * d10     | Modifier: 10  | Breakdown: true
##################################################################

Rolled: 2  		 Running total: 2
Rolled: 3  		 Running total: 5
Applying modifier 10

Total: 15

```
