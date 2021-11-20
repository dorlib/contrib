// Code generated by entc, DO NOT EDIT.

package category

const (
	// Label holds the string label denoting the category type in the database.
	Label = "category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeBlogPosts holds the string denoting the blog_posts edge name in mutations.
	EdgeBlogPosts = "blog_posts"
	// Table holds the table name of the category in the database.
	Table = "categories"
	// BlogPostsTable is the table that holds the blog_posts relation/edge. The primary key declared below.
	BlogPostsTable = "category_blog_posts"
	// BlogPostsInverseTable is the table name for the BlogPost entity.
	// It exists in this package in order to avoid circular dependency with the "blogpost" package.
	BlogPostsInverseTable = "blog_posts"
)

// Columns holds all SQL columns for category fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
}

var (
	// BlogPostsPrimaryKey and BlogPostsColumn2 are the table columns denoting the
	// primary key for the blog_posts relation (M2M).
	BlogPostsPrimaryKey = []string{"category_id", "blog_post_id"}
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