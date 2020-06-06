package resources

type SimpleResource struct {
	Name        string
	Description string
	stackable   bool
}

func (w *SimpleResource) GetName() string {
	return w.Name
}

func (w *SimpleResource) GetDescription() string {
	return w.Description
}

func (w *SimpleResource) IsStackable() bool {
	return w.stackable
}
