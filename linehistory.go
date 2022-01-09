package commandline

// linehistory tracks a series of lines
type linehistory []string

// Previous finds the entry in the list preceeding the given string
// if given string not already in history, it is added.
// if given string is the first entry in the list, an empty string is returned
func (h *linehistory) Previous(s string) string {
	i := indexOf(s, *h)
	if i == 0 {
		// already first entry in the list
		return ""
	}
	if i > 0 {
		// found in the list, return previous
		return (*h)[i-1]
	}
	// not know, add to the list before returning alternative
	var p string
	if len(*h) > 0 {
		p = (*h)[len(*h)-1]
	}
	h.Add(s)
	return p
}

// Next finds the entry in the list following the given string
// if given string not already in the list, it is added and the first entry in the list is returned.
func (h *linehistory) Next(s string) string {
	i := indexOf(s, *h)
	if i >= 0 {
		// found in the list, return next
		if i+1 < len(*h) {
			return (*h)[i+1]
		}
		// already last entry in list
		return ""
	}
	// not know, add and return same
	h.Add(s)
	if len(*h) == 0 {
		return ""
	}
	return (*h)[0]
}

func (h *linehistory) Add(s string) {
	if s == "" {
		return
	}
	// remove any duplicate already in the list
	if i := indexOf(s, *h); i >= 0 {
		if i == len(*h)-1 {
			// already the last entry
			return
		}
		// remove existing entry
		*h = append((*h)[:i], (*h)[i+1:]...)
	}
	*h = append((*h), s)
}

func indexOf(s string, ss []string) int {
	for i, sz := range ss {
		if s == sz {
			return i
		}
	}
	return -1
}
