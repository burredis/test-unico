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

func (s FeiraLivreService) Search(q string) []FeiraLivre {
	return s.repo.Search(q)
}

func (s FeiraLivreService) FindById(id int) FeiraLivre {
	return s.repo.FindById(id)
}

func (s FeiraLivreService) Update(id int, f FeiraLivre) error {
	if err := f.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, f)
}

func (s FeiraLivreService) Remove(id int) error {
	return s.repo.Remove(id)
}
