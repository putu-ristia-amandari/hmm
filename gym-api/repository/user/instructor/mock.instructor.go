package instructor

import (
	"gym-membership/business/instructor"

	"github.com/stretchr/testify/mock"
)

type MockInstructorRepository struct {
	data []*instructor.Instructor
	Mock mock.Mock
}

func (repository *MockInstructorRepository) GetAllInstructor() (instructors []*instructor.Instructor, err error) {
	return repository.data, nil
}
