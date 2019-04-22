package main


type StashList struct {
	NextChangeId 			string		`json:"next_change_id"`
	Stash 					[]Stash		`json:"stashes"`
}

type Stash struct {
	AccountName 		string			`json:"accountName"`
	LastCharacterName 	string			`json:"lastCharacterName"`
	Id 					string			`json:"id"`
	Stash 				string			`json:"stash"`
	StashType 			string			`json:"stashType"`
	Items 				[]Item			`json:"items"`
	Public 				bool			`json:"public"`
}

type Item struct {
	artFilename 			string			`json:"artFilename"`
	Elder 					bool			`json:"elder"`

	Shaper 					bool			`json:"shaper"`
	LockedToCharacter 		bool			`json:"lockedToCharacter"`
	IsRelic 				bool			`json:"isRelic"`
	Duplicated 				bool			`json:"duplicated"`
	Corrupted 				bool			`json:"corrupted"`
	AbyssJewel 				bool			`json:"abyssJewel"`
	Verified 				bool			`json:"verified"`
	W 						int				`json:"w"`
	H 						int				`json:"h"`
	ItemLevel 				int				`json:"ilvl"`
	Icon 					string			`json:"icon"`
	League 					string			`json:"league"`
	Id 						string			`json:"id"`
	Sockets 				[]Socket		`json:"sockets"`
	Name 					string			`json:"name"`
	TypeLine 				string			`json:"typeLine"`
	Note 					string			`json:"note"`
	Properties 				[]Property		`json:"properties"`
	AdditionalProperties 	[]Property		`json:"additionalProperties"`

	Requirements 			[]Property		`json:"requirements"`
	NextLevelRequirements 	[]Property		`json:"nextLevelRequirements"`
	ImplicitMods 			[]string		`json:"implicitMods"`
	ExplicitMods 			[]string		`json:"explicitMods"`
	CosmeticMods 			[]string		`json:"cosmeticMods"`
	CraftedMods 			[]string		`json:"craftedMods"`
	EnchantMods 			[]string		`json:"enchantMods"`
	UtilityMods 			[]string		`json:"utilityMods"`

	FlavourText 			[]string		`json:"flavourText"`
	DescrText 				string			`json:"descrText"`
	FrameType 				int				`json:"frameType"`
	SecDescrText 			string			`json:"secDescrText"`
	ProphecyDiffText 		string			`json:"prophecyDiffText"`
	ProphecyText 			string			`json:"prophecyText"`

	MaxStackSize 			int				`json:"maxStackSize"`
	StackSize 				int				`json:"stackSize"`

	Category 				Category		`json:"category"`
	X 						int				`json:"x"`
	Y 						int				`json:"y"`
	InventoryId 			string			`json:"inventoryId"`
	SocketedItems 			[]Item			`json:"socketedItems"`
	Socket 					int				`json:"socket"`
	Colour 					string			`json:"colour"`
	Support 				bool			`json:"support"`
	TalismanTier 			int				`json:"talismanTier"`

}

type Property struct {
	Name 		string					`json:"name"`
	Values 		[]interface{}			`json:"values"`
	DisplayMode int						`json:"displayMode"`
	Type 		int						`json:"type"`
}

type Category struct {
	Weapons 		[]string			`json:"weapons"`
}

type Socket struct {
	Group 		int						`json:"group"`
	Attr 		string					`json:"attr"`
	SColour 	string				`json:"sColour"`
}
