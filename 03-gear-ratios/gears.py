from collections import defaultdict

with open("input.txt", "r") as f:
    data = f.read().strip()

lines = data.split("\n")
grid = [[c for c in line] for line in lines]
rows = len(grid)
cols = len(grid[0])

part1 = 0
nums = defaultdict(list)

# treversing grid rows
for r in range(len(grid)):
    gears = set()
    n = 0
    has_part = False
    # treversing grid cols
    for c in range(len(grid[r]) + 1):
        # if not in the edge and element is digit
        if c < cols and grid[r][c].isdigit():
            n = n * 10 + int(grid[r][c])
            for rr in [-1, 0, 1]:
                for cc in [-1, 0, 1]:
                    # if not in the edge of grid
                    if 0 <= r + rr < rows and 0 <= c + cc < cols:
                        char = grid[r + rr][c + cc]
                        if not char.isdigit() and char != ".":
                            has_part = True
                        if char == "*":
                            gears.add((r + rr, c + cc))
        # if not digit
        elif n > 0:
            for gear in gears:
                nums[gear].append(n)
            if has_part:
                part1 += n
            n = 0
            has_part = False
            gears = set()

print(part1)
part2 = 0
for k, v in nums.items():
    if len(v) == 2:
        part2 += v[0] * v[1]
print(part2)
