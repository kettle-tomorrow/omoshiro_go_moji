package services

type OmoshiroGoMojiService struct{}

// func (OmoshiroGoMojiService) SetOmoshiroGoMojiService(omoshiroGoMoji *model.OmoshiroGoMojiService) error {
// 	_, err := DbEngine.Insert(omoshiroGoMoji)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (OmoshiroGoMojiService) GetOmoshiroGoMojiList() []model.OmoshiroGoMoji {
// 	tests := make([]model.OmoshiroGoMoji, 0)
// 	err := DbEngine.Distinct("id", "name").Limit(10, 0).Find(&tests)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return tests
// }

// func (OmoshiroGoMojiService) UpdateOmoshiroGoMoji(newOmoshiroGoMoji *model.OmoshiroGoMoji) error {
// 	_, err := DbEngine.Id(newOmoshiroGoMoji.Id).Update(newOmoshiroGoMoji)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (OmoshiroGoMojiService) DeleteOmoshiroGoMoji(id int) error {
// 	omoshiroGoMoji := new(model.OmoshiroGoMoji)
// 	_, err := DbEngine.Id(id).Delete(omoshiroGoMoji)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
