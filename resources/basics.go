package resources

var Wood *SimpleResource = &SimpleResource{
	Name:        "wood",
	Description: "a stack of logs",
	stackable:   true,
}

var Stone *SimpleResource = &SimpleResource{
	Name:        "stone",
	Description: "a collection of stones",
	stackable:   true,
}

var Water *SimpleResource = &SimpleResource{
	Name:        "water",
	Description: "a jug of water",
	stackable:   true,
}
