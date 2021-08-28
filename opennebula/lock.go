package opennebula

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/OpenNebula/one/src/oca/go/src/goca/schemas/shared"
)

var locktypes = []string{"USE", "MANAGE", "ADMIN", "ALL", "UNLOCK"}

type Lockable interface {
	Lock(shared.LockLevel) error
	Unlock() error
}

func lockSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Lock level of the new resource: USE, MANAGE, ADMIN, ALL, UNLOCK",
		ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
			value := v.(string)

			if inArray(value, locktypes) < 0 {
				errors = append(errors, fmt.Errorf("Type %q must be one of: %s", k, strings.Join(locktypes, ",")))
			}

			return
		},
	}
}

func updateLock(lock interface{}, resourceController Lockable) error {
	var err error

	if lock.(string) == "UNLOCK" {
		err = resourceController.Unlock()
	} else {
		var level shared.LockLevel
		err = StringToLockLevel(lock.(string), &level)
		if err != nil {
			return err
		}
		err = resourceController.Lock(level)
	}
	if err != nil {
		return err
	}
}
