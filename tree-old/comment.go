package tree

// Comment are comment in your node tree
type Comment struct {
	*CharacterData
}

func createComment(data string) *Comment {
	return &Comment{
		&CharacterData{
			Node: &Node{
				nodeType: CommentNode,
			},
			data: []rune(data),
		},
	}
}
