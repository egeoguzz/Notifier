package monitorables

var problem_count int = 3 

var current_problem_count [len(tiles)]int 

var types_of_columns [len(tiles)]string
//ID int
func errorOccured (ID int, isErr bool) {
	if isErr == true {
		current_problem_count[ID - 1]++ 
		if current_problem_count[ID-1] == problem_count {
			Notify("will be handled later", isErr)
			fmt.Printf("There is a problem here ")
		}
	} else {
		if current_problem_count[ID-1] >= 3 {
			Notify("Will be handled later", isErr)
			fmt.Printf("Problem is not exist anymore")
		}
		resetErrCount(ID)
	}
	
	
	
}
func resetErrCount (ID int ) {
	current_problem_count[ID - 1] = 0 
}
func Initialization() {
	
	for i := 1; i <= len(tiles); i++ {
		current_problem_count[i-1] = 0 
		types_of_columns[i-1] = tiles.type
	}
}
func Notify (url string, isErr bool) {
	//
}
func printInitialization () {
	for i := 1; i <= len(current_problem_count);i++ {
		fmt.Printf("%d",current_problem_count[i-1])
		fmt.Printf("%s", types_of_columns[i-1])
	}
}


