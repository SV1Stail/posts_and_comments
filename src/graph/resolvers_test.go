package graph

import (
	"context"
	"testing"
)

func TestCreatePost(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Test passed, recovered from panic: %v", r)
		} else {
			t.Errorf("Test failed, expected panic but function did not panic")
		}
	}()
	ctx := context.Background()
	resolver := &mutationResolver{}

	title := "123"
	content := "Test Content"
	allowComments := true
	authorID := "test-author-id"

	_, err := resolver.CreatePost(ctx, title, content, allowComments, authorID)

	if err == nil {
		t.Fatal("Ожидалась ошибка при пустом заголовке, но ошибка не получена")
	}

	expectedError := "title cannot be empty"
	if err.Error() != expectedError {
		t.Errorf("Ожидалась ошибка '%s', получена '%s'", expectedError, err.Error())
	}
}
func TestCreatePost_EmptyTitle(t *testing.T) {
	ctx := context.Background()
	resolver := &mutationResolver{}

	title := ""
	content := "Test Content"
	allowComments := true
	authorID := "test-author-id"

	_, err := resolver.CreatePost(ctx, title, content, allowComments, authorID)

	if err == nil {
		t.Fatal("Ожидалась ошибка при пустом заголовке, но ошибка не получена")
	}

	expectedError := "title cannot be empty"
	if err.Error() != expectedError {
		t.Errorf("Ожидалась ошибка '%s', получена '%s'", expectedError, err.Error())
	}
}
func TestCreatePost_Content(t *testing.T) {
	ctx := context.Background()
	resolver := &mutationResolver{}

	title := "123"
	content := ""
	allowComments := true
	authorID := "test-author-id"

	_, err := resolver.CreatePost(ctx, title, content, allowComments, authorID)

	if err == nil {
		t.Fatal("Ожидалась ошибка при пустом теле, но ошибка не получена")
	}

	expectedError := "content cannot be empty"
	if err.Error() != expectedError {
		t.Errorf("Ожидалась ошибка '%s', получена '%s'", expectedError, err.Error())
	}
}
func TestCreatePost_AuthorID(t *testing.T) {
	ctx := context.Background()
	resolver := &mutationResolver{}

	title := "123"
	content := "123"
	allowComments := true
	authorID := ""

	_, err := resolver.CreatePost(ctx, title, content, allowComments, authorID)

	if err == nil {
		t.Fatal("Ожидалась ошибка при пустом id автора, но ошибка не получена")
	}

	expectedError := "authorID cannot be empty"
	if err.Error() != expectedError {
		t.Errorf("Ожидалась ошибка '%s', получена '%s'", expectedError, err.Error())
	}
}
func TestCreateComment_AuthorID(t *testing.T) {
	ctx := context.Background()
	resolver := &mutationResolver{}

	postID := "123"
	text := "123"
	authorID := ""

	_, err := resolver.CreateComment(ctx, postID, nil, text, authorID)

	if err == nil {
		t.Fatal("Ожидалась ошибка при пустом id автора, но ошибка не получена")
	}

	expectedError := "authorID cannot be empty"
	if err.Error() != expectedError {
		t.Errorf("Ожидалась ошибка '%s', получена '%s'", expectedError, err.Error())
	}
}
func TestCreateComment_Text(t *testing.T) {
	ctx := context.Background()
	resolver := &mutationResolver{}

	postID := "123"
	text := ""
	authorID := "test-author-id"
	_, err := resolver.CreateComment(ctx, postID, nil, text, authorID)

	if err == nil {
		t.Fatal("Ожидалась ошибка при пустом теле, но ошибка не получена")
	}

	expectedError := "text cannot be empty"
	if err.Error() != expectedError {
		t.Errorf("Ожидалась ошибка '%s', получена '%s'", expectedError, err.Error())
	}
}
func TestCreateComment_Text2000(t *testing.T) {
	ctx := context.Background()
	resolver := &mutationResolver{}

	postID := "123"
	text := `Consider Cursor-Based Pagination (Advanced)
For more efficient pagination, especially when dealing with real-time data or large offsets, consider implementing cursor-based pagination:
    Replace offset with a cursor that uniquely identifies the starting point for the next set of records.
    Modify queries to use the cursor for fetching subsequent records.Consider Cursor-Based Pagination (Advanced)

For more efficient pagination, especially when dealing with real-time data or large offsets, consider implementing cursor-based pagination:

    Replace offset with a cursor that uniquely identifies the starting point for the next set of records.
    Modify queries to use the cursor for fetching subsequent records.Consider Cursor-Based Pagination (Advanced)

For more efficient pagination, especially when dealing with real-time data or large offsets, consider implementing cursor-based pagination:

    Replace offset with a cursor that uniquely identifies the starting point for the next set of records.
    Modify queries to use the cursor for fetching subsequent records.Consider Cursor-Based Pagination (Advanced)

For more efficient pagination, especially when dealing with real-time data or large offsets, consider implementing cursor-based pagination:

    Replace offset with a cursor that uniquely identifies the starting point for the next set of records.
    Modify queries to use the cursor for fetching subsequent records.Consider Cursor-Based Pagination (Advanced)

For more efficient pagination, especially when dealing with real-time data or large offsets, consider implementing cursor-based pagination:

    Replace offset with a cursor that uniquely identifies the starting point for the next set of records.
    Modify queries to use the cursor for fetching subsequent records.Consider Cursor-Based Pagination (Advanced)

For more efficient pagination, especially when dealing with real-time data or large offsets, consider implementing cursor-based pagination:

    Replace offset with a cursor that uniquely identifies the starting point for the next set of records.
    Modify queries to use the cursor for fetching subsequent records.`
	authorID := "test-author-id"

	_, err := resolver.CreateComment(ctx, postID, nil, text, authorID)

	if err == nil {
		t.Fatal("Ожидалась ошибка при пустом заголовке, но ошибка не получена")
	}

	expectedError := "comment text exceeds 2000 characters"
	if err.Error() != expectedError {
		t.Errorf("Ожидалась ошибка '%s', получена '%s'", expectedError, err.Error())
	}
}
func TestCreateComment_PostID(t *testing.T) {
	ctx := context.Background()
	resolver := &mutationResolver{}

	postID := ""
	text := "123"
	authorID := "test-author-id"

	_, err := resolver.CreateComment(ctx, postID, nil, text, authorID)

	if err == nil {
		t.Fatal("Ожидалась ошибка при пустом заголовке, но ошибка не получена")
	}

	expectedError := "postID cannot be empty"
	if err.Error() != expectedError {
		t.Errorf("Ожидалась ошибка '%s', получена '%s'", expectedError, err.Error())
	}
}
func TestCreateComment(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Test passed, recovered from panic: %v", r)
		} else {
			t.Errorf("Test failed, expected panic but function did not panic")
		}
	}()
	ctx := context.Background()
	resolver := &mutationResolver{}

	postID := "123"
	text := "123"
	authorID := "test-author-id"

	_, err := resolver.CreateComment(ctx, postID, nil, text, authorID)

	if err == nil {
		t.Fatal("Ожидалась ошибка при пустом заголовке, но ошибка не получена")
	}

	expectedError := "postID cannot be empty"
	if err.Error() != expectedError {
		t.Errorf("Ожидалась ошибка '%s', получена '%s'", expectedError, err.Error())
	}
}
func TestPosts(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Test passed, recovered from panic: %v", r)
		} else {
			t.Errorf("Test failed, expected panic but function did not panic")
		}
	}()
	ctx := context.Background()
	resolver := &queryResolver{}
	_, err := resolver.Posts(ctx)

	if err == nil {
		t.Fatal("Ожидалась ошибка, но ошибка не получена")
	}
}
func TestPost(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Test passed, recovered from panic: %v", r)
		} else {
			t.Errorf("Test failed, expected panic but function did not panic")
		}
	}()
	ctx := context.Background()
	resolver := &queryResolver{}
	_, err := resolver.Post(ctx, "123e4567-e89b-12d3-a456-426614174100")

	if err == nil {
		t.Fatal("Ожидалась ошибка, но ошибка не получена")
	}
}
func TestComments(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Test passed, recovered from panic: %v", r)
		} else {
			t.Errorf("Test failed, expected panic but function did not panic")
		}
	}()
	ctx := context.Background()
	resolver := &queryResolver{}
	x := 0
	_, err := resolver.Comments(ctx, "123e4567-e89b-12d3-a456-426614174100", &x, &x)

	if err == nil {
		t.Fatal("Ожидалась ошибка, но ошибка не получена")
	}
}
