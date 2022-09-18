package format

import (
	"fmt"
	"io"
	"log"
	"sort"

	"github.com/haplesspanda/fe8savereader/parse"
)

type Unit = parse.Unit

type UnitDiff struct {
	New      bool
	Promoted bool
	Died     bool

	LevelDiff    int
	TotalExpDiff int
	HpBonus      int
	PowBonus     int
	SklBonus     int
	SpdBonus     int
	DefBonus     int
	ResBonus     int
	LckBonus     int
	ConBonus     int
	MovBonus     int
}

func Read(file io.ReadSeeker, output io.StringWriter) {
	saveData := parse.ParseSave(file)
	for _, unit := range saveData.Units {
		printUnit(unit, output)
	}
}

func Diff(oldFile io.ReadSeeker, newFile io.ReadSeeker, output io.StringWriter) {
	oldUnits := parse.ParseSave(oldFile).Units
	oldUnitMap := make(map[int]Unit)
	for _, unit := range oldUnits {
		oldUnitMap[unit.CharIndex] = unit
	}

	log.Printf("=== OLD UNITS ===")
	for _, unit := range oldUnits {
		printUnit(unit, nil)
	}
	log.Println()
	log.Println()

	newUnits := parse.ParseSave(newFile).Units
	newUnitMap := make(map[int]Unit)
	for _, unit := range newUnits {
		newUnitMap[unit.CharIndex] = unit
	}

	log.Printf("=== NEW UNITS ===")
	for _, unit := range newUnits {
		printUnit(unit, nil)
	}
	log.Println()
	log.Println()

	// Sort into stable order (by character index)
	newKeys := make([]int, 0)
	for key := range newUnitMap {
		newKeys = append(newKeys, key)
	}
	sort.Ints(newKeys)

	for _, charIndex := range newKeys {
		newUnit := newUnitMap[charIndex]
		oldUnit, oldExists := oldUnitMap[charIndex]

		var diff UnitDiff
		if !oldExists {
			// Zero out diffs so everything isn't increased from 0.
			diff = UnitDiff{
				New:  true,
				Died: newUnit.Dead,
			}
		} else {
			diff = UnitDiff{
				New:      false,
				Promoted: oldUnit.ClassName != newUnit.ClassName,
				Died:     !oldUnit.Dead && newUnit.Dead,

				LevelDiff:    newUnit.Level - oldUnit.Level,
				TotalExpDiff: newUnit.Level*100 + newUnit.Exp - (oldUnit.Level*100 + oldUnit.Exp),
				HpBonus:      newUnit.MaxHp - oldUnit.MaxHp,
				PowBonus:     newUnit.Pow - oldUnit.Pow,
				SklBonus:     newUnit.Skl - oldUnit.Skl,
				SpdBonus:     newUnit.Spd - oldUnit.Spd,
				DefBonus:     newUnit.Def - oldUnit.Def,
				ResBonus:     newUnit.Res - oldUnit.Res,
				LckBonus:     newUnit.Lck - oldUnit.Lck,
				ConBonus:     newUnit.ConBonus - oldUnit.ConBonus,
				MovBonus:     newUnit.MovBonus - oldUnit.MovBonus,
			}
		}

		if !oldUnit.Dead {
			printUnitWithDiff(newUnit, diff, output)
		}
	}
}

func printUnit(unit Unit, output io.StringWriter) {
	var statusString string
	if unit.Dead {
		statusString = ", DEAD"
	}
	writeLine(output, "")
	writeLine(output, fmt.Sprintf("%s, %s, Level %d EXP %d%s", unit.CharName, unit.ClassName, unit.Level, unit.Exp, statusString))
	writeLine(output, fmt.Sprintf("HP %d Pow %d Skl %d Spd %d Lck %d Def %d Res %d", unit.MaxHp, unit.Pow, unit.Skl, unit.Spd, unit.Lck, unit.Def, unit.Res))
}

func printUnitWithDiff(unit Unit, diff UnitDiff, output io.StringWriter) {
	if !diff.Promoted && !diff.Died && !diff.New && diff.TotalExpDiff == 0 && diff.LevelDiff == 0 && diff.HpBonus == 0 && diff.PowBonus == 0 && diff.SklBonus == 0 && diff.SpdBonus == 0 && diff.LckBonus == 0 && diff.DefBonus == 0 && diff.ResBonus == 0 {
		return
	}

	writeLine(output, "")
	defer log.Printf("Wrote unit")

	var statusString string
	if diff.Died {
		statusString = ", Met with a terrible fate :\\("
	} else if diff.New {
		statusString = ", Newly Recruited!"
	} else if diff.Promoted {
		statusString = ", Newly Promoted!"
	}

	var totalExpBonus string
	var levelString string
	if !diff.Died {
		if diff.TotalExpDiff > 0 {
			totalExpBonus = fmt.Sprintf(" **(+%d Total XP)**", diff.TotalExpDiff)
		}
		levelString = fmt.Sprintf(", Level %d%s EXP %d%s", unit.Level, formatBonus(diff.LevelDiff), unit.Exp, totalExpBonus)
	}

	writeLine(output, fmt.Sprintf("**%s**, %s%s%s", unit.CharName, unit.ClassName, levelString, statusString))

	if diff.Died {
		return
	}

	writeLine(output, fmt.Sprintf("HP %d%s Pow %d%s Skl %d%s Spd %d%s Lck %d%s Def %d%s Res %d%s",
		unit.MaxHp, formatBonus(diff.HpBonus),
		unit.Pow, formatBonus(diff.PowBonus),
		unit.Skl, formatBonus(diff.SklBonus),
		unit.Spd, formatBonus(diff.SpdBonus),
		unit.Lck, formatBonus(diff.LckBonus),
		unit.Def, formatBonus(diff.DefBonus),
		unit.Res, formatBonus(diff.ResBonus)))
}

func formatBonus(bonus int) string {
	var result string
	if bonus > 0 {
		result = fmt.Sprintf(" **(+%d)**", bonus)
	}
	return result
}

func writeLine(output io.StringWriter, str string) {
	if output != nil {
		_, err := output.WriteString(str + "\n")
		if err != nil {
			panic(err)
		}
	}

	log.Println(str)
}
