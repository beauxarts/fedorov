package rest

import (
	"encoding/xml"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_pages"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
	"net/http"
	"os"
)

func GetContents(w http.ResponseWriter, r *http.Request) {

	// GET /contXml?id

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, nod.ErrorStr("missing book id"), http.StatusInternalServerError)
		return
	}

	var err error
	if rdx, err = rdx.RefreshReader(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	contentsDir, err := pathways.GetAbsRelDir(data.Contents)
	if err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	contReader, err := kevlar.NewKeyValues(contentsDir, kevlar.XmlExt)
	if err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	var contents *litres_integration.Contents

	contXml, err := contReader.Get(id)
	if err != nil {
		if !os.IsNotExist(err) {
			http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if contXml != nil {
			if err = xml.NewDecoder(contXml).Decode(&contents); err != nil {
				http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	if p := compton_pages.Contents(contents); p != nil {
		if err = p.WriteResponse(w); err != nil {
			http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		}
	}
}
