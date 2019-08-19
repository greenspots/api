package controller

import (
	"fmt"
	"net/http"
	"sync"
)

type (
	DataInitContext struct {
		AppContext
	}
)

func (appCtx *DataInitContext) CreateIndexesHandler(w http.ResponseWriter, r *http.Request) {
	defer fmt.Fprintf(w, "Indexes Created.")
	var wgIndex sync.WaitGroup
	wgIndex.Add(len(config.INDEXES))
	//creates required indexes
	for _, index := range config.INDEXES {
		go func(idx string) {
			defer wgIndex.Done()
			appCtx.EsApi.CreateGspotIndexes(idx)
		}(index)
	}
	wgIndex.Wait()
}

func (appCtx *DataInitContext) CreateTreesHandler(w http.ResponseWriter, r *http.Request) {
	defer fmt.Fprintf(w, "Trees Created")
	var wgTree sync.WaitGroup
	wgTree.Add(len(config.TREE_NAMES))
	//index new document for each tree in the greenspot index
	for idx, tree := range config.TREE_NAMES {
		go func(idx int, tree string) {
			defer wgTree.Done()
			appCtx.EsApi.IndexNewGspot(tree, config.TREE_IDS[idx])
		}(idx, tree)
	}
	wgTree.Wait()
}
