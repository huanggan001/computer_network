package data

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	books "grpc/bookstore/pb"
	"time"
)

const (
	defaultCursor int64 = 0
	defaultSize   int64 = 2
)

// 定义书架
type Shelf struct {
	ID        int64 `gorm:"primarykey"`
	Theme     string
	Size      int64
	Books     []Book
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

type Book struct {
	gorm.Model
	Author  string
	Title   string
	ShelfID int64
}

type BookStore struct {
	books.UnimplementedBookstoreServer
	DB *gorm.DB
}

type Page struct {
	NextID        int64 `json:"next_id"`
	NextTimeAtUTC int64 `json:"next_time_at_utc"`
	PageSize      int64 `json:"page_size"`
}

func (b *BookStore) CreateBooksByShelf(ctx context.Context, req *books.CreateBookRequest) (*books.Book, error) {
	book := Book{
		ShelfID: req.Id,
		Title:   req.Book.Title,
		Author:  req.Book.Author,
	}
	if err := b.DB.WithContext(ctx).Create(&book).Error; err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return req.Book, nil
}

func (b *BookStore) GetShelf(ctx context.Context, req *books.GetShelfRequest) (*books.Shelf, error) {
	var shelf Shelf
	if err := b.DB.WithContext(ctx).First(&shelf, req.Id).Error; err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &books.Shelf{
		Id:    int64(shelf.ID),
		Theme: shelf.Theme,
		Size:  int64(shelf.Size),
	}, nil
}
func (b *BookStore) CreateShelf(ctx context.Context, req *books.CreateShelfRequest) (*books.Shelf, error) {

	shelf := Shelf{
		ID:    req.Shelf.Id,
		Theme: req.Shelf.Theme,
		Size:  req.Shelf.Size,
	}
	if err := b.DB.WithContext(ctx).Create(&shelf).Error; err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return req.Shelf, nil
}
func (b *BookStore) ListBooksByShelf(ctx context.Context, req *books.ListBooksRequest) (*books.ListBooksResponse, error) {
	bs := []Book{}
	cursor, size := defaultCursor, defaultSize
	page := Token(req.GetPageToken()).Decode()
	if req.GetPageToken() != "" {
		cursor = page.NextID
		size = page.PageSize
	}
	if err := b.DB.WithContext(ctx).Where("shelf_id = ? and id > ?", req.Id, cursor).Order("id asc").Limit(int(size + 1)).Find(&bs).Error; err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	realSize := int64(len(bs))
	pageToken := Token("")
	//说明下一页有数据
	if realSize == size+1 {
		p := Page{
			PageSize:      defaultSize,
			NextID:        realSize - 1,
			NextTimeAtUTC: time.Now().UTC().Unix(),
		}
		pageToken = p.Encode()
		realSize -= 1
	}
	var ans []*books.Book
	for i := int64(0); i < realSize; i++ {
		ans = append(ans, &books.Book{
			Id:     int64(bs[i].ID),
			Author: bs[i].Author,
			Title:  bs[i].Title,
		})
	}
	return &books.ListBooksResponse{Books: ans, PageToken: string(pageToken)}, nil
}

type Token string

// Encode 返回分页token
func (p Page) Encode() Token {
	b, err := json.Marshal(p)
	if err != nil {
		return Token("")
	}
	return Token(base64.StdEncoding.EncodeToString(b))
}

// Decode 解析分页信息
func (t Token) Decode() Page {
	var result Page
	if len(t) == 0 {
		return result
	}

	bytes, err := base64.StdEncoding.DecodeString(string(t))
	if err != nil {
		return result
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return result
	}

	return result
}
