package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name                    string              `json:"name,omitempty" validate:"required,min=2,max=128" gorm:"type:varchar(128);not null;"`
	Email                   string              `json:"email,omitempty" validate:"required,email" gorm:"type:varchar(128);uniqueIndex;not null;"`
	Password                string              `json:"password,omitempty" validate:"required,min=8,max=64" gorm:"type:varchar(64);not null;"`
	Role                    string              `json:"role" gorm:"type:varchar(50);"`
	PasswordChangedAt       time.Time           `json:"passwordChangedAt,omitempty"`
	Verified                bool                `json:"verified,omitempty"`
	LinkedIn                string              `json:"linkedIn,omitempty" gorm:"type:varchar(255);"`
	Facebook                string              `json:"facebook,omitempty" gorm:"type:varchar(255);"`
	Instagram               string              `json:"instagram,omitempty" gorm:"type:varchar(255);"`
	ProfilePic              string              `json:"profilePic,omitempty" gorm:"type:varchar(255);"`
	About                   string              `json:"about,omitempty" gorm:"type:text;"`
	SectionScore            []map[string]string `json:"sectionScore,omitempty" gorm:"-"`
	DailyAttemptedQuestions map[string][]string `json:"dailyAttemptedQuestions,omitempty" gorm:"-"`
}

type Question struct {
	ID             string    `gorm:"type:uuid;primaryKey;" json:"id,omitempty"`
	Question       string    `gorm:"type:text;not null;" json:"question,omitempty" validate:"required"`
	Subject        string    `gorm:"type:varchar(255);not null;" json:"subject,omitempty" validate:"required"`
	Tags           []string  `gorm:"type:varchar(255)[]" json:"tags,omitempty"`
	Exam           string    `gorm:"type:varchar(255);" json:"exam,omitempty"`
	Language       string    `gorm:"type:varchar(255);not null;" json:"language,omitempty" validate:"required"`
	Difficulty     int       `gorm:"type:int;" json:"difficulty,omitempty"`
	QuestionType   string    `gorm:"type:varchar(50);not null;" json:"questionType,omitempty" validate:"oneof=m-choice m-select numeric"`
	Options        []string  `gorm:"type:varchar(255)[]" json:"options,omitempty" validate:"required"`
	CorrectOptions int       `gorm:"type:int;not null;" json:"correctOptions,omitempty" validate:"required"`
	Explanation    string    `gorm:"type:text;" json:"explanation,omitempty"`
	CreatedAt      time.Time `gorm:"autoCreateTime;" json:"createdAt,omitempty"`
	EditedAt       time.Time `gorm:"autoUpdateTime;" json:"editedAt,omitempty"`
	CreatedById    string    `gorm:"type:uuid;" json:"createdBy,omitempty"`
	EditedByIds    []string  `gorm:"type:uuid[];" json:"editedBy,omitempty"`
}

type QuestionSet struct {
	ID            string     `gorm:"type:uuid;primaryKey;" json:"id,omitempty"`
	Name          string     `gorm:"type:varchar(255);not null;" json:"name,omitempty"`
	Questions     []Question `gorm:"many2many:question_set_questions;" json:"questions,omitempty"` // Many-to-many relationship
	Mode          string     `gorm:"type:varchar(50);not null;" json:"mode,omitempty" validate:"required,oneof=practice exam timed"`
	Subject       string     `gorm:"type:varchar(255);not null;" json:"subject,omitempty" validate:"required"`
	Tags          []string   `gorm:"type:varchar(255)[]" json:"tags,omitempty" validate:"required"`
	Exam          string     `gorm:"type:varchar(255);" json:"exam,omitempty"`
	Language      string     `gorm:"type:varchar(255);not null;" json:"language,omitempty" validate:"required"`
	TimeDuration  string     `gorm:"type:varchar(50);" json:"time,omitempty"`
	Difficulty    int        `gorm:"type:int;" json:"difficulty,omitempty"`
	Description   string     `gorm:"type:text;" json:"explanation,omitempty"`
	CreatedAt     time.Time  `gorm:"autoCreateTime;" json:"createdAt,omitempty"`
	EditedAt      time.Time  `gorm:"autoUpdateTime;" json:"editedAt,omitempty"`
	CreatedById   string     `gorm:"type:uuid;" json:"createdBy,omitempty"`
	EditedByIds   []string   `gorm:"type:uuid[];" json:"editedBy,omitempty"`
	TotalAttempts int        `gorm:"type:int;" json:"totalAttempts,omitempty"`
	MarksObtained []int      `gorm:"type:int[]" json:"marksObtained,omitempty"`
}

type QTest struct {
	ID                 string           `gorm:"type:uuid;primaryKey;" json:"id,omitempty"`
	Finished           bool             `gorm:"type:boolean;" json:"finished,omitempty"`
	Started            bool             `gorm:"type:boolean;" json:"started,omitempty"`
	Name               string           `gorm:"type:varchar(255);not null;" json:"name,omitempty"`
	Tags               []string         `gorm:"type:varchar(255)[]" json:"tags,omitempty"`
	QuestionSetID      string           `gorm:"type:uuid;not null;" json:"questionSetId,omitempty" validate:"required"`
	QuestionSet        QuestionSet      `gorm:"foreignKey:QuestionSetID;" json:"questionSet,omitempty"` // Belongs-to relationship
	TakenByID          string           `gorm:"type:uuid;not null;" json:"takenById,omitempty" validate:"required"`
	TakenBy            User             `gorm:"foreignKey:TakenByID;" json:"takenBy,omitempty"` // Belongs-to relationship
	NTotalQuestions    int              `gorm:"type:int;not null;" json:"nTotalQuestions,omitempty" validate:"required"`
	AllQuestionsIDs    map[string][]int `gorm:"type:jsonb;" json:"allQuestionsId,omitempty" validate:"required"` // JSONB column type
	CurrentQuestionNum int              `gorm:"type:int;not null;" json:"currentQuestionNum,omitempty" validate:"required"`
	QuestionIDsOrdered []string         `gorm:"type:varchar(255)[]" json:"questionIDsOrdered,omitempty" validate:"required"`
	NCorrectlyAnswered int              `gorm:"type:int;" json:"nCorrectlyAnswered,omitempty"`
	Rank               int              `gorm:"type:int;" json:"rank,omitempty"`
	TakenAtTime        time.Time        `gorm:"autoCreateTime;" json:"takenAt,omitempty"`
	Mode               string           `gorm:"type:varchar(50);not null;" json:"mode,omitempty" validate:"oneof=practice exam timed-practice"`
}
