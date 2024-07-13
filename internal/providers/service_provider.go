package providers

import (
	repPost "github.com/vladislavninetyeight/service/internal/repositories/post"
	repUser "github.com/vladislavninetyeight/service/internal/repositories/user"
	"github.com/vladislavninetyeight/service/internal/services"
	post "github.com/vladislavninetyeight/service/internal/services/post"
	user "github.com/vladislavninetyeight/service/internal/services/user"
	"net/http"
)

type ServiceProvider struct {
	userService services.UserService
	postService services.PostService
	http        *http.Server
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

func (sp *ServiceProvider) GetHTTPServer() *http.Server {
	if sp.http == nil {
		sp.http = &http.Server{
			Addr: ":8080",
		}
	}

	return sp.http
}
