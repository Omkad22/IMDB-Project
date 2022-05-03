// Code generated by entc, DO NOT EDIT.

package director

const (
	// Label holds the string label denoting the director type in the database.
	Label = "director"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldProfileImage holds the string denoting the profileimage field in the database.
	FieldProfileImage = "profile_image"
	// FieldBornAt holds the string denoting the bornat field in the database.
	FieldBornAt = "born_at"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeMovies holds the string denoting the movies edge name in mutations.
	EdgeMovies = "movies"
	// Table holds the table name of the director in the database.
	Table = "directors"
	// MoviesTable is the table that holds the movies relation/edge.
	MoviesTable = "movies"
	// MoviesInverseTable is the table name for the Movie entity.
	// It exists in this package in order to avoid circular dependency with the "movie" package.
	MoviesInverseTable = "movies"
	// MoviesColumn is the table column denoting the movies relation/edge.
	MoviesColumn = "director_id"
)

// Columns holds all SQL columns for director fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldProfileImage,
	FieldBornAt,
	FieldDescription,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
)
