package main

func (s *Storage) persistFolder(folder *Folder) error {
	res := s.db.Create(folder)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (s *Storage) getFolder(id int) (*Folder, error) {
	folder := new(Folder)
	res := s.db.Find(folder, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return folder, nil
}
