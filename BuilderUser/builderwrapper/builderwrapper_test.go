package builderwrapper

import (
	"github.com/tavomartinez88/go/BuilderUser/builders"
	"github.com/tavomartinez88/go/BuilderUser/domain"
	"testing"
)

func get_map() map[string]string {
	data := make(map[string]string)
	data["firstname"]="Gustavo"
	data["lastname"]="Martinez"
	data["email"]="abc@gmail.com"
	data["password"]="123456"
	data["phone"]="351560056"
	return data
}

func check_post_build(user *domain.User, data map[string]string ) bool {
	return user.Email == data["email"] &&
			user.Phone == data["phone"] &&
			user.FirstName == data["firstname"] &&
			user.LastName == data["lastname"]
}

func TestGeneratorUser_BuildUser(t *testing.T) {
	data := get_map()
	generator := &GeneratorUser{}
	user := generator.BuildUser(builders.BuilderUser{}, data)

	if !check_post_build(user, data) {
		t.Error("Expected info", data, "Got", user)
	}
}

func Benchmark(b *testing.B) {
	data := get_map()
	for i := 0; i < b.N; i++ {
		generator := &GeneratorUser{}
		generator.BuildUser(builders.BuilderUser{}, data)
	}
}
