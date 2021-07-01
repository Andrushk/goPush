package inmemory

func UserRepo() *UserMemoryRepo {
	repo := UserMemoryRepo{}
	repo.Init()
	return &repo
}
