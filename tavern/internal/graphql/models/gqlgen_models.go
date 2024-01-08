// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"realm.pub/tavern/internal/ent/host"
)

type ClaimTasksInput struct {
	// The identity the beacon is authenticated as (e.g. 'root')
	Principal string `json:"principal"`
	// The hostname of the system the beacon is running on.
	Hostname string `json:"hostname"`
	// The platform the agent is operating on.
	HostPlatform host.Platform `json:"hostPlatform"`
	// The IP address of the hosts primary interface (if available).
	HostPrimaryIP *string `json:"hostPrimaryIP,omitempty"`
	// Unique identifier of the beacon, each running instance will be different.
	BeaconIdentifier string `json:"beaconIdentifier"`
	// Unique identifier of the underlying host system the beacon is running on.
	HostIdentifier string `json:"hostIdentifier"`
	// Name of the agent program the beacon is running as (e.g. 'imix')
	AgentIdentifier string `json:"agentIdentifier"`
}

type SubmitTaskResultInput struct {
	// ID of the task to submit results for.
	TaskID int `json:"taskID"`
	// Timestamp of when the task execution began. Format as RFC3339Nano.
	ExecStartedAt time.Time `json:"execStartedAt"`
	// Timestamp of when the task execution finished (set only if it has completed). Format as RFC3339Nano.
	ExecFinishedAt *time.Time `json:"execFinishedAt,omitempty"`
	// Output captured as the result of task execution.
	// Submitting multiple outputs will result in appending new output to the previous output.
	Output string `json:"output"`
	// Error message captured as the result of task execution failure.
	Error *string `json:"error,omitempty"`
}

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

var AllRole = []Role{
	RoleAdmin,
	RoleUser,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleUser:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
