package jazz

import "github.com/beevik/etree"

type RootService struct {
	client     *Client
	base       string
	serviceXml *etree.Element
}

func (r *RootService) ServicesXml() (*etree.Element, error) {
	if r.serviceXml == nil {
		_, xml, err := r.client.SimpleGet(
			r.base+"/rootservices",
			"application/rdf+xml",
			"failed to get service XML",
			200)
		if err != nil {
			return nil, err
		}
		r.serviceXml = xml
	}
	return r.serviceXml, nil
}
