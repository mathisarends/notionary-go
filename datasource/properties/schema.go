package datasource

import (
	"encoding/json"
	"fmt"
)

type PropertyType string

const (
	PropertyTypeStatus          PropertyType = "status"
	PropertyTypeSelect          PropertyType = "select"
	PropertyTypeMultiSelect     PropertyType = "multi_select"
	PropertyTypeRelation        PropertyType = "relation"
	PropertyTypeDate            PropertyType = "date"
	PropertyTypeCreatedTime     PropertyType = "created_time"
	PropertyTypeCreatedBy       PropertyType = "created_by"
	PropertyTypeLastEditedTime  PropertyType = "last_edited_time"
	PropertyTypeLastEditedBy    PropertyType = "last_edited_by"
	PropertyTypeLastVisitedTime PropertyType = "last_visited_time"
	PropertyTypeTitle           PropertyType = "title"
	PropertyTypeRichText        PropertyType = "rich_text"
	PropertyTypeURL             PropertyType = "url"
	PropertyTypePeople          PropertyType = "people"
	PropertyTypeNumber          PropertyType = "number"
	PropertyTypeCheckbox        PropertyType = "checkbox"
	PropertyTypeEmail           PropertyType = "email"
	PropertyTypePhoneNumber     PropertyType = "phone_number"
	PropertyTypeFiles           PropertyType = "files"
	PropertyTypeFormula         PropertyType = "formula"
	PropertyTypeRollup          PropertyType = "rollup"
	PropertyTypeUniqueID        PropertyType = "unique_id"
	PropertyTypeButton          PropertyType = "button"
	PropertyTypeLocation        PropertyType = "location"
	PropertyTypePlace           PropertyType = "place"
	PropertyTypeVerification    PropertyType = "verification"
)

type PropertyColor string

const (
	PropertyColorDefault PropertyColor = "default"
	PropertyColorGray    PropertyColor = "gray"
	PropertyColorBrown   PropertyColor = "brown"
	PropertyColorOrange  PropertyColor = "orange"
	PropertyColorYellow  PropertyColor = "yellow"
	PropertyColorGreen   PropertyColor = "green"
	PropertyColorBlue    PropertyColor = "blue"
	PropertyColorPurple  PropertyColor = "purple"
	PropertyColorPink    PropertyColor = "pink"
	PropertyColorRed     PropertyColor = "red"
)

type NumberFormat string

const (
	NumberFormatNumber            NumberFormat = "number"
	NumberFormatNumberWithCommas  NumberFormat = "number_with_commas"
	NumberFormatPercent           NumberFormat = "percent"
	NumberFormatDollar            NumberFormat = "dollar"
	NumberFormatAustralianDollar  NumberFormat = "australian_dollar"
	NumberFormatCanadianDollar    NumberFormat = "canadian_dollar"
	NumberFormatSingaporeDollar   NumberFormat = "singapore_dollar"
	NumberFormatEuro              NumberFormat = "euro"
	NumberFormatPound             NumberFormat = "pound"
	NumberFormatYen               NumberFormat = "yen"
	NumberFormatRuble             NumberFormat = "ruble"
	NumberFormatRupee             NumberFormat = "rupee"
	NumberFormatWon               NumberFormat = "won"
	NumberFormatYuan              NumberFormat = "yuan"
	NumberFormatReal              NumberFormat = "real"
	NumberFormatLira              NumberFormat = "lira"
	NumberFormatRupiah            NumberFormat = "rupiah"
	NumberFormatFranc             NumberFormat = "franc"
	NumberFormatHongKongDollar    NumberFormat = "hong_kong_dollar"
	NumberFormatNewZealandDollar  NumberFormat = "new_zealand_dollar"
	NumberFormatKrona             NumberFormat = "krona"
	NumberFormatNorwegianKrone    NumberFormat = "norwegian_krone"
	NumberFormatMexicanPeso       NumberFormat = "mexican_peso"
	NumberFormatRand              NumberFormat = "rand"
	NumberFormatNewTaiwanDollar   NumberFormat = "new_taiwan_dollar"
	NumberFormatDanishKrone       NumberFormat = "danish_krone"
	NumberFormatZloty             NumberFormat = "zloty"
	NumberFormatBaht              NumberFormat = "baht"
	NumberFormatForint            NumberFormat = "forint"
	NumberFormatKoruna            NumberFormat = "koruna"
	NumberFormatShekel            NumberFormat = "shekel"
	NumberFormatChileanPeso       NumberFormat = "chilean_peso"
	NumberFormatPhilippinePeso    NumberFormat = "philippine_peso"
	NumberFormatDirham            NumberFormat = "dirham"
	NumberFormatColombianPeso     NumberFormat = "colombian_peso"
	NumberFormatRiyal             NumberFormat = "riyal"
	NumberFormatRinggit           NumberFormat = "ringgit"
	NumberFormatLeu               NumberFormat = "leu"
	NumberFormatArgentinePeso     NumberFormat = "argentine_peso"
	NumberFormatUruguayanPeso     NumberFormat = "uruguayan_peso"
	NumberFormatPeruvianSol       NumberFormat = "peruvian_sol"
)

type RelationType string

const (
	RelationTypeSingleProperty RelationType = "single_property"
	RelationTypeDualProperty   RelationType = "dual_property"
)

type RollupFunction string

const (
	RollupFunctionAverage          RollupFunction = "average"
	RollupFunctionChecked          RollupFunction = "checked"
	RollupFunctionCountPerGroup    RollupFunction = "count_per_group"
	RollupFunctionCount            RollupFunction = "count"
	RollupFunctionCountValues      RollupFunction = "count_values"
	RollupFunctionDateRange        RollupFunction = "date_range"
	RollupFunctionEarliestDate     RollupFunction = "earliest_date"
	RollupFunctionEmpty            RollupFunction = "empty"
	RollupFunctionLatestDate       RollupFunction = "latest_date"
	RollupFunctionMax              RollupFunction = "max"
	RollupFunctionMedian           RollupFunction = "median"
	RollupFunctionMin              RollupFunction = "min"
	RollupFunctionNotEmpty         RollupFunction = "not_empty"
	RollupFunctionPercentChecked   RollupFunction = "percent_checked"
	RollupFunctionPercentEmpty     RollupFunction = "percent_empty"
	RollupFunctionPercentNotEmpty  RollupFunction = "percent_not_empty"
	RollupFunctionPercentPerGroup  RollupFunction = "percent_per_group"
	RollupFunctionPercentUnchecked RollupFunction = "percent_unchecked"
	RollupFunctionRange            RollupFunction = "range"
	RollupFunctionUnchecked        RollupFunction = "unchecked"
	RollupFunctionUnique           RollupFunction = "unique"
	RollupFunctionShowOriginal     RollupFunction = "show_original"
	RollupFunctionShowUnique       RollupFunction = "show_unique"
	RollupFunctionSum              RollupFunction = "sum"
)

// ============================================================================
// Shared config models
// ============================================================================

type DataSourcePropertyOption struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Color       PropertyColor `json:"color"`
	Description *string       `json:"description,omitempty"`
}

type DataSourceStatusGroup struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Color     PropertyColor `json:"color"`
	OptionIDs []string      `json:"option_ids"`
}

// ============================================================================
// Base
// ============================================================================

type DataSourceProperty struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description *string      `json:"description,omitempty"`
	Type        PropertyType `json:"type"`
}

// ============================================================================
// Property types
// ============================================================================

// --- Status ---

type DataSourceStatusConfig struct {
	Options []DataSourcePropertyOption `json:"options"`
	Groups  []DataSourceStatusGroup    `json:"groups"`
}

func (c *DataSourceStatusConfig) OptionNames() []string {
	names := make([]string, len(c.Options))
	for i, o := range c.Options {
		names[i] = o.Name
	}
	return names
}

func (c *DataSourceStatusConfig) GroupNames() []string {
	names := make([]string, len(c.Groups))
	for i, g := range c.Groups {
		names[i] = g.Name
	}
	return names
}

type DataSourceStatusProperty struct {
	DataSourceProperty
	Status DataSourceStatusConfig `json:"status"`
}

// --- Select ---

type DataSourceSelectConfig struct {
	Options []DataSourcePropertyOption `json:"options"`
}

func (c *DataSourceSelectConfig) OptionNames() []string {
	names := make([]string, len(c.Options))
	for i, o := range c.Options {
		names[i] = o.Name
	}
	return names
}

type DataSourceSelectProperty struct {
	DataSourceProperty
	Select DataSourceSelectConfig `json:"select"`
}

// --- MultiSelect ---

type DataSourceMultiSelectConfig struct {
	Options []DataSourcePropertyOption `json:"options"`
}

func (c *DataSourceMultiSelectConfig) OptionNames() []string {
	names := make([]string, len(c.Options))
	for i, o := range c.Options {
		names[i] = o.Name
	}
	return names
}

type DataSourceMultiSelectProperty struct {
	DataSourceProperty
	MultiSelect DataSourceMultiSelectConfig `json:"multi_select"`
}

// --- Relation ---

type DataSourceRelationConfig struct {
	DataSourceID   *string                `json:"database_id,omitempty"`
	Type           RelationType           `json:"type"`
	SingleProperty map[string]interface{} `json:"single_property,omitempty"`
}

type DataSourceRelationProperty struct {
	DataSourceProperty
	Relation DataSourceRelationConfig `json:"relation"`
}

func (p *DataSourceRelationProperty) RelatedDataSourceID() *string {
	return p.Relation.DataSourceID
}

// --- Date ---

type DataSourceDateProperty struct {
	DataSourceProperty
	Date struct{} `json:"date"`
}

// --- CreatedTime ---

type DataSourceCreatedTimeProperty struct {
	DataSourceProperty
	CreatedTime struct{} `json:"created_time"`
}

// --- CreatedBy ---

type DataSourceCreatedByProperty struct {
	DataSourceProperty
	CreatedBy struct{} `json:"created_by"`
}

// --- LastEditedTime ---

type DataSourceLastEditedTimeProperty struct {
	DataSourceProperty
	LastEditedTime struct{} `json:"last_edited_time"`
}

// --- LastEditedBy ---

type DataSourceLastEditedByProperty struct {
	DataSourceProperty
	LastEditedBy struct{} `json:"last_edited_by"`
}

// --- LastVisitedTime ---

type DataSourceLastVisitedTimeProperty struct {
	DataSourceProperty
	LastVisitedTime struct{} `json:"last_visited_time"`
}

// --- Title ---

type DataSourceTitleProperty struct {
	DataSourceProperty
	Title struct{} `json:"title"`
}

// --- RichText ---

type DataSourceRichTextProperty struct {
	DataSourceProperty
	RichText struct{} `json:"rich_text"`
}

// --- URL ---

type DataSourceURLProperty struct {
	DataSourceProperty
	URL struct{} `json:"url"`
}

// --- People ---

type DataSourcePeopleProperty struct {
	DataSourceProperty
	People struct{} `json:"people"`
}

// --- Number ---

type DataSourceNumberConfig struct {
	Format NumberFormat `json:"format"`
}

type DataSourceNumberProperty struct {
	DataSourceProperty
	Number DataSourceNumberConfig `json:"number"`
}

func (p *DataSourceNumberProperty) NumberFormat() NumberFormat {
	return p.Number.Format
}

// --- Checkbox ---

type DataSourceCheckboxProperty struct {
	DataSourceProperty
	Checkbox struct{} `json:"checkbox"`
}

// --- Email ---

type DataSourceEmailProperty struct {
	DataSourceProperty
	Email struct{} `json:"email"`
}

// --- PhoneNumber ---

type DataSourcePhoneNumberProperty struct {
	DataSourceProperty
	PhoneNumber struct{} `json:"phone_number"`
}

// --- Files ---

type DataSourceFilesProperty struct {
	DataSourceProperty
	Files struct{} `json:"files"`
}

// --- Formula ---

type DataSourceFormulaConfig struct {
	Expression string `json:"expression"`
}

type DataSourceFormulaProperty struct {
	DataSourceProperty
	Formula DataSourceFormulaConfig `json:"formula"`
}

func (p *DataSourceFormulaProperty) Expression() string {
	return p.Formula.Expression
}

// --- Rollup ---

type DataSourceRollupConfig struct {
	Function             RollupFunction `json:"function"`
	RelationPropertyID   string         `json:"relation_property_id"`
	RelationPropertyName string         `json:"relation_property_name"`
	RollupPropertyID     string         `json:"rollup_property_id"`
	RollupPropertyName   string         `json:"rollup_property_name"`
}

type DataSourceRollupProperty struct {
	DataSourceProperty
	Rollup DataSourceRollupConfig `json:"rollup"`
}

func (p *DataSourceRollupProperty) RollupFunction() RollupFunction {
	return p.Rollup.Function
}

// --- UniqueID ---

type DataSourceUniqueIDConfig struct {
	Prefix *string `json:"prefix,omitempty"`
}

type DataSourceUniqueIDProperty struct {
	DataSourceProperty
	UniqueID DataSourceUniqueIDConfig `json:"unique_id"`
}

func (p *DataSourceUniqueIDProperty) Prefix() *string {
	return p.UniqueID.Prefix
}

// --- Button ---

type DataSourceButtonProperty struct {
	DataSourceProperty
	Button struct{} `json:"button"`
}

// --- Location ---

type DataSourceLocationProperty struct {
	DataSourceProperty
	Location struct{} `json:"location"`
}

// --- Place ---

type DataSourcePlaceProperty struct {
	DataSourceProperty
	Place struct{} `json:"place"`
}

// --- Verification ---

type DataSourceVerificationProperty struct {
	DataSourceProperty
	Verification struct{} `json:"verification"`
}

// --- Unknown (extra fields allowed) ---

type DataSourceUnknownProperty struct {
	Extra map[string]interface{} `json:"-"`
}

func (p *DataSourceUnknownProperty) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &p.Extra)
}

// ============================================================================
// AnyDataSourceProperty — discriminated union
// ============================================================================

// AnyDataSourceProperty is the interface satisfied by every concrete property type.
type AnyDataSourceProperty interface {
	GetType() PropertyType
}

// Implement AnyDataSourceProperty for all concrete types.
func (p *DataSourceStatusProperty) GetType() PropertyType          { return PropertyTypeStatus }
func (p *DataSourceSelectProperty) GetType() PropertyType          { return PropertyTypeSelect }
func (p *DataSourceMultiSelectProperty) GetType() PropertyType     { return PropertyTypeMultiSelect }
func (p *DataSourceRelationProperty) GetType() PropertyType        { return PropertyTypeRelation }
func (p *DataSourceDateProperty) GetType() PropertyType            { return PropertyTypeDate }
func (p *DataSourceCreatedTimeProperty) GetType() PropertyType     { return PropertyTypeCreatedTime }
func (p *DataSourceCreatedByProperty) GetType() PropertyType       { return PropertyTypeCreatedBy }
func (p *DataSourceLastEditedTimeProperty) GetType() PropertyType  { return PropertyTypeLastEditedTime }
func (p *DataSourceLastEditedByProperty) GetType() PropertyType    { return PropertyTypeLastEditedBy }
func (p *DataSourceLastVisitedTimeProperty) GetType() PropertyType { return PropertyTypeLastVisitedTime }
func (p *DataSourceTitleProperty) GetType() PropertyType           { return PropertyTypeTitle }
func (p *DataSourceRichTextProperty) GetType() PropertyType        { return PropertyTypeRichText }
func (p *DataSourceURLProperty) GetType() PropertyType             { return PropertyTypeURL }
func (p *DataSourcePeopleProperty) GetType() PropertyType          { return PropertyTypePeople }
func (p *DataSourceNumberProperty) GetType() PropertyType          { return PropertyTypeNumber }
func (p *DataSourceCheckboxProperty) GetType() PropertyType        { return PropertyTypeCheckbox }
func (p *DataSourceEmailProperty) GetType() PropertyType           { return PropertyTypeEmail }
func (p *DataSourcePhoneNumberProperty) GetType() PropertyType     { return PropertyTypePhoneNumber }
func (p *DataSourceFilesProperty) GetType() PropertyType           { return PropertyTypeFiles }
func (p *DataSourceFormulaProperty) GetType() PropertyType         { return PropertyTypeFormula }
func (p *DataSourceRollupProperty) GetType() PropertyType          { return PropertyTypeRollup }
func (p *DataSourceUniqueIDProperty) GetType() PropertyType        { return PropertyTypeUniqueID }
func (p *DataSourceButtonProperty) GetType() PropertyType          { return PropertyTypeButton }
func (p *DataSourceLocationProperty) GetType() PropertyType        { return PropertyTypeLocation }
func (p *DataSourcePlaceProperty) GetType() PropertyType           { return PropertyTypePlace }
func (p *DataSourceVerificationProperty) GetType() PropertyType    { return PropertyTypeVerification }
func (p *DataSourceUnknownProperty) GetType() PropertyType         { return "" }

// ============================================================================
// DataSourcePropertyWrapper — JSON unmarshaling for AnyDataSourceProperty
// ============================================================================

// DataSourcePropertyWrapper wraps AnyDataSourceProperty for JSON unmarshaling.
// It inspects the "type" discriminator field and deserializes into the correct concrete type.
type DataSourcePropertyWrapper struct {
	Value AnyDataSourceProperty
}

func (w *DataSourcePropertyWrapper) UnmarshalJSON(data []byte) error {
	var discriminator struct {
		Type PropertyType `json:"type"`
	}
	if err := json.Unmarshal(data, &discriminator); err != nil {
		return fmt.Errorf("reading type discriminator: %w", err)
	}

	switch discriminator.Type {
	case PropertyTypeStatus:
		var p DataSourceStatusProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeSelect:
		var p DataSourceSelectProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeMultiSelect:
		var p DataSourceMultiSelectProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeRelation:
		var p DataSourceRelationProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeDate:
		var p DataSourceDateProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeCreatedTime:
		var p DataSourceCreatedTimeProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeCreatedBy:
		var p DataSourceCreatedByProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeLastEditedTime:
		var p DataSourceLastEditedTimeProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeLastEditedBy:
		var p DataSourceLastEditedByProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeLastVisitedTime:
		var p DataSourceLastVisitedTimeProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeTitle:
		var p DataSourceTitleProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeRichText:
		var p DataSourceRichTextProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeURL:
		var p DataSourceURLProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypePeople:
		var p DataSourcePeopleProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeNumber:
		var p DataSourceNumberProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeCheckbox:
		var p DataSourceCheckboxProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeEmail:
		var p DataSourceEmailProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypePhoneNumber:
		var p DataSourcePhoneNumberProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeFiles:
		var p DataSourceFilesProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeFormula:
		var p DataSourceFormulaProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeRollup:
		var p DataSourceRollupProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeUniqueID:
		var p DataSourceUniqueIDProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeButton:
		var p DataSourceButtonProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeLocation:
		var p DataSourceLocationProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypePlace:
		var p DataSourcePlaceProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	case PropertyTypeVerification:
		var p DataSourceVerificationProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	default:
		var p DataSourceUnknownProperty
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		w.Value = &p
	}
	return nil
}

// ============================================================================
// Generic constraint (equivalent to TypeVar bound=DataSourceProperty)
// ============================================================================

// DataSourcePropertyConstraint is the generic type constraint equivalent to
// Python's TypeVar("DataSourcePropertyT", bound=DataSourceProperty).
type DataSourcePropertyConstraint interface {
	AnyDataSourceProperty
	// Embed any shared accessors if needed, e.g.:
	// GetID() string
	// GetName() string
}