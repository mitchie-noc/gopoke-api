package nature

import (
	"github.com/mtslzr/pokeapi-go/structs"
)


type Nature struct {
	Name string
	Decreased_Stat string
	Increased_Stat string
}

func MapNature(original structs.Nature) Nature {
	nature := Nature{}
	nature.Name = original.Name
	nature.Decreased_Stat = getStatName(original.DecreasedStat)
	nature.Increased_Stat = getStatName(original.IncreasedStat)
	return nature
}

func getStatName(stat interface{}) string {
	if statMap, ok := stat.(map[string]interface{}); ok {
		if name, ok := statMap["name"].(string); ok {
			return name
		}
	}
	
	if statStruct, ok := stat.(structs.Stat); ok {
		return statStruct.Name
	}

	return "none" // Default return value if stat is unrecognized
}
