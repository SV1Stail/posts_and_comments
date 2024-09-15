package graph

import (
	"context"
	"testing"

	"github.com/SV1Stail/test_ozon/db"
	"github.com/SV1Stail/test_ozon/graph/model"
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

func TestPost_1(t *testing.T) {
	db.Connect()
	defer db.ClosePool()
	ctx := context.Background()
	resolver := &queryResolver{}

	p, err := resolver.Post(ctx, "123e4567-e89b-12d3-a456-426614174100")
	id := "123e4567-e89b-12d3-a456-426614174100"
	author := "123e4567-e89b-12d3-a456-426614174000"
	title := "Post 1"
	content := "Content of post 1"
	if p.ID != id || p.Author.ID != author || p.Title != title || p.Content != content {
		t.Fatal("получено неверное значение из БД")
	}
	if err != nil {
		t.Fatalf("Получена ошибка при выполнении: %v, ошибка не ожидалась", err)
	}
}
func TestPost_2(t *testing.T) {
	db.Connect()
	defer db.ClosePool()
	ctx := context.Background()
	resolver := &queryResolver{}

	p, err := resolver.Post(ctx, "123e4567-e89b-12d3-a456-426614174101")
	id := "123e4567-e89b-12d3-a456-426614174101"
	author := "123e4567-e89b-12d3-a456-426614174001"
	title := "Post 2"
	content := "Content of post 2"
	if p.ID != id || p.Author.ID != author || p.Title != title || p.Content != content || p.AllowComments != false {
		t.Fatal("получено неверное значение из БД")
	}
	if err != nil {
		t.Fatalf("Получена ошибка при выполнении: %v, ошибка не ожидалась", err)
	}
}
func TestPosts_1(t *testing.T) {
	db.Connect()
	defer db.ClosePool()
	ctx := context.Background()
	resolver := &queryResolver{}

	ps, err := resolver.Posts(ctx)
	id := "123e4567-e89b-12d3-a456-426614174101"
	author := "123e4567-e89b-12d3-a456-426614174001"
	title := "Post 2"
	content := "Content of post 2"
	pMap := make(map[string]*model.Post)
	for _, p := range ps {
		pMap[p.ID] = p
	}
	if p, ok := pMap[id]; ok {
		if p.ID != id || p.Author.ID != author || p.Title != title || p.Content != content || p.AllowComments != false {
			t.Fatal("получено неверное значение из БД")
		}
	} else {
		t.Fatalf("не удалось записать в карту пост")
	}
	id = "123e4567-e89b-12d3-a456-426614174100"
	author = "123e4567-e89b-12d3-a456-426614174000"
	title = "Post 1"
	content = "Content of post 1"
	if p, ok := pMap[id]; ok {
		if p.ID != id || p.Author.ID != author || p.Title != title || p.Content != content || p.AllowComments != true {
			t.Fatal("получено неверное значение из БД")
		}
	} else {
		t.Fatalf("не удалось записать в карту пост")
	}
	if err != nil {
		t.Fatalf("Получена ошибка при выполнении: %v, ошибка не ожидалась", err)
	}
}

func TestComments_1(t *testing.T) {
	db.Connect()
	defer db.ClosePool()
	ctx := context.Background()
	resolver := &queryResolver{}
	postID := "123e4567-e89b-12d3-a456-426614174100"
	commID := "123e4567-e89b-12d3-a456-426614174200"
	lim := 1
	offset := 0
	comms, err := resolver.Comments(ctx, postID, &lim, &offset)
	if err != nil {
		t.Fatalf("Получена ошибка при запрсое комментов к посту %s | error: %v, ошибка не ожидалась", postID, err)
	}
	if comms[0].ID != commID {
		t.Fatal("получено неверное значение из БД")
	}
}
func TestComments_2(t *testing.T) {
	db.Connect()
	defer db.ClosePool()
	ctx := context.Background()
	resolver := &queryResolver{}
	postID := "123e4567-e89b-12d3-a456-426614174100"
	lim := 1
	offset := 1
	comms, err := resolver.Comments(ctx, postID, &lim, &offset)
	if err != nil {
		t.Fatalf("Получена ошибка при запрсое комментов к посту %s | error: %v, ошибка не ожидалась", postID, err)
	}
	if len(comms) != 0 {
		t.Fatal("получено неверное значение из БД, ожидалось пустое значение")
	}
}
func TestComments_3(t *testing.T) {
	db.Connect()
	defer db.ClosePool()
	ctx := context.Background()
	resolver := &queryResolver{}
	postID := "123e4567-e89b-12d3-a456-426614174100"
	commID := "123e4567-e89b-12d3-a456-426614174200"
	childID := "123e4567-e89b-12d3-a456-426614174201"
	lim := 2
	offset := 0
	comms, err := resolver.Comments(ctx, postID, &lim, &offset)
	if err != nil {
		t.Fatalf("Получена ошибка при запрсое комментов к посту %s | error: %v, ошибка не ожидалась", postID, err)
	}
	if comms[0].ID != commID || comms[0].Children[0].ID != childID {
		t.Fatal("получено неверное значение из БД")
	}
}
