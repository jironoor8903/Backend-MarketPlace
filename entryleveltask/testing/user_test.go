package testing

import (
	"entryleveltask/service"
	"testing"


	"github.com/stretchr/testify/require"
)
	


func TestRegister(t *testing.T){
	db := Setup()
	userServiceTest := service.InitUserService(db)
	
	username, err := userServiceTest.RegisterUser("johndoe", "johndoe123", "johndoe@gmail.com")
	require.NoError(t, err)
	require.NotEmpty(t, username)
	require.Equal(t, username, "johndoe")
}

func TestSignin(t *testing.T){
	db := Setup()
	userServiceTest := service.InitUserService(db)
	
	username, err := userServiceTest.RegisterUser("johndoe", "johndoe123", "johndoe@gmail.com")
	require.NoError(t, err)
	require.NotEmpty(t, username)
}
