package druid

import (
	"encoding/json"
	"fmt"
	"github.com/druid-io/druid-operator/pkg/apis/druid/v1alpha1"
	"github.com/druid-io/druid-operator/pkg/controller/druid/ext"
	"reflect"
)

var zkExtTypes = map[string]reflect.Type{}

func init() {
	zkExtTypes["default"] = reflect.TypeOf(ext.DefaultZkManager{})
}

// We might have to add more methods to this interface to enable extensions that completely manage
// deploy, upgrade and termination of zk cluster.
type zookeeperManager interface {
	Configuration() string
}

func createZookeeperManager(spec *v1alpha1.ZookeeperSpec) (zookeeperManager, error) {
	if t, ok := zkExtTypes[spec.Type]; ok {
		v := reflect.New(t).Interface()
		if err := json.Unmarshal(spec.Spec, v); err != nil {
			return nil, fmt.Errorf("Couldn't unmarshall zk type[%s]. Error[%s].", spec.Type, err.Error())
		} else {
			return v.(zookeeperManager), nil
		}
	} else {
		return nil, fmt.Errorf("Can't find type[%s] for Zookeeper Mgmt.", spec.Type)
	}
}
