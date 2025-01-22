package legal_entities

import (
	"time"

	"github.com/google/uuid"
)

type LegalEntities struct {
	UUID      uuid.UUID `json:"uuid" gorm:"type:uuid;default:gen_random_uuid();not null;primary_key:true"`
	Name      string    `json:"name" gorm:"type:varchar(500);default:'';not null"`
	CreatedAt time.Time `gorm:"type:timestamptz;default:now();not null;"`
	UpdatedAt time.Time `gorm:"type:timestamptz;default:now();not null;"`
	DeletedAt time.Time `gorm:"type:timestamptz;default:now();not null;"`
}
