package applicationstage

//Extended methods of applicationStages are descripted here

import (
	"JobHunterBackService/dbcrud"
	"fmt"
	"strconv"
	"time"
)

// Add a new applicationStage into the existing list
func (addingApplicationStage *ApplicationStage) AddApplicationStage() error {
	tx, err := dbcrud.OpenTransaction() //connect to the database and start transaction
	if err != nil {
		return err
	}
	//Get the list of all existing applicationStages
	allExistingApplicationStages, err := GetAllApplicationStagesOfVacancy(addingApplicationStage.VacancyId, tx)
	if err != nil {
		return fmt.Errorf("failed to get a list of existing applicationStages for the vacancy %q %w", addingApplicationStage.VacancyId, err)
	}

	//Add a new applicationStage to the list
	allExistingApplicationStages = append(allExistingApplicationStages[:addingApplicationStage.IndexNumber], addingApplicationStage)
	allExistingApplicationStages = append(allExistingApplicationStages, allExistingApplicationStages[addingApplicationStage.IndexNumber:]...)

	var previousApplicationStageStatus string //Status of the previous applicationStage
	//Save and update applicationStages
	for tempIndex := addingApplicationStage.IndexNumber; tempIndex < len(allExistingApplicationStages); tempIndex++ {
		// Get status of previous applicationStage
		if tempIndex == 0 {
			previousApplicationStageStatus = ""
		} else {
			previousApplicationStageStatus = allExistingApplicationStages[tempIndex-1].Status
		}
		// Renew status of a new appStage and all following.
		// Renew IndexNumber for all applicationStages, which must be after the added one
		// Then save a new one and update all others
		if err = allExistingApplicationStages[tempIndex].DefineStatus(previousApplicationStageStatus); err != nil { //Define a new status for every applicationStage
			return fmt.Errorf("failed to set a status for applicationStages with indexNumber %s %w", strconv.FormatInt(int64(tempIndex), 10), err)
		}
		if tempIndex == addingApplicationStage.IndexNumber { //saving of an addingApplicationStage
			if err = allExistingApplicationStages[tempIndex].Save(tx); err != nil {
				return fmt.Errorf("failed to save a new applicationStage %w", err)
			}
		} else {
			allExistingApplicationStages[tempIndex].IndexNumber = tempIndex //update all following applicationStages
			allExistingApplicationStages[tempIndex].UpdateDate = time.Now()
			if err = allExistingApplicationStages[tempIndex].Update(tx); err != nil {
				return fmt.Errorf("failed to update the existing applicationStage with id= %q %w", allExistingApplicationStages[tempIndex].Id, err)
			}
		}
	}

	//Commit changes
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction %w", err)
	}
	return nil
}

func (deletingApplicationStage *ApplicationStage) DeleteApplicationStage() error {
	tx, err := dbcrud.OpenTransaction() //connect to the database and start transaction
	if err != nil {
		return fmt.Errorf("failed to connect to DB %w", err)
	}

	//Get the list of all existing applicationStages
	allExistingApplicationStages, err := GetAllApplicationStagesOfVacancy(deletingApplicationStage.VacancyId, tx)
	if err != nil {
		return fmt.Errorf("failed to get a list of existing applicationStages for the vacancy %q %w", deletingApplicationStage.VacancyId, err)
	}

	// Delete the chosen applicationStage from slice and DB
	allExistingApplicationStages = append(allExistingApplicationStages[:deletingApplicationStage.IndexNumber], allExistingApplicationStages[deletingApplicationStage.IndexNumber+1:]...)
	if err = deletingApplicationStage.Delete(tx); err != nil {
		return fmt.Errorf("failed to delete the chosen applicationStage with id= %q %w", deletingApplicationStage.Id, err)
	}
	// Get status of previous applicationStage
	var previousApplicationStageStatus string //Status of the previous applicationStage
	// Update status and indexNumber of every following applicationStage
	for tempIndex := deletingApplicationStage.IndexNumber; tempIndex < len(allExistingApplicationStages); tempIndex++ {
		if tempIndex == 0 {
			previousApplicationStageStatus = ""
		} else {
			previousApplicationStageStatus = allExistingApplicationStages[tempIndex-1].Status
		}

		allExistingApplicationStages[tempIndex].IndexNumber = allExistingApplicationStages[tempIndex].IndexNumber - 1
		if err = allExistingApplicationStages[tempIndex].DefineStatus(previousApplicationStageStatus); err != nil { //Define a new status for every applicationStage
			return fmt.Errorf("failed to set an addingApplicationStage.Status %w", err)
		}
		allExistingApplicationStages[tempIndex].UpdateDate = time.Now()
		if err = allExistingApplicationStages[tempIndex].Update(tx); err != nil {
			return fmt.Errorf("failed to update the existing applicationStage with id= %q %w", allExistingApplicationStages[tempIndex].Id, err)
		}
	}
	//Commit changes
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction %w", err)
	}
	return nil
}

// ApplicationStage's status definition based on statuses of the previous and following applicationStages
func (applicationstage *ApplicationStage) DefineStatus(previousApplicationStageStatus string) error {
	statusDecisionMatrix := map[string]string{
		"":        "active",
		"active":  "waiting",
		"waiting": "waiting",
		"blocked": "blocked",
		"passed":  "active",
		"failed":  "blocked",
	}

	result, exists := statusDecisionMatrix[previousApplicationStageStatus]
	if !exists {
		return fmt.Errorf("indefinite status") //impossible status change
	}
	applicationstage.Status = result //set a new status
	return nil
}

func (updatingApplicationStage *ApplicationStage) UpdateApplicationStage() error {
	tx, err := dbcrud.OpenTransaction() //connect to the database and start transaction
	if err != nil {
		return err
	}
	//Get db version of the applicationStage with updatingApplicationStage.id
	dbApplicationStage := NewApplicationStage("", "", "", "", 0, time.Time{}, time.Time{}) //Create an applicationStage for reading from db
	if err := dbApplicationStage.FindById(updatingApplicationStage.Id, tx); err != nil {
		return err
	}
	//if status or indexNumber were changed then we need to update all following applicationStages. Else we need to update only updatingApplicationStage
	if dbApplicationStage.Status != updatingApplicationStage.Status || dbApplicationStage.IndexNumber != updatingApplicationStage.IndexNumber {
		//Get the list of all existing applicationStages
		allExistingApplicationStages, err := GetAllApplicationStagesOfVacancy(updatingApplicationStage.VacancyId, tx)
		if err != nil {
			return fmt.Errorf("failed to get a list of existing applicationStages for the vacancy %q %w", updatingApplicationStage.VacancyId, err)
		}
		//switch position of the updatingApplicationStage
		tempApplicationStage := allExistingApplicationStages[dbApplicationStage.IndexNumber]
		tempSlice := append(allExistingApplicationStages[:dbApplicationStage.IndexNumber], allExistingApplicationStages[dbApplicationStage.IndexNumber+1:]...) // Удаляем элемент oldPos
		tempSlice = append(tempSlice[:updatingApplicationStage.IndexNumber], append([]*ApplicationStage{tempApplicationStage}, tempSlice[updatingApplicationStage.IndexNumber:]...)...)
		allExistingApplicationStages = tempSlice
		tempSlice = nil

		var previousApplicationStageStatus string //Status of the previous applicationStage
		for tempIndex := 0; tempIndex < len(allExistingApplicationStages); tempIndex++ {
			if tempIndex == 0 {
				previousApplicationStageStatus = ""
			} else {
				previousApplicationStageStatus = allExistingApplicationStages[tempIndex-1].Status
			}
			if allExistingApplicationStages[tempIndex].IndexNumber != tempIndex {
				allExistingApplicationStages[tempIndex].IndexNumber = tempIndex
				if err = allExistingApplicationStages[tempIndex].DefineStatus(previousApplicationStageStatus); err != nil {
					return err
				}
				allExistingApplicationStages[tempIndex].UpdateDate = time.Now()
				allExistingApplicationStages[tempIndex].Update(tx)
			}
		}

	} else {
		updatingApplicationStage.UpdateDate = time.Now()
		if err = updatingApplicationStage.Update(tx); err != nil {
			return err
		}
	}
	//Commit changes
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction %w", err)
	}
	return nil
}
