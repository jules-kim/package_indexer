package main 

import "sync"

const (
	OK = "OK\n"
	FAIL = "FAIL\n"
	ERROR = "ERROR\n"
)

// Creates an instance of the PackageIndexer struct 
func CreatePackageIndexer() *PackageIndexer {
	return &PackageIndexer{
		packs: map[string]*Package{},
		mutex: &sync.Mutex{},
	}
}

// Handles requests by converting the packages in the 
// request into a package struct. HandleRequest also 
// directs the control flow to the correct function
// depending on the command given by the client 
func (pi *PackageIndexer) HandleRequest(req Request) string {
	if req.err != "" {
		return ERROR 
	}
	pack := Package{name: req.pack}								/* set the name of the package 	*/
	for _, name := range req.dep {								/* add package dependencies 	*/ 
		pack.deps = append(pack.deps, &Package{name: name})
	}
    switch req.comm {											/* check command type 			*/
    case "INDEX":
        return pi.Index(&pack)
    case "REMOVE":
        return pi.Remove(&pack)
    case "QUERY":
        return pi.Query(pack.name)
    }

    return ERROR 												/* return ERROR response 		*/ 
}

// Checks if a package's dependencies are already indexed.
// If not, the function returns FAIL. If the package
// already exists, the package's list of dependencies
// are updated. 
func (pi *PackageIndexer) Index(pack *Package) string {
	pi.mutex.Lock() 
	defer pi.mutex.Unlock()
	// foreach loop over the package's dependencies 
	for _, dep := range pack.deps {
		// query for each dependency
		if pi.Query(dep.name) == FAIL {
			// dependency not installed, cannot be indexed 
			return FAIL
		}
	}
	// package dependencies exist: update/add package  
	pi.packs[pack.name] = pack

	return OK 
}

// Checks if a package can be removed from the index. And if 
// it can, removes the package from the index. Returns OK
// if the package wasn't indexed. 
func (pi *PackageIndexer) Remove(pack *Package) string {
	pi.mutex.Lock() 
	defer pi.mutex.Unlock()
	// package not indexed, return OK
	if pi.Query(pack.name) == FAIL {
		return OK
	}
	// check if any packages depend on the package to be removed 
	// go through each pack in the index
	for _, p := range pi.packs {
		// go through each pack's dependencies 
		for _, dep := range p.deps {
			// a pack's dependency depends on the pack to be removed
			if pack.name == dep.name {
				return FAIL
			}
		}
	}
	// remove package from index 
	delete(pi.packs, pack.name)

	return OK 
}

// The function Query will search through the package
// indexer to see if the package is already indexed. 
// Not adding or removing anything from the indexer. 
// Only input is the name of the package in order
// to query for it by name. 
func (pi *PackageIndexer) Query(name string) string {
	if _, ok := pi.packs[name]; ok {							/* query indexer for pack 		*/ 
        return OK 												/* package indexed 				*/
    }

    return FAIL													/* package not indexed 			*/ 
}
