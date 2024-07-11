package block_actors

import general "phoenixbuilder/minecraft/protocol/block_actors/general_actors"

// 高炉
type BlastFurnace struct {
	general.FurnaceBlockActor
}

// ID ...
func (*BlastFurnace) ID() string {
	return IDBlastFurnace
}