package resources

var Wood = &SimpleResource{
	Name:        "wood",
	Description: "a stack of logs",
	stackable:   true,
}

var Stone = &SimpleResource{
	Name:        "stone",
	Description: "a collection of stones",
	stackable:   true,
}

var Water = &SimpleResource{
	Name:        "water",
	Description: "a jug of water",
	stackable:   true,
}

var Berries = &SimpleResource{
	Name:        "berries",
	Description: "small wild berries picked from the countryside",
	stackable:   true,
}
