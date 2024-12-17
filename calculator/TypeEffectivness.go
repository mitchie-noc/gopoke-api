package calculator

import (
	"encoding/json"	
	"io/ioutil"
	"os"
	"fmt"
)

type TypeEffectiveness map[string]map[string]float64

type TypeEffectivenessSummary struct {
	Weaknesses  []Effect `json:"weaknesses"`
	Resistances []Effect `json:"resistances"`
	Immunities  []Effect `json:"immunities"`
	Neutral  []Effect `json:"neutral"`
}

type Effect struct {
	TypeName string  `json:"typeName"`
	Value    float64 `json:"value"`
}

var typeEffectiveness TypeEffectiveness

func LoadTypeEffectiveness() (error) {
	file, err := os.Open("./type_defensive_props.json")
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &typeEffectiveness); err != nil {
		return err
	}

	return nil
}

func GetMatchups(types []string) (TypeEffectivenessSummary, error) {
	effectiveness := make(map[string]float64)
	for targetType := range typeEffectiveness {
		effectiveness[targetType] = 1.0
	}

	for _, typeName := range types {
		typeMatchups, exists := typeEffectiveness[typeName]
		if !exists {
			return TypeEffectivenessSummary{}, fmt.Errorf("type %s not found in matchup data", typeName)
		}

		for targetType, value := range typeMatchups {
			if _, ok := effectiveness[targetType]; !ok {
				return TypeEffectivenessSummary{}, fmt.Errorf("unexpected target type %s", targetType)
			}
			effectiveness[targetType] *= value
		}
	}

	var weaknesses, resistances, immunities, neutral []Effect

	for targetType, value := range effectiveness {
		effect := Effect{TypeName: targetType, Value: value}
		switch {
		case value == 0.0:
			immunities = append(immunities, effect)
		case value > 1.0:
			weaknesses = append(weaknesses, effect)
		case value < 1.0:
			resistances = append(resistances, effect)
		default:
			neutral = append(neutral, effect)
		}
	}

	return TypeEffectivenessSummary{
		Weaknesses:  weaknesses,
		Resistances: resistances,
		Immunities:  immunities,
		Neutral:     neutral,
	}, nil
}

