package feiralivre

type FeiraLivreService struct {
	repo FeiraLivreRepository
}

func NewService(repo FeiraLivreRepository) FeiraLivreService {
	return FeiraLivreService{
		repo: repo,
	}
}

func (s FeiraLivreService) Create(f FeiraLivre) error {
	if err := f.Validate(); err != nil {
		return err
	}

	return s.repo.Insert(f)
}

func (s FeiraLivreService) List() []FeiraLivre {
	return s.repo.List()
}
