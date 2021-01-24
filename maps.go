package maps

import (
  "strings"
  "strconv"
  "github.com/google/uuid"
  "github.com/golang/glog"
)


func AppendChildMap(parentMap *map[string]interface{}, child string, childMap map[string]interface{}) {
  for k, v := range childMap {
    (*parentMap)[strings.ToLower(child + "." + k)] = v
  }
}

func UnionMaps(srcMap *map[string]interface{}, newMap *map[string]interface{}) {
  for k, v := range (*newMap) {
    (*srcMap)[strings.ToLower(k)] = v
  }
}

func GetChildMap(parentMap *map[string]interface{}, child string) map[string]interface{} {
  res := make(map[string]interface{})
  zsChild := len(child) + 1
  seach := child + "."
  for k, v := range (*parentMap) {
    if zsChild < len(k) {
      if seach == k[:zsChild] {
        k2 := k[zsChild:]
        res[strings.ToLower(k2)] = v
      }
    }
  }
  return res
}

func GetSizeSlice(parentMap *map[string]interface{}) int {
  res := 0
  for k, _ := range (*parentMap) {
    i := strings.Index(k,".")
    if i >= 0 {
      seach := k[0:i]
      t, err := strconv.Atoi(seach)
      if err == nil {
        if res < t + 1 {
          res = t + 1
        }
      }
    }
  }
  return res
}

func GetMapFieldUUID(data *map[string]interface{}, fieldname string) (uuid.UUID, bool) {
  ids, ok := (*data)[strings.ToLower(fieldname)]
  if !ok {
    return uuid.Nil, false
  }
  id, oks := ids.(string)
  if !oks {
    return uuid.Nil, false
  }
  uid, err := uuid.Parse(id)
  if err != nil {
    glog.Errorf("ERR: getFieldUUID(%s) %v", id, err)
    return uuid.Nil, false
  }
  return uid, true
}

