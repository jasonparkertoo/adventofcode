package day10

/*
 * Attribution:
 * part one: chatgpt
 * part two: https://www.reddit.com/r/adventofcode/comments/1pk87hl/2025_day_10_part_2_bifurcate_your_way_to_victory/
 */

import (
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"

	"adventofcode.dev/utils"
)

// LightMachine represents a light configuration with target state and button masks
// It models a system where buttons can toggle lights, and we need to find minimum presses to reach target
const (
	lightOn             = '#'
	lightOff            = '.'
	bruteForceThreshold = 20
	noSolutionPenalty   = 1_000_000
)

// LightMachine represents a light configuration with target state and button masks
// It models a system where buttons can toggle lights, and we need to find minimum presses to reach target
var buttonPattern = regexp.MustCompile(`\(([^)]*)\)`)

// LightMachine represents a light configuration with target state and button masks
// It models a system where buttons can toggle lights, and we need to find minimum presses to reach target
type LightMachine struct {
	targetState *big.Int
	buttonMasks []*big.Int
}

// StateWithPresses represents a light state along with the number of button presses needed to reach it
type StateWithPresses struct {
	state   *big.Int
	presses int
}

// ParityPattern represents a pattern of parities (0 or 1) for a set of variables
type ParityPattern struct {
	pattern []int
}

// FewestButtonPresses calculates the sum of minimum button presses for all lines
func FewestButtonPresses(d *utils.Data) (int, error) {
	lines := d.Lines()

	total := 0
	for _, line := range lines {
		machine, err := parseLightMachine(line)
		if err != nil {
			return 0, err
		}
		presses, err := findMinimumPresses(machine)
		if err != nil {
			return 0, err
		}
		total += presses
	}
	return total, nil
}

// parseLightMachine parses a line into a LightMachine structure
// It extracts the target light diagram and button masks from the input line
func parseLightMachine(line string) (*LightMachine, error) {
	diagramStart := strings.IndexByte(line, '[')
	diagramEnd := strings.IndexByte(line[diagramStart+1:], ']')
	if diagramEnd != -1 {
		diagramEnd += diagramStart + 1
	}

	if err := validateBrackets(line, diagramStart, diagramEnd); err != nil {
		return nil, err
	}

	diagramSection := line[diagramStart+1 : diagramEnd]
	buttonSection := line[diagramEnd+1:]

	targetState, err := parseLightDiagram(diagramSection)
	if err != nil {
		return nil, err
	}

	buttonMasks, err := parseButtonMasks(buttonSection)
	if err != nil {
		return nil, err
	}

	return &LightMachine{
		targetState: targetState,
		buttonMasks: buttonMasks,
	}, nil
}

// validateBrackets checks that the brackets in a line are properly formed
func validateBrackets(line string, start, end int) error {
	if start == -1 {
		return fmt.Errorf("missing opening '[' in: %s", line)
	}
	if end == -1 {
		return fmt.Errorf("missing closing ']' in: %s", line)
	}
	return nil
}

// parseLightDiagram converts a light diagram string into a big.Int representing the target state
// Light on positions are represented as 1 bits, light off positions as 0 bits
func parseLightDiagram(diagram string) (*big.Int, error) {
	result := new(big.Int)
	for i, ch := range diagram {
		if ch == lightOn {
			bit := new(big.Int).Lsh(big.NewInt(1), uint(i))
			result.Add(result, bit)
		} else if ch != lightOff {
			return nil, fmt.Errorf("invalid character '%c' in diagram", ch)
		}
	}
	return result, nil
}

// parseButtonMasks extracts button masks from a button specification string
// It parses the comma-separated indices within parentheses for each button
func parseButtonMasks(buttonSpec string) ([]*big.Int, error) {
	matches := buttonPattern.FindAllStringSubmatch(buttonSpec, -1)
	masks := make([]*big.Int, 0, len(matches))

	for _, match := range matches {
		if len(match) > 1 {
			mask, err := createButtonMask(strings.TrimSpace(match[1]))
			if err != nil {
				return nil, err
			}
			masks = append(masks, mask)
		}
	}

	return masks, nil
}

// createButtonMask creates a bit mask representing which lights are affected by a button
// It takes a comma-separated string of indices and sets the corresponding bits in a big.Int
func createButtonMask(commaSeparatedIndices string) (*big.Int, error) {
	if commaSeparatedIndices == "" {
		return big.NewInt(0), nil
	}

	mask := new(big.Int)
	indices := strings.SplitSeq(commaSeparatedIndices, ",")

	for indexStr := range indices {
		index, err := strconv.Atoi(strings.TrimSpace(indexStr))
		if err != nil {
			return nil, err
		}
		if index < 0 {
			return nil, fmt.Errorf("negative button index: %d", index)
		}
		mask.SetBit(mask, index, 1)
	}

	return mask, nil
}

// findMinimumPresses determines the minimum number of button presses needed to reach the target state
// Uses brute force for small numbers of buttons or meet-in-the-middle for larger ones
func findMinimumPresses(machine *LightMachine) (int, error) {
	if len(machine.buttonMasks) == 0 {
		if machine.targetState.Cmp(big.NewInt(0)) == 0 {
			return 0, nil
		}
		return 0, fmt.Errorf("no solution found")
	}

	if machine.targetState.Cmp(big.NewInt(0)) == 0 {
		return 0, nil
	}

	if len(machine.buttonMasks) <= bruteForceThreshold {
		return searchBruteForce(machine)
	}
	return searchMeetInMiddle(machine)
}

// searchBruteForce finds the minimum button presses by trying all possible combinations
// This approach is suitable for small numbers of buttons (≤ bruteForceThreshold)
func searchBruteForce(machine *LightMachine) (int, error) {
	totalCombinations := 1 << len(machine.buttonMasks)
	minPresses := -1

	for combo := range totalCombinations {
		pressCount := countBits(combo)
		resultState := computeStateFromCombo(machine.buttonMasks, combo)

		if resultState.Cmp(machine.targetState) == 0 {
			if minPresses == -1 || pressCount < minPresses {
				minPresses = pressCount
			}
		}
	}

	if minPresses == -1 {
		return 0, fmt.Errorf("no solution found")
	}
	return minPresses, nil
}

// searchMeetInMiddle uses the meet-in-the-middle algorithm to find minimum button presses
// Divides the buttons into two halves and combines results for better efficiency on larger inputs
func searchMeetInMiddle(machine *LightMachine) (int, error) {
	halfIndex := len(machine.buttonMasks) / 2

	firstHalfStates := generateAllStates(machine.buttonMasks, 0, halfIndex)
	secondHalfStates := generateAllStates(machine.buttonMasks, halfIndex, len(machine.buttonMasks))

	complementMap := make(map[string]int)
	for _, info := range secondHalfStates {
		neededFromFirst := new(big.Int).Xor(info.state, machine.targetState)
		key := neededFromFirst.String()

		if existing, ok := complementMap[key]; !ok || info.presses < existing {
			complementMap[key] = info.presses
		}
	}

	minPresses := -1
	for _, firstHalf := range firstHalfStates {
		key := firstHalf.state.String()
		if secondPresses, ok := complementMap[key]; ok {
			total := firstHalf.presses + secondPresses
			if minPresses == -1 || total < minPresses {
				minPresses = total
			}
		}
	}

	if minPresses == -1 {
		return 0, fmt.Errorf("no solution found")
	}
	return minPresses, nil
}

// generateAllStates computes all possible states for a given range of buttons
// Returns a slice of StateWithPresses containing the resulting state and number of presses for each combination
func generateAllStates(buttons []*big.Int, startIdx, endIdx int) []StateWithPresses {
	rangeSize := endIdx - startIdx
	totalCombinations := 1 << rangeSize
	states := make([]StateWithPresses, 0, totalCombinations)

	for combo := range totalCombinations {
		state := new(big.Int)
		presses := 0

		for i := range rangeSize {
			if (combo & (1 << i)) != 0 {
				state.Xor(state, buttons[startIdx+i])
				presses++
			}
		}

		states = append(states, StateWithPresses{state: state, presses: presses})
	}

	return states
}

// computeStateFromCombo calculates the resulting state for a given button combination
// It XORs the button masks that are pressed in the combination
func computeStateFromCombo(buttons []*big.Int, combo int) *big.Int {
	result := new(big.Int)
	for i := range buttons {
		if (combo & (1 << i)) != 0 {
			result.Xor(result, buttons[i])
		}
	}
	return result
}

// countBits counts the number of 1 bits in an integer
// This is used to count the number of buttons pressed in a combination
func countBits(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

// CalculateScore computes the total score for all configurations
func CalculateScore(d *utils.Data) (int, error) {
	lines := d.Lines()
	total := 0
	for _, line := range lines {
		score, err := scoreConfiguration(line)
		if err != nil {
			return 0, err
		}
		total += score
	}
	return total, nil
}

// scoreConfiguration calculates the score for a single configuration
// It parses coefficients and goal values, then uses dynamic programming to find the minimum cost
func scoreConfiguration(config string) (int, error) {
	parts := strings.Fields(config)
	if len(parts) < 2 {
		return 0, fmt.Errorf("invalid configuration format")
	}

	goal, err := parseIntegerList(parts[len(parts)-1])
	if err != nil {
		return 0, err
	}

	coefficients, err := parseCoefficients(parts, len(goal))
	if err != nil {
		return 0, err
	}

	patternCosts := buildPatternCostLookup(coefficients)
	memo := make(map[string]int)
	return solveRecursively(goal, patternCosts, memo), nil
}

// parseCoefficients extracts coefficient vectors from the configuration string
// It converts coefficient indices into binary presence vectors of the goal size
func parseCoefficients(parts []string, goalSize int) ([][]int, error) {
	coefficients := make([][]int, 0, len(parts)-2)

	for i := 1; i < len(parts)-1; i++ {
		coefficient, err := parseIntegerList(parts[i])
		if err != nil {
			return nil, err
		}
		vector := toBinaryPresenceVector(coefficient, goalSize)
		coefficients = append(coefficients, vector)
	}

	return coefficients, nil
}

// toBinaryPresenceVector converts a list of indices into a binary vector
// It sets bits at the specified indices to 1, others to 0
func toBinaryPresenceVector(indices []int, size int) []int {
	vector := make([]int, size)
	for _, index := range indices {
		if index < size {
			vector[index] = 1
		}
	}
	return vector
}

// solveRecursively uses dynamic programming with memoization to find the minimum cost
// It recursively breaks down the goal into smaller subproblems using valid patterns
func solveRecursively(goal []int, patternCosts map[string]map[string]int, memo map[string]int) int {
	key := intSliceKey(goal)
	if cached, ok := memo[key]; ok {
		return cached
	}

	if isAllZeros(goal) {
		return 0
	}

	parity := make([]int, len(goal))
	for i, v := range goal {
		parity[i] = v % 2
	}
	parityKey := intSliceKey(parity)

	validPatterns, ok := patternCosts[parityKey]
	if !ok {
		memo[key] = noSolutionPenalty
		return noSolutionPenalty
	}

	minCost := noSolutionPenalty
	for patternStr, patternCost := range validPatterns {
		pattern := parseKeyToIntSlice(patternStr)
		if !canApplyPattern(pattern, goal) {
			continue
		}

		reducedGoal := reduceGoal(goal, pattern)
		cost := patternCost + 2*solveRecursively(reducedGoal, patternCosts, memo)

		if cost < minCost {
			minCost = cost
		}
	}

	memo[key] = minCost
	return minCost
}

// isAllZeros checks if all elements in a slice are zero
func isAllZeros(list []int) bool {
	for _, v := range list {
		if v != 0 {
			return false
		}
	}
	return true
}

// canApplyPattern checks if a pattern can be applied to a goal
// A pattern can be applied if each element of the pattern is less than or equal to the corresponding goal element
func canApplyPattern(pattern, goal []int) bool {
	for i := range pattern {
		if pattern[i] > goal[i] {
			return false
		}
	}
	return true
}

// reduceGoal calculates the new goal after applying a pattern
// It subtracts the pattern from the goal and divides by 2
func reduceGoal(goal, pattern []int) []int {
	result := make([]int, len(goal))
	for i := range goal {
		result[i] = (goal[i] - pattern[i]) / 2
	}
	return result
}

// buildPatternCostLookup creates a lookup table of patterns and their minimum costs
// It generates all possible button combinations and maps their parity patterns to costs
func buildPatternCostLookup(coefficients [][]int) map[string]map[string]int {
	numButtons := len(coefficients)
	numVariables := len(coefficients[0])

	lookup := make(map[string]map[string]int)

	// Initialize all parity patterns
	for _, parityList := range generateAllParityPatterns(numVariables) {
		key := intSliceKey(parityList)
		lookup[key] = make(map[string]int)
	}

	// Generate all button combinations
	for k := 0; k <= numButtons; k++ {
		for _, buttonIndices := range combinations(numButtons, k) {
			pattern := sumCoefficients(coefficients, buttonIndices, numVariables)

			parity := make([]int, len(pattern))
			for i, v := range pattern {
				parity[i] = v % 2
			}
			parityKey := intSliceKey(parity)
			patternKey := intSliceKey(pattern)

			if _, exists := lookup[parityKey][patternKey]; !exists {
				lookup[parityKey][patternKey] = k
			}
		}
	}

	return lookup
}

// sumCoefficients computes the sum of coefficients for selected buttons
// It adds together the coefficient vectors for the specified button indices
func sumCoefficients(coefficients [][]int, buttonIndices []int, numVariables int) []int {
	sum := make([]int, numVariables)

	for _, buttonIdx := range buttonIndices {
		coefficient := coefficients[buttonIdx]
		for i := range numVariables {
			sum[i] += coefficient[i]
		}
	}

	return sum
}

// generateAllParityPatterns generates all possible parity patterns for a given number of bits
// Each pattern represents a different combination of 0s and 1s for the specified number of bits
func generateAllParityPatterns(numBits int) [][]int {
	total := 1 << numBits
	patterns := make([][]int, 0, total)

	for bitmask := range total {
		pattern := make([]int, numBits)
		for bit := range numBits {
			pattern[bit] = (bitmask >> bit) & 1
		}
		patterns = append(patterns, pattern)
	}

	return patterns
}

// combinations generates all combinations of k elements from a set of n elements
// It uses recursion to build all possible combinations
func combinations(n, k int) [][]int {
	result := [][]int{}
	current := []int{}
	generateCombinations(0, n, k, current, &result)
	return result
}

// generateCombinations recursively builds all combinations of k elements from a set
// It's a helper function for the combinations function
func generateCombinations(start, n, k int, current []int, result *[][]int) {
	if len(current) == k {
		combo := make([]int, len(current))
		copy(combo, current)
		*result = append(*result, combo)
		return
	}

	for i := start; i < n; i++ {
		current = append(current, i)
		generateCombinations(i+1, n, k, current, result)
		current = current[:len(current)-1]
	}
}

// parseIntegerList parses a string representation of a list of integers
// It supports formats with [], {}, or () brackets and returns a slice of integers
func parseIntegerList(s string) ([]int, error) {
	if len(s) < 2 {
		return nil, fmt.Errorf("invalid format: %s", s)
	}

	// Handle both [] and {} and ()
	if !((s[0] == '[' && s[len(s)-1] == ']') ||
		(s[0] == '{' && s[len(s)-1] == '}') ||
		(s[0] == '(' && s[len(s)-1] == ')')) {
		return nil, fmt.Errorf("invalid format: %s", s)
	}

	content := s[1 : len(s)-1]
	if content == "" {
		return []int{}, nil
	}

	parts := strings.Split(content, ",")
	result := make([]int, 0, len(parts))

	for _, part := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}

	return result, nil
}

// intSliceKey converts an integer slice to a string key
// Used for memoization and map keys in the dynamic programming approach
func intSliceKey(slice []int) string {
	var sb strings.Builder
	for i, v := range slice {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(strconv.Itoa(v))
	}
	return sb.String()
}

// parseKeyToIntSlice converts a string key back to an integer slice
// This is the inverse operation of intSliceKey and is used to reconstruct patterns
func parseKeyToIntSlice(key string) []int {
	if key == "" {
		return []int{}
	}
	parts := strings.Split(key, ",")
	result := make([]int, len(parts))
	for i, part := range parts {
		result[i], _ = strconv.Atoi(part)
	}
	return result
}
