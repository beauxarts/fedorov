package cli

import (
	"github.com/beauxarts/scrinium/livelib_integration"
	"github.com/boggydigital/kvas"
	"golang.org/x/net/html"
)

func ReduceLiveLibBookDetails(id string, kv kvas.KeyValues) (map[string][]string, error) {
	det, err := kv.Get(id)
	defer det.Close()

	if err != nil {
		return nil, err
	}

	body, err := html.Parse(det)
	if err != nil {
		return nil, err
	}

	return livelib_integration.Reduce(body)
}

func MapLiveLibToFedorov(id string, lrdx map[string][]string, rdx map[string]map[string][]string) {
	//for lp, vals := range lrdx {
	//if p, ok := data.LiveLibPropertyMap[lp]; ok {
	//	rdx[p][id] = vals
	//} else {
	//	nod.Log("unknown LivLib property %s", lp)
	//}
	//}
}
