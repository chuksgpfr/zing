package zing

type Services struct {
	store *Store
}

func NewServices(store *Store) *Services {
	return &Services{
		store: store,
	}
}

func (s *Services) SetCommand(tag, cmd string) (bool, error) {
	// "deploy" --cmd "docker compose build && docker compose push && kubectl rollout restart deploy/{{ .service }} -n {{ .ns }}"
	prompt := false
	// check if tag already exist
	exist, err := s.store.Exist(tag)
	if err != nil {
		return prompt, err
	}
	// if tag exists, prompt user to ask if to update Yes/No
	if exist {
		prompt = true
		return prompt, nil
	}
	// update or store if not exist
	err = s.store.Set(tag, cmd)

	return false, err
}

func (s *Services) ListCommands() (string, error) {
	return s.store.List()
}

func (s *Services) UpdateCommand(tag, cmd string) error {
	// update or store if not exist
	err := s.store.Set(tag, cmd)

	return err
}

func (s *Services) RunCommand(tag string) (string, error) {
	// update or store if not exist
	cmd, err := s.store.Get(tag)

	return cmd, err
}
