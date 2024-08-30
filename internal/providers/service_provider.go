package providers

import (
	"context"
	repPost "github.com/vladislavninetyeight/service/internal/repositories/post"
	repUser "github.com/vladislavninetyeight/service/internal/repositories/user"
	post "github.com/vladislavninetyeight/service/internal/services/post"
	"github.com/vladislavninetyeight/service/internal/services/user"
	"github.com/vladislavninetyeight/service/pkg/client/postgresql"
	"sync"
)

type ServiceProvider struct {
	userService user.Service
	postService post.Service
}

var once sync.Once

var serviceProvider *ServiceProvider

func Container() *ServiceProvider {
	once.Do(func() {
		serviceProvider = &ServiceProvider{}
	})

	return serviceProvider
}

func (sp *ServiceProvider) GetUserService() user.Service {
	if sp.userService == nil {
		conf := postgresql.GetConfig()

		client, err := postgresql.NewClient(context.TODO(), conf)
		if err != nil {
			panic(err)
		}
		sp.userService = user.NewUserService(repUser.NewUserRepository(client))
	}

	return sp.userService
}
func (sp *ServiceProvider) GetPostService() post.Service {
	if sp.postService == nil {
		conf := postgresql.GetConfig()
		client, err := postgresql.NewClient(context.TODO(), conf)
		if err != nil {
			panic(err)
		}
		sp.postService = post.NewPostService(repPost.NewPostRepository(client))
	}

	return sp.postService
}
