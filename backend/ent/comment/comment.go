// Code generated by entc, DO NOT EDIT.

package comment

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTopic holds the string denoting the topic field in the database.
	FieldTopic = "topic"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// EdgeReview holds the string denoting the review edge name in mutations.
	EdgeReview = "review"
	// Table holds the table name of the comment in the database.
	Table = "comments"
	// ReviewTable is the table that holds the review relation/edge. The primary key declared below.
	ReviewTable = "comment_review"
	// ReviewInverseTable is the table name for the Review entity.
	// It exists in this package in order to avoid circular dependency with the "review" package.
	ReviewInverseTable = "reviews"
)

// Columns holds all SQL columns for comment fields.
var Columns = []string{
	FieldID,
	FieldTopic,
	FieldText,
}

var (
	// ReviewPrimaryKey and ReviewColumn2 are the table columns denoting the
	// primary key for the review relation (M2M).
	ReviewPrimaryKey = []string{"comment_id", "review_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
