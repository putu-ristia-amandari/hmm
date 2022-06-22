package modules

import (
	"gym-membership/api"
	authV1Controller "gym-membership/api/v1/auth"
	memberV1Controller "gym-membership/api/v1/membership"
	userV1Controller "gym-membership/api/v1/user"

	memberService "gym-membership/business/membership"
	userService "gym-membership/business/user"
	"gym-membership/config"
	"gym-membership/database"

	memberRepository "gym-membership/repository/membership"
	userRepository "gym-membership/repository/user"
)

func RegisterModules(dbCon *database.DatabaseConnection, config *config.AppConfig) api.Controller {
	userPermitRepository := userRepository.UserRepository(dbCon)
	userPremitService := userService.CreateService(userPermitRepository, config)
	userV1PremitController := userV1Controller.CreateController(userPremitService)

	authV1PremitController := authV1Controller.CreateController(userPremitService, config)

	controllers := api.Controller{
		UserV1Controller: userV1PremitController,
		AuthV1Controller: authV1PremitController,
	}

	return controllers
}

func MembershipModules(dbCon *database.DatabaseConnection, config *config.AppConfig) api.Controller {
	memberPermitRepository := memberRepository.MemberRepository(dbCon)
	memberPremitService := memberService.CreateServiceMembership(memberPermitRepository, config)
	memberV1PremitController := memberV1Controller.CreateController(memberPremitService)

	controllers := api.Controller{
		MemberV1Controller: memberV1PremitController,
	}
}
