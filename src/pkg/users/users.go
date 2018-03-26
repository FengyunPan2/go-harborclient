package users

import (
	"log"

	"fmt"
	"github.com/pkg/errors"
	"hubimage/src/client"
	"hubimage/src/config"
	"hubimage/src/model"
	"hubimage/swagger"
)

const (
	UserEmailExample = "example.com"
)

func requestToUser(request *model.RequestUserCreate) *swagger.User {
	var user swagger.User
	if request.Email == "" {
		email := fmt.Sprintf("%s@%s", request.Username, UserEmailExample)
		request.Email = email
	}
	if request.Realname == "" {
		request.Realname = request.Username
	}
	if request.Comment == "" {
		request.Comment = request.Username
	}

	user.Username = request.Username
	user.Password = request.Password
	user.Email = request.Email
	user.Realname = request.Realname
	user.Comment = request.Comment
	user.RoleName = request.RoleName
	user.HasAdminRole = request.HasAdminRole

	return &user
}

// CreateUser create user
func CreateUser(request *model.RequestUserCreate) (*swagger.User, error) {
	log.Printf("call CreateUser:%#v", request)
	conf := config.GetConfig()
	if conf.Username != "admin" {
		return &swagger.User{}, errors.New("only admin can create user.")
	}
	if request.Username == "" {
		return &swagger.User{}, errors.New("please specified user name.")
	}
	if request.Password == "" {
		return &swagger.User{}, errors.New("please specified user password.")
	}

	userRequest := requestToUser(request)

	harborClient := client.GetHarborClient()
	APIResponse, err := harborClient.UsersPost(*userRequest)
	if err != nil {
		return &swagger.User{}, err
	}
	log.Printf("post APIResponse:%#v", APIResponse.Response)
	if APIResponse.Response.StatusCode/100 != 2 {
		msg := fmt.Sprintf("Failed to create user %s: %v", request.Username, APIResponse.Response.Status)
		return &swagger.User{}, errors.New(msg)
	}

	// get
	user, err := getUser(harborClient, request.Username, request.Email)
	if err != nil {
		return &swagger.User{}, err
	}

	log.Printf("get user:%#v", user)
	return user, nil
}

// GetUserList list user
func GetUserList(request *model.RequestUserList) ([]swagger.User, error) {
	log.Printf("call GetUserList:%#v", request)
	conf := config.GetConfig()
	if conf.Username != "admin" {
		return []swagger.User{}, errors.New("only admin can list user.")
	}
	if request.Page == 0 {
		request.Page = 1
	}
	if request.PageSize == 0 {
		request.PageSize = 100
	}
	harborClient := client.GetHarborClient()

	users, APIResponse, err := harborClient.UsersGet(request.Username, request.Email, request.Page, request.PageSize)
	if err != nil {
		log.Printf("failed to GetUserList:%v", err)
		return []swagger.User{}, err
	}
	log.Printf("list APIResponse:%#v", APIResponse.Response)

	log.Printf("users:%#v", users)
	return users, nil
}

func getUser(harborClient *swagger.ProductsApi, name string, email string) (*swagger.User, error) {
	log.Printf("call getUser")
	users, APIResponse, err := harborClient.UsersGet(name, email, 1, 20)
	if err != nil {
		log.Printf("Failed to getuser:%v", err)
		return &swagger.User{}, err
	}
	log.Printf("get APIResponse:%#v", APIResponse.Response)

	if len(users) == 0 {
		msg := fmt.Sprintf("Error: the user %s not found", name)
		return &swagger.User{}, errors.New(msg)
	}

	for _, item := range users {
		if email != "" {
			if item.Username == name && item.Email == email {
				return &item, nil
			}
		} else {
			if item.Username == name {
				return &item, nil
			}
		}
	}

	msg := fmt.Sprintf("Failed to get user %s", name)
	return &swagger.User{}, errors.New(msg)
}

// HeadUser check if user exist
func HeadUser(name string) (bool, error) {
	log.Printf("call HeadUser:%s", name)
	harborClient := client.GetHarborClient()

	_, err := getUser(harborClient, name, "")
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetUserDetail get detail user
func GetUserDetail(name string) (*swagger.User, error) {
	log.Printf("call GetUserDetail:%s", name)
	conf := config.GetConfig()
	if conf.Username != "admin" {
		return &swagger.User{}, errors.New("only admin can get user.")
	}
	harborClient := client.GetHarborClient()

	user, err := getUser(harborClient, name, "")
	if err != nil {
		return &swagger.User{}, err
	}

	log.Printf("User:%#v", user)
	return user, nil
}

// DeleteUser delete user
func DeleteUser(name string) error {
	log.Printf("call DeleteUser: %s", name)
	conf := config.GetConfig()
	if conf.Username != "admin" {
		return errors.New("only admin can delete user.")
	}
	harborClient := client.GetHarborClient()
	user, err := getUser(harborClient, name, "")
	if err != nil {
		return err
	}

	_, err = harborClient.UsersUserIdDelete(user.UserId)
	if err != nil {
		return err
	}

	log.Printf("success to delete user:%#v", user)
	return nil
}
