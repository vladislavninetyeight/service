package post

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/vladislavninetyeight/service/internal/model"
	"github.com/vladislavninetyeight/service/internal/repositories/post"
	"os"
)

type Service interface {
	Create(ctx context.Context, detail model.PostDetail) error
	GetAll(ctx context.Context) ([]model.Post, error)
	Update(ctx context.Context, post model.PostDetail, id uint) (model.Post, error)
	Delete(ctx context.Context, id uint) error
}

var _ Service = (*service)(nil)

type service struct {
	repo post.Repository
}

func NewPostService(repo post.Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, detail model.PostDetail) error {
	post, err := s.repo.Create(ctx, detail)
	if err != nil {
		return err
	}

	producer, err := sarama.NewSyncProducer([]string{os.Getenv("KAFKA_URL")}, nil)
	if err != nil {
		return err
	}

	postJSON, err := json.Marshal(post)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: os.Getenv("KAFKA_TOPIC"),
		Value: sarama.ByteEncoder(postJSON),
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}
func (s *service) GetAll(ctx context.Context) ([]model.Post, error) {
	posts, err := s.repo.GetAll(ctx)
	if err != nil {
		// TODO
	}

	return posts, nil
}
func (s *service) Get(ctx context.Context, id uint) (model.Post, error) {
	return model.Post{}, nil
}
func (s *service) Update(ctx context.Context, post model.PostDetail, id uint) (model.Post, error) {
	return model.Post{}, nil
}
func (s *service) Delete(ctx context.Context, id uint) error {
	return nil
}
