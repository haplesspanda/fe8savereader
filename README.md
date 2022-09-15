# FE8 Save Reader

This command-line tool can be used to read Fire Emblem: Sacred Stones (aka FE8) savefiles and print them in a human-readable format.

Confirmed to work on savefiles written by the USA version of FE8. It may work on other versions, but I haven't tested them and don't know.

## Usage

Can provide a summary for a single savefile (`read`), or compare two savefiles and summarize diffs (`compare`).

```
> fe8savereader.exe read savefile.sav output.txt

> fe8savereader.exe compare oldsave.sav newsave.sav output.txt
```

## Read Output

Plain text.

```
Eirika, Lord (Eirika), Level 10 EXP 79
HP 22 Pow 7 Skl 14 Spd 14 Lck 11 Def 7 Res 2

Seth, Paladin, Level 2 EXP 94
HP 31 Pow 15 Skl 13 Spd 12 Lck 13 Def 11 Res 9

Gilliam, Knight, Level 8 EXP 10
HP 29 Pow 11 Skl 7 Spd 6 Lck 5 Def 10 Res 3

[...]
```

## Compare Output

Markdown format for easy posting to Discord.

```md
**Seth**, Paladin, Level 8 **(+1)** EXP 37 **(+92 Total XP)**
HP 36 **(+1)** Pow 16 Skl 15 Spd 16 **(+1)** Lck 13 Def 13 Res 10 **(+1)**

**Gilliam**, Knight, Level 16 **(+2)** EXP 86 **(+200 Total XP)**
HP 36 **(+2)** Pow 13 Skl 8 **(+1)** Spd 9 Lck 6 **(+1)** Def 16 **(+1)** Res 4

[...]
```

## Credits

Big thanks to the [Stan doc](https://github.com/StanHash/DOC) that documents the savefile layout, and another of their repos with some in-game data tables.
