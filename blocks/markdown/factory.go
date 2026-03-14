package markdown

func NewLineParser() Parser {
	heading := &HeadingParser{}
	todo := &ToDoParser{}
	bullet := &BulletedListParser{}
	divider := &DividerParser{}
	paragraph := &ParagraphParser{}

	heading.SetNext(todo)
	todo.SetNext(bullet)
	bullet.SetNext(divider)
	divider.SetNext(paragraph)

	return heading
}