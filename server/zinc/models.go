package zinc

type ZincService struct {
	BaseUrl   string
	User      string
	Password  string
	IndexName string
}

type MappingProperty struct {
	Type          string `json:"type"`
	Index         bool   `json:"index"`
	Store         bool   `json:"store"`
	Sortable      bool   `json:"sortable"`
	Aggregatable  bool   `json:"aggregatable"`
	Highlightable bool   `json:"highlightable"`
}

type DateMappingProperty struct {
	Format string `json:"format"`
	MappingProperty
}

type MappingProperties struct {
	MessageID MappingProperty     `json:"message_id"`
	Date      DateMappingProperty `json:"date"`
	From      MappingProperty     `json:"from"`
	To        MappingProperty     `json:"to"`
	CC        MappingProperty     `json:"cc"`
	BCC       MappingProperty     `json:"bcc"`
	Subject   MappingProperty     `json:"subject"`
	Body      MappingProperty     `json:"body"`
}

type Properties struct {
	Properties MappingProperties `json:"properties"`
}

type IndexMapping struct {
	Name        string     `json:"name"`
	StorageType string     `json:"storage_type"`
	Mappings    Properties `json:"mappings"`
}
