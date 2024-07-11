package providers

import (
	"github.com/vladislavninetyeight/service/tree/main/internal/model/internal/app/server"
	repPost "github.com/vladislavninetyeight/service/tree/main/internal/model/internal/repositories/post"
	repUser "github.com/vladislavninetyeight/service/tree/main/internal/model/internal/repositories/user"
	"github.com/vladislavninetyeight/service/tree/main/internal/model/internal/services"
	post "github.com/vladislavninetyeight/service/tree/main/internal/model/internal/services/post"
	user "github.com/vladislavninetyeight/service/tree/main/internal/model/internal/services/user"
)

type ServiceProvider struct {
	userService services.UserService
	postService services.PostService
	server      *server.Server
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (sp *ServiceProvider) GetUserService() services.UserService {
	if sp.userService == nil {
		sp.userService = user.NewUserService(repUser.NewUserRepository())
	}

	return sp.userService
}
func (sp *ServiceProvider) GetPostService() services.PostService {
	if sp.postService == nil {
		sp.postService = post.NewPostService(repPost.NewPostRepository())
	}

	return sp.postService
}

func (sp *ServiceProvider) GetHTTPServer() *server.Server {
	if sp.server == nil {
		sp.server = server.NewServer(sp)
	}

	return sp.server
}
