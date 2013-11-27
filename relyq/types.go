package relyq

import (
  "github.com/extemporalgenome/uuid"
)

// An arbitrary task object that can be directly used by applications
type ArbitraryTask map[string]interface{}

func (t ArbitraryTask) Id() []byte {
  if id, ok := t["id"]; ok {
    return []byte(id.(string))
  }
  id := uuid.NewRandom().String()
  t["id"] = id
  return []byte(id)
}

// A struct that implements Ider to be used in task objects for applications.
// Use like so:
//
//    type MyTask struct {
//      StructuredTask
//      OtherFields string
//    }
type StructuredTask struct {
  RqId []byte `json:"id"`
}

func (t *StructuredTask) Id() []byte {
  if t == nil {
    t = new(StructuredTask)
  }

  if t.RqId == nil {
    t.RqId = uuid.NewRandom().Bytes()
  }
  return t.RqId
}
