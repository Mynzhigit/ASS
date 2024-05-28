package usecase


type AdminUsecase interface {
	GetAdminByID(adminID int) (*entity.Admin, error)
}

type adminUsecase struct {
	adminRepo repository.AdminRepository
}

func NewAdminUsecase(adminRepo repository.AdminRepository) AdminUsecase {
	return &adminUsecase{
		adminRepo: adminRepo,
	}
}

func (u *adminUsecase) GetAdminByID(adminID int) (*entity.Admin, error) {
	admin, err := u.adminRepo.GetByID(adminID)
	if err != nil {
		return nil, err
	}
	return admin, nil
}
