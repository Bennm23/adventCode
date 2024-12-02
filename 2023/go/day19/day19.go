package main


import (
    "advent/lib"
    "fmt"
    "regexp"
    "strconv"
    "strings"
)

type Condition struct {
    which string;
    lessThan bool;
    num int;
    outcome string;
    isElse bool;
}

type Workflow struct {
    conditions []Condition;
}

type WorkflowMap map[string]Workflow
type Part map[string]int

var intFind regexp.Regexp = *regexp.MustCompile(`[\d]+`)

func NewWorkflow(line string) (string, Workflow) {
    split := strings.Split(line, "{")
    key := split[0]
    values := split[1][:len(split[1])-1]

    commands := strings.Split(values, ",")


    workflow := Workflow{}

    for _, command := range commands {
        if !strings.Contains(command, ":") {
            //Then this is the else condition
            cond := Condition{which: command, isElse:true}
            workflow.conditions = append(workflow.conditions, cond)
            continue
        }
        condition, action := strings.Split(command, ":")[0], strings.Split(command, ":")[1]
        lessThan := strings.Contains(condition, "<")

        numLoc := intFind.FindStringIndex(condition)
        num, _ := strconv.Atoi(condition[numLoc[0]:numLoc[1]])
        c := Condition {
            which: string(condition[0]),
            lessThan: lessThan,
            num: num,
            outcome: action,
        }
        workflow.conditions = append(workflow.conditions, c)
    }
    return key, workflow
}

func p1(parts []Part, workflows WorkflowMap) int {
    sum := 0

    for _, part := range parts {

        location := "in"

        for location != "R" && location != "A" {
            work := workflows[location]
            location = work.evaluate(part)
        }
        
        if location == "A" {
            for _, v := range part {
                sum += v
            }
        }
    }
	return sum
}

func (flow Workflow) evaluate(part Part) string {
    for _, condition := range flow.conditions {

        partVal, found := part[condition.which]
        //If not found, then A, R or another mapping
        if !found {
            return condition.which
        }

        if condition.lessThan && partVal < condition.num {
            return condition.outcome
        } else if !condition.lessThan && partVal > condition.num {
            return condition.outcome
        }
    }
    panic("NO MATCH FOUND!!")
}

func solve() {
    groups := lib.ReadFileToGroups("day19.txt", "")
    workflows := make(WorkflowMap)
    parts := make([]Part, 0)

    for _, workflowString := range groups[0] {
        key, workflow := NewWorkflow(workflowString)
        workflows[key] = workflow
    }

    for _, rating := range groups[1] {
        scores := intFind.FindAllStringIndex(rating, -1)
        x, _ := strconv.Atoi(rating[scores[0][0]:scores[0][1]])
        m, _ := strconv.Atoi(rating[scores[1][0]:scores[1][1]])
        a, _ := strconv.Atoi(rating[scores[2][0]:scores[2][1]])
        s, _ := strconv.Atoi(rating[scores[3][0]:scores[3][1]])

        part := make(Part)
        part["x"] = x
        part["m"] = m
        part["a"] = a
        part["s"] = s
    
        parts = append(parts, part)
    }


    fmt.Println("Part 1 = ", p1(parts, workflows))//397134

	intervals := IntervalMap{}
	intervals["x"] = Interval{1, 4000}
	intervals["m"] = Interval{1, 4000}
	intervals["a"] = Interval{1, 4000}
	intervals["s"] = Interval{1, 4000}

	fmt.Println("Part 2 = ", p2(intervals, &workflows, "in"))//127517902575337

	fmt.Println("Intervals = ", intervals)

}
func main() {

	lib.RunAndPrintDuration(func() {
		solve()
	})//240-310
}

type IntervalMap map[string]Interval
type Interval struct {
	start, end int
}

func multiplyIntervals(intervals IntervalMap) int64 {
	product := int64(1)

	for _, interval := range intervals {
		product *= int64(interval.end - interval.start + 1)
	}

	return product
}

func handleOutcome(intervals IntervalMap, workflows *WorkflowMap, outcome string) int64 {
	if outcome == "A" {
		return multiplyIntervals(intervals)
	} else if outcome != "R" {
		return p2(intervals, workflows, outcome)
	}

	return 0
}

func p2(intervals IntervalMap, workflows *WorkflowMap, workflowId string) int64 {

	var combined int64 = 0

	workflow := (*workflows)[workflowId]

	for _, condition := range workflow.conditions {
		newIntervals := lib.CopyMap(intervals)
		fmt.Println("New Intervals = ", newIntervals)
		currInterval, found := newIntervals[condition.which]

		//Found is true if which == 'x|m|a|s'
		if found && condition.lessThan {
			prevInterval := intervals[condition.which]
			prevInterval.start = condition.num
			intervals[condition.which] = prevInterval

			currInterval.end = condition.num - 1
			newIntervals[condition.which] = currInterval

			fmt.Println("Less Than Condition Outcome = ", condition.outcome)
			combined += handleOutcome(newIntervals, workflows, condition.outcome)
			
		} else if found && !condition.lessThan {
			fmt.Println("Greater Than Condition Outcome = ", condition.outcome)
			prevInterval := intervals[condition.which]
			prevInterval.end = condition.num
			intervals[condition.which] = prevInterval

			currInterval.start = condition.num + 1
			newIntervals[condition.which] = currInterval

			combined += handleOutcome(newIntervals, workflows, condition.outcome)

		} else if condition.which != "R" {
			combined += handleOutcome(intervals, workflows, condition.which)
		}
	}

	return combined
}