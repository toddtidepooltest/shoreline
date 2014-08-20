package clients

import (
	"./../models"
	"errors"
)

type MockStoreClient struct {
	salt            string
	doBad           bool
	returnDifferent bool
}

func NewMockStoreClient(salt string, returnDifferent, doBad bool) *MockStoreClient {
	return &MockStoreClient{salt: salt, doBad: doBad, returnDifferent: returnDifferent}
}

func (d MockStoreClient) Close() {}

func (d MockStoreClient) Ping() error {
	if d.doBad {
		return errors.New("Session failure")
	}
	return nil
}

func (d MockStoreClient) UpsertUser(user *models.User) error {
	if d.doBad {
		return errors.New("UpsertUser failure")
	}
	return nil
}

func (d MockStoreClient) FindUsers(user *models.User) (found []*models.User, err error) {
	//`find` a pretend one we just made

	if d.doBad {
		return found, errors.New("FindUsers failure")
	}

	if d.returnDifferent {
		return []*models.User{}, nil
	}

	if user.Pw != "" && user.Name != "" {
		found, _ := models.NewUser(&models.UserDetail{Name: user.Name, Pw: user.Pw, Emails: []string{}}, d.salt)
		return []*models.User{found}, nil
	}

	return []*models.User{user}, nil

}

func (d MockStoreClient) FindUser(user *models.User) (found *models.User, err error) {

	if d.doBad {
		return found, errors.New("FindUser failure")
	}
	//`find` a pretend one we just made

	if d.returnDifferent {
		other, _ := models.NewUser(&models.UserDetail{Name: "Some One Else", Pw: "s0m30n33ls3", Emails: []string{}}, d.salt)
		return other, nil
	}

	if user.Pw != "" && user.Name != "" {
		found, _ := models.NewUser(&models.UserDetail{Name: user.Name, Pw: user.Pw, Emails: []string{}}, d.salt)
		return found, nil
	}
	return user, nil
}

func (d MockStoreClient) RemoveUser(user *models.User) error {
	if d.doBad {
		return errors.New("RemoveUser failure")
	}
	return nil
}

func (d MockStoreClient) AddToken(token *models.SessionToken) error {
	if d.doBad {
		return errors.New("AddToken failure")
	}
	return nil
}

func (d MockStoreClient) FindToken(token *models.SessionToken) (*models.SessionToken, error) {
	if d.doBad {
		return token, errors.New("FindToken failure")
	}
	//`find` a pretend one we just made
	return token, nil
}

func (d MockStoreClient) RemoveToken(token *models.SessionToken) error {
	if d.doBad {
		return errors.New("RemoveToken failure")
	}
	return nil
}
