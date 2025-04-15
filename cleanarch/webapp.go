package main

import (
	"cleanarch/controller"
	"cleanarch/repository"
	"cleanarch/router"
	"cleanarch/service"
	"fmt"
	"net/http"
)

var (
	// Dependency injection and separation of concern!!
	// https://www.youtube.com/watch?v=Yg_ae0UvCv4&list=PL7Bs8ngpweC6KN8g1_LS4Be0bWXB23UKB&index=3

	// Repository is an independent implementation (i.e. Firebase or Mongo)
	// Service is independent of the repository implementation.
	// Controller takes the service, where service can have different implementation.

	// Application independent of http frameworks, databases, UI (due to REST APIs).
	// Application is testable since we implement all layers with an interface.
	// Interfaces can be mocked independent of their implementation.
	// 1. Test the repository using an in-memory database.
	// 2. Unit test the service by mocking the repository.
	// 3. Unit test the controller by mocking the service.
	// 4. Service is where the use case or business logic is implemented.
	// 		In order to test the business logic without depending on external dependencies,
	// 		we mock the repository interface.
	postRepository repository.PostRepository = repository.NewSQLiteRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	const port string = ":8080"

	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPosts)
	httpRouter.SERVE(port)
}
