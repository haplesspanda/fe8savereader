package parse

import (
	"io"
	"os"
)

// Character and class lists from here with some manual edits.
// https://github.com/StanHash/FE-CHAX/tree/master/GameData/Tables
var characterNames = []string{"(Unused)",
	"Eirika",
	"Seth",
	"Gilliam",
	"Franz",
	"Moulder",
	"Vanessa",
	"Ross",
	"Neimi",
	"Colm",
	"Garcia",
	"Innes",
	"Lute",
	"Natasha",
	"Cormag",
	"Ephraim",
	"Forde",
	"Kyle",
	"Amelia",
	"Artur",
	"Gerik",
	"Tethys",
	"Marisa",
	"Saleh",
	"Ewan",
	"L'Arachel",
	"Dozla",
	"-",
	"Rennac",
	"Duessel",
	"Myrrh",
	"Knoll",
	"Joshua",
	"Syrene",
	"Tana",
	"Lyon",
	"Orson",
	"Glen",
	"Selena",
	"Valter",
	"Riev",
	"Caellach",
	"Fado",
	"Ismaire",
	"Hayden"}

var classNames = []string{"(Unused))",
	"Lord (Ephraim)",
	"Lord (Eirika)",
	"Great Lord (Ephraim)",
	"Great Lord (Eirika)",
	"Cavalier",
	"Cavalier (F)",
	"Paladin",
	"Paladin (F)",
	"Knight",
	"Knight (F)",
	"General",
	"General (F)",
	"Thief",
	"Manakete",
	"Mercenary",
	"Mercenary (F)",
	"Hero",
	"Hero (F)",
	"Myrmidon",
	"Myrmidon (F)",
	"Swordmaster",
	"Swordmaster (F)",
	"Assassin",
	"Assassin (F)",
	"Archer",
	"Archer (F)",
	"Sniper",
	"Sniper (F)",
	"Ranger",
	"Ranger (F)",
	"Wyvern Rider",
	"Wyvern Rider (F)",
	"Wyvern Lord",
	"Wyvern Lord (F)",
	"Wyvern Knight",
	"Wyvern Knight (F)",
	"Mage",
	"Mage (F)",
	"Sage",
	"Sage (F)",
	"Mage Knight",
	"Mage Knight (F)",
	"Bishop",
	"Bishop (F)",
	"Shaman",
	"Shaman (F)",
	"Druid",
	"Druid (F)",
	"Summoner",
	"Summoner (F)",
	"Rogue",
	"Gorgon Egg",
	"Great Knight",
	"Great Knight (F)",
	"Recruit (2)",
	"Journeyman (3)",
	"Pupil (3)",
	"Recruit (3)",
	"Manakete",
	"Manakete (F)",
	"Journeyman (1)",
	"Pupil (1)",
	"Fighter",
	"Warrior",
	"Brigand",
	"Pirate",
	"Berserker",
	"Monk",
	"Priest",
	"Bard",
	"Recruit (1)",
	"Pegasus Knight",
	"Falcoknight",
	"Cleric",
	"Troubadour",
	"Valkyrie",
	"Dancer",
	"Soldier",
	"Necromancer",
	"Fleet",
	"Ghost Fighter",
	"Revenant",
	"Entombed",
	"Bonewalker",
	"Bonewalker (Bow)",
	"Wight",
	"Wight (Bow)",
	"Bael",
	"Elder Bael",
	"Cyclops",
	"Mauthedoog",
	"Gwyllgi",
	"Tarvos",
	"Maelduin",
	"Mogall",
	"ArchMogall",
	"Gorgon",
	"Gorgon Egg",
	"Gargoyle",
	"Deathgoyle",
	"Dracozombie",
	"Demon King",
	"Archer on Ballista",
	"Archer on Iron Ballista",
	"Archer on Killer Ballista",
	"Ballista",
	"Iron Ballista",
	"Killer Ballista",
	"Civilian",
	"Civilian (F)",
	"Civilian",
	"Civilian (F)",
	"Civilian",
	"Civilian (F)",
	"Peer",
	"Queen",
	"Prince",
	"Queen",
	"--",
	"Fallen Prince",
	"Tent",
	"Pontifex",
	"Dead Peer",
	"Cyclops",
	"Elder Bael",
	"Journeyman (2)",
	"Pupil (2)",
}

type ChapterData struct {
	Name  string
	Route string
}

var chapterData = map[int]ChapterData{
	0x00: {Name: "Prologue: The Fall of Renais", Route: "None"},
	0x01: {Name: "Ch.1: Escape!", Route: "None"},
	0x02: {Name: "Ch.2: The Protected", Route: "None"},
	0x03: {Name: "Ch.3: The Bandits of Borgo", Route: "None"},
	0x04: {Name: "Ch.4: Ancient Horrors", Route: "None"},
	0x05: {Name: "Ch.5x: Unbroken Heart", Route: "None"},
	0x06: {Name: "Ch.5: The Empire's Reach", Route: "None"},
	0x07: {Name: "Ch.6: Victims of War", Route: "None"},
	0x08: {Name: "Ch.7: Waterside Renvall", Route: "None"},
	0x09: {Name: "Ch.8: It's a Trap!", Route: "None"},
	0x0A: {Name: "Ch.9: Distant Blade", Route: "Eirika"},
	0x0B: {Name: "Ch.10: Revolt at Carcino", Route: "Eirika"},
	0x0C: {Name: "Ch.12: Village of Silence", Route: "Eirika"},
	0x0D: {Name: "Ch.13: Hamill Canyon", Route: "Eirika"},
	0x0E: {Name: "Ch.14: Queen of White Dunes", Route: "Eirika"},
	0x0F: {Name: "Ch.15: Scorched Sand", Route: "Eirika"},
	0x10: {Name: "Ch.16: Ruled by Madness", Route: "Eirika"},
	0x11: {Name: "Ch.17: River of Regrets", Route: "Eirika"},
	0x12: {Name: "Ch.18: Two Faces of Evil", Route: "Eirika"},
	0x13: {Name: "Ch.19: Last Hope", Route: "Eirika"},
	0x14: {Name: "Ch.20: Darkling Woods", Route: "Eirika"},
	0x15: {Name: "Final Chapter: Sacred Stone (Part 1)", Route: "Eirika"}, // Unused?
	0x16: {Name: "Final Chapter: Sacred Stone", Route: "Eirika"},
	0x17: {Name: "Ch.9: Fort Rigwald", Route: "Ephraim"},
	0x18: {Name: "Ch.10: Turning Traitor", Route: "Ephraim"},
	0x19: {Name: "Ch.12: Landing at Taizel", Route: "Ephraim"},
	0x1A: {Name: "Ch.13: Fluorspar's Oath", Route: "Ephraim"},
	0x1B: {Name: "Ch.14: Father and Son", Route: "Ephraim"},
	0x1C: {Name: "Ch.15: Scorched Sand", Route: "Ephraim"},
	0x1D: {Name: "Ch.16: Ruled by Madness", Route: "Ephraim"},
	0x1E: {Name: "Ch.17: River of Regrets", Route: "Ephraim"},
	0x1F: {Name: "Ch.18: Two Faces of Evil", Route: "Ephraim"},
	0x20: {Name: "Ch.19: Last Hope", Route: "Ephraim"},
	0x21: {Name: "Ch.20: Darkling Woods", Route: "Ephraim"},
	0x22: {Name: "Final Chapter: Sacred Stone (Part 1)", Route: "Ephraim"}, // Unused?
	0x23: {Name: "Final Chapter: Sacred Stone", Route: "Ephraim"},
	0x3D: {Name: "Ch.11: Creeping Darkness", Route: "Eirika"},
	0x3E: {Name: "Ch.11: Phantom Ship", Route: "Ephraim"},
}

type SaveData struct {
	Units    []Unit
	Chapters []Chapter

	Route      string
	TotalTurns int
}

type Unit struct {
	// Fields from save data.
	ClassIndex int
	Level      int
	Exp        int
	Dead       bool
	MetisTome  bool
	MaxHp      int
	Pow        int
	Skl        int
	Spd        int
	Def        int
	Res        int
	Lck        int
	ConBonus   int
	MovBonus   int
	CharIndex  int

	// Fields inferred from ROM data.
	CharName  string
	ClassName string
}

type Chapter struct {
	ChapterId int
	TurnCount int
	Time      int

	// Fields inferred from ROM data.
	ChapterName string
}

func ParseSave(f io.ReadSeeker) SaveData {
	// Save file layout:
	// https://github.com/StanHash/DOC/blob/master/RealSaveData.txt

	saveStart := 0x3FC4
	// saveLength := 0xDC8
	unitOffset := 0x04C
	unitSectionLength := 0x72C
	unitLength := 0x24
	offset, err := f.Seek(int64(saveStart+unitOffset), io.SeekStart)
	if err != nil {
		panic(err)
	}

	units := make([]Unit, 256)
	for i := 0; offset < int64(saveStart+unitOffset+unitSectionLength); i++ {
		unit := Unit{}
		buf := make([]byte, unitLength)
		n, err := f.Read(buf[:cap(buf)])
		buf = buf[:n]
		if err != nil {
			panic(err)
		}
		unit.ClassIndex = int(buf[0] & 0x7f)
		unit.Level = int(buf[0]&0x80>>7 + buf[1]&0x0f<<1)
		unit.Exp = int(buf[1]&0xf0>>4 + buf[2]&0x07<<4)
		unit.Dead = buf[3]&0x80 == 0x80
		unit.MaxHp = int(buf[5]&0xf0>>4 + buf[6]&0x03<<4)
		unit.Pow = int(buf[6] & 0x7c >> 2)
		unit.Skl = int(buf[6]&0x80>>7 + buf[7]&0x0f<<1)
		unit.Spd = int(buf[7]&0xf0>>4 + buf[8]&0x01<<4)
		unit.Def = int(buf[8] & 0x3e >> 1)
		unit.Res = int(buf[8]&0xc0>>6 + buf[9]&0x07<<2)
		unit.Lck = int(buf[9] & 0xf8 >> 3)
		unit.ConBonus = int(buf[10] & 0x1f)
		unit.MovBonus = int(buf[10]&0xe0>>5 + buf[11]&0x03<<3)
		unit.CharIndex = int(buf[0x14])
		if unit.CharIndex != 0 {
			unit.CharName = characterNames[unit.CharIndex]
			unit.ClassName = classNames[unit.ClassIndex]

			units[unit.CharIndex] = unit
		}

		offset, err = f.Seek(0, os.SEEK_CUR)
		if err != nil {
			panic(err)
		}
	}

	returnedUnits := make([]Unit, 0)
	for _, unit := range units {
		if unit.CharIndex != 0 {
			returnedUnits = append(returnedUnits, unit)
		}
	}

	chapterOffset := 0xCAC
	chapterSectionLength := 0xC0

	chapters := make([]Chapter, 0)
	offset, err = f.Seek(int64(saveStart+chapterOffset), io.SeekStart)
	if err != nil {
		panic(err)
	}

	route := "None"
	totalTurns := 0

	for i := 0; offset < int64(saveStart+chapterOffset+chapterSectionLength); i++ {
		chapter := Chapter{}
		buf := make([]byte, 4)
		n, err := f.Read(buf[:cap(buf)])
		buf = buf[:n]
		if err != nil {
			panic(err)
		}

		chapter.ChapterId = int(buf[0] & 0x7f)
		chapter.TurnCount = int(buf[0]&0x80>>7 + buf[1]<<1)
		// TODO: Not sure about byte ordering here
		// chapter.Time = int(buf[2]) * 256 + int(buf[3])

		if chapter.TurnCount != 0 {
			entry, exists := chapterData[chapter.ChapterId]
			if exists {
				chapter.ChapterName = entry.Name
				if entry.Route != "None" {
					route = entry.Route
				}
			} else {
				chapter.ChapterName = "Unknown chapter"
			}
			chapters = append(chapters, chapter)
			totalTurns += chapter.TurnCount
		}

		offset, err = f.Seek(0, os.SEEK_CUR)
		if err != nil {
			panic(err)
		}
	}

	return SaveData{
		Units:      returnedUnits,
		Chapters:   chapters,
		Route:      route,
		TotalTurns: totalTurns,
	}
}
