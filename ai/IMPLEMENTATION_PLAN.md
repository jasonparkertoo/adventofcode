# Implementation Plan for utils Package Refactoring

## 1. Error Handling Improvements

### Problem:
- The `ReadLines` function uses `panic` instead of returning errors
- The `NewData` function uses `panic` instead of returning errors  
- The condition in `Line` method is incorrect
- `os.Getwd()` error is ignored

### Solution:
1. Replace all `panic` calls with proper error returns
2. Fix the condition in `Line` method from `if n < len(d.lines) || n < 1` to `if n > len(d.lines) || n < 1`
3. Handle `os.Getwd()` error properly
4. Add error checking for scanner after reading
5. Update `NewData` to return `(Data, error)` instead of just `*Data`

## 2. Documentation Enhancements

### Problem:
- Missing package comment
- Missing comments for exported identifiers

### Solution:
1. Add package comment explaining the purpose
2. Add comments for all exported constants (`MsgPanic`, `MsgExpected`, `Example`, `Challenge`, `Year2024`, `Year2025`)
3. Add comments for all exported functions (`ReadLines`, `NewData`, `Lines`, `Line`, `TransformData`, `AsGrid`)

## 3. Resource Management

### Problem:
- File closing errors are ignored
- No explicit error checking after reading

### Solution:
1. Properly handle file closing with error checking
2. Add scanner error checking after reading
3. Use proper defer pattern that doesn't mask errors

## 4. Code Quality Improvements

### Problem:
- Inconsistent use of error handling patterns
- Lack of proper error wrapping
- Some code patterns don't follow idiomatic Go

### Solution:
1. Use `fmt.Errorf("context: %w", err)` for proper error wrapping
2. Ensure all functions that can fail properly return errors
3. Make sure defer statements properly handle cleanup
4. Maintain consistent code style with gofmt compliance

## 5. Test Impact Assessment

### Problem:
- Tests likely depend on current panic behavior
- Tests may fail with new error handling

### Solution:
1. Identify all tests that use `utils` package
2. Update tests to handle returned errors properly
3. Verify that existing test behavior is preserved
4. Create test cases for error conditions

## 6. Implementation Order

1. Fix `os.Getwd()` error handling and add proper error returns
2. Fix `Line` method condition
3. Add scanner error checking
4. Update `NewData` to return errors
5. Fix defer statement for proper file closing
6. Add missing documentation
7. Verify gofmt compliance
8. Update tests to handle new error patterns

## 7. Specific Code Changes Required

### File: `utils/utils.go`

1. **Package Comment**:
   ```go
   // Package utils provides utility functions for reading and processing input data
   // for Advent of Code solutions.
   ```

2. **Constants Documentation**:
   - Add comments for `MsgPanic`, `MsgExpected`
   - Add comments for `Example`, `Challenge`
   - Add comment for `Year2024`, `Year2025`

3. **Function Signatures**:
   - `NewData(ds DataSet, y string) (*Data, error)` instead of `*Data`
   - `ReadLines(y string, ds DataSet) ([]string, error)` - already correct

4. **Error Handling Fixes**:
   - Replace `panic(err.Error())` with proper error returns
   - Fix condition in `Line` method
   - Add error checking for `os.Getwd()`
   - Add error checking for scanner

5. **Resource Management**:
   - Improve defer statement for file closing
   - Proper error handling for file operations

## 8. Testing Considerations

### Tests That Will Be Affected:
- All tests that call `NewData()` will need to handle the new error return
- Tests that depend on `Line()` method may need updated assertions
- Tests that depend on panic behavior will need to be updated to expect errors instead

### Test Updates Required:
1. Update `NewData` calls in tests to handle errors
2. Update assertions to check for expected errors
3. Ensure existing test logic is preserved
4. Add test cases for error conditions