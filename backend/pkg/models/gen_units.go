/////////////////////////////
/// GENERATED DO NOT EDIT ///
/////////////////////////////

package models

import (
	"github.com/Codename-Recon/Codename-Recon/backend/pkg/enums/movement"
	"github.com/Codename-Recon/Codename-Recon/backend/pkg/enums/unit"
)

var Units = [unit.SIZE]Unit{
	{
		Name:             "ANTI_AIR",
		Description:      "ANTI_AIR_DESCRIPTION",
		Cost:             8000,
		Health:           100,
		Mp:               6,
		MovementType:     movement.TREADS,
		Fuel:             60,
		TurnFuel:         0,
		HiddenTurnFuel:   0,
		Vision:           2,
		Ammo:             9,
		Weapons:          [2]int{7, -1},
		MinRange:         1,
		MaxRange:         1,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: true,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     0,
	},
	{
		Name:             "APC",
		Description:      "APC_DESCRIPTION",
		Cost:             5000,
		Health:           100,
		Mp:               6,
		MovementType:     movement.TREADS,
		Fuel:             70,
		TurnFuel:         0,
		HiddenTurnFuel:   0,
		Vision:           1,
		Ammo:             0,
		Weapons:          [2]int{-1, -1},
		MinRange:         0,
		MaxRange:         0,
		CanSupply:        true,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: false,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         true,
		CarryingTypes:    []unit.EUnit{unit.INFANTRY, unit.MECH},
		CarryingSize:     1,
	},
	{
		Name:             "ARTILLERY",
		Description:      "ARTILLERY_DESCRIPTION",
		Cost:             6000,
		Health:           100,
		Mp:               5,
		MovementType:     movement.TREADS,
		Fuel:             50,
		TurnFuel:         0,
		HiddenTurnFuel:   0,
		Vision:           1,
		Ammo:             9,
		Weapons:          [2]int{0, -1},
		MinRange:         2,
		MaxRange:         3,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: false,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     0,
	},
	{
		Name:             "BCOPTER",
		Description:      "B-COPTER_DESCRIPTION",
		Cost:             9000,
		Health:           100,
		Mp:               6,
		MovementType:     movement.AIR,
		Fuel:             99,
		TurnFuel:         2,
		HiddenTurnFuel:   0,
		Vision:           3,
		Ammo:             6,
		Weapons:          [2]int{-1, -1},
		MinRange:         1,
		MaxRange:         1,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: false,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     1,
	},
	{
		Name:             "BATTLESHIP",
		Description:      "BATTLESHIP_DESCRIPTION",
		Cost:             28000,
		Health:           100,
		Mp:               5,
		MovementType:     movement.SEA,
		Fuel:             99,
		TurnFuel:         1,
		HiddenTurnFuel:   0,
		Vision:           2,
		Ammo:             9,
		Weapons:          [2]int{-1, -1},
		MinRange:         2,
		MaxRange:         6,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: false,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     0,
	},
	{
		Name:             "BLACK_BOAT",
		Description:      "BLACK_BOAT_DESCRIPTION",
		Cost:             7500,
		Health:           100,
		Mp:               7,
		MovementType:     movement.LANDER,
		Fuel:             60,
		TurnFuel:         1,
		HiddenTurnFuel:   0,
		Vision:           1,
		Ammo:             0,
		Weapons:          [2]int{-1, -1},
		MinRange:         0,
		MaxRange:         0,
		CanSupply:        false,
		CanRepair:        true,
		CanCapture:       false,
		CanMoveAndAttack: false,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         true,
		CarryingTypes:    []unit.EUnit{unit.INFANTRY, unit.MECH},
		CarryingSize:     1,
	},
	{
		Name:             "BLACK_BOMB",
		Description:      "BLACK_BOMB_DESCRIPTION",
		Cost:             25000,
		Health:           100,
		Mp:               9,
		MovementType:     movement.AIR,
		Fuel:             45,
		TurnFuel:         5,
		HiddenTurnFuel:   0,
		Vision:           1,
		Ammo:             0,
		Weapons:          [2]int{-1, -1},
		MinRange:         0,
		MaxRange:         0,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: false,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     1,
	},
	{
		Name:             "BOMBER",
		Description:      "BOMBER_DESCRIPTION",
		Cost:             22000,
		Health:           100,
		Mp:               7,
		MovementType:     movement.AIR,
		Fuel:             99,
		TurnFuel:         5,
		HiddenTurnFuel:   0,
		Vision:           2,
		Ammo:             9,
		Weapons:          [2]int{-1, -1},
		MinRange:         1,
		MaxRange:         1,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: true,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     1,
	},
	{
		Name:             "CARRIER",
		Description:      "CARRIER_DESCRIPTION",
		Cost:             30000,
		Health:           100,
		Mp:               5,
		MovementType:     movement.SEA,
		Fuel:             99,
		TurnFuel:         1,
		HiddenTurnFuel:   0,
		Vision:           4,
		Ammo:             9,
		Weapons:          [2]int{-1, -1},
		MinRange:         3,
		MaxRange:         8,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: false,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         true,
		CarryingTypes:    []unit.EUnit{unit.BCOPTER},
		CarryingSize:     2,
	},
	{
		Name:             "CRUISER",
		Description:      "CRUISER_DESCRIPTION",
		Cost:             18000,
		Health:           100,
		Mp:               6,
		MovementType:     movement.SEA,
		Fuel:             99,
		TurnFuel:         1,
		HiddenTurnFuel:   0,
		Vision:           3,
		Ammo:             9,
		Weapons:          [2]int{-1, -1},
		MinRange:         1,
		MaxRange:         1,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: true,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     0,
	},
	{
		Name:             "FIGHTER",
		Description:      "FIGHTER_DESCRIPTION",
		Cost:             20000,
		Health:           100,
		Mp:               9,
		MovementType:     movement.AIR,
		Fuel:             99,
		TurnFuel:         5,
		HiddenTurnFuel:   0,
		Vision:           2,
		Ammo:             9,
		Weapons:          [2]int{-1, -1},
		MinRange:         1,
		MaxRange:         1,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: true,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     1,
	},
	{
		Name:             "INFANTRY",
		Description:      "INFANTRY_DESCRIPTION",
		Cost:             1000,
		Health:           100,
		Mp:               3,
		MovementType:     movement.FOOT,
		Fuel:             99,
		TurnFuel:         0,
		HiddenTurnFuel:   0,
		Vision:           2,
		Ammo:             -1,
		Weapons:          [2]int{3, -1},
		MinRange:         1,
		MaxRange:         1,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       true,
		CanMoveAndAttack: true,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     0,
	},
	{
		Name:             "LANDER",
		Description:      "LANDER_DESCRIPTION",
		Cost:             12000,
		Health:           100,
		Mp:               6,
		MovementType:     movement.LANDER,
		Fuel:             99,
		TurnFuel:         1,
		HiddenTurnFuel:   0,
		Vision:           1,
		Ammo:             0,
		Weapons:          [2]int{-1, -1},
		MinRange:         0,
		MaxRange:         0,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: false,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         true,
		CarryingTypes:    []unit.EUnit{unit.ANTI_AIR, unit.APC, unit.ARTILLERY, unit.INFANTRY, unit.MD_TANK, unit.MECH, unit.MEGA_TANK, unit.MISSILE, unit.NEOTANK, unit.RECON, unit.TCOPTER},
		CarryingSize:     1,
	},
	{
		Name:             "MD_TANK",
		Description:      "MEDIUM_TANK_DESCRIPTION",
		Cost:             16000,
		Health:           100,
		Mp:               5,
		MovementType:     movement.TREADS,
		Fuel:             50,
		TurnFuel:         0,
		HiddenTurnFuel:   0,
		Vision:           1,
		Ammo:             8,
		Weapons:          [2]int{4, 6},
		MinRange:         1,
		MaxRange:         1,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: true,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     0,
	},
	{
		Name:             "MECH",
		Description:      "MECH_DESCRIPTION",
		Cost:             2500,
		Health:           100,
		Mp:               2,
		MovementType:     movement.BAZOOKA,
		Fuel:             70,
		TurnFuel:         0,
		HiddenTurnFuel:   0,
		Vision:           2,
		Ammo:             3,
		Weapons:          [2]int{1, 3},
		MinRange:         1,
		MaxRange:         1,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       true,
		CanMoveAndAttack: true,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     0,
	},
	{
		Name:             "MEGA_TANK",
		Description:      "MEGA_TANK_DESCRIPTION",
		Cost:             28000,
		Health:           100,
		Mp:               4,
		MovementType:     movement.TREADS,
		Fuel:             50,
		TurnFuel:         0,
		HiddenTurnFuel:   0,
		Vision:           1,
		Ammo:             3,
		Weapons:          [2]int{-1, -1},
		MinRange:         1,
		MaxRange:         1,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: true,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     0,
	},
	{
		Name:             "MISSILE",
		Description:      "MISSILE_DESCRIPTION",
		Cost:             12000,
		Health:           100,
		Mp:               4,
		MovementType:     movement.WHEELS,
		Fuel:             50,
		TurnFuel:         0,
		HiddenTurnFuel:   0,
		Vision:           1,
		Ammo:             6,
		Weapons:          [2]int{-1, -1},
		MinRange:         3,
		MaxRange:         5,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: false,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     0,
	},
	{
		Name:             "NEOTANK",
		Description:      "NEOTANK_DESCRIPTION",
		Cost:             22000,
		Health:           100,
		Mp:               6,
		MovementType:     movement.TREADS,
		Fuel:             99,
		TurnFuel:         0,
		HiddenTurnFuel:   0,
		Vision:           1,
		Ammo:             9,
		Weapons:          [2]int{-1, -1},
		MinRange:         1,
		MaxRange:         1,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: true,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     0,
	},
	{
		Name:             "PIPERUNNER",
		Description:      "PIPERUNNER_DESCRIPTION",
		Cost:             20000,
		Health:           100,
		Mp:               9,
		MovementType:     movement.PIPERUNNER,
		Fuel:             99,
		TurnFuel:         0,
		HiddenTurnFuel:   0,
		Vision:           4,
		Ammo:             9,
		Weapons:          [2]int{-1, -1},
		MinRange:         2,
		MaxRange:         5,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: false,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     0,
	},
	{
		Name:             "RECON",
		Description:      "RECON_DESCRIPTION",
		Cost:             4000,
		Health:           100,
		Mp:               8,
		MovementType:     movement.WHEELS,
		Fuel:             80,
		TurnFuel:         0,
		HiddenTurnFuel:   0,
		Vision:           5,
		Ammo:             -1,
		Weapons:          [2]int{6, -1},
		MinRange:         1,
		MaxRange:         1,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: true,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     0,
	},
	{
		Name:             "ROCKET",
		Description:      "ROCKET_DESCRIPTION",
		Cost:             15000,
		Health:           100,
		Mp:               5,
		MovementType:     movement.WHEELS,
		Fuel:             50,
		TurnFuel:         0,
		HiddenTurnFuel:   0,
		Vision:           1,
		Ammo:             6,
		Weapons:          [2]int{5, -1},
		MinRange:         3,
		MaxRange:         5,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: false,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     0,
	},
	{
		Name:             "STEALTH",
		Description:      "STEALTH_DESCRIPTION",
		Cost:             24000,
		Health:           100,
		Mp:               6,
		MovementType:     movement.AIR,
		Fuel:             60,
		TurnFuel:         5,
		HiddenTurnFuel:   8,
		Vision:           4,
		Ammo:             6,
		Weapons:          [2]int{-1, -1},
		MinRange:         1,
		MaxRange:         1,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: true,
		CanHide:          true,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     1,
	},
	{
		Name:             "CRUISER",
		Description:      "CRUISER_DESCRIPTION",
		Cost:             20000,
		Health:           100,
		Mp:               5,
		MovementType:     movement.SEA,
		Fuel:             60,
		TurnFuel:         1,
		HiddenTurnFuel:   5,
		Vision:           5,
		Ammo:             6,
		Weapons:          [2]int{-1, -1},
		MinRange:         1,
		MaxRange:         1,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: true,
		CanHide:          false,
		CanDive:          true,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     0,
	},
	{
		Name:             "TCOPTER",
		Description:      "T-COPTER_DESCRIPTION",
		Cost:             5000,
		Health:           100,
		Mp:               6,
		MovementType:     movement.AIR,
		Fuel:             99,
		TurnFuel:         2,
		HiddenTurnFuel:   0,
		Vision:           2,
		Ammo:             0,
		Weapons:          [2]int{-1, -1},
		MinRange:         0,
		MaxRange:         0,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: false,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{unit.INFANTRY, unit.MECH},
		CarryingSize:     1,
	},
	{
		Name:             "TANK",
		Description:      "TANK_DESCRIPTION",
		Cost:             7000,
		Health:           100,
		Mp:               6,
		MovementType:     movement.TREADS,
		Fuel:             70,
		TurnFuel:         0,
		HiddenTurnFuel:   0,
		Vision:           3,
		Ammo:             9,
		Weapons:          [2]int{2, 6},
		MinRange:         1,
		MaxRange:         1,
		CanSupply:        false,
		CanRepair:        false,
		CanCapture:       false,
		CanMoveAndAttack: true,
		CanHide:          false,
		CanDive:          false,
		CanCarry:         false,
		CarryingTypes:    []unit.EUnit{},
		CarryingSize:     0,
	},
}