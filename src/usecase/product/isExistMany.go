package productusecase

func (i *sProductUsecase) IsExistMany(details []*string) error {
	_, err := i.productRepository.IsExistsMany(details)
	if err != nil {
		return err
	}

	return nil
}
