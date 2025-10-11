package services

import (
	"api/dto"
	"api/entities"
	usersRepo "api/repositories/users"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func CreateUser(usersRepo usersRepo.UsersRepository, userDTO dto.CreateUserDTO) (*entities.User, error) {
	if userDTO.DiscordID == "" {
		return nil, errors.New("discord ID é obrigatório")
	}

	if userDTO.Email == "" {
		return nil, errors.New("email é obrogatório")
	}

	discordExists, err := usersRepo.ExistsByDicord(userDTO.DiscordID)
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar Discord ID: %v", err)
	}
	if discordExists {
		return nil, errors.New("já existe um usuário com essa ID do Discord")
	}

	emailExists, err := usersRepo.ExistsByEmail(userDTO.Email)
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar email: %v", err)
	}
	if emailExists {
		return nil, errors.New("já existe um usuário com esse email")
	}

	user := &entities.User{
		DiscordID: userDTO.DiscordID,
		Email:     userDTO.Email,
		FullName:  userDTO.FullName,
	}

	if err := usersRepo.Save(user); err != nil {
		return nil, fmt.Errorf("erro ao salvar usuário: %v", err)
	}

	return user, nil
}

func GetUserProfile(usersRepo usersRepo.UsersRepository, userID string) (*entities.User, error) {
	return usersRepo.GetByID(userID)
}

func ListUsers(usersRepo usersRepo.UsersRepository) ([]*entities.User, error) {
	users, err := usersRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuários: %v", err)
	}

	if len(users) == 0 {
		return nil, errors.New("nenhum usuário encontrado")
	}

	return users, nil
}

func UpdateUser(usersRepo usersRepo.UsersRepository, updateRequest dto.UpdateUserDTO, targetUserID string) error {
	user, err := usersRepo.GetByID(targetUserID)
	if err != nil {
		return errors.New("usuário não encontrado")
	}

	if updateRequest.DiscordID != nil {
		discordExists, err := usersRepo.ExistsByDicord(*updateRequest.DiscordID)
		if err != nil {
			return fmt.Errorf("erro ao verificar Discord ID: %v", err)
		}
		// evitar falso positivo ao atualizar com o mesmo valor
		if discordExists && user.DiscordID != *updateRequest.DiscordID {
			return errors.New("já existe um usuário com essa ID do Discord")
		}
		user.DiscordID = *updateRequest.DiscordID
	}

	if updateRequest.Email != nil {
		emailExists, err := usersRepo.ExistsByEmail(*updateRequest.Email)
		if err != nil {
			return fmt.Errorf("erro ao verificar email: %v", err)
		}
		if emailExists && user.Email != *updateRequest.Email {
			return errors.New("já existe um usuário com esse email")
		}
		user.Email = *updateRequest.Email
	}

	if updateRequest.FullName != nil {
		user.FullName = *updateRequest.FullName
	}

	if err := usersRepo.Save(user); err != nil {
		return errors.New("erro ao atualizar usuário")
	}

	return nil
}

func DeleteUser(usersRepo usersRepo.UsersRepository, userID string) error {
	if userID == "" {
		return errors.New("id é obrigatório")
	}

	err := usersRepo.DeleteByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("usuário não encontrado")
		}
		return err
	}

	return nil
}
